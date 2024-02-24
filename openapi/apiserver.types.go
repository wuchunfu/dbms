// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.15.0 DO NOT EDIT.
package openapi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/url"
	"path"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// CaseFieldRule defines model for CaseFieldRule.
type CaseFieldRule struct {
	CaseFieldRuleS *string `json:"caseFieldRuleS,omitempty"`
	CaseFieldRuleT *string `json:"caseFieldRuleT,omitempty"`
}

// ColumnStructRule defines model for ColumnStructRule.
type ColumnStructRule struct {
	ColumnNameS   *string `json:"columnNameS,omitempty"`
	ColumnTypeS   *string `json:"columnTypeS,omitempty"`
	ColumnTypeT   *string `json:"columnTypeT,omitempty"`
	DefaultValueS *string `json:"defaultValueS,omitempty"`
	DefaultValueT *string `json:"defaultValueT,omitempty"`
	SchemaNameS   *string `json:"schemaNameS,omitempty"`
	TableNameS    *string `json:"tableNameS,omitempty"`
}

// DataMigrateRule defines model for DataMigrateRule.
type DataMigrateRule struct {
	EnableChunkStrategy *bool   `json:"enableChunkStrategy,omitempty"`
	SqlHintS            *string `json:"sqlHintS,omitempty"`
	TableNameS          *string `json:"tableNameS,omitempty"`
	WhereRange          *string `json:"whereRange,omitempty"`
}

// DataMigrateTask defines model for DataMigrateTask.
type DataMigrateTask struct {
	CaseFieldRule         *CaseFieldRule         `json:"caseFieldRule,omitempty"`
	Comment               *string                `json:"comment,omitempty"`
	DatasourceNameS       *string                `json:"datasourceNameS,omitempty"`
	DatasourceNameT       *string                `json:"datasourceNameT,omitempty"`
	SchemaRouteRule       *SchemaRouteRule       `json:"schemaRouteRule,omitempty"`
	StatementMigrateParam *StatementMigrateParam `json:"statementMigrateParam,omitempty"`
	TaskName              *string                `json:"taskName,omitempty"`
}

// Database defines model for Database.
type Database struct {
	Host          *string `json:"host,omitempty"`
	Password      *string `json:"password,omitempty"`
	Port          *uint64 `json:"port,omitempty"`
	Schema        *string `json:"schema,omitempty"`
	SlowThreshold *uint64 `json:"slowThreshold,omitempty"`
	Username      *string `json:"username,omitempty"`
}

// Datasource defines model for Datasource.
type Datasource struct {
	Comment        *string `json:"comment,omitempty"`
	ConnectCharset *string `json:"connectCharset,omitempty"`
	ConnectParams  *string `json:"connectParams,omitempty"`
	ConnectStatus  *string `json:"connectStatus,omitempty"`
	DatasourceName *string `json:"datasourceName,omitempty"`
	DbType         *string `json:"dbType,omitempty"`
	Host           *string `json:"host,omitempty"`
	Password       *string `json:"password,omitempty"`
	PdbName        *string `json:"pdbName,omitempty"`
	Port           *uint64 `json:"port,omitempty"`
	ServiceName    *string `json:"serviceName,omitempty"`
	Username       *string `json:"username,omitempty"`
}

// NewDatasource defines model for NewDatasource.
type NewDatasource struct {
	Datasource *[]Datasource `json:"datasource,omitempty"`
}

// RequestDeleteParam defines model for RequestDeleteParam.
type RequestDeleteParam struct {
	Param *[]string `json:"param,omitempty"`
}

// RequestPostParam defines model for RequestPostParam.
type RequestPostParam struct {
	Page     *uint64 `json:"page,omitempty"`
	PageSize *uint64 `json:"pageSize,omitempty"`
	Param    *string `json:"param,omitempty"`
}

// Response defines model for Response.
type Response struct {
	Code  uint64 `json:"code"`
	Data  string `json:"data"`
	Error string `json:"error"`
}

// SchemaRouteRule defines model for SchemaRouteRule.
type SchemaRouteRule struct {
	DataMigrateRules *[]DataMigrateRule `json:"DataMigrateRules,omitempty"`
	ExcludeTableS    *[]string          `json:"excludeTableS,omitempty"`
	IncludeTableS    *[]string          `json:"includeTableS,omitempty"`
	SchemaNameS      *string            `json:"schemaNameS,omitempty"`
	SchemaNameT      *string            `json:"schemaNameT,omitempty"`
	TableRouteRules  *[]TableRouteRule  `json:"tableRouteRules,omitempty"`
}

// SchemaStructRule defines model for SchemaStructRule.
type SchemaStructRule struct {
	ColumnTypeS   *string `json:"columnTypeS,omitempty"`
	ColumnTypeT   *string `json:"columnTypeT,omitempty"`
	DefaultValueS *string `json:"defaultValueS,omitempty"`
	DefaultValueT *string `json:"defaultValueT,omitempty"`
	SchemaNameS   *string `json:"schemaNameS,omitempty"`
}

// SqlMigrateParam defines model for SqlMigrateParam.
type SqlMigrateParam struct {
	BatchSize            *uint64 `json:"batchSize,omitempty"`
	CallTimeout          *uint64 `json:"callTimeout,omitempty"`
	EnableConsistentRead *bool   `json:"enableConsistentRead,omitempty"`
	SqlHintT             *string `json:"sqlHintT,omitempty"`
	SqlThreadS           *uint64 `json:"sqlThreadS,omitempty"`
	SqlThreadT           *uint64 `json:"sqlThreadT,omitempty"`
}

// SqlMigrateRule defines model for SqlMigrateRule.
type SqlMigrateRule struct {
	CaseFieldRuleT   *string            `json:"caseFieldRuleT,omitempty"`
	ColumnRouteRules *map[string]string `json:"columnRouteRules,omitempty"`
	SchemaNameT      *string            `json:"schemaNameT,omitempty"`
	SqlHintT         *string            `json:"sqlHintT,omitempty"`
	SqlQueryS        *string            `json:"sqlQueryS,omitempty"`
	TableNameT       *string            `json:"tableNameT,omitempty"`
}

// SqlMigrateTask defines model for SqlMigrateTask.
type SqlMigrateTask struct {
	CaseFieldRule   *CaseFieldRule    `json:"caseFieldRule,omitempty"`
	Comment         *string           `json:"comment,omitempty"`
	DatasourceNameS *string           `json:"datasourceNameS,omitempty"`
	DatasourceNameT *string           `json:"datasourceNameT,omitempty"`
	SqlMigrateParam *SqlMigrateParam  `json:"sqlMigrateParam,omitempty"`
	SqlMigrateRules *[]SqlMigrateRule `json:"sqlMigrateRules,omitempty"`
	TaskName        *string           `json:"taskName,omitempty"`
}

// StatementMigrateParam defines model for StatementMigrateParam.
type StatementMigrateParam struct {
	BatchSize            *uint64 `json:"batchSize,omitempty"`
	CallTimeout          *uint64 `json:"callTimeout,omitempty"`
	ChunkSize            *uint64 `json:"chunkSize,omitempty"`
	EnableCheckpoint     *bool   `json:"enableCheckpoint,omitempty"`
	EnableConsistentRead *bool   `json:"enableConsistentRead,omitempty"`
	SqlHintS             *string `json:"sqlHintS,omitempty"`
	SqlHintT             *string `json:"sqlHintT,omitempty"`
	SqlThreadS           *uint64 `json:"sqlThreadS,omitempty"`
	SqlThreadT           *uint64 `json:"sqlThreadT,omitempty"`
	TableThread          *uint64 `json:"tableThread,omitempty"`
}

// StructMigrateParam defines model for StructMigrateParam.
type StructMigrateParam struct {
	CreateIfNotExist *bool   `json:"createIfNotExist,omitempty"`
	DirectWrite      *bool   `json:"directWrite,omitempty"`
	MigrateThread    *uint64 `json:"migrateThread,omitempty"`
	OutputDir        *string `json:"outputDir,omitempty"`
}

// StructMigrateRule defines model for StructMigrateRule.
type StructMigrateRule struct {
	ColumnStructRules *[]ColumnStructRule `json:"columnStructRules,omitempty"`
	SchemaStructRules *[]SchemaStructRule `json:"schemaStructRules,omitempty"`
	TableAttrsRules   *[]TableAttrsRule   `json:"tableAttrsRules,omitempty"`
	TableStructRules  *[]TableStructRule  `json:"tableStructRules,omitempty"`
	TaskStructRules   *[]TaskStructRule   `json:"taskStructRules,omitempty"`
}

// StructMigrateTask defines model for StructMigrateTask.
type StructMigrateTask struct {
	CaseFieldRule      *CaseFieldRule      `json:"caseFieldRule,omitempty"`
	Comment            *string             `json:"comment,omitempty"`
	DatasourceNameS    *string             `json:"datasourceNameS,omitempty"`
	DatasourceNameT    *string             `json:"datasourceNameT,omitempty"`
	SchemaRouteRule    *SchemaRouteRule    `json:"schemaRouteRule,omitempty"`
	StructMigrateParam *StructMigrateParam `json:"structMigrateParam,omitempty"`
	StructMigrateRule  *StructMigrateRule  `json:"structMigrateRule,omitempty"`
	TaskName           *string             `json:"taskName,omitempty"`
}

// TableAttrsRule defines model for TableAttrsRule.
type TableAttrsRule struct {
	SchemaNameS *string   `json:"schemaNameS,omitempty"`
	TableAttrsT *string   `json:"tableAttrsT,omitempty"`
	TableNamesS *[]string `json:"tableNamesS,omitempty"`
}

// TableRouteRule defines model for TableRouteRule.
type TableRouteRule struct {
	ColumnRouteRules *map[string]string `json:"columnRouteRules,omitempty"`
	TableNameS       *string            `json:"tableNameS,omitempty"`
	TableNameT       *string            `json:"tableNameT,omitempty"`
}

// TableStructRule defines model for TableStructRule.
type TableStructRule struct {
	ColumnTypeS   *string `json:"columnTypeS,omitempty"`
	ColumnTypeT   *string `json:"columnTypeT,omitempty"`
	DefaultValueS *string `json:"defaultValueS,omitempty"`
	DefaultValueT *string `json:"defaultValueT,omitempty"`
	SchemaNameS   *string `json:"schemaNameS,omitempty"`
	TableNameS    *string `json:"tableNameS,omitempty"`
}

// Task defines model for Task.
type Task struct {
	Express  *string `json:"express,omitempty"`
	Operate  *string `json:"operate,omitempty"`
	TaskName *string `json:"taskName,omitempty"`
}

// TaskStructRule defines model for TaskStructRule.
type TaskStructRule struct {
	ColumnTypeS   *string `json:"columnTypeS,omitempty"`
	ColumnTypeT   *string `json:"columnTypeT,omitempty"`
	DefaultValueS *string `json:"defaultValueS,omitempty"`
	DefaultValueT *string `json:"defaultValueT,omitempty"`
}

// APIPutDatabaseJSONRequestBody defines body for APIPutDatabase for application/json ContentType.
type APIPutDatabaseJSONRequestBody = Database

// APIDeleteDatasourceJSONRequestBody defines body for APIDeleteDatasource for application/json ContentType.
type APIDeleteDatasourceJSONRequestBody = RequestDeleteParam

// APIListDatasourceJSONRequestBody defines body for APIListDatasource for application/json ContentType.
type APIListDatasourceJSONRequestBody = RequestPostParam

// APIPutDatasourceJSONRequestBody defines body for APIPutDatasource for application/json ContentType.
type APIPutDatasourceJSONRequestBody = NewDatasource

// APIPostTaskJSONRequestBody defines body for APIPostTask for application/json ContentType.
type APIPostTaskJSONRequestBody = Task

// APIDeleteDataMigrateJSONRequestBody defines body for APIDeleteDataMigrate for application/json ContentType.
type APIDeleteDataMigrateJSONRequestBody = RequestDeleteParam

// APIListDataMigrateJSONRequestBody defines body for APIListDataMigrate for application/json ContentType.
type APIListDataMigrateJSONRequestBody = RequestPostParam

// APIPutDataMigrateJSONRequestBody defines body for APIPutDataMigrate for application/json ContentType.
type APIPutDataMigrateJSONRequestBody = DataMigrateTask

// APIDeleteSqlMigrateJSONRequestBody defines body for APIDeleteSqlMigrate for application/json ContentType.
type APIDeleteSqlMigrateJSONRequestBody = RequestDeleteParam

// APIListSqlMigrateJSONRequestBody defines body for APIListSqlMigrate for application/json ContentType.
type APIListSqlMigrateJSONRequestBody = RequestPostParam

// APIPutSqlMigrateJSONRequestBody defines body for APIPutSqlMigrate for application/json ContentType.
type APIPutSqlMigrateJSONRequestBody = SqlMigrateTask

// APIDeleteStructMigrateJSONRequestBody defines body for APIDeleteStructMigrate for application/json ContentType.
type APIDeleteStructMigrateJSONRequestBody = RequestDeleteParam

// APIListStructMigrateJSONRequestBody defines body for APIListStructMigrate for application/json ContentType.
type APIListStructMigrateJSONRequestBody = RequestPostParam

// APIPutStructMigrateJSONRequestBody defines body for APIPutStructMigrate for application/json ContentType.
type APIPutStructMigrateJSONRequestBody = StructMigrateTask

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"fwMAAP//DYQrNZZBAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %w", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %w", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	res := make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	resolvePath := PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		pathToFile := url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
