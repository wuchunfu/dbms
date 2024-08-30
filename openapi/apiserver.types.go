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
	Separator              *string   `json:"separator,omitempty"`
	SqlHintS               *string   `json:"sqlHintS,omitempty"`
	SqlHintT               *string   `json:"sqlHintT,omitempty"`
	SqlThread              *uint64   `json:"sqlThread,omitempty"`
	TableThread            *uint64   `json:"tableThread,omitempty"`
	WriteThread            *uint64   `json:"writeThread,omitempty"`
}

// DataCompareRule defines model for DataCompareRule.
type DataCompareRule struct {
	CompareConditionField  *string   `json:"compareConditionField,omitempty"`
	CompareConditionRangeS *string   `json:"compareConditionRangeS,omitempty"`
	CompareConditionRangeT *string   `json:"compareConditionRangeT,omitempty"`
	IgnoreConditionFields  *[]string `json:"ignoreConditionFields,omitempty"`
	IgnoreSelectFields     *[]string `json:"ignoreSelectFields,omitempty"`
	SqlHintS               *string   `json:"sqlHintS,omitempty"`
	SqlHintT               *string   `json:"sqlHintT,omitempty"`
	TableNameS             *string   `json:"tableNameS,omitempty"`
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
	ExcludeSequenceS *[]string         `json:"excludeSequenceS"`
	ExcludeTableS    *[]string         `json:"excludeTableS"`
	IncludeSequenceS *[]string         `json:"includeSequenceS"`
	IncludeTableS    *[]string         `json:"includeTableS"`
	SchemaNameS      *string           `json:"schemaNameS,omitempty"`
	SchemaNameT      *string           `json:"schemaNameT,omitempty"`
	TableRouteRules  *[]TableRouteRule `json:"tableRouteRules"`
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
	EnableCheckpoint     *bool   `json:"enableCheckpoint,omitempty"`
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
	CallTimeout       *uint64 `json:"callTimeout,omitempty"`
	CompareThread     *uint64 `json:"compareThread,omitempty"`
	EnableCheckpoint  *bool   `json:"enableCheckpoint,omitempty"`
	IgnoreCaseCompare *bool   `json:"ignoreCaseCompare,omitempty"`
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
