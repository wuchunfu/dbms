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
	CallTimeout    *uint64 `json:"callTimeout,omitempty"`
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
	WriteThread          *uint64 `json:"writeThread,omitempty"`
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
	DisableMd5Checksum     *bool     `json:"disableMd5Checksum,omitempty"`
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
	SqlHintS              *string   `json:"sqlHintS,omitempty"`
	SqlHintT              *string   `json:"sqlHintT,omitempty"`
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
	WriteThread          *uint64 `json:"writeThread,omitempty"`
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
	WriteThread          *uint64 `json:"writeThread,omitempty"`
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
	WriteThread          *uint64 `json:"writeThread,omitempty"`
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
	CallTimeout      *uint64 `json:"callTimeout,omitempty"`
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
	CallTimeout        *uint64 `json:"callTimeout,omitempty"`
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
	"nwEAAP//izNAU7F9AAA=",
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
