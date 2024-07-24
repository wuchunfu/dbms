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
	BatchSize              *uint64   `json:"batchSize,omitempty"`
	CallTimeout            *uint64   `json:"callTimeout,omitempty"`
	ChunkSize              *uint64   `json:"chunkSize,omitempty"`
	ConsistentReadPointS   *string   `json:"consistentReadPointS,omitempty"`
	ConsistentReadPointT   *string   `json:"consistentReadPointT,omitempty"`
	EnableCheckpoint       *bool     `json:"enableCheckpoint,omitempty"`
	EnableCollationSetting *bool     `json:"enableCollationSetting,omitempty"`
	EnableConsistentRead   *bool     `json:"enableConsistentRead,omitempty"`
	IgnoreConditionFields  *[]string `json:"ignoreConditionFields,omitempty"`
	OnlyCompareRow         *bool     `json:"onlyCompareRow,omitempty"`
	RepairStmtFlow         *string   `json:"repairStmtFlow,omitempty"`
	SqlHintS               *string   `json:"sqlHintS,omitempty"`
	SqlHintT               *string   `json:"sqlHintT,omitempty"`
	SqlThread              *uint64   `json:"sqlThread,omitempty"`
	TableThread            *uint64   `json:"tableThread,omitempty"`
	WriteThread            *uint64   `json:"writeThread,omitempty"`
}

// DataCompareRule defines model for DataCompareRule.
type DataCompareRule struct {
	CompareConditionField *string   `json:"compareConditionField,omitempty"`
	CompareConditionRange *string   `json:"compareConditionRange,omitempty"`
	IgnoreConditionFields *[]string `json:"ignoreConditionFields,omitempty"`
	IgnoreSelectFields    *[]string `json:"ignoreSelectFields,omitempty"`
	TableNameS            *string   `json:"tableNameS,omitempty"`
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
	DbName         *string `json:"dbName"`
	DbType         *string `json:"dbType,omitempty"`
	Host           *string `json:"host,omitempty"`
	Password       *string `json:"password,omitempty"`
	PdbName        *string `json:"pdbName"`
	Port           *uint64 `json:"port,omitempty"`
	ServiceName    *string `json:"serviceName"`
	SessionParams  *string `json:"sessionParams"`
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
