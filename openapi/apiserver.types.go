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
	"/1LamxvlfAAA",
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
