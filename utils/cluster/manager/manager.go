/*
Copyright © 2020 Marvin

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package manager

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/wentaojin/dbms/utils/cluster"
	"gopkg.in/yaml.v2"

	"github.com/fatih/color"
	"github.com/wentaojin/dbms/utils/executor"

	"github.com/wentaojin/dbms/utils/cluster/operator"

	"github.com/wentaojin/dbms/logger/printer"

	"github.com/wentaojin/dbms/utils/ctxt"

	"github.com/wentaojin/dbms/utils/cluster/task"

	"go.uber.org/zap"

	"github.com/wentaojin/dbms/utils/stringutil"
)

type Controller struct {
	BasePath string
	Meta     cluster.IMetadata
	Logger   *printer.Logger
}

func New(basePath string, logger *printer.Logger) *Controller {
	return &Controller{
		BasePath: basePath,
		Meta:     new(cluster.Metadata),
		Logger:   logger,
	}
}

func (m *Controller) NewMetadata() cluster.IMetadata {
	return m.Meta
}

// GetMetaFilePath get meta file path
func (m *Controller) GetMetaFilePath(clusterName string) string {
	return filepath.Join(m.BasePath, clusterName, cluster.MetaFileName)
}

// CheckClusterNameConflict check if the cluster exist by checking the meta file.
func (m *Controller) CheckClusterNameConflict(clusterName string) (exist bool, err error) {
	fname := m.Path(clusterName, cluster.MetaFileName)

	_, err = os.Stat(fname)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// CheckClusterPortConflict check if the cluster exist by checking the meta file.
func (m *Controller) CheckClusterPortConflict(clusters map[string]cluster.Metadata, clusterName string, topo *cluster.Topology) error {
	type Entry struct {
		clusterName   string
		componentName string
		port          int
		instance      cluster.Instance
	}

	var currentEntries []Entry
	var existingEntries []Entry

	for name, metadata := range clusters {
		if name == clusterName {
			continue
		}

		metadata.GetTopology().IterInstance(func(inst cluster.Instance) {
			for _, port := range inst.UsedPort() {
				existingEntries = append(existingEntries, Entry{
					clusterName:   name,
					componentName: inst.ComponentName(),
					port:          port,
					instance:      inst,
				})
			}
		})
	}

	topo.IterInstance(func(inst cluster.Instance) {
		for _, port := range inst.UsedPort() {
			currentEntries = append(currentEntries, Entry{
				componentName: inst.ComponentName(),
				port:          port,
				instance:      inst,
			})
		}
	})

	for _, p1 := range currentEntries {
		for _, p2 := range existingEntries {
			if p1.port == p2.port {
				// build the conflict info
				properties := map[string]string{
					"ThisPort":       strconv.Itoa(p1.port),
					"ThisComponent":  p1.componentName,
					"ThisHost":       p1.instance.InstanceHost(),
					"ExistCluster":   p2.clusterName,
					"ExistPort":      strconv.Itoa(p2.port),
					"ExistComponent": p2.componentName,
					"ExistHost":      p2.instance.InstanceHost(),
				}

				// build error message
				zap.L().Info("Meet deploy port conflict", zap.Any("info", properties))
				return fmt.Errorf(`deploy port conflicts to an existing cluster
The port you specified in the topology file is:
  Port:      %v
  Component: %v

It conflicts to a port in the existing cluster:
  Existing Cluster Name: %v
  Existing Port:         %v
  Existing Component:    %v

Please change to use another port or another host.`,
					properties["ThisPort"], properties["ThisComponent"], properties["ExistCluster"], properties["ExistPort"], properties["ExistComponent"])
			}
		}
	}

	return nil
}

// SaveMetadata save the meta with specified cluster name.
func (m *Controller) SaveMetadata(clusterName string, meta *cluster.Metadata) error {
	data, err := yaml.Marshal(meta)
	if err != nil {
		return err
	}
	err = stringutil.SaveFileWithBackup(m.Path(clusterName, cluster.MetaFileName), data, m.Path(clusterName, cluster.BackupDirName))
	if err != nil {
		return err
	}
	return nil
}

// CheckClusterDirConflict checks cluster dir conflict or overlap
func (m *Controller) CheckClusterDirConflict(clusters map[string]cluster.Metadata, clusterName string, topo *cluster.Topology) error {
	instanceDirAccessor := dirAccessors()
	var currentEntries []DirEntry
	var existingEntries []DirEntry

	// rebuild existing disk status
	for name, metadata := range clusters {
		if name == clusterName {
			continue
		}

		etopo := metadata.GetTopology()

		etopo.IterInstance(func(inst cluster.Instance) {
			for _, dirAccessor := range instanceDirAccessor {
				existingEntries = appendEntries(name, topo, inst, dirAccessor, existingEntries)
			}
		})
	}

	topo.IterInstance(func(inst cluster.Instance) {
		for _, dirAccessor := range instanceDirAccessor {
			currentEntries = appendEntries(clusterName, topo, inst, dirAccessor, currentEntries)
		}
	})

	for _, d1 := range currentEntries {
		// data_dir is relative to deploy_dir by default, so they can be with
		// same (sub) paths as long as the deploy_dirs are different
		if d1.dirKind == "data directory" && !strings.HasPrefix(d1.dir, "/") {
			continue
		}
		for _, d2 := range existingEntries {
			if d1.instance.InstanceHost() != d2.instance.InstanceHost() {
				continue
			}

			if d1.dir == d2.dir && d1.dir != "" {
				properties := map[string]string{
					"ThisDirKind":    d1.dirKind,
					"ThisDir":        d1.dir,
					"ThisComponent":  d1.instance.ComponentName(),
					"ThisHost":       d1.instance.InstanceHost(),
					"ExistCluster":   d2.clusterName,
					"ExistDirKind":   d2.dirKind,
					"ExistDir":       d2.dir,
					"ExistComponent": d2.instance.ComponentName(),
					"ExistHost":      d2.instance.InstanceHost(),
				}
				zap.L().Info("Meet deploy directory conflict", zap.Any("info", properties))
				return fmt.Errorf(`deploy directory conflicts to an existing cluster
The directory you specified in the topology file is:
  Directory: %v %v
  Component: %v %v

It conflicts to a directory in the existing cluster:
  Existing Cluster Name: %v
  Existing Directory:    %v %v
  Existing Component:    %v %v

Please change to use another directory or another host.
`, properties["ThisDirKind"], properties["ThisDir"], properties["ThisComponent"], properties["ThisHost"],
					properties["ExistCluster"], properties["ExistDirKind"], properties["ExistDir"], properties["ExistComponent"], properties["ExistHost"])
			}
		}
	}
	return CheckClusterDirOverlap(currentEntries)
}

// GetAllClusters get a metadata list of all clusters deployed by current user
func (m *Controller) GetAllClusters() (map[string]cluster.Metadata, error) {
	clusters := make(map[string]cluster.Metadata)
	names, err := m.ListClusterNameAll()
	if err != nil {
		return nil, err
	}
	for _, name := range names {
		var metadata cluster.Metadata
		err = m.Metadata(name, metadata)
		if err != nil {
			return nil, err
		}
		clusters[name] = metadata
	}
	return clusters, nil
}

// Metadata tries to read the metadata of a cluster from file
func (m *Controller) Metadata(clusterName string, meta any) error {
	fname := m.Path(clusterName, cluster.MetaFileName)

	yamlFile, err := os.ReadFile(fname)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, meta)
	if err != nil {
		return err
	}
	return nil
}

// ListClusterNameAll list the all of cluster
func (m *Controller) ListClusterNameAll() ([]string, error) {
	return GetClusterNameList(m.BasePath)
}

// ScaleOutLockedErr Determine whether there is a lock, and report an error if it exists
func (m *Controller) ScaleOutLockedErr(clusterName string) error {
	if locked, err := m.IsScaleOutLocked(clusterName); locked {
		return fmt.Errorf("scale-out file lock already exists, please waitting and retry, error detail: %v", err)
	}
	return nil
}

// IsScaleOutLocked judge the cluster scale-out file lock status
func (m *Controller) IsScaleOutLocked(clusterName string) (locked bool, err error) {
	fname := m.Path(clusterName, cluster.ScaleOutLockName)

	_, err = os.Stat(fname)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return true, nil
}

// NewScaleOutLock save the meta with specified cluster name.
func (m *Controller) NewScaleOutLock(clusterName string, topo *cluster.Topology) error {
	locked, err := m.IsScaleOutLocked(clusterName)
	if err != nil {
		return err
	}
	if locked {
		return m.ScaleOutLockedErr(clusterName)
	}

	lockFile := m.Path(clusterName, cluster.ScaleOutLockName)

	data, err := yaml.Marshal(topo)
	if err != nil {
		return err
	}

	err = stringutil.WriteFile(lockFile, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

// ReleaseScaleOutLock remove the scale-out file lock with specified cluster
func (m *Controller) ReleaseScaleOutLock(clusterName string) error {
	return os.Remove(m.Path(clusterName, cluster.ScaleOutLockName))
}

// Remove remove the data with specified cluster name.
func (m *Controller) Remove(clusterName string) error {
	return os.RemoveAll(m.Path(clusterName))
}

// If the flag --topology-file is specified, the first 2 steps will be skipped.
// 1. Write Topology to a temporary file.
// 2. Open file in editor.
// 3. Check and update Topology.
// 4. Save meta file.
func (m *Controller) EditTopology(origTopo *cluster.Topology, data []byte, newTopoFile string, skipConfirm bool) (*cluster.Topology, error) {
	var name string
	if newTopoFile == "" {
		file, err := os.CreateTemp(os.TempDir(), "*")
		if err != nil {
			return nil, err
		}

		name = file.Name()

		_, err = io.Copy(file, bytes.NewReader(data))
		if err != nil {
			return nil, err
		}

		err = file.Close()
		if err != nil {
			return nil, err
		}

		err = stringutil.OpenFileInEditor(name)
		if err != nil {
			return nil, err
		}
	} else {
		name = newTopoFile
	}

	// Now user finish editing the file or user has provided the new topology file
	newData, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	newTopo := m.NewMetadata().GetTopology()
	err = yaml.UnmarshalStrict(newData, newTopo)
	if err != nil {
		fmt.Print(color.RedString("New topology could not be saved: "))
		m.Logger.Infof("Failed to parse topology file: %v", err)
		if newTopoFile == "" {
			if pass, _ := stringutil.PromptForConfirmNo("Do you want to continue editing? [Y/n]: "); !pass {
				return m.EditTopology(origTopo, newData, newTopoFile, skipConfirm)
			}
		}
		m.Logger.Infof("Nothing changed.")
		return nil, nil
	}

	// report error if immutable field has been changed
	if err := stringutil.ValidateSpecDiff(origTopo, newTopo); err != nil {
		fmt.Print(color.RedString("New topology could not be saved: "))
		m.Logger.Errorf("%s", err)
		if newTopoFile == "" {
			if pass, _ := stringutil.PromptForConfirmNo("Do you want to continue editing? [Y/n]: "); !pass {
				return m.EditTopology(origTopo, newData, newTopoFile, skipConfirm)
			}
		}
		m.Logger.Infof("Nothing changed.")
		return nil, nil
	}

	origData, err := yaml.Marshal(origTopo)
	if err != nil {
		return nil, err
	}

	if bytes.Equal(origData, newData) {
		m.Logger.Infof("The file has nothing changed")
		return nil, nil
	}

	stringutil.ShowDiff(string(origData), string(newData), os.Stdout)

	if !skipConfirm {
		if err := stringutil.PromptForConfirmOrAbortError(
			color.HiYellowString("Please check change highlight above, do you want to apply the change? [y/N]:"),
		); err != nil {
			return nil, err
		}
	}

	return newTopo, nil
}

func GetClusterNameList(basePath string) ([]string, error) {
	var clusterNames []string
	fileInfos, err := os.ReadDir(basePath)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	for _, info := range fileInfos {
		if info.IsDir() {
			if stringutil.IsPathNotExist(filepath.Join(append([]string{
				basePath,
				info.Name(),
				cluster.MetaFileName,
			})...)) {
				continue
			}
			clusterNames = append(clusterNames, info.Name())
		}
	}
	return clusterNames, nil
}

// InitClusterMetadataDir init cluster metadata dir
func (m *Controller) InitClusterMetadataDir(clusterName string) error {
	err := stringutil.PathNotExistOrCreate(m.Path(clusterName, cluster.CacheDirName))
	if err != nil {
		return err
	}
	err = stringutil.PathNotExistOrCreate(m.Path(clusterName, cluster.PatchDirName))
	if err != nil {
		return err
	}
	err = stringutil.PathNotExistOrCreate(m.Path(clusterName, cluster.SshDirName))
	if err != nil {
		return err
	}
	err = stringutil.PathNotExistOrCreate(m.Path(clusterName, cluster.AuditDirName))
	if err != nil {
		return err
	}
	err = stringutil.PathNotExistOrCreate(m.Path(clusterName, cluster.BackupDirName))
	if err != nil {
		return err
	}
	err = stringutil.PathNotExistOrCreate(m.Path(clusterName, cluster.LockDirName))
	if err != nil {
		return err
	}
	return nil
}

// ConfirmTopology confirm topology
func (m *Controller) ConfirmTopology(clusterName, clusterVersion string, topo *cluster.Topology, patchedRoles stringutil.StringSet) error {
	m.Logger.Infof("Please confirm your topology:")

	cyan := color.New(color.FgCyan, color.Bold)
	fmt.Printf("Cluster type:    %s\n", "DBMS")
	fmt.Printf("Cluster name:    %s\n", cyan.Sprint(clusterName))
	fmt.Printf("Cluster version: %s\n", cyan.Sprint(clusterVersion))

	clusterTable := [][]string{
		// Header
		{"Role", "Host"},
	}

	clusterTable[0] = append(clusterTable[0], "Ports", "OS/Arch", "Directories")

	topo.IterInstance(func(inst cluster.Instance) {
		comp := inst.ComponentName()
		if patchedRoles.Exist(comp) || inst.IsPatched() {
			comp += " (patched)"
		}
		instInfo := []string{comp, inst.InstanceHost()}
		instInfo = append(instInfo,
			stringutil.JoinInt(inst.UsedPort(), "/"),
			stringutil.OsArch(inst.OS(), inst.Arch()),
			strings.Join(inst.UsedDir(), ","))

		clusterTable = append(clusterTable, instInfo)
	})

	stringutil.PrintTable(clusterTable, true)

	m.Logger.Warnf("Attention:")
	m.Logger.Warnf("    1. If the topology is not what you expected, check your yaml file.")
	m.Logger.Warnf("    2. Please confirm there is no port/directory conflicts in same host.")
	if len(patchedRoles) != 0 {
		m.Logger.Errorf("    3. The component marked as `patched` has been replaced by previous patch commanm.")
	}

	return stringutil.PromptForConfirmOrAbortError("Do you want to continue? [y/N]: ")
}

// Path returns the full path to a subpath (file or directory) of a
// cluster, it is a subdir in the profile dir of the user, with the cluster name
// as its name.
func (m *Controller) Path(cluster string, subpath ...string) string {
	if cluster == "" {
		// keep the same behavior with legacy version of TiUP, we could change
		// it in the future if needed.
		cluster = "default-cluster"
	}

	return filepath.Join(append([]string{
		m.BasePath,
		cluster,
	}, subpath...)...)
}

// FillHost fill full host cpu-arch and kernel-name
func (m *Controller) FillHost(s, p *operator.SSHConnectionProps, topo *cluster.Topology, gOpt *operator.Options, user string, sudo bool) error {
	if err := m.fillHostArchOrOS(s, p, topo, gOpt, user, cluster.FullArchType, sudo); err != nil {
		return err
	}
	return m.fillHostArchOrOS(s, p, topo, gOpt, user, cluster.FullOSType, sudo)
}

// fillHostArchOrOS full host cpu-arch or kernel-name
func (m *Controller) fillHostArchOrOS(s, p *operator.SSHConnectionProps, topo *cluster.Topology, gOpt *operator.Options, user string, fullType cluster.FullHostType, sudo bool) error {
	globalSSHType := topo.GlobalOptions.SSHType
	hostArchOrOS := map[string]string{}
	var detectTasks []*task.StepDisplay

	topo.IterInstance(func(inst cluster.Instance) {
		if fullType == cluster.FullOSType {
			if inst.OS() != "" {
				return
			}
		} else if inst.Arch() != "" {
			return
		}

		if _, ok := hostArchOrOS[inst.InstanceHost()]; ok {
			return
		}
		hostArchOrOS[inst.InstanceHost()] = ""

		tf := task.NewBuilder(m.Logger).
			RootSSH(
				inst.InstanceHost(),
				inst.InstanceSshPort(),
				user,
				s.Password,
				s.IdentityFile,
				s.IdentityFilePassphrase,
				gOpt.SSHTimeout,
				gOpt.OptTimeout,
				gOpt.SSHProxyHost,
				gOpt.SSHProxyPort,
				gOpt.SSHProxyUser,
				p.Password,
				p.IdentityFile,
				p.IdentityFilePassphrase,
				gOpt.SSHProxyTimeout,
				executor.SSHType(globalSSHType),
				gOpt.SSHType,
				sudo,
			)

		switch fullType {
		case cluster.FullOSType:
			tf = tf.Shell(inst.InstanceHost(), "uname -s", "", false)
		default:
			tf = tf.Shell(inst.InstanceHost(), "uname -m", "", false)
		}
		detectTasks = append(detectTasks, tf.BuildAsStep(fmt.Sprintf("  - Detecting node %s %s info", inst.InstanceHost(), string(fullType))))
	})
	if len(detectTasks) == 0 {
		return nil
	}

	ctx := ctxt.New(
		context.Background(),
		gOpt.Concurrency,
		m.Logger,
	)
	t := task.NewBuilder(m.Logger).
		ParallelStep(fmt.Sprintf("+ Detect CPU %s Name", string(fullType)), false, detectTasks...).
		Build()

	if err := t.Execute(ctx); err != nil {
		return fmt.Errorf("failed to fetch cpu-arch or kernel-name, error detail: %v", err)
	}

	for host := range hostArchOrOS {
		stdout, _, ok := ctxt.GetInner(ctx).GetOutputs(host)
		if !ok {
			return fmt.Errorf("no check results found for %s", host)
		}
		hostArchOrOS[host] = strings.Trim(string(stdout), "\n")
	}
	return topo.FillHostArchOrOS(hostArchOrOS, fullType)
}