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
package processor

import (
	"context"
	"fmt"
	"strings"

	"github.com/wentaojin/dbms/database"
	"github.com/wentaojin/dbms/utils/constant"
	"github.com/wentaojin/dbms/utils/stringutil"
)

type SqlMigrateRule struct {
	Ctx             context.Context    `json:"-"`
	TaskName        string             `json:"taskName"`
	TaskMode        string             `json:"taskMode"`
	TaskFlow        string             `json:"taskFlow"`
	SchemaNameT     string             `json:"schemaNameT"`
	TableNameT      string             `json:"tableNameT"`
	SqlHintT        string             `json:"sqlHintT"`
	GlobalSqlHintT  string             `json:"globalSqlHintT"`
	DatabaseS       database.IDatabase `json:"databaseS"`
	DBCharsetS      string             `json:"DBCharsetS"`
	SqlQueryS       string             `json:"sqlQueryS"`
	ColumnRouteRule map[string]string  `json:"columnRouteRule"`
	CaseFieldRuleS  string             `json:"caseFieldRuleS"`
	CaseFieldRuleT  string             `json:"caseFieldRuleT"`
}

func (r *SqlMigrateRule) GenSqlMigrateSchemaNameRule() (string, error) {
	var schemaNameT string
	switch {
	case strings.EqualFold(r.TaskFlow, constant.TaskFlowOracleToTiDB) || strings.EqualFold(r.TaskFlow, constant.TaskFlowOracleToMySQL):
		convertUtf8Raw, err := stringutil.CharsetConvert([]byte(r.SchemaNameT), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
		if err != nil {
			return "", fmt.Errorf("[GetSchemaNameRule] oracle schema [%s] charset convert failed, %v", r.SchemaNameT, err)
		}
		schemaNameT = stringutil.BytesToString(convertUtf8Raw)
	default:
		return "", fmt.Errorf("the task_name [%s] task_mode [%v] taskflow [%s] schema_name_t [%v] sql migrate sql schema rule isn't support, please contact author", r.TaskName, r.TaskMode, r.TaskFlow, r.SchemaNameT)
	}

	if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueSqlMigrateCaseFieldRuleLower) {
		schemaNameT = strings.ToLower(schemaNameT)
	}
	if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueSqlMigrateCaseFieldRuleUpper) {
		schemaNameT = strings.ToUpper(schemaNameT)
	}

	return schemaNameT, nil
}

func (r *SqlMigrateRule) GenSqlMigrateTableNameRule() (string, error) {
	var tableNameT string
	switch {
	case strings.EqualFold(r.TaskFlow, constant.TaskFlowOracleToTiDB) || strings.EqualFold(r.TaskFlow, constant.TaskFlowOracleToMySQL):
		convertUtf8Raw, err := stringutil.CharsetConvert([]byte(r.TableNameT), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
		if err != nil {
			return "", fmt.Errorf("[GetTableNameRule] oracle schema [%s] table [%v] charset convert failed, %v", r.SchemaNameT, r.TableNameT, err)
		}
		tableNameT = stringutil.BytesToString(convertUtf8Raw)
	default:
		return "", fmt.Errorf("the task_name [%s] task_mode [%v] taskflow [%s] schema_name_t [%v] sql migrate sql table rule isn't support, please contact author", r.TaskName, r.TaskMode, r.TaskFlow, r.SchemaNameT)
	}

	if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueSqlMigrateCaseFieldRuleLower) {
		tableNameT = strings.ToLower(tableNameT)
	}
	if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueSqlMigrateCaseFieldRuleUpper) {
		tableNameT = strings.ToUpper(tableNameT)
	}
	return tableNameT, nil
}

func (r *SqlMigrateRule) GenSqlMigrateTableColumnRule() (string, string, string, error) {
	// column name rule
	sqlText, err := stringutil.Decrypt(r.SqlQueryS, []byte(constant.DefaultDataEncryptDecryptKey))
	if err != nil {
		return "", "", "", err
	}
	columnNames, columnTypeMap, columnScaleMap, err := r.DatabaseS.GetDatabaseTableColumnNameSqlDimensions(sqlText)
	if err != nil {
		return "", "", "", err
	}

	var (
		columnNameSliO []string
		columnNameSliS []string
		columnNameSliT []string
	)
	for _, c := range columnNames {
		var columnName string
		switch {
		case strings.EqualFold(r.TaskFlow, constant.TaskFlowOracleToTiDB) || strings.EqualFold(r.TaskFlow, constant.TaskFlowOracleToMySQL):
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(c), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return "", "", "", fmt.Errorf("[GenTableColumnRule] oracle sql migrate task sql column [%v] charset convert [UTFMB4] failed, error: %v", c, err)
			}
			columnName = stringutil.BytesToString(columnNameUtf8Raw)
		default:
			return "", "", "", fmt.Errorf("the task_name [%s] task_mode [%v] taskflow [%s] schema_name_t [%v] sql migrate sql column rule isn't support, please contact author", r.TaskName, r.TaskMode, r.TaskFlow, r.SchemaNameT)
		}

		columnNameSliO = append(columnNameSliO, columnName)

		columnNameS, err := OptimizerOracleDataMigrateColumnS(stringutil.StringBuilder(constant.StringSeparatorDoubleQuotes, columnName, constant.StringSeparatorDoubleQuotes), columnTypeMap[c], columnScaleMap[c])
		if err != nil {
			return "", "", "", err
		}
		columnNameSliS = append(columnNameSliS, columnNameS)

		// column name caseFieldRule
		var (
			columnNameSNew string
			columnNameTNew string
		)
		if strings.EqualFold(r.CaseFieldRuleS, constant.ParamValueSqlMigrateCaseFieldRuleLower) {
			columnNameSNew = strings.ToLower(columnName)
		}
		if strings.EqualFold(r.CaseFieldRuleS, constant.ParamValueSqlMigrateCaseFieldRuleUpper) {
			columnNameSNew = strings.ToUpper(columnName)
		}
		if strings.EqualFold(r.CaseFieldRuleS, constant.ParamValueSqlMigrateCaseFieldRuleOrigin) {
			columnNameSNew = columnName
		}

		if val, ok := r.ColumnRouteRule[columnNameSNew]; ok {
			columnNameTNew = val
		} else {
			columnNameTNew = columnNameSNew
		}
		switch {
		case strings.EqualFold(r.TaskFlow, constant.TaskFlowOracleToTiDB) || strings.EqualFold(r.TaskFlow, constant.TaskFlowOracleToMySQL):
			if strings.EqualFold(r.TaskMode, constant.TaskModeSqlMigrate) {
				columnNameSliT = append(columnNameSliT, fmt.Sprintf("`%s`", columnNameTNew))
			} else {
				return "", "", "", fmt.Errorf("the task_name [%s] task_mode [%v] taskflow [%s] schema_name_t [%v] sql migrate sql column rule isn't support, please contact author", r.TaskName, r.TaskMode, r.TaskFlow, r.SchemaNameT)
			}
		default:
			return "", "", "", fmt.Errorf("the task_name [%s] task_mode [%v] taskflow [%s] schema_name_t [%v] sql migrate sql column rule isn't support, please contact author", r.TaskName, r.TaskMode, r.TaskFlow, r.SchemaNameT)
		}
	}
	return stringutil.StringJoin(columnNameSliO, constant.StringSeparatorComma), stringutil.StringJoin(columnNameSliS, constant.StringSeparatorComma), stringutil.StringJoin(columnNameSliT, constant.StringSeparatorComma), nil
}

func (r *SqlMigrateRule) GenSqlMigrateTableCustomRule() (string, string) {
	if strings.EqualFold(r.SqlHintT, "") {
		return r.GlobalSqlHintT, r.SqlQueryS
	}
	return r.SqlHintT, r.SqlQueryS
}
