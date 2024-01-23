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
package taskflow

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/wentaojin/dbms/logger"
	"github.com/wentaojin/dbms/model/rule"
	"go.uber.org/zap"

	"github.com/wentaojin/dbms/database"

	"github.com/wentaojin/dbms/database/oracle"
	"github.com/wentaojin/dbms/model"
	"github.com/wentaojin/dbms/model/datasource"
	"github.com/wentaojin/dbms/model/task"
	"github.com/wentaojin/dbms/pool"
	"github.com/wentaojin/dbms/proto/pb"
	"github.com/wentaojin/dbms/utils/constant"
	"github.com/wentaojin/dbms/utils/stringutil"
)

type StructMigrateTask struct {
	Ctx          context.Context
	TaskName     string
	TaskRuleName string
	TaskFlow     string
	MigrateTasks []*task.StructMigrateTask
	DatasourceS  *datasource.Datasource
	DatasourceT  *datasource.Datasource
	TaskParams   *pb.StructMigrateParam
}

func (st *StructMigrateTask) Start() error {
	logger.Info("struct migrate task get rules",
		zap.String("task_name", st.TaskName),
		zap.String("task_flow", st.TaskFlow))

	dbTypeSli := stringutil.StringSplit(st.TaskFlow, constant.StringSeparatorAite)
	buildInDatatypeRules, err := model.GetIBuildInDatatypeRuleRW().QueryBuildInDatatypeRule(st.Ctx, dbTypeSli[0], dbTypeSli[1])
	if err != nil {
		return err
	}
	buildInDefaultValueRules, err := model.GetBuildInDefaultValueRuleRW().QueryBuildInDefaultValueRule(st.Ctx, dbTypeSli[0], dbTypeSli[1])
	if err != nil {
		return err
	}

	groupSchemaTasks := make(map[string][]*task.StructMigrateTask)
	// the according schemaName, split task group for the migrateTasks
	groupSchemas := make(map[string]struct{})
	for _, m := range st.MigrateTasks {
		groupSchemas[m.SchemaNameS] = struct{}{}
	}

	for s, _ := range groupSchemas {
		var tasks []*task.StructMigrateTask
		for _, m := range st.MigrateTasks {
			if strings.EqualFold(s, m.SchemaNameS) {
				tasks = append(tasks, m)
			}
		}
		groupSchemaTasks[s] = tasks
	}

	// init database conn
	logger.Info("struct migrate task init connection",
		zap.String("task_name", st.TaskName),
		zap.String("task_flow", st.TaskFlow))
	targetSource, err := database.NewDatabase(st.Ctx, st.DatasourceT, "")
	if err != nil {
		return err
	}

	for schema, tasks := range groupSchemaTasks {
		schemaTaskTime := time.Now()

		sourceSource, err := database.NewDatabase(st.Ctx, st.DatasourceS, schema)
		if err != nil {
			return err
		}

		orac := sourceSource.(*oracle.Database)

		logger.Info("struct migrate task inspect task",
			zap.String("task_name", st.TaskName),
			zap.String("task_flow", st.TaskFlow))
		dbCharsetS, schemaCollationS, nlsComp, tableCollationS, dbCollationS, dbCharsetT, err := InspectStructMigrateTask(st.TaskName, st.TaskFlow, schema, orac, st.DatasourceS.ConnectCharset, st.DatasourceT.ConnectCharset)
		if err != nil {
			return err
		}

		// write schema
		logger.Info("struct migrate task get route",
			zap.String("task_name", st.TaskName),
			zap.String("task_flow", st.TaskFlow))
		schemaStartTime := time.Now()
		var (
			createSchema string
			schemaNameT  string
		)
		schemaRoute, err := model.GetIMigrateSchemaRouteRW().GetSchemaRouteRule(st.Ctx, &rule.SchemaRouteRule{TaskRuleName: st.TaskRuleName, SchemaNameS: schema})
		if err != nil {
			return err
		}

		switch stringutil.StringUpper(st.TaskParams.CaseFieldRule) {
		case constant.ParamValueStructMigrateCaseFieldNameLower:
			schemaNameT = stringutil.StringLower(schemaRoute.SchemaNameT)
		case constant.ParamValueStructMigrateCaseFieldNameUpper:
			schemaNameT = stringutil.StringUpper(schemaRoute.SchemaNameT)
		case constant.ParamValueStructMigrateCaseFieldNameOrigin:
			schemaNameT = schemaRoute.SchemaNameT
		default:
			return fmt.Errorf("the task_name [%v] task_flow [%v] case field rule [%s] is not support, please contact author or double check", st.TaskName, st.TaskFlow, st.TaskParams.CaseFieldRule)
		}

		logger.Info("struct migrate task process schema",
			zap.String("task_name", st.TaskName),
			zap.String("task_flow", st.TaskFlow),
			zap.String("schema_name_s", schemaRoute.SchemaNameS),
			zap.String("schema_name_t", schemaNameT))
		switch {
		case strings.EqualFold(st.TaskFlow, constant.TaskFlowOracleToTiDB) || strings.EqualFold(st.TaskFlow, constant.TaskFlowOracleToMySQL):
			if dbCollationS {
				targetSchemaCollation, ok := constant.MigrateTableStructureDatabaseCollationMap[st.TaskFlow][stringutil.StringUpper(schemaCollationS)][dbCharsetT]
				if !ok {
					return fmt.Errorf("oracle current task [%s] taskflow [%s] schema [%s] collation [%s] isn't support", st.TaskName, st.TaskFlow, schema, schemaCollationS)
				}
				createSchema = fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET %s COLLATE %s;", schemaNameT, dbCharsetT, targetSchemaCollation)
			} else {
				targetSchemaCollation, ok := constant.MigrateTableStructureDatabaseCollationMap[st.TaskFlow][stringutil.StringUpper(nlsComp)][dbCharsetT]
				if !ok {
					return fmt.Errorf("oracle current task [%s] taskflow [%s] schema [%s] nls_comp collation [%s] isn't support", st.TaskName, st.TaskFlow, schema, nlsComp)
				}
				createSchema = fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET %s COLLATE %s;", schemaNameT, dbCharsetT, targetSchemaCollation)
			}
		default:
			return fmt.Errorf("oracle current task [%s] taskflow [%s] schema [%s] isn't support, please contact author or reselect", st.TaskName, st.TaskFlow, schema)
		}

		encryptCreateSchema, err := stringutil.Encrypt(createSchema, []byte(constant.DefaultDataEncryptDecryptKey))
		if err != nil {
			return err
		}

		// schema create failed, return
		_, err = model.GetIStructMigrateTaskRW().CreateStructMigrateTask(st.Ctx, &task.StructMigrateTask{
			TaskName:        st.TaskName,
			SchemaNameS:     schema,
			TableTypeS:      constant.OracleDatabaseTableTypeSchema,
			SchemaNameT:     schemaNameT,
			TaskStatus:      constant.TaskDatabaseStatusSuccess,
			TargetSqlDigest: encryptCreateSchema,
			IsSchemaCreate:  constant.DatabaseIsSchemaCreateSqlYES,
			Duration:        time.Now().Sub(schemaStartTime).Seconds(),
		})
		if err != nil {
			return err
		}

		// direct write database -> schema
		if st.TaskParams.DirectWrite {
			_, err = targetSource.ExecContext(createSchema)
			if err != nil {
				return err
			}
		}
		logger.Info("struct migrate task process tables",
			zap.String("task_name", st.TaskName),
			zap.String("task_flow", st.TaskFlow))

		p := pool.NewPool(st.Ctx, int(st.TaskParams.MigrateThread),
			pool.WithTaskQueueSize(int(st.TaskParams.TaskQueueSize)),
			pool.WithPanicHandle(true),
			pool.WithResultCallback(func(r pool.Result) {
				smt := r.Task.Job.(*task.StructMigrateTask)
				if r.Error != nil {
					logger.Warn("struct migrate task",
						zap.String("task_name", st.TaskName),
						zap.String("task_flow", st.TaskFlow),
						zap.String("schema_name_s", smt.SchemaNameS),
						zap.String("table_name_s", smt.TableNameS),
						zap.Error(r.Error))

					errW := model.Transaction(st.Ctx, func(txnCtx context.Context) error {
						_, err = model.GetIStructMigrateTaskRW().UpdateStructMigrateTask(txnCtx,
							&task.StructMigrateTask{TaskName: smt.TaskName, SchemaNameS: smt.SchemaNameS, TableNameS: smt.TableNameS},
							map[string]string{
								"TaskStatus":  constant.TaskDatabaseStatusFailed,
								"ErrorDetail": r.Error.Error(),
							})
						if err != nil {
							return err
						}
						_, err = model.GetITaskLogRW().CreateLog(txnCtx, &task.Log{
							TaskName: smt.TaskName,
							LogDetail: fmt.Sprintf("%v [%v] struct migrate task [%v] source table [%v.%v] failed, please see [struct_migrate_task] detail",
								time.Now(),
								constant.TaskModeStructMigrate,
								smt.TaskName,
								smt.SchemaNameS,
								smt.TableNameS),
						})
						if err != nil {
							return err
						}
						return nil
					})
					if errW != nil {
						panic(fmt.Sprintf("oracle current task [%s] taskflow [%s] schema [%s] table [%s] struct migrate task panic: %v", r.Task.Group, st.TaskFlow, smt.SchemaNameS, smt.TableNameS, errW))
					}
				}
			}),
			pool.WithExecuteTask(func(ctx context.Context, t pool.Task) error {
				startTime := time.Now()
				smt := t.Job.(*task.StructMigrateTask)

				// if the schema table success, skip
				if strings.EqualFold(smt.TaskStatus, constant.TaskDatabaseStatusSuccess) {
					logger.Info("struct migrate task",
						zap.String("task_name", st.TaskName),
						zap.String("task_flow", st.TaskFlow),
						zap.String("schema_name_s", smt.SchemaNameS),
						zap.String("table_name_s", smt.TableNameS),
						zap.String("task_status", constant.TaskDatabaseStatusSuccess),
						zap.String("table task had done", "skip migrate"),
						zap.String("cost", time.Now().Sub(startTime).String()))
					return nil
				}
				if strings.EqualFold(smt.TaskStatus, constant.TaskDatabaseStatusRunning) {
					logger.Info("struct migrate task",
						zap.String("task_name", st.TaskName),
						zap.String("task_flow", st.TaskFlow),
						zap.String("schema_name_s", smt.SchemaNameS),
						zap.String("table_name_s", smt.TableNameS),
						zap.String("task_status", constant.TaskDatabaseStatusRunning),
						zap.String("table task has running", "current status may panic, skip migrate, please double check"),
						zap.String("cost", time.Now().Sub(startTime).String()))
					return nil
				}

				err = model.Transaction(ctx, func(txnCtx context.Context) error {
					_, err = model.GetIStructMigrateTaskRW().UpdateStructMigrateTask(txnCtx, &task.StructMigrateTask{
						TaskName:    t.Name,
						SchemaNameS: smt.SchemaNameS,
						TableNameS:  smt.TableNameS},
						map[string]string{
							"TaskStatus": constant.TaskDatabaseStatusRunning,
						})
					if err != nil {
						return err
					}
					_, err = model.GetITaskLogRW().CreateLog(txnCtx, &task.Log{
						TaskName: t.Name,
						LogDetail: fmt.Sprintf("%v [%v] the worker task [%v] source table [%v.%v] starting",
							stringutil.CurrentTimeFormatString(),
							stringutil.StringLower(constant.TaskModeStructMigrate),
							t.Name,
							smt.SchemaNameS,
							smt.TableNameS),
					})
					if err != nil {
						return err
					}
					return nil
				})
				if err != nil {
					return err
				}

				sourceTime := time.Now()
				dataSource := &Datasource{
					DatabaseS:        orac,
					SchemaNameS:      smt.SchemaNameS,
					TableNameS:       smt.TableNameS,
					TableTypeS:       smt.TableTypeS,
					CollationS:       dbCollationS,
					DBCharsetS:       dbCharsetS,
					SchemaCollationS: schemaCollationS,
					TableCollationS:  tableCollationS[smt.TableNameS],
					DBNlsCompS:       nlsComp,
				}

				attributes, err := database.IDatabaseTableAttributes(dataSource)
				if err != nil {
					return err
				}
				logger.Info("struct migrate task",
					zap.String("task_name", st.TaskName),
					zap.String("task_flow", st.TaskFlow),
					zap.String("schema_name_s", smt.SchemaNameS),
					zap.String("table_name_s", smt.TableNameS),
					zap.String("task_stage", "datasource reader"),
					zap.String("cost", time.Now().Sub(sourceTime).String()))
				ruleTime := time.Now()
				dataRule := &Rule{
					Ctx:                      ctx,
					TaskName:                 st.TaskName,
					TaskFlow:                 st.TaskFlow,
					TaskRuleName:             st.TaskRuleName,
					SchemaNameS:              smt.SchemaNameS,
					TableNameS:               smt.TableNameS,
					TablePrimaryKeyS:         attributes.PrimaryKey,
					TableColumnsS:            attributes.TableColumns,
					TableCommentS:            attributes.TableComment,
					CaseFieldRule:            st.TaskParams.CaseFieldRule,
					DBCollationS:             dbCollationS,
					DBCharsetS:               dbCharsetS,
					DBCharsetT:               dbCharsetT,
					BuildinDatatypeRules:     buildInDatatypeRules,
					BuildinDefaultValueRules: buildInDefaultValueRules,
				}

				rules, err := database.IDatabaseTableAttributesRule(dataRule)
				if err != nil {
					return err
				}
				logger.Info("struct migrate task",
					zap.String("task_name", st.TaskName),
					zap.String("task_flow", st.TaskFlow),
					zap.String("schema_name_s", smt.SchemaNameS),
					zap.String("table_name_s", smt.TableNameS),
					zap.String("task_stage", "rule generate"),
					zap.String("cost", time.Now().Sub(ruleTime).String()))

				tableTime := time.Now()
				dataTable := &Table{
					TaskName:            st.TaskName,
					TaskFlow:            st.TaskFlow,
					Datasource:          dataSource,
					TableAttributes:     attributes,
					TableAttributesRule: rules,
				}

				tableStruct, err := database.IDatabaseTableStruct(dataTable)
				if err != nil {
					return err
				}

				logger.Info("struct migrate task",
					zap.String("task_name", st.TaskName),
					zap.String("task_flow", st.TaskFlow),
					zap.String("schema_name_s", smt.SchemaNameS),
					zap.String("table_name_s", smt.TableNameS),
					zap.String("task_stage", "struct generate"),
					zap.String("table detail", dataTable.String()),
					zap.String("cost", time.Now().Sub(tableTime).String()))

				writerTime := time.Now()
				var w database.ITableStructDatabaseWriter
				w = NewStructMigrateDatabase(ctx, st.TaskName, st.TaskFlow, targetSource, startTime, tableStruct)

				if st.TaskParams.DirectWrite {
					err = w.SyncStructDatabase()
					if err != nil {
						return err
					}
					logger.Info("struct migrate task",
						zap.String("task_name", st.TaskName),
						zap.String("task_flow", st.TaskFlow),
						zap.String("schema_name_s", smt.SchemaNameS),
						zap.String("table_name_s", smt.TableNameS),
						zap.String("task_stage", "struct sync database"),
						zap.String("cost", time.Now().Sub(writerTime).String()))
				} else {
					err = w.WriteStructDatabase()
					if err != nil {
						return err
					}
					logger.Info("struct migrate task",
						zap.String("task_name", st.TaskName),
						zap.String("task_flow", st.TaskFlow),
						zap.String("schema_name_s", smt.SchemaNameS),
						zap.String("table_name_s", smt.TableNameS),
						zap.String("task_stage", "struct write database"),
						zap.String("cost", time.Now().Sub(writerTime).String()))
				}

				logger.Info("struct migrate task",
					zap.String("task_name", st.TaskName),
					zap.String("task_flow", st.TaskFlow),
					zap.String("schema_name_s", smt.SchemaNameS),
					zap.String("table_name_s", smt.TableNameS),
					zap.String("task_status", constant.TaskDatabaseStatusSuccess),
					zap.String("cost", time.Now().Sub(startTime).String()))

				return nil
			}))

		for _, taskJob := range tasks {
			p.AddTask(pool.Task{
				Name:  st.TaskName,
				Group: st.TaskName,
				Job:   taskJob,
			})
		}

		p.Wait()

		p.Release()

		err = sourceSource.Close()
		if err != nil {
			return err
		}

		// write file
		if !st.TaskParams.DirectWrite {
			writerTime := time.Now()
			var w database.ITableStructFileWriter
			w = NewStructMigrateFile(st.Ctx, st.TaskName, st.TaskFlow, schema, st.TaskParams.OutputDir)
			err = w.InitOutputFile()
			if err != nil {
				return err
			}
			err = w.SyncStructFile()
			if err != nil {
				return err
			}
			logger.Info("struct migrate task",
				zap.String("task_name", st.TaskName),
				zap.String("task_flow", st.TaskFlow),
				zap.String("schema_name_s", schema),
				zap.String("task_stage", "struct sync file"),
				zap.String("cost", time.Now().Sub(writerTime).String()))
		}
		logger.Info("struct migrate task",
			zap.String("task_name", st.TaskName),
			zap.String("task_flow", st.TaskFlow),
			zap.String("schema_name_s", schema),
			zap.String("cost", time.Now().Sub(schemaTaskTime).String()))
	}
	return nil
}

func InspectStructMigrateTask(taskName, taskFlow string, schemaNameS string, orac *oracle.Database, connectDBCharsetS, connectDBCharsetT string) (string, string, string, map[string]string, bool, string, error) {
	var (
		dbCharsetS       string
		dbCharsetT       string
		schemaCollationS string
	)
	tableCollationS := make(map[string]string)

	dbCharsetS = stringutil.StringUpper(connectDBCharsetS)

	dbCharsetT = constant.MigrateTableStructureDatabaseCharsetMap[taskFlow][dbCharsetS]
	if !strings.EqualFold(connectDBCharsetT, dbCharsetT) {
		return dbCharsetS, schemaCollationS, "", tableCollationS, false, dbCharsetT, fmt.Errorf("oracle current subtask [%s] taskflow [%s] schema [%s] mapping charset [%s] isn't equal with database connect charset [%s], please adjust database connect charset", taskName, taskFlow, schemaNameS, dbCharsetS, connectDBCharsetT)
	}

	dbCharset, err := orac.GetDatabaseCharset()
	if err != nil {
		return dbCharsetS, schemaCollationS, "", tableCollationS, false, dbCharsetT, err
	}
	nlsComp, nlsSort, err := orac.GetDatabaseCharsetCollation()
	if err != nil {
		return dbCharsetS, schemaCollationS, "", tableCollationS, false, dbCharsetT, err
	}
	if _, ok := constant.MigrateTableStructureDatabaseCollationMap[taskFlow][stringutil.StringUpper(nlsComp)][constant.MigrateTableStructureDatabaseCharsetMap[taskFlow][dbCharset]]; !ok {
		return dbCharsetS, schemaCollationS, nlsComp, tableCollationS, false, dbCharsetT,
			fmt.Errorf("oracle database nls comp [%s] , mysql db isn't support", nlsComp)
	}
	if _, ok := constant.MigrateTableStructureDatabaseCollationMap[taskFlow][stringutil.StringUpper(nlsSort)][constant.MigrateTableStructureDatabaseCharsetMap[taskFlow][dbCharset]]; !ok {
		return dbCharsetS, schemaCollationS, nlsComp, tableCollationS, false, dbCharsetT, fmt.Errorf("oracle database nls sort [%s] , mysql db isn't support", nlsSort)
	}

	if !strings.EqualFold(nlsSort, nlsComp) {
		return dbCharsetS, schemaCollationS, nlsComp, tableCollationS, false, dbCharsetT, fmt.Errorf("oracle database nls_sort [%s] and nls_comp [%s] isn't different, need be equal; because mysql db isn't support", nlsSort, nlsComp)
	}

	// whether the oracle version can specify table and field collation，if the oracle database version is 12.2 and the above version, it's specify table and field collation, otherwise can't specify
	// oracle database nls_sort/nls_comp value need to be equal, USING_NLS_COMP value is nls_comp
	oracleDBVersion, err := orac.GetDatabaseVersion()
	if err != nil {
		return dbCharsetS, schemaCollationS, nlsComp, tableCollationS, false, dbCharsetT, err

	}

	oracleCollation := false
	if stringutil.VersionOrdinal(oracleDBVersion) >= stringutil.VersionOrdinal(constant.OracleDatabaseTableAndColumnSupportVersion) {
		oracleCollation = true
	}

	if oracleCollation {
		schemaCollationS, err = orac.GetDatabaseSchemaCollation(schemaNameS)
		if err != nil {
			return dbCharsetS, schemaCollationS, nlsComp, tableCollationS, false, dbCharsetT, err

		}
		tableCollationS, err = orac.GetDatabaseSchemaTableCollation(schemaNameS, schemaCollationS)
		if err != nil {
			return dbCharsetS, schemaCollationS, nlsComp, tableCollationS, false, dbCharsetT, err

		}
	}
	return dbCharsetS, schemaCollationS, nlsComp, tableCollationS, false, dbCharsetT, nil
}
