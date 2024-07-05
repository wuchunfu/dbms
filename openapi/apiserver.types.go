// Package openapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
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

// AssessMigrateParam defines model for AssessMigrateParam.
type AssessMigrateParam struct {
	CaseFieldRuleS *string `json:"caseFieldRuleS,omitempty"`
	SchemaNameS    *string `json:"schemaNameS,omitempty"`
}

// AssessMigrateTask defines model for AssessMigrateTask.
type AssessMigrateTask struct {
	AssessMigrateParam *AssessMigrateParam `json:"assessMigrateParam,omitempty"`
	Comment            *string             `json:"comment"`
	DatasourceNameS    *string             `json:"datasourceNameS,omitempty"`
	DatasourceNameT    *string             `json:"datasourceNameT,omitempty"`
	TaskName           *string             `json:"taskName,omitempty"`
}

// CaseFieldRule defines model for CaseFieldRule.
type CaseFieldRule struct {
	CaseFieldRuleS *string `json:"caseFieldRuleS,omitempty"`
	CaseFieldRuleT *string `json:"caseFieldRuleT,omitempty"`
}

// ColumnStructRule defines model for ColumnStructRule.
type ColumnStructRule struct {
	ColumnNameS   *string `json:"columnNameS,omitempty"`
	ColumnTypeS   *string `json:"columnTypeS"`
	ColumnTypeT   *string `json:"columnTypeT"`
	DefaultValueS *string `json:"defaultValueS"`
	DefaultValueT *string `json:"defaultValueT"`
	SchemaNameS   *string `json:"schemaNameS,omitempty"`
	TableNameS    *string `json:"tableNameS,omitempty"`
}

// CsvMigrateParam defines model for CsvMigrateParam.
type CsvMigrateParam struct {
	BatchSize            *uint64 `json:"batchSize,omitempty"`
	CallTimeout          *uint64 `json:"callTimeout,omitempty"`
	ChunkSize            *uint64 `json:"chunkSize,omitempty"`
	DataCharsetT         *string `json:"dataCharsetT,omitempty"`
	Delimiter            *string `json:"delimiter,omitempty"`
	DiskUsageFactor      *string `json:"diskUsageFactor,omitempty"`
	EnableCheckpoint     *bool   `json:"enableCheckpoint,omitempty"`
	EnableConsistentRead *bool   `json:"enableConsistentRead,omitempty"`
	EscapeBackslash      *bool   `json:"escapeBackslash,omitempty"`
	Header               *bool   `json:"header,omitempty"`
	NullValue            *string `json:"nullValue,omitempty"`
	OutputDir            *string `json:"outputDir,omitempty"`
	Separator            *string `json:"separator,omitempty"`
	SqlHintS             *string `json:"sqlHintS,omitempty"`
	SqlThreadS           *uint64 `json:"sqlThreadS,omitempty"`
	TableThread          *uint64 `json:"tableThread,omitempty"`
	Terminator           *string `json:"terminator,omitempty"`
}

// CsvMigrateTask defines model for CsvMigrateTask.
type CsvMigrateTask struct {
	CaseFieldRule    *CaseFieldRule     `json:"caseFieldRule,omitempty"`
	Comment          *string            `json:"comment"`
	CsvMigrateParam  *CsvMigrateParam   `json:"csvMigrateParam,omitempty"`
	DataMigrateRules *[]DataMigrateRule `json:"dataMigrateRules"`
	DatasourceNameS  *string            `json:"datasourceNameS,omitempty"`
	DatasourceNameT  *string            `json:"datasourceNameT,omitempty"`
	SchemaRouteRule  *SchemaRouteRule   `json:"schemaRouteRule,omitempty"`
	TaskName         *string            `json:"taskName,omitempty"`
}

// DataCompareParam defines model for DataCompareParam.
type DataCompareParam struct {
	BatchSize            *uint64 `json:"batchSize,omitempty"`
	CallTimeout          *uint64 `json:"callTimeout,omitempty"`
	ConsistentReadPointS *string `json:"consistentReadPointS,omitempty"`
	ConsistentReadPointT *string `json:"consistentReadPointT,omitempty"`
	EnableCheckpoint     *bool   `json:"enableCheckpoint,omitempty"`
	EnableConsistentRead *bool   `json:"enableConsistentRead,omitempty"`
	OnlyCompareRow       *bool   `json:"onlyCompareRow,omitempty"`
	SqlHintS             *string `json:"sqlHintS,omitempty"`
	SqlHintT             *string `json:"sqlHintT,omitempty"`
	SqlThread            *uint64 `json:"sqlThread,omitempty"`
	TableThread          *uint64 `json:"tableThread,omitempty"`
	WriteThread          *uint64 `json:"writeThread,omitempty"`
}

// DataCompareRule defines model for DataCompareRule.
type DataCompareRule struct {
	CompareField *string   `json:"compareField,omitempty"`
	CompareRange *string   `json:"compareRange,omitempty"`
	IgnoreFields *[]string `json:"ignoreFields,omitempty"`
	TableNameS   *string   `json:"tableNameS,omitempty"`
}

// DataCompareTask defines model for DataCompareTask.
type DataCompareTask struct {
	CaseFieldRule    *CaseFieldRule     `json:"caseFieldRule,omitempty"`
	Comment          *string            `json:"comment"`
	DataCompareParam *DataCompareParam  `json:"dataCompareParam,omitempty"`
	DataCompareRules *[]DataCompareRule `json:"dataCompareRules"`
	DatasourceNameS  *string            `json:"datasourceNameS,omitempty"`
	DatasourceNameT  *string            `json:"datasourceNameT,omitempty"`
	SchemaRouteRule  *SchemaRouteRule   `json:"schemaRouteRule,omitempty"`
	TaskName         *string            `json:"taskName,omitempty"`
}

// DataMigrateRule defines model for DataMigrateRule.
type DataMigrateRule struct {
	EnableChunkStrategy *bool   `json:"enableChunkStrategy,omitempty"`
	SqlHintS            *string `json:"sqlHintS,omitempty"`
	TableNameS          *string `json:"tableNameS,omitempty"`
	WhereRange          *string `json:"whereRange,omitempty"`
}

// DataScanParam defines model for DataScanParam.
type DataScanParam struct {
	BatchSize            *uint64 `json:"batchSize,omitempty"`
	CallTimeout          *uint64 `json:"callTimeout,omitempty"`
	ChunkSize            *uint64 `json:"chunkSize,omitempty"`
	EnableCheckpoint     *bool   `json:"enableCheckpoint,omitempty"`
	EnableConsistentRead *bool   `json:"enableConsistentRead,omitempty"`
	SqlHintS             *string `json:"sqlHintS,omitempty"`
	SqlThreadS           *uint64 `json:"sqlThreadS,omitempty"`
	TableSamplerateS     *uint64 `json:"tableSamplerateS,omitempty"`
	TableThread          *uint64 `json:"tableThread,omitempty"`
}

// DataScanRule defines model for DataScanRule.
type DataScanRule struct {
	SqlHintS         *string `json:"sqlHintS,omitempty"`
	TableNameS       *string `json:"tableNameS,omitempty"`
	TableSamplerateS *uint64 `json:"tableSamplerateS,omitempty"`
}

// DataScanTask defines model for DataScanTask.
type DataScanTask struct {
	CaseFieldRule   *CaseFieldRule   `json:"caseFieldRule,omitempty"`
	Comment         *string          `json:"comment"`
	DataScanParam   *DataScanParam   `json:"dataScanParam,omitempty"`
	DataScanRules   *[]DataScanRule  `json:"dataScanRules"`
	DatasourceNameS *string          `json:"datasourceNameS,omitempty"`
	DatasourceNameT *string          `json:"datasourceNameT,omitempty"`
	SchemaRouteRule *SchemaRouteRule `json:"schemaRouteRule,omitempty"`
	TaskName        *string          `json:"taskName,omitempty"`
}

// Database defines model for Database.
type Database struct {
	Host          *string `json:"host,omitempty"`
	InitThread    *uint64 `json:"initThread,omitempty"`
	Password      *string `json:"password,omitempty"`
	Port          *uint64 `json:"port,omitempty"`
	Schema        *string `json:"schema,omitempty"`
	SlowThreshold *uint64 `json:"slowThreshold,omitempty"`
	Username      *string `json:"username,omitempty"`
}

// Datasource defines model for Datasource.
type Datasource struct {
	Comment        *string `json:"comment"`
	ConnectCharset *string `json:"connectCharset,omitempty"`
	ConnectParams  *string `json:"connectParams"`
	ConnectStatus  *string `json:"connectStatus"`
	DatasourceName *string `json:"datasourceName,omitempty"`
	DbType         *string `json:"dbType,omitempty"`
	Host           *string `json:"host,omitempty"`
	Password       *string `json:"password,omitempty"`
	PdbName        *string `json:"pdbName"`
	Port           *uint64 `json:"port,omitempty"`
	ServiceName    *string `json:"serviceName"`
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
	ExcludeTableS   *[]string         `json:"excludeTableS"`
	IncludeTableS   *[]string         `json:"includeTableS"`
	SchemaNameS     *string           `json:"schemaNameS,omitempty"`
	SchemaNameT     *string           `json:"schemaNameT,omitempty"`
	TableRouteRules *[]TableRouteRule `json:"tableRouteRules"`
}

// SchemaStructRule defines model for SchemaStructRule.
type SchemaStructRule struct {
	ColumnTypeS   *string `json:"columnTypeS"`
	ColumnTypeT   *string `json:"columnTypeT"`
	DefaultValueS *string `json:"defaultValueS"`
	DefaultValueT *string `json:"defaultValueT"`
	SchemaNameS   *string `json:"schemaNameS,omitempty"`
}

// SqlMigrateParam defines model for SqlMigrateParam.
type SqlMigrateParam struct {
	BatchSize            *uint64 `json:"batchSize,omitempty"`
	CallTimeout          *uint64 `json:"callTimeout,omitempty"`
	EnableConsistentRead *bool   `json:"enableConsistentRead,omitempty"`
	EnableSafeMode       *bool   `json:"enableSafeMode,omitempty"`
	SqlHintT             *string `json:"sqlHintT,omitempty"`
	SqlThreadS           *uint64 `json:"sqlThreadS,omitempty"`
	SqlThreadT           *uint64 `json:"sqlThreadT,omitempty"`
}

// SqlMigrateRule defines model for SqlMigrateRule.
type SqlMigrateRule struct {
	CaseFieldRuleT   *string            `json:"caseFieldRuleT,omitempty"`
	ColumnRouteRules *map[string]string `json:"columnRouteRules"`
	SchemaNameT      *string            `json:"schemaNameT,omitempty"`
	SqlHintT         *string            `json:"sqlHintT,omitempty"`
	SqlQueryS        *string            `json:"sqlQueryS,omitempty"`
	TableNameT       *string            `json:"tableNameT,omitempty"`
}

// SqlMigrateTask defines model for SqlMigrateTask.
type SqlMigrateTask struct {
	CaseFieldRule   *CaseFieldRule    `json:"caseFieldRule,omitempty"`
	Comment         *string           `json:"comment"`
	DatasourceNameS *string           `json:"datasourceNameS,omitempty"`
	DatasourceNameT *string           `json:"datasourceNameT,omitempty"`
	SqlMigrateParam *SqlMigrateParam  `json:"sqlMigrateParam,omitempty"`
	SqlMigrateRules *[]SqlMigrateRule `json:"sqlMigrateRules"`
	TaskName        *string           `json:"taskName,omitempty"`
}

// StatementMigrateParam defines model for StatementMigrateParam.
type StatementMigrateParam struct {
	BatchSize            *uint64 `json:"batchSize,omitempty"`
	CallTimeout          *uint64 `json:"callTimeout,omitempty"`
	ChunkSize            *uint64 `json:"chunkSize,omitempty"`
	EnableCheckpoint     *bool   `json:"enableCheckpoint,omitempty"`
	EnableConsistentRead *bool   `json:"enableConsistentRead,omitempty"`
	EnableSafeMode       *bool   `json:"enableSafeMode,omitempty"`
	SqlHintS             *string `json:"sqlHintS,omitempty"`
	SqlHintT             *string `json:"sqlHintT,omitempty"`
	SqlThreadS           *uint64 `json:"sqlThreadS,omitempty"`
	SqlThreadT           *uint64 `json:"sqlThreadT,omitempty"`
	TableThread          *uint64 `json:"tableThread,omitempty"`
}

// StmtMigrateTask defines model for StmtMigrateTask.
type StmtMigrateTask struct {
	DataMigrateRules      *[]DataMigrateRule     `json:"DataMigrateRules"`
	CaseFieldRule         *CaseFieldRule         `json:"caseFieldRule,omitempty"`
	Comment               *string                `json:"comment"`
	DatasourceNameS       *string                `json:"datasourceNameS,omitempty"`
	DatasourceNameT       *string                `json:"datasourceNameT,omitempty"`
	SchemaRouteRule       *SchemaRouteRule       `json:"schemaRouteRule,omitempty"`
	StatementMigrateParam *StatementMigrateParam `json:"statementMigrateParam,omitempty"`
	TaskName              *string                `json:"taskName,omitempty"`
}

// StructCompareParam defines model for StructCompareParam.
type StructCompareParam struct {
	CompareThread    *uint64 `json:"compareThread,omitempty"`
	EnableCheckpoint *bool   `json:"enableCheckpoint,omitempty"`
}

// StructCompareRule defines model for StructCompareRule.
type StructCompareRule struct {
	ColumnStructRules *[]ColumnStructRule `json:"columnStructRules"`
	SchemaStructRules *[]SchemaStructRule `json:"schemaStructRules"`
	TableStructRules  *[]TableStructRule  `json:"tableStructRules"`
	TaskStructRules   *[]TaskStructRule   `json:"taskStructRules"`
}

// StructCompareTask defines model for StructCompareTask.
type StructCompareTask struct {
	CaseFieldRule      *CaseFieldRule      `json:"caseFieldRule,omitempty"`
	Comment            *string             `json:"comment"`
	DatasourceNameS    *string             `json:"datasourceNameS,omitempty"`
	DatasourceNameT    *string             `json:"datasourceNameT,omitempty"`
	SchemaRouteRule    *SchemaRouteRule    `json:"schemaRouteRule,omitempty"`
	StructCompareParam *StructCompareParam `json:"structCompareParam,omitempty"`
	StructCompareRule  *StructCompareRule  `json:"structCompareRule,omitempty"`
	TaskName           *string             `json:"taskName,omitempty"`
}

// StructMigrateParam defines model for StructMigrateParam.
type StructMigrateParam struct {
	CreateIfNotExist   *bool   `json:"createIfNotExist,omitempty"`
	EnableCheckpoint   *bool   `json:"enableCheckpoint,omitempty"`
	EnableDirectCreate *bool   `json:"enableDirectCreate,omitempty"`
	MigrateThread      *uint64 `json:"migrateThread,omitempty"`
}

// StructMigrateRule defines model for StructMigrateRule.
type StructMigrateRule struct {
	ColumnStructRules *[]ColumnStructRule `json:"columnStructRules"`
	SchemaStructRules *[]SchemaStructRule `json:"schemaStructRules"`
	TableAttrsRules   *[]TableAttrsRule   `json:"tableAttrsRules"`
	TableStructRules  *[]TableStructRule  `json:"tableStructRules"`
	TaskStructRules   *[]TaskStructRule   `json:"taskStructRules"`
}

// StructMigrateTask defines model for StructMigrateTask.
type StructMigrateTask struct {
	CaseFieldRule      *CaseFieldRule      `json:"caseFieldRule,omitempty"`
	Comment            *string             `json:"comment"`
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
	TableAttrsT *string   `json:"tableAttrsT"`
	TableNamesS *[]string `json:"tableNamesS"`
}

// TableRouteRule defines model for TableRouteRule.
type TableRouteRule struct {
	ColumnRouteRules *map[string]string `json:"columnRouteRules"`
	TableNameS       *string            `json:"tableNameS,omitempty"`
	TableNameT       *string            `json:"tableNameT,omitempty"`
}

// TableStructRule defines model for TableStructRule.
type TableStructRule struct {
	ColumnTypeS   *string `json:"columnTypeS"`
	ColumnTypeT   *string `json:"columnTypeT"`
	DefaultValueS *string `json:"defaultValueS"`
	DefaultValueT *string `json:"defaultValueT"`
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
	ColumnTypeS   *string `json:"columnTypeS"`
	ColumnTypeT   *string `json:"columnTypeT"`
	DefaultValueS *string `json:"defaultValueS"`
	DefaultValueT *string `json:"defaultValueT"`
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

// APIDeleteAssessMigrateJSONRequestBody defines body for APIDeleteAssessMigrate for application/json ContentType.
type APIDeleteAssessMigrateJSONRequestBody = RequestDeleteParam

// APIListAssessMigrateJSONRequestBody defines body for APIListAssessMigrate for application/json ContentType.
type APIListAssessMigrateJSONRequestBody = RequestPostParam

// APIPutAssessMigrateJSONRequestBody defines body for APIPutAssessMigrate for application/json ContentType.
type APIPutAssessMigrateJSONRequestBody = AssessMigrateTask

// APIDeleteCsvMigrateJSONRequestBody defines body for APIDeleteCsvMigrate for application/json ContentType.
type APIDeleteCsvMigrateJSONRequestBody = RequestDeleteParam

// APIListCsvMigrateJSONRequestBody defines body for APIListCsvMigrate for application/json ContentType.
type APIListCsvMigrateJSONRequestBody = RequestPostParam

// APIPutCsvMigrateJSONRequestBody defines body for APIPutCsvMigrate for application/json ContentType.
type APIPutCsvMigrateJSONRequestBody = CsvMigrateTask

// APIDeleteDataCompareJSONRequestBody defines body for APIDeleteDataCompare for application/json ContentType.
type APIDeleteDataCompareJSONRequestBody = RequestDeleteParam

// APIListDataCompareJSONRequestBody defines body for APIListDataCompare for application/json ContentType.
type APIListDataCompareJSONRequestBody = RequestPostParam

// APIPutDataCompareJSONRequestBody defines body for APIPutDataCompare for application/json ContentType.
type APIPutDataCompareJSONRequestBody = DataCompareTask

// APIDeleteDataScanJSONRequestBody defines body for APIDeleteDataScan for application/json ContentType.
type APIDeleteDataScanJSONRequestBody = RequestDeleteParam

// APIListDataScanJSONRequestBody defines body for APIListDataScan for application/json ContentType.
type APIListDataScanJSONRequestBody = RequestPostParam

// APIPutDataScanJSONRequestBody defines body for APIPutDataScan for application/json ContentType.
type APIPutDataScanJSONRequestBody = DataScanTask

// APIDeleteSqlMigrateJSONRequestBody defines body for APIDeleteSqlMigrate for application/json ContentType.
type APIDeleteSqlMigrateJSONRequestBody = RequestDeleteParam

// APIListSqlMigrateJSONRequestBody defines body for APIListSqlMigrate for application/json ContentType.
type APIListSqlMigrateJSONRequestBody = RequestPostParam

// APIPutSqlMigrateJSONRequestBody defines body for APIPutSqlMigrate for application/json ContentType.
type APIPutSqlMigrateJSONRequestBody = SqlMigrateTask

// APIDeleteStmtMigrateJSONRequestBody defines body for APIDeleteStmtMigrate for application/json ContentType.
type APIDeleteStmtMigrateJSONRequestBody = RequestDeleteParam

// APIListStmtMigrateJSONRequestBody defines body for APIListStmtMigrate for application/json ContentType.
type APIListStmtMigrateJSONRequestBody = RequestPostParam

// APIPutStmtMigrateJSONRequestBody defines body for APIPutStmtMigrate for application/json ContentType.
type APIPutStmtMigrateJSONRequestBody = StmtMigrateTask

// APIDeleteStructCompareJSONRequestBody defines body for APIDeleteStructCompare for application/json ContentType.
type APIDeleteStructCompareJSONRequestBody = RequestDeleteParam

// APIListStructCompareJSONRequestBody defines body for APIListStructCompare for application/json ContentType.
type APIListStructCompareJSONRequestBody = RequestPostParam

// APIPutStructCompareJSONRequestBody defines body for APIPutStructCompare for application/json ContentType.
type APIPutStructCompareJSONRequestBody = StructCompareTask

// APIDeleteStructMigrateJSONRequestBody defines body for APIDeleteStructMigrate for application/json ContentType.
type APIDeleteStructMigrateJSONRequestBody = RequestDeleteParam

// APIListStructMigrateJSONRequestBody defines body for APIListStructMigrate for application/json ContentType.
type APIListStructMigrateJSONRequestBody = RequestPostParam

// APIPutStructMigrateJSONRequestBody defines body for APIPutStructMigrate for application/json ContentType.
type APIPutStructMigrateJSONRequestBody = StructMigrateTask

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
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
