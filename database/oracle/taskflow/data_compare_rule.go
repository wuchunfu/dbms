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

	"github.com/wentaojin/dbms/database"

	"github.com/wentaojin/dbms/model"
	"github.com/wentaojin/dbms/model/rule"
	"github.com/wentaojin/dbms/utils/constant"
	"github.com/wentaojin/dbms/utils/stringutil"
)

type DataCompareRule struct {
	Ctx            context.Context    `json:"-"`
	TaskName       string             `json:"taskName"`
	TaskMode       string             `json:"taskMode"`
	TaskFlow       string             `json:"taskFlow"`
	SchemaNameS    string             `json:"schemaNameS"`
	TableNameS     string             `json:"tableNameS"`
	GlobalSqlHintS string             `json:"globalSqlHintS"`
	TableTypeS     map[string]string  `json:"tableTypeS"`
	OnlyCompareRow bool               `json:"onlyCompareRow"`
	DatabaseS      database.IDatabase `json:"databaseS"`
	DBCharsetS     string             `json:"DBCharsetS"`
	DBCharsetT     string             `json:"DBCharsetT"`
	DBCollationS   bool               `json:"DBCollationS"`
	CaseFieldRuleS string             `json:"caseFieldRuleS"`
	CaseFieldRuleT string             `json:"caseFieldRuleT"`
}

func (r *DataCompareRule) GenSchemaNameRule() (string, string, error) {
	routeRule, err := model.GetIMigrateSchemaRouteRW().GetSchemaRouteRule(r.Ctx, &rule.SchemaRouteRule{
		TaskName: r.TaskName, SchemaNameS: r.SchemaNameS})
	if err != nil {
		return "", "", err
	}
	convertUtf8Raw, err := stringutil.CharsetConvert([]byte(r.SchemaNameS), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
	if err != nil {
		return "", "", fmt.Errorf("[GetSchemaNameRule] oracle schema [%s] charset convert failed, %v", r.SchemaNameS, err)
	}
	schemaNameS := stringutil.BytesToString(convertUtf8Raw)

	var schemaNameTNew string

	if !strings.EqualFold(routeRule.SchemaNameT, "") {
		schemaNameTNew = routeRule.SchemaNameT
	} else {
		schemaNameTNew = schemaNameS
	}

	if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueDataMigrateCaseFieldRuleLower) {
		schemaNameTNew = strings.ToLower(schemaNameTNew)
	}
	if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueDataMigrateCaseFieldRuleUpper) {
		schemaNameTNew = strings.ToUpper(schemaNameTNew)
	}

	return schemaNameS, schemaNameTNew, nil
}

func (r *DataCompareRule) GenSchemaTableNameRule() (string, string, error) {
	routeRule, err := model.GetIMigrateTableRouteRW().GetTableRouteRule(r.Ctx, &rule.TableRouteRule{
		TaskName: r.TaskName, SchemaNameS: r.SchemaNameS, TableNameS: r.TableNameS})
	if err != nil {
		return "", "", err
	}

	convertUtf8Raw, err := stringutil.CharsetConvert([]byte(r.TableNameS), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
	if err != nil {
		return "", "", fmt.Errorf("[GetTableNameRule] oracle schema [%s] table [%v] charset convert failed, %v", r.SchemaNameS, r.TableNameS, err)
	}
	tableNameS := stringutil.BytesToString(convertUtf8Raw)

	var tableNameTNew string
	if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueDataMigrateCaseFieldRuleLower) {
		tableNameTNew = strings.ToLower(tableNameS)
	}
	if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueDataMigrateCaseFieldRuleUpper) {
		tableNameTNew = strings.ToUpper(tableNameS)
	}
	if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueDataMigrateCaseFieldRuleOrigin) {
		tableNameTNew = tableNameS
	}

	if !strings.EqualFold(routeRule.TableNameT, "") {
		tableNameTNew = routeRule.TableNameT
	}
	return tableNameS, tableNameTNew, nil
}

func (r *DataCompareRule) GenSchemaTableColumnRule() (string, string, string, string, error) {
	columnRules := make(map[string]string)

	columnRoutes, err := model.GetIMigrateColumnRouteRW().FindColumnRouteRule(r.Ctx, &rule.ColumnRouteRule{
		TaskName:    r.TaskName,
		SchemaNameS: r.SchemaNameS,
		TableNameS:  r.TableNameS,
	})
	if err != nil {
		return "", "", "", "", err
	}

	sourceColumnNameS, err := r.DatabaseS.GetDatabaseTableColumnNameTableDimensions(r.SchemaNameS, r.TableNameS)
	if err != nil {
		return "", "", "", "", err
	}

	for _, c := range sourceColumnNameS {
		columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(c), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
		if err != nil {
			return "", "", "", "", fmt.Errorf("[GetTableColumnRule] oracle schema [%s] table [%s] column [%s] charset convert [UTFMB4] failed, error: %v", r.SchemaNameS, r.TableNameS, c, err)
		}

		columnName := stringutil.BytesToString(columnNameUtf8Raw)

		// column name caseFieldRule
		var (
			columnNameSNew string
			columnNameTNew string
		)
		if strings.EqualFold(r.CaseFieldRuleS, constant.ParamValueDataMigrateCaseFieldRuleLower) {
			columnNameSNew = strings.ToLower(columnName)
		}
		if strings.EqualFold(r.CaseFieldRuleS, constant.ParamValueDataMigrateCaseFieldRuleUpper) {
			columnNameSNew = strings.ToUpper(columnName)
		}
		if strings.EqualFold(r.CaseFieldRuleS, constant.ParamValueDataMigrateCaseFieldRuleOrigin) {
			columnNameSNew = columnName
		}

		if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueDataMigrateCaseFieldRuleLower) {
			columnNameTNew = strings.ToLower(columnName)
		}
		if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueDataMigrateCaseFieldRuleUpper) {
			columnNameTNew = strings.ToUpper(columnName)
		}
		if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueDataMigrateCaseFieldRuleOrigin) {
			columnNameTNew = columnName
		}

		columnRules[columnNameSNew] = columnNameTNew
	}

	for _, c := range columnRoutes {
		if _, exist := columnRules[c.ColumnNameS]; exist {
			columnRules[c.ColumnNameS] = c.ColumnNameT
		}
	}

	var (
		columnNameSilS, columnNameSilSO, columnNameSliT, columnNameSliTO []string
	)

	sourceColumnInfos, err := r.DatabaseS.GetDatabaseTableColumns(r.SchemaNameS, r.TableNameS, r.DBCollationS)
	if err != nil {
		return "", "", "", "", err
	}

	for _, rowCol := range sourceColumnInfos {
		columnNameS := rowCol["COLUMN_NAME"]

		columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(columnNameS), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
		if err != nil {
			return "", "", "", "", fmt.Errorf("[GenTableQueryColumnRule] oracle schema [%s] table [%s] column [%s] charset convert [UTFMB4] failed, error: %v", r.SchemaNameS, r.TableNameS, columnNameS, err)
		}
		columnNameS = stringutil.BytesToString(columnNameUtf8Raw)

		columnNameSO, err := optimizerDataMigrateColumnS(columnNameS, rowCol["DATA_TYPE"], rowCol["DATA_SCALE"])
		if err != nil {
			return "", "", "", "", err
		}
		columnNameSilSO = append(columnNameSilSO, columnNameSO)

		var (
			columnNameSNew string
			columnNameT    string
		)
		if strings.EqualFold(r.CaseFieldRuleS, constant.ParamValueDataMigrateCaseFieldRuleLower) {
			columnNameSNew = strings.ToLower(columnNameS)
		}
		if strings.EqualFold(r.CaseFieldRuleS, constant.ParamValueDataMigrateCaseFieldRuleUpper) {
			columnNameSNew = strings.ToUpper(columnNameS)
		}
		if strings.EqualFold(r.CaseFieldRuleS, constant.ParamValueDataMigrateCaseFieldRuleOrigin) {
			columnNameSNew = columnNameS
		}

		if val, ok := columnRules[columnNameSNew]; ok {
			switch {
			case strings.EqualFold(r.TaskFlow, constant.TaskFlowOracleToTiDB) || strings.EqualFold(r.TaskFlow, constant.TaskFlowOracleToMySQL):
				columnNameT = fmt.Sprintf("%s%s%s", constant.StringSeparatorBacktick, val, constant.StringSeparatorBacktick)
			default:
				return "", "", "", "", fmt.Errorf("oracle current task [%s] schema [%s] taskflow [%s] column rule isn't support, please contact author", r.TaskName, r.SchemaNameS, r.TaskFlow)
			}
		} else {
			return "", "", "", "", fmt.Errorf("[GetTableColumnRule] oracle schema [%s] table [%s] column [%s] isn't exist, please contact author or double check again", r.SchemaNameS, r.TableNameS, columnNameS)
		}

		columnNameSliTO = append(columnNameSliTO, columnNameT)
		columnNameS, columnNameT, err = optimizerDataCompareColumnST(columnNameS, rowCol["DATA_TYPE"], rowCol["DATA_SCALE"], columnNameT)
		if err != nil {
			return "", "", "", "", err
		}
		columnNameSilS = append(columnNameSilS, columnNameS)
		columnNameSliT = append(columnNameSliT, columnNameT)
	}

	if r.OnlyCompareRow {
		return stringutil.StringJoin(columnNameSilSO, constant.StringSeparatorComma), "COUNT(1) AS ROWSCOUNT", stringutil.StringJoin(columnNameSliTO, constant.StringSeparatorComma), "COUNT(1) AS ROWSCOUNT", nil
	}
	return stringutil.StringJoin(columnNameSilSO, constant.StringSeparatorComma),
		fmt.Sprintf(`UPPER(DBMS_CRYPTO.HASH(UTL_I18N.STRING_TO_RAW(%s,'%s'), 2 /*DBMS_CRYPTO.HASH_MD5*/)) AS ROWSCHECKSUM`,
			stringutil.StringJoin(columnNameSilS, constant.StringSplicingSymbol), constant.ORACLECharsetAL32UTF8),
		stringutil.StringJoin(columnNameSliTO, constant.StringSeparatorComma),
		fmt.Sprintf(`UPPER(MD5(CONVERT(CONCAT(%s) USING '%s'))) AS ROWSCHECKSUM`,
			stringutil.StringJoin(columnNameSliT, constant.StringSeparatorComma), constant.MYSQLCharsetUTF8MB4),
		nil
}

func (r *DataCompareRule) GenSchemaTableTypeRule() string {
	return r.TableTypeS[r.TableNameS]
}

func (r *DataCompareRule) GenSchemaTableCompareMethodRule() string {
	if r.OnlyCompareRow {
		return constant.DataCompareMethodCheckRows
	}
	return constant.DataCompareMethodCheckSum
}
