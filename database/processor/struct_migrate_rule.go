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

	"github.com/wentaojin/dbms/database/mapping"

	"github.com/wentaojin/dbms/model/buildin"

	"github.com/wentaojin/dbms/logger"
	"github.com/wentaojin/dbms/model"
	"github.com/wentaojin/dbms/model/migrate"
	"github.com/wentaojin/dbms/model/rule"
	"github.com/wentaojin/dbms/utils/constant"
	"github.com/wentaojin/dbms/utils/stringutil"
	"go.uber.org/zap"
)

type StructMigrateRule struct {
	Ctx                      context.Context                  `json:"-"`
	TaskName                 string                           `json:"taskName"`
	TaskMode                 string                           `json:"taskMode"`
	TaskFlow                 string                           `json:"taskFlow"`
	SchemaNameS              string                           `json:"schemaNameS"`
	TableNameS               string                           `json:"tableNameS"`
	TablePrimaryAttrs        []map[string]string              `json:"tablePrimaryAttrs"`
	TableColumnsAttrs        []map[string]string              `json:"tableColumnsAttrs"`
	TableCommentAttrs        []map[string]string              `json:"tableCommentAttrs"`
	TableCharsetAttr         string                           `json:"tableCharsetAttr"`
	TableCollationAttr       string                           `json:"tableCollationAttr"`
	CaseFieldRuleT           string                           `json:"caseFieldRuleT"`
	CreateIfNotExist         bool                             `json:"createIfNotExist"`
	DBVersionS               string                           `json:"dbVersionS"`
	DBCharsetS               string                           `json:"dbCharsetS"` // the string data charset conversation
	DBCharsetT               string                           `json:"dbCharsetT"` // the string data charset conversation
	BuildinDatatypeRules     []*buildin.BuildinDatatypeRule   `json:"-"`
	BuildinDefaultValueRules []*buildin.BuildinDefaultvalRule `json:"-"`
}

func (r *StructMigrateRule) GetCreatePrefixRule() string {
	switch r.TaskFlow {
	case constant.TaskFlowOracleToTiDB, constant.TaskFlowOracleToMySQL, constant.TaskFlowPostgresToMySQL, constant.TaskFlowPostgresToTiDB:
		if r.CreateIfNotExist {
			return `CREATE TABLE IF NOT EXISTS`
		}
		return `CREATE TABLE`
	default:
		return `CREATE TABLE`
	}
}

func (r *StructMigrateRule) GetSchemaNameRule() (map[string]string, error) {
	schemaRoute := make(map[string]string)

	routeRule, err := model.GetIMigrateSchemaRouteRW().GetSchemaRouteRule(r.Ctx, &rule.SchemaRouteRule{
		TaskName: r.TaskName, SchemaNameS: r.SchemaNameS})
	if err != nil {
		return schemaRoute, err
	}

	var schemaNameS string
	switch r.TaskFlow {
	case constant.TaskFlowOracleToMySQL, constant.TaskFlowOracleToTiDB:
		convertUtf8Raw, err := stringutil.CharsetConvert([]byte(r.SchemaNameS), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
		if err != nil {
			return nil, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GetSchemaNameRule] schema [%s] charset convert failed, %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, err)
		}
		schemaNameS = stringutil.BytesToString(convertUtf8Raw)
	case constant.TaskFlowPostgresToMySQL, constant.TaskFlowPostgresToTiDB:
		convertUtf8Raw, err := stringutil.CharsetConvert([]byte(r.SchemaNameS), constant.MigratePostgreSQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
		if err != nil {
			return nil, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GetSchemaNameRule] schema [%s] charset convert failed, %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, err)
		}
		schemaNameS = stringutil.BytesToString(convertUtf8Raw)
	default:
		return nil, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] isn't support, please contact author or reselect", r.TaskName, r.TaskFlow, r.TaskMode)
	}

	var schemaNameSNew string

	if !strings.EqualFold(routeRule.SchemaNameT, "") {
		schemaNameSNew = routeRule.SchemaNameT
	} else {
		schemaNameSNew = schemaNameS
	}

	if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueStructMigrateCaseFieldRuleLower) {
		schemaNameSNew = strings.ToLower(schemaNameSNew)
	}
	if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueStructMigrateCaseFieldRuleUpper) {
		schemaNameSNew = strings.ToUpper(schemaNameSNew)
	}

	schemaRoute[schemaNameS] = schemaNameSNew

	return schemaRoute, nil
}

func (r *StructMigrateRule) GetTableNameRule() (map[string]string, error) {
	tableRoute := make(map[string]string)
	routeRule, err := model.GetIMigrateTableRouteRW().GetTableRouteRule(r.Ctx, &rule.TableRouteRule{
		TaskName: r.TaskName, SchemaNameS: r.SchemaNameS, TableNameS: r.TableNameS})
	if err != nil {
		return tableRoute, err
	}

	var tableNameS string
	switch r.TaskFlow {
	case constant.TaskFlowOracleToMySQL, constant.TaskFlowOracleToTiDB:
		convertUtf8Raw, err := stringutil.CharsetConvert([]byte(r.TableNameS), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
		if err != nil {
			return nil, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GetTableNameRule] schema [%s] table [%v] charset convert failed, %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, r.TableNameS, err)
		}
		tableNameS = stringutil.BytesToString(convertUtf8Raw)
	case constant.TaskFlowPostgresToMySQL, constant.TaskFlowPostgresToTiDB:
		convertUtf8Raw, err := stringutil.CharsetConvert([]byte(r.TableNameS), constant.MigratePostgreSQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
		if err != nil {
			return nil, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] [GetTableNameRule] schema [%s] table [%v] charset convert failed, %v", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, r.TableNameS, err)
		}
		tableNameS = stringutil.BytesToString(convertUtf8Raw)
	default:
		return nil, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] isn't support, please contact author or reselect", r.TaskName, r.TaskFlow, r.TaskMode)
	}

	var tableNameSNew string
	if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueStructMigrateCaseFieldRuleLower) {
		tableNameSNew = strings.ToLower(tableNameS)
	}
	if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueStructMigrateCaseFieldRuleUpper) {
		tableNameSNew = strings.ToUpper(tableNameS)
	}
	if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueStructMigrateCaseFieldRuleOrigin) {
		tableNameSNew = tableNameS
	}

	if !strings.EqualFold(routeRule.TableNameT, "") {
		tableRoute[tableNameS] = routeRule.TableNameT
	} else {
		tableRoute[tableNameS] = tableNameSNew
	}
	return tableRoute, nil
}

func (r *StructMigrateRule) GetCaseFieldRule() string {
	return r.CaseFieldRuleT
}

// GetTableColumnRule used for get custom table column rule
// column datatype rule priority:
// - column level
// - table level
// - task level
// - default level
func (r *StructMigrateRule) GetTableColumnRule() (map[string]string, map[string]string, map[string]string, error) {
	columnRules := make(map[string]string)
	columnDatatypeRules := make(map[string]string)
	columnDefaultValueRules := make(map[string]string)

	columnRoutes, err := model.GetIMigrateColumnRouteRW().FindColumnRouteRule(r.Ctx, &rule.ColumnRouteRule{
		TaskName:    r.TaskName,
		SchemaNameS: r.SchemaNameS,
		TableNameS:  r.TableNameS,
	})
	if err != nil {
		return columnRules, columnDatatypeRules, columnDefaultValueRules, err
	}
	structTaskRules, err := model.GetIStructMigrateTaskRuleRW().QueryTaskStructRule(r.Ctx, &migrate.TaskStructRule{TaskName: r.TaskName})
	if err != nil {
		return columnRules, columnDatatypeRules, columnDefaultValueRules, err
	}
	structSchemaRules, err := model.GetIStructMigrateSchemaRuleRW().QuerySchemaStructRule(r.Ctx, &migrate.SchemaStructRule{
		TaskName:    r.TaskName,
		SchemaNameS: r.SchemaNameS})
	if err != nil {
		return columnRules, columnDatatypeRules, columnDefaultValueRules, err
	}
	structTableRules, err := model.GetIStructMigrateTableRuleRW().QueryTableStructRule(r.Ctx, &migrate.TableStructRule{
		TaskName:    r.TaskName,
		SchemaNameS: r.SchemaNameS,
		TableNameS:  r.TableNameS})
	if err != nil {
		return columnRules, columnDatatypeRules, columnDefaultValueRules, err
	}
	structColumnRules, err := model.GetIStructMigrateColumnRuleRW().QueryColumnStructRule(r.Ctx, &migrate.ColumnStructRule{
		TaskName:    r.TaskName,
		SchemaNameS: r.SchemaNameS,
		TableNameS:  r.TableNameS})
	if err != nil {
		return columnRules, columnDatatypeRules, columnDefaultValueRules, err
	}

	for _, c := range r.TableColumnsAttrs {
		var columnName string
		switch r.TaskFlow {
		case constant.TaskFlowOracleToMySQL, constant.TaskFlowOracleToTiDB:
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(c["COLUMN_NAME"]), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return columnRules, columnDatatypeRules, columnDefaultValueRules, fmt.Errorf("[GetTableColumnRule] the upstream database schema [%s] table [%s] column [%s] charset [%v] convert [UTF8MB4] failed, error: %v", r.SchemaNameS, r.TableNameS, r.DBCharsetS, c["COLUMN_NAME"], err)
			}

			columnName = stringutil.BytesToString(columnNameUtf8Raw)
		case constant.TaskFlowPostgresToMySQL, constant.TaskFlowPostgresToTiDB:
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(c["COLUMN_NAME"]), constant.MigratePostgreSQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return columnRules, columnDatatypeRules, columnDefaultValueRules, fmt.Errorf("[GetTableColumnRule] the upstream database schema [%s] table [%s] column [%s] charset [%v] convert [UTF8MB4] failed, error: %v", r.SchemaNameS, r.TableNameS, r.DBCharsetS, c["COLUMN_NAME"], err)
			}

			columnName = stringutil.BytesToString(columnNameUtf8Raw)
		default:
			return columnRules, columnDatatypeRules, columnDefaultValueRules, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] isn't support, please contact author or reselect", r.TaskName, r.TaskFlow, r.TaskMode)
		}

		// column name caseFieldRule
		var columnNameTNew string

		if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueStructMigrateCaseFieldRuleLower) {
			columnNameTNew = strings.ToLower(columnName)
		}
		if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueStructMigrateCaseFieldRuleUpper) {
			columnNameTNew = strings.ToUpper(columnName)
		}
		if strings.EqualFold(r.CaseFieldRuleT, constant.ParamValueStructMigrateCaseFieldRuleOrigin) {
			columnNameTNew = columnName
		}

		columnRules[columnName] = columnNameTNew

		var columnDefaultValues, columnComment string
		switch r.TaskFlow {
		case constant.TaskFlowOracleToMySQL, constant.TaskFlowOracleToTiDB:
			defaultValUtf8Raw, err := stringutil.CharsetConvert([]byte(c["DATA_DEFAULT"]), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return columnRules, columnDatatypeRules, columnDefaultValueRules, fmt.Errorf("[GetTableColumnRule] the upstream database schema [%s] table [%s] column [%s] default value [%s] charset [%v] convert [UTF8MB4] failed, error: %v", r.SchemaNameS, r.TableNameS, r.DBCharsetS, c["COLUMN_NAME"], c["DATA_DEFAULT"], err)
			}
			columnDefaultValues = stringutil.BytesToString(defaultValUtf8Raw)

			commentUtf8Raw, err := stringutil.CharsetConvert([]byte(c["COMMENTS"]), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return columnRules, columnDatatypeRules, columnDefaultValueRules, fmt.Errorf("[GetTableColumnRule] the upstream database schema [%s] table [%s] column [%s] comment [%s] charset [%v] convert [UTF8MB4] failed, error: %v", r.SchemaNameS, r.TableNameS, r.DBCharsetS, c["COLUMN_NAME"], c["COMMENTS"], err)
			}
			columnComment = stringutil.BytesToString(commentUtf8Raw)
		case constant.TaskFlowPostgresToMySQL, constant.TaskFlowPostgresToTiDB:
			defaultValUtf8Raw, err := stringutil.CharsetConvert([]byte(c["DATA_DEFAULT"]), constant.MigratePostgreSQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return columnRules, columnDatatypeRules, columnDefaultValueRules, fmt.Errorf("[GetTableColumnRule] the upstream database schema [%s] table [%s] column [%s] default value [%s] charset [%v] convert [UTF8MB4] failed, error: %v", r.SchemaNameS, r.TableNameS, r.DBCharsetS, c["COLUMN_NAME"], c["DATA_DEFAULT"], err)
			}
			columnDefaultValues = stringutil.BytesToString(defaultValUtf8Raw)

			commentUtf8Raw, err := stringutil.CharsetConvert([]byte(c["COMMENTS"]), constant.MigratePostgreSQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return columnRules, columnDatatypeRules, columnDefaultValueRules, fmt.Errorf("[GetTableColumnRule] the upstream database schema [%s] table [%s] column [%s] comment [%s] charset [%v] convert [UTF8MB4] failed, error: %v", r.SchemaNameS, r.TableNameS, r.DBCharsetS, c["COLUMN_NAME"], c["COMMENTS"], err)
			}
			columnComment = stringutil.BytesToString(commentUtf8Raw)
		default:
			return columnRules, columnDatatypeRules, columnDefaultValueRules, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] isn't support, please contact author or reselect", r.TaskName, r.TaskFlow, r.TaskMode)
		}

		var (
			originColumnType, buildInColumnType string
		)
		// task flow
		switch r.TaskFlow {
		case constant.TaskFlowOracleToMySQL, constant.TaskFlowOracleToTiDB:
			originColumnType, buildInColumnType, err = mapping.OracleDatabaseTableColumnMapMYSQLCompatibleDatatypeRule(
				r.TaskFlow,
				&mapping.Column{
					ColumnName:    c["COLUMN_NAME"],
					Datatype:      c["DATA_TYPE"],
					CharUsed:      c["CHAR_USED"],
					CharLength:    c["CHAR_LENGTH"],
					DataPrecision: c["DATA_PRECISION"],
					DataLength:    c["DATA_LENGTH"],
					DataScale:     c["DATA_SCALE"],
					DataDefault:   columnDefaultValues,
					Nullable:      c["NULLABLE"],
					Comment:       columnComment,
				}, r.BuildinDatatypeRules)
			if err != nil {
				return nil, nil, nil, err
			}
			// priority, return target database table column datatype
			convertColumnDatatype, convertColumnDefaultValue, err := mapping.OracleHandleColumnRuleWithPriority(
				r.TableNameS,
				c["COLUMN_NAME"],
				originColumnType,
				buildInColumnType,
				columnDefaultValues,
				constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)],
				constant.MigrateMySQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetT)],
				r.BuildinDefaultValueRules,
				structTaskRules,
				structSchemaRules,
				structTableRules,
				structColumnRules)
			if err != nil {
				return nil, nil, nil, err
			}

			// column datatype upper case field rule
			columnDatatypeRules[columnName] = stringutil.StringUpper(convertColumnDatatype)
			columnDefaultValueRules[columnName] = convertColumnDefaultValue
		case constant.TaskFlowPostgresToMySQL, constant.TaskFlowPostgresToTiDB:
			originColumnType, buildInColumnType, err = mapping.PostgresDatabaseTableColumnMapMYSQLCompatibleDatatypeRule(
				r.TaskFlow,
				&mapping.Column{
					ColumnName:        c["COLUMN_NAME"],
					Datatype:          c["DATA_TYPE"],
					DataLength:        c["DATA_LENGTH"],
					DataPrecision:     c["DATA_PRECISION"],
					DataScale:         c["DATA_SCALE"],
					DataDefault:       columnDefaultValues,
					Nullable:          c["NULLABLE"],
					DatetimePrecision: c["DATETIME_PRECISION"],
					Comment:           columnComment,
				}, r.BuildinDatatypeRules)
			if err != nil {
				return nil, nil, nil, err
			}
			// priority, return target database table column datatype
			convertColumnDatatype, convertColumnDefaultValue, err := mapping.PostgresHandleColumnRuleWithPriority(
				r.TableNameS,
				c["COLUMN_NAME"],
				originColumnType,
				buildInColumnType,
				columnDefaultValues,
				constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)],
				constant.MigrateMySQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetT)],
				r.BuildinDefaultValueRules,
				structTaskRules,
				structSchemaRules,
				structTableRules,
				structColumnRules)
			if err != nil {
				return nil, nil, nil, err
			}
			// column datatype upper case field rule
			columnDatatypeRules[columnName] = stringutil.StringUpper(convertColumnDatatype)
			columnDefaultValueRules[columnName] = convertColumnDefaultValue
		default:
			return nil, nil, nil, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] schema [%s] taskflow [%s] column rule isn't support, please contact author", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, r.TaskFlow)
		}
	}

	for _, c := range columnRoutes {
		if _, exist := columnRules[c.ColumnNameS]; exist {
			columnRules[c.ColumnNameS] = c.ColumnNameT
		}
	}
	return columnRules, columnDatatypeRules, columnDefaultValueRules, nil
}

func (r *StructMigrateRule) GetTableAttributesRule() (string, error) {
	attrs, err := model.GetIStructMigrateTableAttrsRuleRW().GetTableAttrsRule(r.Ctx, &migrate.TableAttrsRule{
		TaskName:    r.TaskName,
		SchemaNameS: r.SchemaNameS,
		TableNameS:  r.TableNameS,
	})
	if err != nil {
		return "", err
	}

	switch r.TaskFlow {
	case constant.TaskFlowOracleToTiDB, constant.TaskFlowPostgresToTiDB:
		tableNameAttr := make(map[string]string)
		for _, attr := range attrs {
			tableNameAttr[attr.TableNameS] = attr.TableAttrT
		}

		if len(tableNameAttr) > 0 {
			// prority table -> *
			if val, ok := tableNameAttr[r.TableNameS]; ok {
				return val, nil
			}
			if val, ok := tableNameAttr[constant.StringSeparatorAsterisk]; ok {
				return val, nil
			}
			return "", fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] schema_name_s [%s] taskflow [%s] table_name_s [%s] attributes rule get failed, please contact author", r.TaskName, r.TaskFlow, r.TaskMode, r.SchemaNameS, r.TaskFlow, r.TableNameS)
		}
		return "", nil
	default:
		logger.Warn("get table rule",
			zap.String("task_name", r.TableNameS),
			zap.String("task_flow", r.TaskFlow),
			zap.String("task_notes", fmt.Sprintf("the task_name [%s] schema_name_s [%s] taskflow [%s] attributes rule isn't support, only the downstream tidb database is support,skip operator", r.TaskName, r.SchemaNameS, r.TaskFlow)),
			zap.String("operator", "ignore"))
		return "", nil
	}
}

func (r *StructMigrateRule) GetTableCommentRule() (string, error) {
	var tableComment string
	if len(r.TableCommentAttrs) > 1 {
		return tableComment, fmt.Errorf("the upstream database schema [%s] table [%s] comments [%s] records are over one, current value is [%d]", r.SchemaNameS, r.TableNameS, r.TableCommentAttrs[0]["COMMENTS"], len(r.TableCommentAttrs))
	}
	if len(r.TableCommentAttrs) == 0 || strings.EqualFold(r.TableCommentAttrs[0]["COMMENTS"], "") {
		return tableComment, nil
	}

	switch r.TaskFlow {
	case constant.TaskFlowOracleToTiDB, constant.TaskFlowOracleToMySQL:
		convertUtf8Raw, err := stringutil.CharsetConvert([]byte(r.TableCommentAttrs[0]["COMMENTS"]), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
		if err != nil {
			return tableComment, fmt.Errorf("[GetTableCommentRule] the upstream database schema [%s] table [%s] comments [%s] charset [%v] convert [UTF8MB4] failed, error: %v", r.SchemaNameS, r.TableNameS, r.TableCommentAttrs[0]["COMMENTS"], r.DBCharsetS, err)
		}
		tableComment = stringutil.EscapeDatabaseSingleQuotesSpecialLetters(convertUtf8Raw, '\'')
		return tableComment, nil
	case constant.TaskFlowPostgresToTiDB, constant.TaskFlowPostgresToMySQL:
		convertUtf8Raw, err := stringutil.CharsetConvert([]byte(r.TableCommentAttrs[0]["COMMENTS"]), constant.MigratePostgreSQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
		if err != nil {
			return tableComment, fmt.Errorf("[GetTableCommentRule] the upstream database schema [%s] table [%s] comments [%s] charset [%v] convert [UTF8MB4] failed, error: %v", r.SchemaNameS, r.TableNameS, r.TableCommentAttrs[0]["COMMENTS"], r.DBCharsetS, err)
		}
		tableComment = stringutil.EscapeDatabaseSingleQuotesSpecialLetters(convertUtf8Raw, '\'')
		return tableComment, nil
	default:
		return tableComment, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] isn't support, please contact author or reselect", r.TaskName, r.TaskFlow, r.TaskMode)
	}
}

func (r *StructMigrateRule) GetTableCharsetRule() (string, error) {
	if val, ok := constant.MigrateTableStructureDatabaseCharsetMap[r.TaskFlow][stringutil.StringUpper(r.TableCharsetAttr)]; ok {
		return val, nil
	}
	return "", fmt.Errorf("[GetTableCharsetRule] the upstream database schema [%s] table [%s] charset [%s] mapping failed, please checking", r.SchemaNameS, r.TableNameS, r.TableCharsetAttr)
}

func (r *StructMigrateRule) GetTableCollationRule() (string, error) {
	if val, ok := constant.MigrateTableStructureDatabaseCollationMap[r.TaskFlow][stringutil.StringUpper(r.TableCollationAttr)][constant.MigrateTableStructureDatabaseCharsetMap[r.TaskFlow][stringutil.StringUpper(r.TableCharsetAttr)]]; ok {
		return val, nil
	}
	return "", fmt.Errorf("[GetTableCharsetRule] the upstream database schema [%s] table [%s] collation [%s] mapping failed, please checking", r.SchemaNameS, r.TableNameS, r.TableCollationAttr)
}

func (r *StructMigrateRule) GetTableColumnCharsetRule() (map[string]string, error) {
	columnCharsetMap := make(map[string]string)
	for _, rowCol := range r.TableColumnsAttrs {
		var (
			columnName    string
			columnCharset string
		)
		switch r.TaskFlow {
		case constant.TaskFlowOracleToTiDB, constant.TaskFlowOracleToMySQL:
			if stringutil.IsContainedStringIgnoreCase(constant.OracleCompatibleDatabaseTableColumnSupportCharsetCollationDatatype, rowCol["DATA_TYPE"]) {
				// oracle database isnot setting column charset
				tableCharset, err := r.GetTableCharsetRule()
				if err != nil {
					return nil, err
				}
				columnCharset = tableCharset
			} else {
				columnCharset = ""
			}
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(rowCol["COLUMN_NAME"]), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return columnCharsetMap, fmt.Errorf("[GetTableColumnCharsetRule] the upstream database schema [%s] table [%s] column [%s] charset convert [UTF8MB4] failed, error: %v", r.SchemaNameS, r.TableNameS, rowCol["COLUMN_NAME"], err)
			}
			columnName = stringutil.BytesToString(columnNameUtf8Raw)
		case constant.TaskFlowPostgresToTiDB, constant.TaskFlowPostgresToMySQL:
			// postgres support collation version >= 9.1
			// check column sort collation
			if stringutil.IsContainedStringIgnoreCase(constant.PostgresDatabaseTableColumnSupportCharsetCollationDatatype, rowCol["DATA_TYPE"]) {
				if val, ok := constant.MigrateTableStructureDatabaseCharsetMap[r.TaskFlow][stringutil.StringUpper(rowCol["CHARSET"])]; ok {
					columnCharset = val
				} else {
					if strings.EqualFold(stringutil.StringUpper(rowCol["CHARSET"]), "UNKNOWN") {
						// return the table charset
						tableCharset, err := r.GetTableCharsetRule()
						if err != nil {
							return nil, err
						}
						columnCharset = tableCharset
					} else {
						return columnCharsetMap, fmt.Errorf("[GetTableColumnCharsetRule] the upstream database schema [%s] table [%s] column [%s] datatype [%s] charset [%s] mapping failed, please checking", r.SchemaNameS, r.TableNameS, rowCol["COLUMN_NAME"], rowCol["DATA_TYPE"], rowCol["CHARSET"])
					}
				}
			} else {
				columnCharset = ""
			}
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(rowCol["COLUMN_NAME"]), constant.MigratePostgreSQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return columnCharsetMap, fmt.Errorf("[GetTableColumnCollationRule] the upstream database schema [%s] table [%s] column [%s] charset convert [UTF8MB4] failed, error: %v", r.SchemaNameS, r.TableNameS, rowCol["COLUMN_NAME"], err)
			}
			columnName = stringutil.BytesToString(columnNameUtf8Raw)
		default:
			return nil, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] isn't support, please contact author or reselect", r.TaskName, r.TaskFlow, r.TaskMode)
		}

		columnCharsetMap[columnName] = columnCharset
	}
	return columnCharsetMap, nil
}

func (r *StructMigrateRule) GetTableColumnCollationRule() (map[string]string, error) {
	columnCollationMap := make(map[string]string)
	for _, rowCol := range r.TableColumnsAttrs {
		var (
			columnName      string
			columnCollation string
		)
		switch r.TaskFlow {
		case constant.TaskFlowOracleToTiDB, constant.TaskFlowOracleToMySQL:
			// the oracle 12.2 and the above version support column collation
			// the oracle 12.2 and the below version isn't support column collation, and set ""
			// check column sort collation
			if stringutil.IsContainedStringIgnoreCase(constant.OracleCompatibleDatabaseTableColumnSupportCharsetCollationDatatype, rowCol["DATA_TYPE"]) {
				if stringutil.VersionOrdinal(r.DBVersionS) >= stringutil.VersionOrdinal(constant.OracleDatabaseTableAndColumnSupportVersion) {
					if val, ok := constant.MigrateTableStructureDatabaseCollationMap[r.TaskFlow][stringutil.StringUpper(rowCol["COLLATION"])][constant.MigrateTableStructureDatabaseCharsetMap[r.TaskFlow][stringutil.StringUpper(r.TableCharsetAttr)]]; ok {
						columnCollation = val
					} else {
						return columnCollationMap, fmt.Errorf("the upstream database schema [%s] table [%s] column [%s] datatype [%s] charset [%s] collation [%s] mapping failed, please checking", r.SchemaNameS, r.TableNameS, rowCol["COLUMN_NAME"], rowCol["DATA_TYPE"], r.TableCharsetAttr, rowCol["COLLATION"])
					}
				} else {
					// return the table charset collation
					tableCollation, err := r.GetTableCollationRule()
					if err != nil {
						return nil, err
					}
					columnCollation = tableCollation
				}
			} else {
				// exclude the column with the integer datatype
				columnCollation = ""
			}
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(rowCol["COLUMN_NAME"]), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return columnCollationMap, fmt.Errorf("[GetTableColumnCollationRule] the upstream database schema [%s] table [%s] column [%s] charset convert [UTF8MB4] failed, error: %v", r.SchemaNameS, r.TableNameS, rowCol["COLUMN_NAME"], err)
			}
			columnName = stringutil.BytesToString(columnNameUtf8Raw)
		case constant.TaskFlowPostgresToTiDB, constant.TaskFlowPostgresToMySQL:
			// postgres support collation version >= 9.1
			// check column sort collation
			if stringutil.IsContainedStringIgnoreCase(constant.PostgresDatabaseTableColumnSupportCharsetCollationDatatype, rowCol["DATA_TYPE"]) {
				if val, ok := constant.MigrateTableStructureDatabaseCollationMap[r.TaskFlow][stringutil.StringUpper(rowCol["COLLATION"])][constant.MigrateTableStructureDatabaseCharsetMap[r.TaskFlow][stringutil.StringUpper(r.TableCharsetAttr)]]; ok {
					columnCollation = val
				} else {
					if strings.EqualFold(stringutil.StringUpper(rowCol["CHARSET"]), "UNKNOWN") {
						// return the table charset collation
						tableCollation, err := r.GetTableCollationRule()
						if err != nil {
							return nil, err
						}
						columnCollation = tableCollation
					} else {
						return columnCollationMap, fmt.Errorf("[GetTableColumnCollationRule] the upstream database schema [%s] table [%s] column [%s] datatype [%s] charset [%s] collation [%s] mapping failed, please checking", r.SchemaNameS, r.TableNameS, rowCol["COLUMN_NAME"], rowCol["DATA_TYPE"], r.TableCharsetAttr, rowCol["COLLATION"])
					}
				}
			} else {
				columnCollation = ""
			}
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(rowCol["COLUMN_NAME"]), constant.MigratePostgreSQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return columnCollationMap, fmt.Errorf("[GetTableColumnCollationRule] the upstream database schema [%s] table [%s] column [%s] charset convert [UTF8MB4] failed, error: %v", r.SchemaNameS, r.TableNameS, rowCol["COLUMN_NAME"], err)
			}
			columnName = stringutil.BytesToString(columnNameUtf8Raw)
		default:
			return nil, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] isn't support, please contact author or reselect", r.TaskName, r.TaskFlow, r.TaskMode)
		}
		columnCollationMap[columnName] = columnCollation
	}
	return columnCollationMap, nil
}

func (r *StructMigrateRule) GetTableColumnCommentRule() (map[string]string, error) {
	columnCommentMap := make(map[string]string)
	for _, rowCol := range r.TableColumnsAttrs {
		var columnName, columnComment string
		switch r.TaskFlow {
		case constant.TaskFlowOracleToTiDB, constant.TaskFlowOracleToMySQL:
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(rowCol["COLUMN_NAME"]), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return columnCommentMap, fmt.Errorf("[GetTableColumnCommentRule] the upstream database schema [%s] table [%s] column [%s] charset convert [UTF8MB4] failed, error: %v", r.SchemaNameS, r.TableNameS, rowCol["COLUMN_NAME"], err)
			}
			columnName = stringutil.BytesToString(columnNameUtf8Raw)

			commentUtf8Raw, err := stringutil.CharsetConvert([]byte(rowCol["COMMENTS"]), constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return columnCommentMap, fmt.Errorf("[GetTableColumnCommentRule] the upstream database schema [%s] table [%s] column [%s] comment [%s] charset convert failed, %v", r.SchemaNameS, r.TableNameS, rowCol["COLUMN_NAME"], rowCol["COMMENTS"], err)
			}

			columnComment = stringutil.EscapeDatabaseSingleQuotesSpecialLetters(commentUtf8Raw, '\'')
		case constant.TaskFlowPostgresToTiDB, constant.TaskFlowPostgresToMySQL:
			columnNameUtf8Raw, err := stringutil.CharsetConvert([]byte(rowCol["COLUMN_NAME"]), constant.MigratePostgreSQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return columnCommentMap, fmt.Errorf("[GetTableColumnCommentRule] the upstream database schema [%s] table [%s] column [%s] charset convert [UTF8MB4] failed, error: %v", r.SchemaNameS, r.TableNameS, rowCol["COLUMN_NAME"], err)
			}
			columnName = stringutil.BytesToString(columnNameUtf8Raw)

			commentUtf8Raw, err := stringutil.CharsetConvert([]byte(rowCol["COMMENTS"]), constant.MigratePostgreSQLCompatibleCharsetStringConvertMapping[stringutil.StringUpper(r.DBCharsetS)], constant.CharsetUTF8MB4)
			if err != nil {
				return columnCommentMap, fmt.Errorf("[GetTableColumnCommentRule] the upstream database schema [%s] table [%s] column [%s] comment [%s] charset convert failed, %v", r.SchemaNameS, r.TableNameS, rowCol["COLUMN_NAME"], rowCol["COMMENTS"], err)
			}

			columnComment = stringutil.EscapeDatabaseSingleQuotesSpecialLetters(commentUtf8Raw, '\'')
		default:
			return nil, fmt.Errorf("the task_name [%s] task_flow [%s] and task_mode [%s] isn't support, please contact author or reselect", r.TaskName, r.TaskFlow, r.TaskMode)
		}
		columnCommentMap[columnName] = columnComment
	}
	return columnCommentMap, nil
}

func (r *StructMigrateRule) String() string {
	jsonStr, _ := stringutil.MarshalJSON(r)
	return jsonStr
}
