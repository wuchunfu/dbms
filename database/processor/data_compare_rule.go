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
	"strconv"
	"strings"

	"github.com/wentaojin/dbms/database"
	"github.com/wentaojin/dbms/logger"
	"go.uber.org/zap"

	"github.com/wentaojin/dbms/model"
	"github.com/wentaojin/dbms/model/rule"
	"github.com/wentaojin/dbms/utils/constant"
	"github.com/wentaojin/dbms/utils/stringutil"
)

type DataCompareRule struct {
	Ctx                         context.Context    `json:"-"`
	TaskName                    string             `json:"taskName"`
	TaskMode                    string             `json:"taskMode"`
	TaskFlow                    string             `json:"taskFlow"`
	SchemaNameS                 string             `json:"schemaNameS"`
	TableNameS                  string             `json:"tableNameS"`
	SchemaNameT                 string             `json:"schemNameT"`
	TableNameT                  string             `json:"tableNameT"`
	GlobalSqlHintS              string             `json:"globalSqlHintS"`
	GlobalSqlHintT              string             `json:"globalSqlHintT"`
	TableTypeS                  map[string]string  `json:"tableTypeS"`
	ColumnNameSliS              []string           `json:"columnNameSliS"`
	IgnoreSelectFields          []string           `json:"ignoreSelectFields"`
	GlobalIgnoreConditionFields []string           `json:"globalIgnoreConditionFields"`
	OnlyDatabaseCompareRow      bool               `json:"onlyDatabaseCompareRow"`
	OnlyProgramCompareCRC32     bool               `json:"onlyProgramCompareCRC32"`
	DisableDatabaseCompareMd5   bool               `json:"disableDatabaseCompareMd5"`
	DatabaseS                   database.IDatabase `json:"-"`
	DatabaseT                   database.IDatabase `json:"-"`
	DBCharsetS                  string             `json:"DBCharsetS"`
	DBCharsetT                  string             `json:"DBCharsetT"`
	CaseFieldRuleS              string             `json:"caseFieldRuleS"`
	CaseFieldRuleT              string             `json:"caseFieldRuleT"`
}

func (r *DataCompareRule) GenSchemaNameRule() (string, string, error) {
	routeRule, err := model.GetIMigrateSchemaRouteRW().GetSchemaRouteRule(r.Ctx, &rule.SchemaRouteRule{
		TaskName: r.TaskName, SchemaNameS: r.SchemaNameS})
	if err != nil {
		return "", "", err
	}

	var schemaNameS string
	switch r.TaskFlow {
	case constant.TaskFlowOracleToMySQL, constant.TaskFlowOracleToTiDB:
		convertUtf8Raw, err := stringutil.CharsetConvert([]byte(r.SchemaNameS), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
		if err != nil {
			return "", "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GetSchemaNameRule] schema [%s] charset convert failed, %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, err)
		}
		schemaNameS = stringutil.BytesToString(convertUtf8Raw)
	case constant.TaskFlowTiDBToOracle, constant.TaskFlowTiDBToPostgres:
		convertUtf8Raw, err := stringutil.CharsetConvert([]byte(r.SchemaNameS), constant.MigrateMySQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
		if err != nil {
			return "", "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GetSchemaNameRule] schema [%s] charset convert failed, %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, err)
		}
		schemaNameS = stringutil.BytesToString(convertUtf8Raw)
	case constant.TaskFlowPostgresToMySQL, constant.TaskFlowPostgresToTiDB:
		convertUtf8Raw, err := stringutil.CharsetConvert([]byte(r.SchemaNameS), constant.MigratePostgreSQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
		if err != nil {
			return "", "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GetSchemaNameRule] schema [%s] charset convert failed, %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, err)
		}
		schemaNameS = stringutil.BytesToString(convertUtf8Raw)
	default:
		return "", "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] isn't support, please contact author or reselect", r.TaskName, r.TaskFlow, r.TaskMode)
	}

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

	// grant
	r.SchemaNameT = schemaNameTNew
	return schemaNameS, schemaNameTNew, nil
}

func (r *DataCompareRule) GenSchemaTableNameRule() (string, string, error) {
	routeRule, err := model.GetIMigrateTableRouteRW().GetTableRouteRule(r.Ctx, &rule.TableRouteRule{
		TaskName: r.TaskName, SchemaNameS: r.SchemaNameS, TableNameS: r.TableNameS})
	if err != nil {
		return "", "", err
	}

	var tableNameS string
	switch r.TaskFlow {
	case constant.TaskFlowOracleToMySQL, constant.TaskFlowOracleToTiDB:
		convertUtf8Raw, err := stringutil.CharsetConvert([]byte(r.TableNameS), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
		if err != nil {
			return "", "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GetTableNameRule] schema [%s] table [%v] charset convert failed, %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, r.TableNameS, err)
		}
		tableNameS = stringutil.BytesToString(convertUtf8Raw)
	case constant.TaskFlowTiDBToOracle, constant.TaskFlowTiDBToPostgres:
		convertUtf8Raw, err := stringutil.CharsetConvert([]byte(r.TableNameS), constant.MigrateMySQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
		if err != nil {
			return "", "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GetTableNameRule] schema [%s] table [%v] charset convert failed, %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, r.TableNameS, err)
		}
		tableNameS = stringutil.BytesToString(convertUtf8Raw)
	case constant.TaskFlowPostgresToMySQL, constant.TaskFlowPostgresToTiDB:
		convertUtf8Raw, err := stringutil.CharsetConvert([]byte(r.TableNameS), constant.MigratePostgreSQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
		if err != nil {
			return "", "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GetTableNameRule] schema [%s] table [%v] charset convert failed, %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, r.TableNameS, err)
		}
		tableNameS = stringutil.BytesToString(convertUtf8Raw)
	default:
		return "", "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] isn't support, please contact author or reselect", r.TaskName, r.TaskFlow, r.TaskMode)
	}

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

	// grant
	r.TableNameT = tableNameTNew
	return tableNameS, tableNameTNew, nil
}

func (r *DataCompareRule) GetSchemaTableColumnNameRule() (map[string]string, error) {
	columnRoutes, err := model.GetIMigrateColumnRouteRW().FindColumnRouteRule(r.Ctx, &rule.ColumnRouteRule{
		TaskName:    r.TaskName,
		SchemaNameS: r.SchemaNameS,
		TableNameS:  r.TableNameS,
	})
	if err != nil {
		return nil, err
	}
	columnRouteRules := make(map[string]string)
	for _, routeRule := range columnRoutes {
		columnRouteRules[routeRule.ColumnNameS] = routeRule.ColumnNameT
	}

	for _, c := range stringutil.StringItemsFilterDifference(r.ColumnNameSliS, r.IgnoreSelectFields) {
		var columnNameS string
		switch r.TaskFlow {
		case constant.TaskFlowOracleToMySQL, constant.TaskFlowOracleToTiDB:
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(c), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return nil, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GetTableColumnRule] schema [%s] table [%s] column [%s] charset convert [UTFMB4] failed, error: %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, r.TableNameS, c, err)
			}
			columnNameS = stringutil.BytesToString(columnNameUtf8Raw)
		case constant.TaskFlowTiDBToOracle, constant.TaskFlowTiDBToPostgres:
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(c), constant.MigrateMySQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return nil, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GetTableColumnRule] schema [%s] table [%s] column [%s] charset convert [UTFMB4] failed, error: %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, r.TableNameS, c, err)
			}
			columnNameS = stringutil.BytesToString(columnNameUtf8Raw)
		case constant.TaskFlowPostgresToMySQL, constant.TaskFlowPostgresToTiDB:
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(c), constant.MigratePostgreSQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return nil, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GetTableColumnRule] schema [%s] table [%s] column [%s] charset convert [UTFMB4] failed, error: %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, r.TableNameS, c, err)
			}
			columnNameS = stringutil.BytesToString(columnNameUtf8Raw)
		default:
			return nil, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] isn't support, please contact author or reselect", r.TaskName, r.TaskFlow, r.TaskMode)
		}

		var columnName string
		if strings.EqualFold(r.CaseFieldRuleS, constant.ParamValueDataCompareCaseFieldRuleLower) {
			columnName = strings.ToLower(columnNameS)
		}
		if strings.EqualFold(r.CaseFieldRuleS, constant.ParamValueDataCompareCaseFieldRuleUpper) {
			columnName = strings.ToUpper(columnNameS)
		}
		if strings.EqualFold(r.CaseFieldRuleS, constant.ParamValueDataCompareCaseFieldRuleOrigin) {
			columnName = columnNameS
		}

		var columnNameT string
		if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueDataCompareCaseFieldRuleLower) {
			columnNameT = strings.ToLower(columnNameS)
		}
		if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueDataCompareCaseFieldRuleUpper) {
			columnNameT = strings.ToUpper(columnNameS)
		}
		if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueDataCompareCaseFieldRuleOrigin) {
			columnNameT = columnNameS
		}

		if val, ok := columnRouteRules[columnName]; ok {
			columnRouteRules[columnName] = val
		} else {
			columnRouteRules[columnName] = columnNameT
		}
	}
	return columnRouteRules, nil
}

func (r *DataCompareRule) GenSchemaTableColumnSelectRule() (string, string, string, string, string, string, error) {
	columnRules := make(map[string]string)

	columnRoutes, err := model.GetIMigrateColumnRouteRW().FindColumnRouteRule(r.Ctx, &rule.ColumnRouteRule{
		TaskName:    r.TaskName,
		SchemaNameS: r.SchemaNameS,
		TableNameS:  r.TableNameS,
	})
	if err != nil {
		return "", "", "", "", "", "", err
	}

	sourceColumnNameS, err := r.DatabaseS.GetDatabaseTableColumnNameTableDimensions(r.SchemaNameS, r.TableNameS)
	if err != nil {
		return "", "", "", "", "", "", err
	}

	// assignment
	r.ColumnNameSliS = sourceColumnNameS

	for _, c := range stringutil.StringItemsFilterDifference(sourceColumnNameS, r.IgnoreSelectFields) {
		var columnName string
		switch r.TaskFlow {
		case constant.TaskFlowOracleToMySQL, constant.TaskFlowOracleToTiDB:
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(c), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return "", "", "", "", "", "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GenSchemaTableColumnSelectRule] schema [%s] table [%s] column [%s] charset convert [UTFMB4] failed, error: %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, r.TableNameS, c, err)
			}
			columnName = stringutil.BytesToString(columnNameUtf8Raw)
		case constant.TaskFlowTiDBToOracle, constant.TaskFlowTiDBToPostgres:
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(c), constant.MigrateMySQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return "", "", "", "", "", "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GenSchemaTableColumnSelectRule] schema [%s] table [%s] column [%s] charset convert [UTFMB4] failed, error: %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, r.TableNameS, c, err)
			}
			columnName = stringutil.BytesToString(columnNameUtf8Raw)
		case constant.TaskFlowPostgresToMySQL, constant.TaskFlowPostgresToTiDB:
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(c), constant.MigratePostgreSQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return "", "", "", "", "", "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GenSchemaTableColumnSelectRule] schema [%s] table [%s] column [%s] charset convert [UTFMB4] failed, error: %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, r.TableNameS, c, err)
			}
			columnName = stringutil.BytesToString(columnNameUtf8Raw)
		default:
			return "", "", "", "", "", "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] isn't support, please contact author or reselect", r.TaskName, r.TaskFlow, r.TaskMode)
		}

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
		columnNameSilSC, columnNameSilSO, columnNameSliTC, columnNameSliTO, ignoreTypes []string
	)

	// downstream oracle
	// find the oracle long/ long raw / bfile datatype, rollback to crc32
	downstreamDBDatatypeRollback := make(map[string]string)
	downstreamDBDateDatatype := make(map[string]string)
	downstreamDBCharDatatype := make(map[string]string)
	downstreamDBTimestampDatatype := make(map[string]string)

	// downstream postgres
	downstreamDBBooleanDatatype := make(map[string]string)

	switch r.TaskFlow {
	case constant.TaskFlowTiDBToOracle, constant.TaskFlowMySQLToOracle:
		targetColumnInfos, err := r.DatabaseT.GetDatabaseTableColumnInfo(r.SchemaNameT, r.TableNameT)
		if err != nil {
			return "", "", "", "", "", "", err
		}
		for _, c := range targetColumnInfos {
			if strings.EqualFold(c["DATA_TYPE"], constant.BuildInOracleDatatypeLong) || strings.EqualFold(c["DATA_TYPE"], constant.BuildInOracleDatatypeLongRAW) || strings.EqualFold(c["DATA_TYPE"], constant.BuildInOracleDatatypeBfile) {
				downstreamDBDatatypeRollback[c["COLUMN_NAME"]] = c["DATA_TYPE"]
			}
			if strings.EqualFold(c["DATA_TYPE"], constant.BuildInOracleDatatypeChar) || strings.EqualFold(c["DATA_TYPE"], constant.BuildInOracleDatatypeCharacter) {
				downstreamDBCharDatatype[c["COLUMN_NAME"]] = c["DATA_TYPE"]
			}
			if stringutil.IsContainedString(constant.DataCompareOracleDatabaseSupportDateSubtypes, c["DATA_TYPE"]) {
				downstreamDBDateDatatype[c["COLUMN_NAME"]] = c["DATA_TYPE"]
			}
			if stringutil.IsContainedString(constant.DataCompareOracleDatabaseSupportTimestampSubtypes, c["DATA_TYPE"]) {
				downstreamDBTimestampDatatype[c["COLUMN_NAME"]] = c["DATA_TYPE"]
			}
		}
	case constant.TaskFlowTiDBToPostgres, constant.TaskFlowMySQLToPostgres:
		targetColumnInfos, err := r.DatabaseT.GetDatabaseTableColumnInfo(r.SchemaNameT, r.TableNameT)
		if err != nil {
			return "", "", "", "", "", "", err
		}
		for _, c := range targetColumnInfos {
			if strings.EqualFold(c["DATA_TYPE"], constant.BuildInPostgresDatatypeBoolean) {
				downstreamDBBooleanDatatype[c["COLUMN_NAME"]] = c["DATA_TYPE"]
			}
			if strings.EqualFold(c["DATA_TYPE"], constant.BuildInPostgresDatatypeCharacter) {
				downstreamDBBooleanDatatype[c["COLUMN_NAME"]] = c["DATA_TYPE"]
			}
		}
	case constant.TaskFlowOracleToTiDB, constant.TaskFlowOracleToMySQL, constant.TaskFlowPostgresToTiDB, constant.TaskFlowPostgresToMySQL:
		targetColumnInfos, err := r.DatabaseT.GetDatabaseTableColumnInfo(r.SchemaNameT, r.TableNameT)
		if err != nil {
			return "", "", "", "", "", "", err
		}
		for _, c := range targetColumnInfos {
			if strings.EqualFold(c["DATA_TYPE"], constant.BuildInMySQLDatatypeChar) {
				downstreamDBCharDatatype[c["COLUMN_NAME"]] = c["DATA_TYPE"]
			}
		}
	}

	sourceColumnInfos, err := r.DatabaseS.GetDatabaseTableColumnInfo(r.SchemaNameS, r.TableNameS)
	if err != nil {
		return "", "", "", "", "", "", err
	}

	for _, rowCol := range sourceColumnInfos {
		columnNameS := rowCol["COLUMN_NAME"]
		datatypeS := rowCol["DATA_TYPE"]
		dataPrecisionS := rowCol["DATA_PRECISION"]
		dataScaleS := rowCol["DATA_SCALE"]
		dataLen := rowCol["DATA_LENGTH"]

		dataLenC, err := strconv.ParseInt(dataLen, 10, 64)
		if err != nil {
			return "", "", "", "", "", "", err
		}

		switch r.TaskFlow {
		case constant.TaskFlowOracleToMySQL, constant.TaskFlowOracleToTiDB:
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(columnNameS), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return "", "", "", "", "", "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GenTableQueryColumnRule] schema [%s] table [%s] column [%s] charset convert [UTFMB4] failed, error: %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, r.TableNameS, columnNameS, err)
			}
			columnNameS = stringutil.BytesToString(columnNameUtf8Raw)
		case constant.TaskFlowTiDBToOracle, constant.TaskFlowTiDBToPostgres:
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(columnNameS), constant.MigrateMySQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return "", "", "", "", "", "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GenTableQueryColumnRule] schema [%s] table [%s] column [%s] charset convert [UTFMB4] failed, error: %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, r.TableNameS, columnNameS, err)
			}
			columnNameS = stringutil.BytesToString(columnNameUtf8Raw)
		case constant.TaskFlowPostgresToMySQL, constant.TaskFlowPostgresToTiDB:
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(columnNameS), constant.MigratePostgreSQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return "", "", "", "", "", "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GenTableQueryColumnRule] schema [%s] table [%s] column [%s] charset convert [UTFMB4] failed, error: %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, r.TableNameS, columnNameS, err)
			}
			columnNameS = stringutil.BytesToString(columnNameUtf8Raw)
		default:
			return "", "", "", "", "", "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] isn't support, please contact author or reselect", r.TaskName, r.TaskFlow, r.TaskMode)
		}

		// skip ignore field
		if stringutil.IsContainedStringIgnoreCase(r.IgnoreSelectFields, columnNameS) {
			continue
		}

		switch r.TaskFlow {
		case constant.TaskFlowOracleToMySQL, constant.TaskFlowOracleToTiDB:
			var columnNameSO string
			if strings.EqualFold(datatypeS, constant.BuildInOracleDatatypeChar) || strings.EqualFold(datatypeS, constant.BuildInOracleDatatypeCharacter) {
				columnNameSO, err = OptimizerOracleDataMigrateColumnS(fmt.Sprintf("RTRIM(%s%s%s)", constant.StringSeparatorDoubleQuotes, columnNameS, constant.StringSeparatorDoubleQuotes), datatypeS, dataScaleS)
			} else {
				columnNameSO, err = OptimizerOracleDataMigrateColumnS(stringutil.StringBuilder(constant.StringSeparatorDoubleQuotes, columnNameS, constant.StringSeparatorDoubleQuotes), datatypeS, dataScaleS)
			}
			if err != nil {
				return "", "", "", "", "", "", err
			}
			columnNameSilSO = append(columnNameSilSO, columnNameSO)
		case constant.TaskFlowTiDBToOracle, constant.TaskFlowTiDBToPostgres:
			columnNameSO, err := OptimizerMYSQLCompatibleDataMigrateColumnS(stringutil.StringBuilder(constant.StringSeparatorBacktick, columnNameS, constant.StringSeparatorBacktick), datatypeS, rowCol["DATETIME_PRECISION"])
			if err != nil {
				return "", "", "", "", "", "", err
			}
			columnNameSilSO = append(columnNameSilSO, columnNameSO)
		case constant.TaskFlowPostgresToMySQL, constant.TaskFlowPostgresToTiDB:
			var columnNameSO string
			if strings.EqualFold(datatypeS, constant.BuildInPostgresDatatypeCharacter) {
				columnNameSO, err = OptimizerPostgresDataMigrateColumnS(fmt.Sprintf("RTRIM(%s%s%s)", constant.StringSeparatorDoubleQuotes, columnNameS, constant.StringSeparatorDoubleQuotes), datatypeS, dataScaleS)
			} else {
				columnNameSO, err = OptimizerPostgresDataMigrateColumnS(stringutil.StringBuilder(constant.StringSeparatorDoubleQuotes, columnNameS, constant.StringSeparatorDoubleQuotes), datatypeS, dataScaleS)
			}
			if err != nil {
				return "", "", "", "", "", "", err
			}
			columnNameSilSO = append(columnNameSilSO, columnNameSO)
		default:
			return "", "", "", "", "", "", fmt.Errorf("the task_flow [%s] isn't support, please contact author or reselect", r.TaskFlow)
		}

		var (
			columnNameSNew    string
			columnNameT       string
			isRollbackOracle  bool
			isBooleanPostgres bool
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
			switch r.TaskFlow {
			case constant.TaskFlowOracleToTiDB, constant.TaskFlowOracleToMySQL, constant.TaskFlowPostgresToTiDB, constant.TaskFlowPostgresToMySQL:
				columnNameT = fmt.Sprintf("%s%s%s", constant.StringSeparatorBacktick, val, constant.StringSeparatorBacktick)
			case constant.TaskFlowTiDBToOracle, constant.TaskFlowMySQLToOracle:
				if _, rollback := downstreamDBDatatypeRollback[val]; rollback {
					isRollbackOracle = true
				}
				if _, okTime := downstreamDBTimestampDatatype[val]; okTime {
					datetimePrecision := rowCol["DATETIME_PRECISION"]

					datetimePrecisionS, err := strconv.ParseInt(datetimePrecision, 10, 64)
					if err != nil {
						return "", "", "", "", "", "", err
					}
					columnNameT = stringutil.StringBuilder(`NVL(TO_CHAR("`, val, `",'YYYY-MM-DD HH24:MI:SS.FF`, strconv.FormatInt(datetimePrecisionS, 10), `'),'0')`)
				} else if _, okDate := downstreamDBDateDatatype[val]; okDate {
					columnNameT = stringutil.StringBuilder(`NVL(TO_CHAR("`, val, `",'YYYY-MM-DD HH24:MI:SS'),'0')`)
				} else {
					if _, okChar := downstreamDBCharDatatype[val]; okChar {
						// check whether the data type of the tidb field is char. If it is char, the data verification oracle needs to RTRIM the spaces（Leading spaces are not truncated at Oracle and TIDB/MYSQL）
						// 1. Oracle char data is written in fixed length. If the length is not enough, spaces are added. For query conditions, spaces are automatically added to the fixed length, and the data is automatically filled with spaces.
						// 2. Tidb char data is written in fixed length. If the length is not enough, spaces are added. For query conditions, spaces are not automatically added, and the data is not filled with spaces.
						// 3. Varchar2 vs varchar, both behave the same, no difference
						if strings.EqualFold(datatypeS, constant.BuildInMySQLDatatypeChar) {
							columnNameT = fmt.Sprintf("RTRIM(%s%s%s)", constant.StringSeparatorDoubleQuotes, val, constant.StringSeparatorDoubleQuotes)
						} else {
							columnNameT = fmt.Sprintf("%s%s%s", constant.StringSeparatorDoubleQuotes, val, constant.StringSeparatorDoubleQuotes)
						}
					} else {
						columnNameT = fmt.Sprintf("%s%s%s", constant.StringSeparatorDoubleQuotes, val, constant.StringSeparatorDoubleQuotes)
					}
				}
			case constant.TaskFlowTiDBToPostgres, constant.TaskFlowMySQLToPostgres:
				if _, isBoolean := downstreamDBBooleanDatatype[val]; isBoolean {
					isBooleanPostgres = true
				}
				if _, okChar := downstreamDBCharDatatype[val]; okChar {
					// check whether the data type of the tidb field is char. If it is char, the data verification postgres needs to RTRIM the spaces（Leading spaces are not truncated at Postgres and TIDB/MYSQL）
					// 1. Postgres char data is written in fixed length. If the length is not enough, spaces are added. For query conditions, spaces are automatically added to the fixed length, and the data is automatically filled with spaces.
					// 2. Tidb char data is written in fixed length. If the length is not enough, spaces are added. For query conditions, spaces are not automatically added, and the data is not filled with spaces.
					// 3. Varchar2 vs varchar, both behave the same, no difference
					if strings.EqualFold(datatypeS, constant.BuildInMySQLDatatypeChar) {
						columnNameT = fmt.Sprintf("RTRIM(%s%s%s)", constant.StringSeparatorDoubleQuotes, val, constant.StringSeparatorDoubleQuotes)
					} else {
						columnNameT = fmt.Sprintf("%s%s%s", constant.StringSeparatorDoubleQuotes, val, constant.StringSeparatorDoubleQuotes)
					}
				} else {
					columnNameT = fmt.Sprintf("%s%s%s", constant.StringSeparatorDoubleQuotes, val, constant.StringSeparatorDoubleQuotes)
				}
			default:
				return "", "", "", "", "", "", fmt.Errorf("the task_name [%s] schema [%s] taskflow [%s] column rule isn't support, please contact author", r.TaskName, r.SchemaNameS, r.TaskFlow)
			}
		} else {
			return "", "", "", "", "", "", fmt.Errorf("[GetTableColumnRule] get schema [%s] table [%s] column [%s] isn't exist, please contact author or double check again", r.SchemaNameS, r.TableNameS, columnNameS)
		}

		columnNameSliTO = append(columnNameSliTO, columnNameT)

		switch r.TaskFlow {
		case constant.TaskFlowOracleToMySQL, constant.TaskFlowOracleToTiDB:
			charLen := rowCol["CHAR_LENGTH"]

			charLenC, err := strconv.ParseInt(charLen, 10, 64)
			if err != nil {
				return "", "", "", "", "", "", err
			}
			columnNameSC, columnNameTC, err := OptimizerOracleMigrateMYSQLCompatibleDataCompareColumnST(columnNameS, datatypeS, stringutil.Min(charLenC, dataLenC), dataPrecisionS, dataScaleS, stringutil.StringUpper(r.DBCharsetS), constant.ORACLECharsetAL32UTF8, columnNameT, constant.MYSQLCharsetUTF8MB4)
			if err != nil {
				return "", "", "", "", "", "", err
			}
			columnNameSilSC = append(columnNameSilSC, columnNameSC)
			columnNameSliTC = append(columnNameSliTC, columnNameTC)

			if strings.EqualFold(datatypeS, constant.BuildInOracleDatatypeLong) || strings.EqualFold(datatypeS, constant.BuildInOracleDatatypeLongRAW) || strings.EqualFold(datatypeS, constant.BuildInOracleDatatypeBfile) {
				ignoreTypes = append(ignoreTypes, columnNameS)
			}
		case constant.TaskFlowPostgresToMySQL, constant.TaskFlowPostgresToTiDB:
			datetimePrecision := rowCol["DATETIME_PRECISION"]

			datetimePrecisionS, err := strconv.ParseInt(datetimePrecision, 10, 64)
			if err != nil {
				return "", "", "", "", "", "", err
			}
			columnNameSC, columnNameTC, err := OptimizerPostgresCompatibleMigrateMySQLCompatibleDataCompareColumnST(columnNameS, datatypeS, datetimePrecisionS, dataPrecisionS, dataScaleS, constant.PostgreSQLCharsetUTF8, columnNameT, constant.MYSQLCharsetUTF8MB4)
			if err != nil {
				return "", "", "", "", "", "", err
			}
			columnNameSilSC = append(columnNameSilSC, columnNameSC)
			columnNameSliTC = append(columnNameSliTC, columnNameTC)
		case constant.TaskFlowTiDBToOracle:
			datetimePrecision := rowCol["DATETIME_PRECISION"]

			datetimePrecisionS, err := strconv.ParseInt(datetimePrecision, 10, 64)
			if err != nil {
				return "", "", "", "", "", "", err
			}
			columnNameSC, columnNameTC, err := OptimizerMYSQLCompatibleMigrateOracleDataCompareColumnST(columnNameS, datatypeS, datetimePrecisionS, dataLenC, dataPrecisionS, dataScaleS, constant.MYSQLCharsetUTF8MB4, columnNameT, stringutil.StringUpper(r.DBCharsetT), constant.ORACLECharsetAL32UTF8)
			if err != nil {
				return "", "", "", "", "", "", err
			}
			columnNameSilSC = append(columnNameSilSC, columnNameSC)
			columnNameSliTC = append(columnNameSliTC, columnNameTC)

			// rollback
			if isRollbackOracle {
				ignoreTypes = append(ignoreTypes, columnNameS)
				logger.Warn("data compare task column datatype rollback tips",
					zap.String("task_name", r.TaskName),
					zap.String("task_mode", r.TaskMode),
					zap.String("task_flow", r.TaskFlow),
					zap.String("schema_name_s", r.SchemaNameS),
					zap.String("table_name_s", r.TableNameS),
					zap.Strings("column_name_s", ignoreTypes),
					zap.String("rollback tips", "the downstream oracle column have long、long raw and bfile datatype, rollback compare_crc32"))
			}
		case constant.TaskFlowTiDBToPostgres:
			datetimePrecision := rowCol["DATETIME_PRECISION"]

			datetimePrecisionS, err := strconv.ParseInt(datetimePrecision, 10, 64)
			if err != nil {
				return "", "", "", "", "", "", err
			}
			columnNameSC, columnNameTC, err := OptimizerMYSQLCompatibleMigratePostgresDataCompareColumnST(columnNameS, datatypeS, isBooleanPostgres, datetimePrecisionS, dataPrecisionS, dataScaleS, constant.MYSQLCharsetUTF8MB4, columnNameT, constant.PostgreSQLCharsetUTF8)
			if err != nil {
				return "", "", "", "", "", "", err
			}
			columnNameSilSC = append(columnNameSilSC, columnNameSC)
			columnNameSliTC = append(columnNameSliTC, columnNameTC)
		default:
			return "", "", "", "", "", "", fmt.Errorf("the task_flow [%s] isn't support, please contact author or reselect", r.TaskFlow)
		}
	}

	if r.OnlyDatabaseCompareRow {
		return stringutil.StringJoin(columnNameSilSO, constant.StringSeparatorComma), "COUNT(1) AS ROWSCOUNT",
			stringutil.StringJoin(columnNameSliTO, constant.StringSeparatorComma), "COUNT(1) AS ROWSCOUNT",
			stringutil.StringJoin(columnNameSilSO, constant.StringSeparatorComplexSymbol),
			stringutil.StringJoin(columnNameSliTO, constant.StringSeparatorComplexSymbol), nil
	}

	// if the database table column is existed in long and long raw datatype, the data compare task degenerate into program CRC32 data check
	// you can circumvent this by setting ignore-fields
	if len(ignoreTypes) > 0 {
		r.OnlyProgramCompareCRC32 = true
		logger.Warn("data compare task column datatype rollback action",
			zap.String("task_name", r.TaskName),
			zap.String("task_mode", r.TaskMode),
			zap.String("task_flow", r.TaskFlow),
			zap.String("schema_name_s", r.SchemaNameS),
			zap.String("table_name_s", r.TableNameS),
			zap.Strings("column_name_s", ignoreTypes),
			zap.String("rollback action", "compare_crc32"))
		return stringutil.StringJoin(columnNameSilSO, constant.StringSeparatorComma), "",
			stringutil.StringJoin(columnNameSliTO, constant.StringSeparatorComma), "",
			stringutil.StringJoin(columnNameSilSO, constant.StringSeparatorComplexSymbol),
			stringutil.StringJoin(columnNameSliTO, constant.StringSeparatorComplexSymbol), nil
	}

	switch r.TaskFlow {
	case constant.TaskFlowOracleToMySQL, constant.TaskFlowOracleToTiDB:
		return stringutil.StringJoin(columnNameSilSO, constant.StringSeparatorComma),
			stringutil.StringJoin(columnNameSilSC, constant.StringSplicingSymbol),
			stringutil.StringJoin(columnNameSliTO, constant.StringSeparatorComma),
			stringutil.StringJoin(columnNameSliTC, constant.StringSeparatorComma),
			stringutil.StringJoin(columnNameSilSC, constant.StringSeparatorComplexSymbol),
			stringutil.StringJoin(columnNameSliTC, constant.StringSeparatorComplexSymbol), nil
	case constant.TaskFlowTiDBToOracle, constant.TaskFlowMySQLToOracle:
		return stringutil.StringJoin(columnNameSilSO, constant.StringSeparatorComma),
			stringutil.StringJoin(columnNameSilSC, constant.StringSeparatorComma),
			stringutil.StringJoin(columnNameSliTO, constant.StringSeparatorComma),
			stringutil.StringJoin(columnNameSliTC, constant.StringSplicingSymbol),
			stringutil.StringJoin(columnNameSilSC, constant.StringSeparatorComplexSymbol),
			stringutil.StringJoin(columnNameSliTC, constant.StringSeparatorComplexSymbol), nil
	case constant.TaskFlowPostgresToMySQL, constant.TaskFlowPostgresToTiDB:
		return stringutil.StringJoin(columnNameSilSO, constant.StringSeparatorComma),
			stringutil.StringJoin(columnNameSilSC, constant.StringSplicingSymbol),
			stringutil.StringJoin(columnNameSliTO, constant.StringSeparatorComma),
			stringutil.StringJoin(columnNameSliTC, constant.StringSeparatorComma),
			stringutil.StringJoin(columnNameSilSC, constant.StringSeparatorComplexSymbol),
			stringutil.StringJoin(columnNameSliTC, constant.StringSeparatorComplexSymbol), nil
	case constant.TaskFlowTiDBToPostgres, constant.TaskFlowMySQLToPostgres:
		return stringutil.StringJoin(columnNameSilSO, constant.StringSeparatorComma),
			stringutil.StringJoin(columnNameSilSC, constant.StringSeparatorComma),
			stringutil.StringJoin(columnNameSliTO, constant.StringSeparatorComma),
			stringutil.StringJoin(columnNameSliTC, constant.StringSplicingSymbol),
			stringutil.StringJoin(columnNameSilSC, constant.StringSeparatorComplexSymbol),
			stringutil.StringJoin(columnNameSliTC, constant.StringSeparatorComplexSymbol), nil
	default:
		return "", "", "", "", "", "", fmt.Errorf("the task_name [%s] schema [%s] taskflow [%s] return isn't support, please contact author", r.TaskName, r.SchemaNameS, r.TaskFlow)
	}
}

func (r *DataCompareRule) GenSchemaTableTypeRule() string {
	return r.TableTypeS[r.TableNameS]
}

func (r *DataCompareRule) GenSchemaTableCompareMethodRule() string {
	if r.OnlyDatabaseCompareRow {
		return constant.DataCompareMethodDatabaseCheckRows
	}
	if r.OnlyProgramCompareCRC32 {
		return constant.DataCompareMethodProgramCheckCRC32
	}
	if !r.DisableDatabaseCompareMd5 {
		return constant.DataCompareMethodDatabaseCheckMD5
	}
	return constant.DataCompareMethodDatabaseCheckCRC32
}

func (r *DataCompareRule) GenSchemaTableCustomRule() (string, string, string, []string, string, string, error) {
	cr, err := model.GetIDataCompareRuleRW().GetDataCompareRule(r.Ctx, &rule.DataCompareRule{
		TaskName: r.TaskName, SchemaNameS: r.SchemaNameS, TableNameS: r.TableNameS})
	if err != nil {
		return "", "", "", nil, "", "", err
	}
	if !strings.EqualFold(cr.IgnoreSelectFields, "") {
		r.IgnoreSelectFields = stringutil.StringSplit(cr.IgnoreSelectFields, constant.StringSeparatorComma)
	}
	var ignoreColumnConditionFields []string
	if !strings.EqualFold(cr.IgnoreConditionFields, "") {
		ignoreColumnConditionFields = stringutil.StringSplit(cr.IgnoreConditionFields, constant.StringSeparatorComma)
	} else {
		ignoreColumnConditionFields = r.GlobalIgnoreConditionFields
	}

	var sqlHintS, sqlHintT string
	if strings.EqualFold(cr.SqlHintS, "") {
		sqlHintS = r.GlobalSqlHintS
	} else {
		sqlHintS = cr.SqlHintS
	}
	if strings.EqualFold(cr.SqlHintT, "") {
		sqlHintT = r.GlobalSqlHintT
	} else {
		sqlHintT = cr.SqlHintT
	}

	// replace compareConditionRangeT
	if strings.EqualFold(cr.CompareConditionRangeT, "") {
		cr.CompareConditionRangeT = cr.CompareConditionRangeS
	}

	return cr.CompareConditionField, cr.CompareConditionRangeS, cr.CompareConditionRangeT, ignoreColumnConditionFields, sqlHintS, sqlHintT, nil
}
