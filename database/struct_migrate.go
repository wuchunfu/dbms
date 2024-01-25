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
package database

// ITableAttributesReader used for database table attributes
type ITableAttributesReader interface {
	GetTablePrimaryKey() ([]map[string]string, error)
	GetTableUniqueKey() ([]map[string]string, error)
	GetTableForeignKey() ([]map[string]string, error)
	GetTableCheckKey() ([]map[string]string, error)
	GetTableUniqueIndex() ([]map[string]string, error)
	GetTableNormalIndex() ([]map[string]string, error)
	GetTableComment() ([]map[string]string, error)
	GetTableColumns() ([]map[string]string, error)
	GetTableColumnComment() ([]map[string]string, error)

	GetTableOriginStruct() (string, error)
}

type TableAttributes struct {
	PrimaryKey    []map[string]string `json:"primary_key"`
	UniqueKey     []map[string]string `json:"unique_key"`
	ForeignKey    []map[string]string `json:"foreign_key"`
	CheckKey      []map[string]string `json:"check_key"`
	UniqueIndex   []map[string]string `json:"unique_index"`
	NormalIndex   []map[string]string `json:"normal_index"`
	TableComment  []map[string]string `json:"table_comment"`
	TableColumns  []map[string]string `json:"table_columns"`
	ColumnComment []map[string]string `json:"column_comment"`
	OriginStruct  string              `json:"origin_struct"`
}

func IDatabaseTableAttributes(t ITableAttributesReader) (*TableAttributes, error) {
	primaryKey, err := t.GetTablePrimaryKey()
	if err != nil {
		return &TableAttributes{}, err
	}
	uniqueKey, err := t.GetTableUniqueKey()
	if err != nil {
		return &TableAttributes{}, err
	}
	foreignKey, err := t.GetTableForeignKey()
	if err != nil {
		return &TableAttributes{}, err
	}
	checkKey, err := t.GetTableCheckKey()
	if err != nil {
		return &TableAttributes{}, err
	}
	uniqueIndex, err := t.GetTableUniqueIndex()
	if err != nil {
		return &TableAttributes{}, err
	}
	normalIndex, err := t.GetTableNormalIndex()
	if err != nil {
		return &TableAttributes{}, err
	}
	tableComment, err := t.GetTableComment()
	if err != nil {
		return &TableAttributes{}, err
	}
	tableColumns, err := t.GetTableColumns()
	if err != nil {
		return &TableAttributes{}, err
	}
	columnComment, err := t.GetTableColumnComment()
	if err != nil {
		return &TableAttributes{}, err
	}

	return &TableAttributes{
		PrimaryKey:    primaryKey,
		UniqueKey:     uniqueKey,
		ForeignKey:    foreignKey,
		CheckKey:      checkKey,
		UniqueIndex:   uniqueIndex,
		NormalIndex:   normalIndex,
		TableComment:  tableComment,
		TableColumns:  tableColumns,
		ColumnComment: columnComment,
	}, nil
}

type ITableAttributesRuleReader interface {
	GetCreatePrefixRule() string
	GetCaseFieldRule() string
	GetSchemaNameRule() (map[string]string, error)
	GetTableNameRule() (map[string]string, error)
	GetTableColumnRule() (map[string]string, map[string]string, map[string]string, error)
	GetTableAttributesRule() (string, error)
	GetTableCommentRule() (string, error)
	GetTableColumnCollationRule() (map[string]string, error)
	GetTableColumnCommentRule() (map[string]string, error)
}

type TableAttributesRule struct {
	CreatePrefixRule       string            `json:"createPrefixRule"`
	CaseFieldRule          string            `json:"caseFieldRule"`
	SchemaNameRule         map[string]string `json:"schemaNameRule"`
	TableNameRule          map[string]string `json:"tableNameRule"`
	ColumnNameRule         map[string]string `json:"columnNameRule"`
	ColumnDatatypeRule     map[string]string `json:"columnDatatypeRule"`
	ColumnDefaultValueRule map[string]string `json:"columnDefaultValueRule"`
	ColumnCollationRule    map[string]string `json:"columnCollationRule"`
	ColumnCommentRule      map[string]string `json:"columnCommentRule"`
	TableAttrRule          string            `json:"tableAttrRule"`
	TableCommentRule       string            `json:"tableCommentRule"`
}

func IDatabaseTableAttributesRule(t ITableAttributesRuleReader) (*TableAttributesRule, error) {
	schemaNameRule, err := t.GetSchemaNameRule()
	if err != nil {
		return &TableAttributesRule{}, err
	}
	tableNameRule, err := t.GetTableNameRule()
	if err != nil {
		return &TableAttributesRule{}, err
	}
	columnNameRule, datatypeRule, defaultValRule, err := t.GetTableColumnRule()
	if err != nil {
		return &TableAttributesRule{}, err
	}
	rule, err := t.GetTableCommentRule()
	if err != nil {
		return &TableAttributesRule{}, err
	}
	attr, err := t.GetTableAttributesRule()
	if err != nil {
		return &TableAttributesRule{}, err
	}
	collationRule, err := t.GetTableColumnCollationRule()
	if err != nil {
		return &TableAttributesRule{}, err
	}
	commentRule, err := t.GetTableColumnCommentRule()
	if err != nil {
		return &TableAttributesRule{}, err
	}
	return &TableAttributesRule{
			CreatePrefixRule:       t.GetCreatePrefixRule(),
			CaseFieldRule:          t.GetCaseFieldRule(),
			SchemaNameRule:         schemaNameRule,
			TableNameRule:          tableNameRule,
			ColumnNameRule:         columnNameRule,
			ColumnDatatypeRule:     datatypeRule,
			ColumnDefaultValueRule: defaultValRule,
			ColumnCollationRule:    collationRule,
			ColumnCommentRule:      commentRule,
			TableAttrRule:          attr,
			TableCommentRule:       rule,
		},
		nil
}

type ITableAttributesProcessor interface {
	GenSchemaNameS() string
	GenTableNameS() string
	GenTableTypeS() string
	GenTableOriginDDlS() string
	GenSchemaNameT() (string, error)
	GenTableNameT() (string, error)
	GenTableCreatePrefixT() string
	GenTableSuffix() (string, error)
	GenTablePrimaryKey() (string, error)
	GenTableUniqueKey() ([]string, error)
	GenTableForeignKey() ([]string, error)
	GenTableCheckKey() ([]string, error)
	GenTableUniqueIndex() ([]string, []string, error)
	GenTableNormalIndex() ([]string, []string, error)
	GenTableComment() (string, error)
	GenTableColumns() ([]string, error)
	GenTableColumnComment() ([]string, error)
}

type TableStruct struct {
	SchemaNameS          string   `json:"schemaNameS"`
	TableNameS           string   `json:"tableNameS"`
	TableTypeS           string   `json:"tableTypeS"`
	OriginDdlS           string   `json:"originDdlS"`
	SchemaNameT          string   `json:"schemaNameT"`
	TableNameT           string   `json:"tableNameT"`
	TableCreatePrefixT   string   `json:"tableCreatePrefixT"`
	TablePrimaryKey      string   `json:"tablePrimaryKey"`
	TableColumns         []string `json:"tableColumns"`
	TableUniqueKeys      []string `json:"tableUniqueKeys"`
	TableNormalIndexes   []string `json:"tableNormalIndexes"`
	TableUniqueIndexes   []string `json:"tableUniqueIndexes"`
	TableSuffix          string   `json:"tableSuffix"`
	TableComment         string   `json:"tableComment"`
	TableColumnComment   []string `json:"tableColumnComment"`
	TableCheckKeys       []string `json:"tableCheckKeys"`
	TableForeignKeys     []string `json:"tableForeignKeys"`
	TableIncompatibleDDL []string `json:"tableIncompatibleDDL"`
}

func IDatabaseTableStruct(p ITableAttributesProcessor) (*TableStruct, error) {
	var incompatibleSqls []string
	schemaNameT, err := p.GenSchemaNameT()
	if err != nil {
		return &TableStruct{}, err
	}
	tableNameT, err := p.GenTableNameT()
	if err != nil {
		return &TableStruct{}, err
	}
	columns, err := p.GenTableColumns()
	if err != nil {
		return &TableStruct{}, err
	}
	primaryKey, err := p.GenTablePrimaryKey()
	if err != nil {
		return &TableStruct{}, err
	}
	uniqueKeys, err := p.GenTableUniqueKey()
	if err != nil {
		return &TableStruct{}, err
	}
	normalIndexes, normalIndexCompSQL, err := p.GenTableNormalIndex()
	if err != nil {
		return &TableStruct{}, err
	}
	uniqueIndexes, uniqueIndexCompSQL, err := p.GenTableUniqueIndex()
	if err != nil {
		return &TableStruct{}, err
	}
	tableSuffix, err := p.GenTableSuffix()
	if err != nil {
		return &TableStruct{}, err
	}
	tableComment, err := p.GenTableComment()
	if err != nil {
		return &TableStruct{}, err
	}
	checkKeys, err := p.GenTableCheckKey()
	if err != nil {
		return &TableStruct{}, err
	}
	foreignKeys, err := p.GenTableForeignKey()
	if err != nil {
		return &TableStruct{}, err
	}
	columnComments, err := p.GenTableColumnComment()
	if err != nil {
		return &TableStruct{}, err
	}

	if len(normalIndexCompSQL) > 0 {
		incompatibleSqls = append(incompatibleSqls, normalIndexCompSQL...)
	}
	if len(uniqueIndexCompSQL) > 0 {
		incompatibleSqls = append(incompatibleSqls, uniqueIndexCompSQL...)
	}

	return &TableStruct{
		SchemaNameS:          p.GenSchemaNameS(),
		TableNameS:           p.GenTableNameS(),
		TableTypeS:           p.GenTableTypeS(),
		OriginDdlS:           p.GenTableOriginDDlS(),
		SchemaNameT:          schemaNameT,
		TableNameT:           tableNameT,
		TableCreatePrefixT:   p.GenTableCreatePrefixT(),
		TablePrimaryKey:      primaryKey,
		TableColumns:         columns,
		TableUniqueKeys:      uniqueKeys,
		TableNormalIndexes:   normalIndexes,
		TableUniqueIndexes:   uniqueIndexes,
		TableSuffix:          tableSuffix,
		TableComment:         tableComment,
		TableColumnComment:   columnComments,
		TableCheckKeys:       checkKeys,
		TableForeignKeys:     foreignKeys,
		TableIncompatibleDDL: incompatibleSqls,
	}, nil
}

type ITableStructDatabaseWriter interface {
	WriteStructDatabase() error
	SyncStructDatabase() error
}

type ITableStructFileWriter interface {
	InitOutputFile() error
	SyncStructFile() error
}
