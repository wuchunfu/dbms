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
package tidb

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/wentaojin/dbms/message"
	"github.com/wentaojin/dbms/model/rule"
	"github.com/wentaojin/dbms/utils/constant"
	"github.com/wentaojin/dbms/utils/stringutil"
)

/*
TypeTinyBlob   ColumnType = 249 // TINYTEXT/TINYBLOB -> 249
TypeMediumBlob ColumnType = 250 // MEDIUMTEXT/MEDIUMBLOB -> 250
TypeLongBlob   ColumnType = 251 // LONGTEXT/LONGBLOB -> 251
TypeBlob       ColumnType = 252 // TEXT/BLOB -> 252

The same field type id of the message event generated by tidb ticdc may represent different data types, and different data types correspond to different downstream database data types, such as text -> clob, blob -> blob. The consumption process cannot identify the specific downstream data type and needs to identify it based on downstream metadata to determine whether it is passed in string or []byte format.
*/
// RowChangedEvent store the ddl and dml event
type RowChangedEvent struct {
	SchemaName string `json:"schemaName"`
	TableName  string `json:"tableName"`
	QueryType  string `json:"queryType"`
	CommitTs   uint64 `json:"commitTS"`

	IsDDL    bool   `json:"isDDL"`
	DdlQuery string `json:"ddlQuery"`

	// The table synchronized by TiCDC needs to have at least one valid index. The definition of a valid index is as follows:
	// 1,	the primary key (PRIMARY KEY) is a valid index.
	// 2,	each column in the unique index (UNIQUE INDEX) is explicitly defined as NOT NULL in the table structure and there are no virtual generated columns (VIRTUAL GENERATED COLUMNS).
	// Data synchronization TiCDC will select a valid index as the Handle Index. The HandleKeyFlag of the columns included in the Handle Index is set to 1.
	ValidUniqColumns map[string]interface{} `json:"validUniqColumns"`

	// Represents all field names, but Kafka information does not carry field offsets, and is not guaranteed to be strictly consistent with the order in which the fields are created in the table structure. Sort by field name
	Columns    []*columnAttr     `json:"columns"`
	ColumnType map[string]string `json:"columnType"`

	NewColumnData map[string]interface{} `json:"newColumnData"`
	// Only when this message is generated by an Update type event, record the name of each column and the data value before Update
	OldColumnData map[string]interface{} `json:"oldColumnData"`
}

type columnAttr struct {
	ColumnName        string `json:"columnName"`
	ColumnType        string `json:"columnType"`
	IsGeneratedColumn bool   `json:"isGeneratedColumn"`
}

func (e *RowChangedEvent) String() string {
	js, _ := json.Marshal(e)
	return string(js)
}

func (e *RowChangedEvent) Delete(
	dbTypeT string,
	tableRoute []*rule.TableRouteRule,
	columnRoute []*rule.ColumnRouteRule,
	caseFieldRuleT string) (string, []interface{}, error) {
	var schemaName, tableName string
	schemaName, tableName = e.rewriteSchemaTable(tableRoute)
	switch caseFieldRuleT {
	case constant.ParamValueRuleCaseFieldNameUpper:
		schemaName = strings.ToUpper(schemaName)
		tableName = strings.ToUpper(tableName)
	case constant.ParamValueRuleCaseFieldNameLower:
		schemaName = strings.ToLower(schemaName)
		tableName = strings.ToLower(tableName)
	}

	ms := e.metadata(schemaName, tableName)
	if ms == nil {
		return "", nil, fmt.Errorf("the schema_name_t [%s] table_name_t [%s] metadata not existed, normally should be existed, please contact author or retry", schemaName, tableName)
	}

	columnRouteR := e.rewriteTableColumn(columnRoute)

	uniqColumnData := make(map[string]interface{})

	for c, p := range e.ValidUniqColumns {
		if val, ok := columnRouteR[c]; ok {
			uniqColumnData[val] = p
		} else {
			switch caseFieldRuleT {
			case constant.ParamValueRuleCaseFieldNameUpper:
				uniqColumnData[strings.ToUpper(c)] = p
			case constant.ParamValueRuleCaseFieldNameLower:
				uniqColumnData[strings.ToLower(c)] = p
			default:
				uniqColumnData[c] = p
			}
		}
	}

	upstreamColumnType := make(map[string]string)
	for k, v := range e.ColumnType {
		if val, ok := columnRouteR[k]; ok {
			upstreamColumnType[val] = v
		} else {
			switch caseFieldRuleT {
			case constant.ParamValueRuleCaseFieldNameUpper:
				upstreamColumnType[strings.ToUpper(k)] = v
			case constant.ParamValueRuleCaseFieldNameLower:
				upstreamColumnType[strings.ToLower(k)] = v
			default:
				upstreamColumnType[k] = v
			}
		}
	}

	// all downstream databases except MySQL and TiDB use the unified adaptive field length to automatically fill spaces in DELETE statements of char type to avoid data processing errors due to missing spaces (for example, Oracle database queries or DML operations do not automatically fill spaces), while INSERT statements use the original value of the upstream TiDB char data type to consume synchronously.
	switch strings.ToUpper(dbTypeT) {
	case constant.DatabaseTypeOracle:
		delPrefix := fmt.Sprintf(`DELETE FROM "%s"."%s" WHERE`, schemaName, tableName)
		var (
			cols   []string
			params []interface{}
		)
		for c, p := range uniqColumnData {
			if ds, ok := ms.tableColumns[c]; ok {
				if tys, ok := upstreamColumnType[c]; ok {
					if p != nil {
						switch {
						case stringutil.IsContainedStringIgnoreCase(message.MYSQLCompatibleMsgColumnStringCharacterDatatype, tys) &&
							stringutil.IsContainedStringIgnoreCase(
								constant.OracleCompatibleDatabaseTableColumnBinaryDatatype, ds.columnType):
							cols = append(cols, fmt.Sprintf("%s = :%s", c, c))
							params = append(params, []byte(p.(string)))
							continue
						case stringutil.IsContainedStringIgnoreCase(message.MYSQLCompatibleMsgColumnDatetimeCharacterDatatype, tys):
							if strings.EqualFold(ds.columnType, constant.BuildInOracleDatatypeDate) {
								cols = append(cols, fmt.Sprintf("%s = TO_DATE(:%s,'YYYY-MM-DD HH24:MI:SS')", c, c))
								params = append(params, p)
								continue
							} else if strings.Contains(strings.ToUpper(ds.columnType), constant.BuildInOracleDatatypeTimestamp) {
								cols = append(cols, fmt.Sprintf("%s = TO_TIMESTAMP(:%s,'YYYY-MM-DD HH24:MI:SS.FF6')", c, c))
								params = append(params, p)
								continue
							}
						case stringutil.IsContainedStringIgnoreCase(message.MYSQLCompatibleMsgColumnYearCharacterDatatype, tys):
							if strings.EqualFold(ds.columnType, constant.BuildInOracleDatatypeDate) {
								cols = append(cols, fmt.Sprintf("%s = TO_DATE(:%s,'YYYY')", c, c))
								params = append(params, p)
								continue
							} else if strings.Contains(strings.ToUpper(ds.columnType), constant.BuildInOracleDatatypeTimestamp) {
								cols = append(cols, fmt.Sprintf("%s = TO_TIMESTAMP(:%s,'YYYY')", c, c))
								params = append(params, p)
								continue
							}
						case stringutil.IsContainedStringIgnoreCase(message.MYSQLCompatibleMsgColumnTimeCharacterDatatype, tys):
							if strings.EqualFold(ds.columnType, constant.BuildInOracleDatatypeDate) {
								cols = append(cols, fmt.Sprintf("%s = TO_DATE(:%s,'HH24:MI:SS')", c, c))
								params = append(params, p)
								continue
							} else if strings.Contains(strings.ToUpper(ds.columnType), constant.BuildInOracleDatatypeTimestamp) {
								cols = append(cols, fmt.Sprintf("%s = TO_TIMESTAMP(:%s,'HH24:MI:SS')", c, c))
								params = append(params, p)
								continue
							} else if strings.Contains(strings.ToUpper(ds.columnType), "INTERVAL DAY") {
								sec, err := stringutil.ConvertTimeToSeconds(p.(string))
								if err != nil {
									return "", nil, err
								}
								cols = append(cols, fmt.Sprintf("%s = NUMTODSINTERVAL(:%s, 'SECOND')", c, c))
								params = append(params, sec)
								continue
							}
						case stringutil.IsContainedStringIgnoreCase(message.MYSQLCompatibleMsgColumnCharCharacterDatatype, tys):
							if strings.EqualFold(ds.columnType, constant.BuildInOracleDatatypeChar) ||
								strings.EqualFold(ds.columnType, constant.BuildInOracleDatatypeCharacter) ||
								strings.EqualFold(ds.columnType, constant.BuildInOracleDatatypeNchar) {
								cols = append(cols, fmt.Sprintf("%s = :%s", c, c))
								str := p.(string)
								inter := ds.dataLength - len(str)
								if inter > 0 {
									params = append(params, fmt.Sprintf("%s%s", str, stringutil.PaddingString(inter, " ", " ")))
								} else {
									params = append(params, p)
								}
								continue
							}
						default:
							cols = append(cols, fmt.Sprintf("%s = :%s", c, c))
							params = append(params, p)
							continue
						}
					}
					// nil == NULL values
					cols = append(cols, fmt.Sprintf("%s = :%s", c, c))
					params = append(params, p)
					continue
				}
				cols = append(cols, fmt.Sprintf("%s = :%s", c, c))
				params = append(params, p)
			} else {
				return "", nil, fmt.Errorf("the schema_name_t [%s] table_name_t [%s] column_name [%s] metadata information not existed, normally should be existed, please contact author or retry", schemaName, tableName, c)
			}
		}
		return fmt.Sprintf("%s %s", delPrefix, strings.Join(cols, " AND ")), params, nil
	case constant.DatabaseTypePostgresql:
		delPrefix := fmt.Sprintf(`DELETE FROM "%s"."%s" WHERE`, schemaName, tableName)
		var (
			cols   []string
			params []interface{}
		)
		placeholder := 1
		for c, p := range uniqColumnData {
			if ds, ok := ms.tableColumns[c]; ok {
				if tys, ok := upstreamColumnType[c]; ok {
					if p != nil {
						switch {
						case stringutil.IsContainedStringIgnoreCase(message.MYSQLCompatibleMsgColumnStringCharacterDatatype, tys) && stringutil.IsContainedStringIgnoreCase(
							constant.PostgresDatabaseTableColumnBinaryDatatype, ds.columnType):
							cols = append(cols, fmt.Sprintf("%s = $%d", c, placeholder))
							placeholder++
							params = append(params, []byte(p.(string)))
							continue
						case stringutil.IsContainedStringIgnoreCase(message.MYSQLCompatibleMsgColumnTimeCharacterDatatype, tys):
							if strings.EqualFold(ds.columnType, constant.BuildInPostgresDatatypeTimeWithoutTimeZone) {
								cols = append(cols, fmt.Sprintf("%s = $%d", c, placeholder))
								placeholder++
								params = append(params, p)
								continue
							} else if strings.EqualFold(ds.columnType, constant.BuildInPostgresDatatypeInterval) {
								sec, err := stringutil.ConvertTimeToSeconds(p.(string))
								if err != nil {
									return "", nil, err
								}
								cols = append(cols, fmt.Sprintf("%s = MAKE_INTERVAL(secs => $%d)", c, placeholder))
								placeholder++
								params = append(params, sec)
								continue
							}
						case stringutil.IsContainedStringIgnoreCase(message.MYSQLCompatibleMsgColumnCharCharacterDatatype, tys):
							if strings.EqualFold(ds.columnType, constant.BuildInPostgresDatatypeCharacter) {
								cols = append(cols, fmt.Sprintf("%s = :%s", c, c))
								str := p.(string)
								inter := ds.dataLength - len(str)
								if inter > 0 {
									params = append(params, fmt.Sprintf("%s%s", str, stringutil.PaddingString(inter, " ", " ")))
								} else {
									params = append(params, p)
								}
								continue
							}
						default:
							cols = append(cols, fmt.Sprintf("%s = $%d", c, placeholder))
							placeholder++
							params = append(params, p)
							continue
						}
					}
					// nil == NULL values
					cols = append(cols, fmt.Sprintf("%s = $%d", c, placeholder))
					placeholder++
					params = append(params, p)
					continue
				}
				cols = append(cols, fmt.Sprintf("%s = $%d", c, placeholder))
				placeholder++
				params = append(params, p)
			} else {
				return "", nil, fmt.Errorf("the schema_name_t [%s] table_name_t [%s] column_name [%s] metadata information not existed, normally should be existed, please contact author or retry", schemaName, tableName, c)
			}
		}
		return fmt.Sprintf("%s %s", delPrefix, strings.Join(cols, " AND ")), params, nil
	case constant.DatabaseTypeMySQL, constant.DatabaseTypeTiDB:
		delPrefix := fmt.Sprintf("DELETE FROM `%s`.`%s` WHERE", schemaName, tableName)
		var (
			cols   []string
			params []interface{}
		)
		for c, p := range uniqColumnData {
			cols = append(cols, fmt.Sprintf("%s = ?", c))
			params = append(params, p)
		}
		return fmt.Sprintf("%s %s", delPrefix, strings.Join(cols, " AND ")), params, nil
	default:
		panic(fmt.Errorf("delete statament not support database type [%s]", dbTypeT))
	}

}

func (e *RowChangedEvent) Insert(dbTypeT string,
	tableRoute []*rule.TableRouteRule,
	columnRoute []*rule.ColumnRouteRule,
	caseFieldRuleT string) (string, []interface{}, error) {
	var (
		schemaName, tableName string
		// validUniqColumns      []string
	)
	schemaName, tableName = e.rewriteSchemaTable(tableRoute)
	switch caseFieldRuleT {
	case constant.ParamValueRuleCaseFieldNameUpper:
		schemaName = strings.ToUpper(schemaName)
		tableName = strings.ToUpper(tableName)
	case constant.ParamValueRuleCaseFieldNameLower:
		schemaName = strings.ToLower(schemaName)
		tableName = strings.ToLower(tableName)
	}
	columnRouteR := e.rewriteTableColumn(columnRoute)

	ms := e.metadata(schemaName, tableName)
	if ms == nil {
		return "", nil, fmt.Errorf("the schema_name_t [%s] table_name_t [%s] metadata not existed, normally should be existed, please contact author or retry", schemaName, tableName)
	}

	newColumnData := make(map[string]interface{})
	for c, p := range e.NewColumnData {
		var colName string
		if val, ok := columnRouteR[c]; ok {
			colName = val
		} else {
			switch caseFieldRuleT {
			case constant.ParamValueRuleCaseFieldNameUpper:
				colName = strings.ToUpper(c)
			case constant.ParamValueRuleCaseFieldNameLower:
				colName = strings.ToLower(c)
			default:
				colName = c
			}
		}
		if meta, existd := ms.tableColumns[colName]; existd {
			if e.isGeneratedColumn(c) && meta.isGeneraed {
				continue
			} else {
				newColumnData[colName] = p
			}
		} else {
			return "", nil, fmt.Errorf("the schema_name_t [%s] table_name_t [%s] column_name_t [%s] metadata not existed, normally should be existed, please contact author or retry", schemaName, tableName, colName)
		}
	}

	upstreamColumnType := make(map[string]string)
	for c, v := range e.ColumnType {
		var colName string
		if val, ok := columnRouteR[c]; ok {
			colName = val
		} else {
			switch caseFieldRuleT {
			case constant.ParamValueRuleCaseFieldNameUpper:
				colName = strings.ToUpper(c)
			case constant.ParamValueRuleCaseFieldNameLower:
				colName = strings.ToLower(c)
			default:
				colName = c
			}
		}
		if meta, existd := ms.tableColumns[colName]; existd {
			if e.isGeneratedColumn(c) && meta.isGeneraed {
				continue
			} else {
				upstreamColumnType[colName] = v
			}
		} else {
			return "", nil, fmt.Errorf("the schema_name_t [%s] table_name_t [%s] column_name_t [%s] metadata not existed, normally should be existed, please contact author or retry", schemaName, tableName, colName)
		}
	}

	// for c, _ := range e.ValidUniqColumns {
	// 	switch caseFieldRuleT {
	// 	case constant.ParamValueRuleCaseFieldNameUpper:
	// 		validUniqColumns = append(validUniqColumns, strings.ToUpper(c))
	// 	case constant.ParamValueRuleCaseFieldNameLower:
	// 		validUniqColumns = append(validUniqColumns, strings.ToLower(c))
	// 	default:
	// 		validUniqColumns = append(validUniqColumns, c)
	// 	}
	// }

	switch strings.ToUpper(dbTypeT) {
	case constant.DatabaseTypeOracle:
		insPrefix := fmt.Sprintf(`INSERT INTO "%s"."%s" `, schemaName, tableName)
		var (
			cols   []string
			pals   []string
			params []interface{}
		)
		for c, p := range newColumnData {
			cols = append(cols, fmt.Sprintf(`"%s"`, c))
			if ds, ok := ms.tableColumns[c]; ok {
				if tys, ok := upstreamColumnType[c]; ok {
					if p != nil {
						switch {
						case stringutil.IsContainedStringIgnoreCase(message.MYSQLCompatibleMsgColumnStringCharacterDatatype, tys) &&
							stringutil.IsContainedStringIgnoreCase(
								constant.OracleCompatibleDatabaseTableColumnBinaryDatatype, ds.columnType):
							pals = append(pals, fmt.Sprintf(":%s", c))
							params = append(params, []byte(p.(string)))
							continue
						case stringutil.IsContainedStringIgnoreCase(message.MYSQLCompatibleMsgColumnDatetimeCharacterDatatype, tys):
							if strings.EqualFold(ds.columnType, constant.BuildInOracleDatatypeDate) {
								pals = append(pals, fmt.Sprintf("TO_DATE(:%s,'YYYY-MM-DD HH24:MI:SS')", c))
								params = append(params, p)
								continue
							} else if strings.Contains(strings.ToUpper(ds.columnType), constant.BuildInOracleDatatypeTimestamp) {
								pals = append(pals, fmt.Sprintf("TO_TIMESTAMP(:%s,'YYYY-MM-DD HH24:MI:SS.FF6')", c))
								params = append(params, p)
								continue
							}
						case stringutil.IsContainedStringIgnoreCase(message.MYSQLCompatibleMsgColumnYearCharacterDatatype, tys):
							if strings.EqualFold(ds.columnType, constant.BuildInOracleDatatypeDate) {
								pals = append(pals, fmt.Sprintf("TO_DATE(:%s,'YYYY')", c))
								params = append(params, p)
								continue
							} else if strings.Contains(strings.ToUpper(ds.columnType), constant.BuildInOracleDatatypeTimestamp) {
								pals = append(pals, fmt.Sprintf("TO_TIMESTAMP(:%s,'YYYY')", c))
								params = append(params, p)
								continue
							}
						case stringutil.IsContainedStringIgnoreCase(message.MYSQLCompatibleMsgColumnTimeCharacterDatatype, tys):
							if strings.EqualFold(ds.columnType, constant.BuildInOracleDatatypeDate) {
								pals = append(pals, fmt.Sprintf("TO_DATE(:%s,'HH24:MI:SS')", c))
								params = append(params, p)
								continue
							} else if strings.Contains(strings.ToUpper(ds.columnType), constant.BuildInOracleDatatypeTimestamp) {
								pals = append(pals, fmt.Sprintf("TO_TIMESTAMP(:%s,'HH24:MI:SS')", c))
								params = append(params, p)
								continue
							} else if strings.Contains(strings.ToUpper(ds.columnType), "INTERVAL DAY") {
								sec, err := stringutil.ConvertTimeToSeconds(p.(string))
								if err != nil {
									return "", nil, err
								}
								pals = append(pals, fmt.Sprintf("NUMTODSINTERVAL(:%s, 'SECOND')", c))
								params = append(params, sec)
								continue
							}
						default:
							pals = append(pals, fmt.Sprintf(":%s", c))
							params = append(params, p)
							continue
						}
					}
					// nil == NULL values
					pals = append(pals, fmt.Sprintf(":%s", c))
					params = append(params, p)
					continue
				}
				pals = append(pals, fmt.Sprintf(":%s", c))
				params = append(params, p)
			} else {
				return "", nil, fmt.Errorf("the schema_name_t [%s] table_name_t [%s] column_name [%s] metadata information not existed, normally should be existed, please contact author or retry", schemaName, tableName, c)
			}
		}
		return fmt.Sprintf("%s (%s) VALUES (%s)", insPrefix,
			strings.Join(cols, constant.StringSeparatorComma), strings.Join(pals, constant.StringSeparatorComma)), params, nil
		// MERGE INTO BLOB / LONG RAW / RAW datatype data error ORA-00927: missing equal sign , changed delete + insert
		// insPrefix := fmt.Sprintf(`MERGE INTO "%s"."%s" t`, schemaName, tableName)
		// var (
		// 	dualS         []string
		// 	onS           []string
		// 	inCols        []string
		// 	valS          []string
		// 	setS          []string
		// 	params        []interface{}
		// 	originColumns []string
		// )
		// for c, p := range newColumnData {
		// 	dualS = append(dualS, fmt.Sprintf(`:%s AS "%s"`, c, c))
		// 	inCols = append(inCols, fmt.Sprintf(`"%s"`, c))
		// 	valS = append(valS, fmt.Sprintf(`s."%s"`, c))

		// 	if ds.columnType, ok := ms.tableColumnsT[c]; ok {
		// 		if tys, ok := upstreamColumnType[c]; ok {
		// 			if stringutil.IsContainedStringIgnoreCase(MessageEventColumnStringCharacterDatatype, tys) &&
		// 				stringutil.IsContainedStringIgnoreCase(constant.OracleCompatibleDatabaseTableColumnBinaryDatatype, ds) {
		// 				// nil == NULL values
		// 				if p != nil {
		// 					params = append(params, []byte(p.(string)))
		// 				} else {
		// 					params = append(params, p)
		// 				}
		// 				continue
		// 			}
		// 		}
		// 		params = append(params, p)
		// 	} else {
		// 		return "", nil, fmt.Errorf("the schema_name_t [%s] table_name_t [%s] column_name [%s] metadata information not existed, normally should be existed, please contact author or retry", schemaName, tableName, c)
		// 	}

		// 	originColumns = append(originColumns, c)
		// }

		// for _, p := range validUniqColumns {
		// 	onS = append(onS, fmt.Sprintf(`t."%s" = s."%s"`, p, p))
		// }

		// for _, s := range difference(originColumns, validUniqColumns) {
		// 	setS = append(setS, fmt.Sprintf(`t."%s" = s."%s"`, s, s))
		// }

		// return fmt.Sprintf("%s USING (SELECT %s FROM DUAL) s ON (%s) WHEN MATCHED THEN UPDATE SET %s WHEN NOT MATCHED THEN INSERT (%s) VALUES (%s)",
		// 	insPrefix,
		// 	strings.Join(dualS, constant.StringSeparatorComma),
		// 	strings.Join(onS, constant.StringSeparatorComma),
		// 	strings.Join(setS, constant.StringSeparatorComma),
		// 	strings.Join(inCols, constant.StringSeparatorComma),
		// 	strings.Join(valS, constant.StringSeparatorComma)), params, nil
	case constant.DatabaseTypePostgresql:
		insPrefix := fmt.Sprintf(`INSERT INTO "%s"."%s"`, schemaName, tableName)
		var (
			cols   []string
			pals   []string
			params []interface{}
		)
		placeholder := 1
		for c, p := range newColumnData {
			cols = append(cols, fmt.Sprintf(`"%s"`, c))
			pals = append(pals, fmt.Sprintf("$%d", placeholder))

			placeholder++

			if ds, ok := ms.tableColumns[c]; ok {
				if tys, ok := upstreamColumnType[c]; ok {
					if p != nil {
						switch {
						case stringutil.IsContainedStringIgnoreCase(message.MYSQLCompatibleMsgColumnStringCharacterDatatype, tys) && stringutil.IsContainedStringIgnoreCase(
							constant.PostgresDatabaseTableColumnBinaryDatatype, ds.columnType):
							params = append(params, []byte(p.(string)))
							continue
						case stringutil.IsContainedStringIgnoreCase(message.MYSQLCompatibleMsgColumnTimeCharacterDatatype, tys):
							if strings.EqualFold(ds.columnType, constant.BuildInPostgresDatatypeTimeWithoutTimeZone) {
								params = append(params, p)
								continue
							} else if strings.EqualFold(ds.columnType, constant.BuildInPostgresDatatypeInterval) {
								sec, err := stringutil.ConvertTimeToSeconds(p.(string))
								if err != nil {
									return "", nil, err
								}
								params = append(params, fmt.Sprintf("MAKE_INTERVAL(secs => %d)", sec))
								continue
							}
						default:
							params = append(params, p)
							continue
						}
					}
					// nil == NULL values
					params = append(params, p)
					continue
				}
				params = append(params, p)
			} else {
				return "", nil, fmt.Errorf("the schema_name_t [%s] table_name_t [%s] column_name [%s] metadata information not existed, normally should be existed, please contact author or retry", schemaName, tableName, c)
			}
		}
		return fmt.Sprintf("%s (%s) VALUES (%s)", insPrefix,
			strings.Join(cols, constant.StringSeparatorComma), strings.Join(pals, constant.StringSeparatorComma)), params, nil

		// insPrefix := fmt.Sprintf(`INSERT INTO "%s"."%s"`, schemaName, tableName)
		// var (
		// 	cols   []string
		// 	pals   []string
		// 	params []interface{}
		// 	onS    []string
		// 	setS   []string
		// )
		// placeholder := 1
		// for c, p := range newColumnData {
		// 	cols = append(cols, fmt.Sprintf(`"%s"`, c))
		// 	pals = append(pals, fmt.Sprintf("$%d", placeholder))

		// 	if ds.columnType, ok := ms.tableColumnsT[c]; ok {
		// 		if tys, ok := upstreamColumnType[c]; ok {
		// 			if stringutil.IsContainedStringIgnoreCase(MessageEventColumnStringCharacterDatatype, tys) &&
		// 				stringutil.IsContainedStringIgnoreCase(constant.PostgresDatabaseTableColumnBinaryDatatype, ds) {
		// 				// nil == NULL values
		// 				if p != nil {
		// 					params = append(params, []byte(p.(string)))
		// 				} else {
		// 					params = append(params, p)
		// 				}
		// 				continue
		// 			}
		// 		}
		// 		params = append(params, p)
		// 	} else {
		// 		return "", nil, fmt.Errorf("the schema_name_t [%s] table_name_t [%s] column_name [%s] metadata information not existed, normally should be existed, please contact author or retry", schemaName, tableName, c)
		// 	}

		// 	setS = append(setS, fmt.Sprintf(`"%s" =  EXCLUDED."%s"`, c, c))
		// 	placeholder++
		// }
		// for _, p := range validUniqColumns {
		// 	onS = append(onS, fmt.Sprintf(`"%s"`, p))
		// }
		// return fmt.Sprintf("%s (%s) VALUES (%s) ON CONFLICT (%s) DO UPDATE SET %s",
		// 	insPrefix,
		// 	strings.Join(cols, constant.StringSeparatorComma),
		// 	strings.Join(pals, constant.StringSeparatorComma),
		// 	strings.Join(onS, constant.StringSeparatorComma),
		// 	strings.Join(setS, constant.StringSeparatorComma),
		// ), params, nil
	case constant.DatabaseTypeMySQL, constant.DatabaseTypeTiDB:
		// insPrefix := fmt.Sprintf("REPLACE INTO `%s`.`%s` ", schemaName, tableName)
		insPrefix := fmt.Sprintf("INSERT INTO `%s`.`%s` ", schemaName, tableName)
		var (
			cols   []string
			pals   []string
			params []interface{}
		)
		for c, p := range newColumnData {
			cols = append(cols, fmt.Sprintf("`%s`", c))
			pals = append(pals, "?")
			params = append(params, p)
		}
		return fmt.Sprintf("%s (%s) VALUES (%s)", insPrefix,
			strings.Join(cols, constant.StringSeparatorComma), strings.Join(pals, constant.StringSeparatorComma)), params, nil
	default:
		panic(fmt.Errorf("replace statament not support database type [%s]", dbTypeT))
	}
}

func (e *RowChangedEvent) rewriteSchemaTable(tableRoute []*rule.TableRouteRule) (string, string) {
	for _, t := range tableRoute {
		if e.SchemaName == t.SchemaNameS && e.TableName == t.TableNameS {
			return t.SchemaNameT, t.TableNameT
		}
	}
	return e.SchemaName, e.TableName
}

func (e *RowChangedEvent) rewriteTableColumn(columnRoute []*rule.ColumnRouteRule) map[string]string {
	columnR := make(map[string]string)
	for _, t := range columnRoute {
		if e.SchemaName == t.SchemaNameS && e.TableName == t.TableNameS {
			columnR[t.ColumnNameS] = t.ColumnNameT
		}
	}
	return columnR
}

func (e *RowChangedEvent) metadata(schemaName, tableName string) *metadata {
	metadata, existed := metaCache.Get(schemaName, tableName)
	if existed {
		return metadata
	}
	return nil
}

func (e *RowChangedEvent) isGeneratedColumn(columnName string) bool {
	for _, c := range e.Columns {
		if columnName == c.ColumnName && c.IsGeneratedColumn {
			return true
		}
	}
	return false
}

type DDLChangedEvent struct {
	CommitTs   uint64  `json:"commitTs"`
	SchemaName string  `json:"schemaName"`
	TableName  string  `json:"tableName"`
	DdlQuery   string  `json:"ddlQuery"`
	DdlType    DDLType `json:"ddlType"`
}

func (d *DDLChangedEvent) String() string {
	jsb, _ := json.Marshal(d)
	return string(jsb)
}

// DDLChangedEvents is a slice of DDLChangedEvent and implements the sort.Interface interface.
type DDLChangedEvents []*DDLChangedEvent

func (d DDLChangedEvents) Len() int           { return len(d) }
func (d DDLChangedEvents) Less(i, j int) bool { return d[i].CommitTs < d[j].CommitTs }
func (d DDLChangedEvents) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }

// Add a new DDLChangedEvent and keep the ddls sorted by CommitTs
func (d *DDLChangedEvents) Add(event *DDLChangedEvent) {
	*d = append(*d, event)
	sort.Sort(*d)
}

// EventGroup could store change ddl and dml event message.
type EventGroup struct {
	events []*RowChangedEvent
}

// NewEventsGroup will create new event group.
func NewEventGroup() *EventGroup {
	return &EventGroup{
		events: make([]*RowChangedEvent, 0),
	}
}

// Append will append an event to event groups.
func (g *EventGroup) Append(e *RowChangedEvent) {
	g.events = append(g.events, e)
}

// OrderSortedCommitTs extract and sort all Commits
func (g *EventGroup) OrderSortedCommitTs() []uint64 {
	var commitTsSli []uint64

	for _, event := range g.events {
		if event != nil {
			commitTsSli = append(commitTsSli, event.CommitTs)
		}
	}

	sort.Slice(commitTsSli, func(i, j int) bool {
		return commitTsSli[i] < commitTsSli[j]
	})
	return commitTsSli
}

// ResolvedTs will get events whose CommitTs is less than or equal to resolveTs,
// and at the same time remove events whose CommitTs is less than or equal to resolveTs from the original queue
func (g *EventGroup) ResolvedTs(resolveTs uint64) []*RowChangedEvent {
	sort.Slice(g.events, func(i, j int) bool {
		return g.events[i].CommitTs < g.events[j].CommitTs
	})

	i := sort.Search(len(g.events), func(i int) bool {
		return g.events[i].CommitTs > resolveTs
	})

	result := g.events[:i]
	g.events = g.events[i:]
	return result
}

// DDLCommitTs returns all events strictly < ddlCommitTs
func (g *EventGroup) DDLCommitTs(ddlCommitTs uint64) []*RowChangedEvent {
	// Sort the events by CommitTs first
	sort.Slice(g.events, func(i, j int) bool {
		return g.events[i].CommitTs < g.events[j].CommitTs
	})

	// Find the first CommitTs >= resolveTs
	i := sort.Search(len(g.events), func(i int) bool {
		return g.events[i].CommitTs >= ddlCommitTs
	})

	// Return all events where CommitTs < ddlCommitTs and update g.events
	result := g.events[:i]
	g.events = g.events[i:]
	return result
}

// RemoveDDLCommitTs remove the ddl event equal to ddlCommit ts
func (g *EventGroup) RemoveDDLCommitTs(ddlCommitTs uint64) {
	var retainedEvents []*RowChangedEvent

	// Iterate through the events and separate them into removed and retained slices
	for _, event := range g.events {
		if event.CommitTs == ddlCommitTs && event.IsDDL {
			continue
		} else {
			retainedEvents = append(retainedEvents, event)
		}
	}

	// Update g.events to contain only retained events
	g.events = retainedEvents
}
