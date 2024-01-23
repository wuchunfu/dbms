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
package migrate

import (
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/wentaojin/dbms/openapi"
	"github.com/wentaojin/dbms/utils/stringutil"
)

type StructConfig struct {
	TaskName     string `toml:"task-name" json:"taskName"`
	TaskRuleName string `toml:"task-rule-name" json:"taskRuleName"`

	StructMigrateParam StructMigrateParam `toml:"struct-migrate-param" json:"structMigrateParam"`
	StructMigrateRule  StructMigrateRule  `toml:"struct-migrate-rule" json:"structMigrateRule"`
}

type StructMigrateParam struct {
	CaseFieldRule string `toml:"case-field-rule" json:"caseFieldRule"`
	MigrateThread int64  `toml:"migrate-thread" json:"migrateThread"`
	TaskQueueSize int64  `toml:"task-queue-size" json:"taskQueueSize"`
	DirectWrite   bool   `toml:"direct-write" json:"directWrite"`
	OutputDir     string `toml:"output-dir" json:"outputDir"`
}

type StructMigrateRule struct {
	TaskStructRules   []TaskStructRule   `toml:"task-struct-rules" json:"taskStructRules"`
	SchemaStructRules []SchemaStructRule `toml:"schema-struct-rules" json:"schemaStructRules"`
	TableStructRules  []TableStructRule  `toml:"table-struct-rules" json:"tableStructRules"`
	ColumnStructRules []ColumnStructRule `toml:"column-struct-rules" json:"columnStructRules"`
	TableAttrsRules   []TableAttrsRule   `toml:"table-attrs-rules" json:"tableAttrsRules"`
}

type TaskStructRule struct {
	ColumnTypeS   string `toml:"column-type-s" json:"columnTypeS"`
	ColumnTypeT   string `toml:"column-type-t" json:"columnTypeT"`
	DefaultValueS string `toml:"default-value-s" json:"defaultValueS"`
	DefaultValueT string `toml:"default-value-t" json:"defaultValueT"`
}

type SchemaStructRule struct {
	SourceSchema  string `toml:"source-schema" json:"sourceSchema"`
	ColumnTypeS   string `toml:"column-type-s" json:"columnTypeS"`
	ColumnTypeT   string `toml:"column-type-t" json:"columnTypeT"`
	DefaultValueS string `toml:"default-value-s" json:"defaultValueS"`
	DefaultValueT string `toml:"default-value-t" json:"defaultValueT"`
}

type TableStructRule struct {
	SourceSchema  string `toml:"source-schema" json:"sourceSchema"`
	SourceTable   string `toml:"source-table" json:"sourceTable"`
	ColumnTypeS   string `toml:"column-type-s" json:"columnTypeS"`
	ColumnTypeT   string `toml:"column-type-t" json:"columnTypeT"`
	DefaultValueS string `toml:"default-value-s" json:"defaultValueS"`
	DefaultValueT string `toml:"default-value-t" json:"defaultValueT"`
}

type ColumnStructRule struct {
	SourceSchema  string `toml:"source-schema" json:"sourceSchema"`
	SourceTable   string `toml:"source-table" json:"sourceTable"`
	SourceColumn  string `toml:"source-column" json:"sourceColumn"`
	ColumnTypeS   string `toml:"column-type-s" json:"columnTypeS"`
	ColumnTypeT   string `toml:"column-type-t" json:"columnTypeT"`
	DefaultValueS string `toml:"default-value-s" json:"defaultValueS"`
	DefaultValueT string `toml:"default-value-t" json:"defaultValueT"`
}

type TableAttrsRule struct {
	SourceSchema string   `toml:"source-schema" json:"sourceSchema"`
	SourceTables []string `toml:"source-tables" json:"sourceTables"`
	TableAttrsT  string   `toml:"table-attrs-t" json:"tableAttrsT"`
}

func (s *StructConfig) String() string {
	jsonStr, _ := stringutil.MarshalJSON(s)
	return jsonStr
}

func UpsertStructMigrate(serverAddr string, file string) error {
	var cfg = &StructConfig{}
	if _, err := toml.DecodeFile(file, cfg); err != nil {
		return fmt.Errorf("failed decode toml config file %s: %v", file, err)
	}

	resp, err := openapi.Request(openapi.RequestPUTMethod, stringutil.StringBuilder(stringutil.WrapScheme(serverAddr, false), openapi.DBMSAPIBasePath, openapi.APITaskPath, "/", openapi.APIStructMigratePath), []byte(cfg.String()))
	if err != nil {
		return err
	}

	var jsonData map[string]interface{}
	err = stringutil.UnmarshalJSON(resp, &jsonData)
	if err != nil {
		return fmt.Errorf("error decoding JSON: %v", err)
	}

	formattedJSON, err := stringutil.MarshalIndentJSON(stringutil.FormatJSONFields(jsonData))
	if err != nil {
		return fmt.Errorf("error encoding JSON: %v", err)
	}
	fmt.Println(formattedJSON)
	return nil
}

func DeleteStructMigrate(serverAddr string, name string) error {
	resp, err := openapi.Request(openapi.RequestDELETEMethod, stringutil.StringBuilder(stringutil.WrapScheme(serverAddr, false), openapi.DBMSAPIBasePath, openapi.APITaskPath, "/", openapi.APIStructMigratePath), []byte(name))
	if err != nil {
		return err
	}
	var jsonData map[string]interface{}
	err = stringutil.UnmarshalJSON(resp, &jsonData)
	if err != nil {
		return fmt.Errorf("error decoding JSON: %v", err)
	}

	formattedJSON, err := stringutil.MarshalIndentJSON(stringutil.FormatJSONFields(jsonData))
	if err != nil {
		return fmt.Errorf("error encoding JSON: %v", err)
	}
	fmt.Println(formattedJSON)
	return nil
}

func GetStructMigrate(serverAddr string, name string) error {
	resp, err := openapi.Request(openapi.RequestGETMethod, stringutil.StringBuilder(stringutil.WrapScheme(serverAddr, false), openapi.DBMSAPIBasePath, openapi.APITaskPath, "/", openapi.APIStructMigratePath), []byte(name))
	if err != nil {
		return err
	}

	var jsonData map[string]interface{}
	err = stringutil.UnmarshalJSON(resp, &jsonData)
	if err != nil {
		return fmt.Errorf("error decoding JSON: %v", err)
	}

	formattedJSON, err := stringutil.MarshalIndentJSON(stringutil.FormatJSONFields(jsonData))
	if err != nil {
		return fmt.Errorf("error encoding JSON: %v", err)
	}
	fmt.Println(formattedJSON)
	return nil
}
