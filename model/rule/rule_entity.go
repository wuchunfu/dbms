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
package rule

import "github.com/wentaojin/dbms/model/common"

type MigrateTaskRule struct {
	ID              uint64 `gorm:"primarykey;autoIncrement;comment:id" json:"id"`
	TaskRuleName    string `gorm:"type:varchar(300);not null;uniqueIndex:uniq_datasource_rule_name;comment:migrate task datasource name" json:"taskRuleName"`
	DatasourceNameS string `gorm:"type:varchar(300);not null;comment:source datasource of task" json:"datasourceNameS"`
	DatasourceNameT string `gorm:"type:varchar(300);not null;comment:target datasource of task" json:"datasourceNameT"`
	*common.Entity
}

type SchemaRouteRule struct {
	ID           uint64 `gorm:"primary_key;autoIncrement;comment:id" json:"id"`
	TaskRuleName string `gorm:"type:varchar(300);not null;uniqueIndex:uniq_schema_table_name;comment:migrate task datasource name" json:"taskRuleName"`
	SchemaNameS  string `gorm:"type:varchar(120);not null;uniqueIndex:uniq_schema_table_name;comment:source schema name" json:"schemaNameS"`
	SchemaNameT  string `gorm:"type:varchar(120);not null;comment:target schema name" json:"schemaNameT"`
	//IncludeTable string `gorm:"comment:source schema include table list" json:"includeTable"` // used for task record
	//ExcludeTable string `gorm:"comment:source schema exclude table list" json:"excludeTable"` // used for task record
	*common.Entity
}

// MigrateTaskTable used for configure includeTable and excludeTable
type MigrateTaskTable struct {
	ID           uint64 `gorm:"primary_key;autoIncrement;comment:id" json:"id"`
	TaskRuleName string `gorm:"type:varchar(300);not null;uniqueIndex:uniq_schema_table_name;comment:migrate task datasource name" json:"taskRuleName"`
	SchemaNameS  string `gorm:"type:varchar(120);not null;uniqueIndex:uniq_schema_table_name;comment:source schema name" json:"schemaNameS"`
	TableNameS   string `gorm:"type:varchar(120);not null;uniqueIndex:uniq_schema_table_name;comment:source schema table list" json:"tableNameS"`
	IsExclude    string `gorm:"type:varchar(5);not null;comment:source schema table is whether exclude" json:"isExclude"`
}

type TableRouteRule struct {
	ID           uint64 `gorm:"primary_key;autoIncrement;comment:id" json:"id"`
	TaskRuleName string `gorm:"type:varchar(300);not null;uniqueIndex:uniq_schema_table_name;comment:migrate task datasource name" json:"taskRuleName"`
	SchemaNameS  string `gorm:"type:varchar(120);not null;uniqueIndex:uniq_schema_table_name;comment:source schema name" json:"schemaNameS"`
	TableNameS   string `gorm:"type:varchar(120);not null;uniqueIndex:uniq_schema_table_name;comment:source table name" json:"tableNameS"`
	SchemaNameT  string `gorm:"type:varchar(120);not null;comment:target schema name" json:"schemaNameT"`
	TableNameT   string `gorm:"type:varchar(120);not null;comment:target table name" json:"tableNameT"`
	*common.Entity
}

type ColumnRouteRule struct {
	ID           uint64 `gorm:"primary_key;autoIncrement;comment:id" json:"id"`
	TaskRuleName string `gorm:"type:varchar(300);not null;uniqueIndex:uniq_schema_table_name;comment:migrate task datasource name" json:"taskRuleName"`
	SchemaNameS  string `gorm:"type:varchar(120);not null;uniqueIndex:uniq_schema_table_name;comment:source schema name" json:"schemaNameS"`
	TableNameS   string `gorm:"type:varchar(120);not null;uniqueIndex:uniq_schema_table_name;comment:source table name" json:"tableNameS"`
	ColumnNameS  string `gorm:"type:varchar(120);not null;uniqueIndex:uniq_schema_table_name;comment:source table name" json:"columnNameS"`
	SchemaNameT  string `gorm:"type:varchar(120);not null;comment:target schema name" json:"schemaNameT"`
	TableNameT   string `gorm:"type:varchar(120);not null;comment:target table name" json:"tableNameT"`
	ColumnNameT  string `gorm:"type:varchar(120);not null;comment:source table name" json:"columnNameT"`
	*common.Entity
}
