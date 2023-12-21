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
package master

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/wentaojin/dbms/utils/configutil"

	"github.com/BurntSushi/toml"
	"github.com/wentaojin/dbms/logger"
	"github.com/wentaojin/dbms/version"
	"go.uber.org/zap"
)

// Config is the configuration for dbms-master
type Config struct {
	FlagSet       *flag.FlagSet             `json:"-"`
	ConfigFile    string                    `toml:"config-file" json:"config-file"`
	OpenAPI       bool                      `toml:"openapi" json:"openapi"`
	MasterOptions *configutil.MasterOptions `toml:"master" json:"master"`
	LogConfig     *logger.Config            `toml:"log" json:"log"`

	PrintVersion bool `json:"-"`
}

func NewConfig() *Config {
	cfg := &Config{}
	cfg.FlagSet = flag.NewFlagSet("dbms master", flag.ContinueOnError)
	fs := cfg.FlagSet
	fs.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage of dbms master:")
		fs.PrintDefaults()
	}
	fs.BoolVar(&cfg.PrintVersion, "V", false, "prints version and exit")
	fs.BoolVar(&cfg.OpenAPI, "openapi", true, "enable openapi")
	fs.StringVar(&cfg.ConfigFile, "config", "", "path to config file")
	return cfg
}

func (c *Config) Parse(args []string) error {
	err := c.FlagSet.Parse(args)
	switch err {
	case nil:
	case flag.ErrHelp:
		os.Exit(0)
	default:
		os.Exit(2)
	}

	if c.PrintVersion {
		fmt.Println(version.GetRawVersionInfo())
		os.Exit(0)
	}

	if c.ConfigFile != "" {
		if err = c.configFromFile(c.ConfigFile); err != nil {
			return err
		}
	}

	// Parse again to replace with command line options.
	err = c.FlagSet.Parse(args)
	if err != nil {
		return err
	}

	if len(c.FlagSet.Args()) != 0 {
		return fmt.Errorf("master config invalid flag: [%v]", c.FlagSet.Args())
	}

	return c.adjust()
}

// adjust configs
func (c *Config) adjust() error {
	if c.OpenAPI {
		logger.Warn("openapi is a GA feature and removed from experimental features, so this configuration may have no affect in feature release, please set openapi=true in dm-master config file")
	}
	return nil
}

// configFromFile loads config from file.
func (c *Config) configFromFile(path string) error {
	_, err := toml.DecodeFile(path, c)
	if err != nil {
		return fmt.Errorf("config decode from file failed: %v", err)
	}
	return nil
}

func (c *Config) String() string {
	cfg, err := json.Marshal(c)
	if err != nil {
		logger.Error("marshal to json", zap.Reflect("master config", c), zap.Error(err))
	}

	return string(cfg)
}
