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
	"bytes"
	"fmt"
	"github.com/fatih/color"

	"github.com/BurntSushi/toml"
	"github.com/wentaojin/dbms/openapi"
	"github.com/wentaojin/dbms/utils/stringutil"
)

type SqlConfig struct {
	TaskName        string `toml:"task-name" json:"taskName"`
	DatasourceNameS string `toml:"datasource-name-s" json:"datasourceNameS"`
	DatasourceNameT string `toml:"datasource-name-t" json:"datasourceNameT"`
	Comment         string `toml:"comment" json:"comment"`

	CaseFieldRule   CaseFieldRule    `toml:"case-field-rule" json:"caseFieldRule"`
	SqlMigrateRules []SqlMigrateRule `toml:"sql-migrate-rules" json:"sqlMigrateRules"`
	SqlMigrateParam SqlMigrateParam  `toml:"sql-migrate-param" json:"sqlMigrateParam"`
}

type SqlMigrateParam struct {
	BatchSize            uint64 `toml:"batch-size" json:"batchSize"`
	SqlThreadS           uint64 `toml:"sql-thread-s" json:"sqlThreadS"`
	SqlThreadT           uint64 `toml:"sql-thread-t" json:"sqlThreadT"`
	SqlHintT             string `toml:"sql-hint-t" json:"sqlHintT"`
	CallTimeout          uint64 `toml:"call-timeout" json:"callTimeout"`
	EnableConsistentRead bool   `toml:"enable-consistent-read" json:"enableConsistentRead"`
	EnableSafeMode       bool   `toml:"enable-safe-mode" json:"enableSafeMode"`
}

func (s *SqlConfig) String() string {
	jsonStr, _ := stringutil.MarshalJSON(s)
	return jsonStr
}

func UpsertSqlMigrate(serverAddr string, file string) error {
	var cfg = &SqlConfig{}

	cyan := color.New(color.FgCyan, color.Bold)
	fmt.Printf("Component:    %s\n", cyan.Sprint("dbms-ctl"))
	fmt.Printf("Command:      %s\n", cyan.Sprint("sql"))
	fmt.Printf("File:         %s\n", cyan.Sprint(file))
	fmt.Printf("Action:       %s\n", cyan.Sprint("upsert"))

	if _, err := toml.DecodeFile(file, cfg); err != nil {
		fmt.Printf("Status:       %s\n", cyan.Sprint("failed"))
		fmt.Printf("Response:     %s\n", color.RedString("failed decode toml config file %s: %v", file, err))
		return nil
	}
	resp, err := openapi.Request(openapi.RequestPUTMethod, stringutil.StringBuilder(stringutil.WrapScheme(serverAddr, false), openapi.DBMSAPIBasePath, openapi.APITaskPath, "/", openapi.APISqlMigratePath), []byte(cfg.String()))
	if err != nil {
		fmt.Printf("Status:       %s\n", cyan.Sprint("failed"))
		fmt.Printf("Response:     %s\n", color.RedString("the request failed: %v", err))
		return nil
	}

	var jsonData map[string]interface{}
	err = stringutil.UnmarshalJSON(resp, &jsonData)
	if err != nil {
		fmt.Printf("Status:       %s\n", cyan.Sprint("failed"))
		fmt.Printf("Response:     %s\n", color.RedString("error encoding JSON: %v", err))
		return nil
	}

	formattedJSON, err := stringutil.MarshalIndentJSON(stringutil.FormatJSONFields(jsonData))
	if err != nil {
		fmt.Printf("Status:       %s\n", cyan.Sprint("failed"))
		fmt.Printf("Response:     %s\n", color.RedString("error encoding JSON: %v", err))
		return nil
	}
	fmt.Printf("Status:       %s\n", cyan.Sprint("success"))
	fmt.Printf("Response:     %s\n", formattedJSON)
	return nil
}

func DeleteSqlMigrate(serverAddr string, name string) error {
	bodyReq := make(map[string]interface{})
	bodyReq["param"] = []string{name}

	cyan := color.New(color.FgCyan, color.Bold)
	fmt.Printf("Component:    %s\n", cyan.Sprint("dbms-ctl"))
	fmt.Printf("Command:      %s\n", cyan.Sprint("scan"))
	fmt.Printf("Task:         %s\n", cyan.Sprint(name))
	fmt.Printf("Action:       %s\n", cyan.Sprint("delete"))

	jsonStr, err := stringutil.MarshalJSON(bodyReq)
	if err != nil {
		fmt.Printf("Status:       %s\n", cyan.Sprint("failed"))
		fmt.Printf("Response:     %s\n", color.RedString("error marshal JSON: %v", err))
		return nil
	}
	resp, err := openapi.Request(openapi.RequestDELETEMethod, stringutil.StringBuilder(stringutil.WrapScheme(serverAddr, false), openapi.DBMSAPIBasePath, openapi.APITaskPath, "/", openapi.APISqlMigratePath), []byte(jsonStr))
	if err != nil {
		fmt.Printf("Status:       %s\n", cyan.Sprint("failed"))
		fmt.Printf("Response:     %s\n", color.RedString("the request failed: %v", err))
		return nil
	}
	if bytes.Equal(resp, []byte("")) {
		fmt.Printf("Status:       %s\n", cyan.Sprint("success"))
		fmt.Printf("Response:     %s\n", color.GreenString("the sql migrate task has been deleted or not existed, return the response null"))
		return nil
	}
	var jsonData map[string]interface{}
	err = stringutil.UnmarshalJSON(resp, &jsonData)
	if err != nil {
		fmt.Printf("Status:       %s\n", cyan.Sprint("failed"))
		fmt.Printf("Response:     %s\n", color.RedString("error decoding JSON: %v", err))
		return nil
	}

	formattedJSON, err := stringutil.MarshalIndentJSON(stringutil.FormatJSONFields(jsonData))
	if err != nil {
		fmt.Printf("Status:       %s\n", cyan.Sprint("failed"))
		fmt.Printf("Response:     %s\n", color.RedString("error decoding JSON: %v", err))
		return nil
	}
	fmt.Printf("Status:       %s\n", cyan.Sprint("success"))
	fmt.Printf("Response:     %s\n", formattedJSON)
	return nil
}

func GetSqlMigrate(serverAddr string, name string) error {
	bodyReq := make(map[string]interface{})
	bodyReq["param"] = name
	bodyReq["page"] = 1
	bodyReq["pageSize"] = 100

	cyan := color.New(color.FgCyan, color.Bold)
	fmt.Printf("Component:    %s\n", cyan.Sprint("dbms-ctl"))
	fmt.Printf("Command:      %s\n", cyan.Sprint("sql"))
	fmt.Printf("Task:         %s\n", cyan.Sprint(name))
	fmt.Printf("Action:       %s\n", cyan.Sprint("get"))

	jsonStr, err := stringutil.MarshalJSON(bodyReq)
	if err != nil {
		fmt.Printf("Status:       %s\n", cyan.Sprint("failed"))
		fmt.Printf("Response:     %s\n", color.RedString("error marshal JSON: %v", err))
		return nil
	}
	resp, err := openapi.Request(openapi.RequestPOSTMethod, stringutil.StringBuilder(stringutil.WrapScheme(serverAddr, false), openapi.DBMSAPIBasePath, openapi.APITaskPath, "/", openapi.APISqlMigratePath), []byte(jsonStr))
	if err != nil {
		fmt.Printf("Status:       %s\n", cyan.Sprint("failed"))
		fmt.Printf("Response:     %s\n", color.RedString("the request failed: %v", err))
		return nil
	}

	var jsonData map[string]interface{}
	err = stringutil.UnmarshalJSON(resp, &jsonData)
	if err != nil {
		fmt.Printf("Status:       %s\n", cyan.Sprint("failed"))
		fmt.Printf("Response:     %s\n", color.RedString("error encoding JSON: %v", err))
		return nil
	}
	formattedJSON, err := stringutil.MarshalIndentJSON(stringutil.FormatJSONFields(jsonData))
	if err != nil {
		fmt.Printf("Status:       %s\n", cyan.Sprint("failed"))
		fmt.Printf("Response:     %s\n", color.RedString("error encoding JSON: %v", err))
		return nil
	}
	fmt.Printf("Status:       %s\n", cyan.Sprint("success"))
	fmt.Printf("Response:     %s\n", formattedJSON)
	return nil
}
