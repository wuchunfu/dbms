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
package constant

var (
	DataCompareOracleDatabaseSupportNumberSubtypes = []string{
		BuildInOracleDatatypeNumber,
		BuildInOracleDatatypeDecimal,
		BuildInOracleDatatypeDec,
		BuildInOracleDatatypeDoublePrecision,
		BuildInOracleDatatypeNumeric,
		BuildInOracleDatatypeReal,
		BuildInOracleDatatypeInteger,
		BuildInOracleDatatypeInt,
		BuildInOracleDatatypeSmallint,
		BuildInOracleDatatypeFloat,
		BuildInOracleDatatypeBinaryDouble,
		BuildInOracleDatatypeBinaryFloat,
	}
	DataCompareOracleDatabaseSupportVarcharSubtypes = []string{
		BuildInOracleDatatypeVarchar,
		BuildInOracleDatatypeVarchar2,
		BuildInOracleDatatypeChar,
		BuildInOracleDatatypeNchar,
		BuildInOracleDatatypeNvarchar2,
		BuildInOracleDatatypeCharacter,
		BuildInOracleDatatypeNcharVarying,
	}
	DataCompareOracleDatabaseSupportDateSubtypes = []string{
		BuildInOracleDatatypeDate,
	}
	DataCompareOracleDatabaseSupportTimestampSubtypes = []string{
		BuildInOracleDatatypeTimestamp,
		BuildInOracleDatatypeTimestamp0,
		BuildInOracleDatatypeTimestamp1,
		BuildInOracleDatatypeTimestamp2,
		BuildInOracleDatatypeTimestamp3,
		BuildInOracleDatatypeTimestamp4,
		BuildInOracleDatatypeTimestamp5,
		BuildInOracleDatatypeTimestamp6,
		BuildInOracleDatatypeTimestamp7,
		BuildInOracleDatatypeTimestamp8,
		BuildInOracleDatatypeTimestamp9,
		BuildInOracleDatatypeTimestampWithTimeZone0,
		BuildInOracleDatatypeTimestampWithTimeZone1,
		BuildInOracleDatatypeTimestampWithTimeZone2,
		BuildInOracleDatatypeTimestampWithTimeZone3,
		BuildInOracleDatatypeTimestampWithTimeZone4,
		BuildInOracleDatatypeTimestampWithTimeZone5,
		BuildInOracleDatatypeTimestampWithTimeZone6,
		BuildInOracleDatatypeTimestampWithTimeZone7,
		BuildInOracleDatatypeTimestampWithTimeZone8,
		BuildInOracleDatatypeTimestampWithTimeZone9,
		BuildInOracleDatatypeTimestampWithLocalTimeZone0,
		BuildInOracleDatatypeTimestampWithLocalTimeZone1,
		BuildInOracleDatatypeTimestampWithLocalTimeZone2,
		BuildInOracleDatatypeTimestampWithLocalTimeZone3,
		BuildInOracleDatatypeTimestampWithLocalTimeZone4,
		BuildInOracleDatatypeTimestampWithLocalTimeZone5,
		BuildInOracleDatatypeTimestampWithLocalTimeZone6,
		BuildInOracleDatatypeTimestampWithLocalTimeZone7,
		BuildInOracleDatatypeTimestampWithLocalTimeZone8,
		BuildInOracleDatatypeTimestampWithLocalTimeZone9,
	}

	DataCompareORACLECompatibleDatabaseColumnDatatypeSupportCollation = []string{
		BuildInOracleDatatypeChar,
		BuildInOracleDatatypeNchar,
		BuildInOracleDatatypeVarchar,
		BuildInOracleDatatypeVarchar2,
		BuildInOracleDatatypeNvarchar2,
		BuildInOracleDatatypeCharacter,
		BuildInOracleDatatypeNcharVarying}

	DataCompareSymbolLt                 = "<"
	DataCompareSymbolLte                = "<="
	DataCompareSymbolGt                 = ">"
	DataCompareMethodDatabaseCheckMD5   = "DATABASE_MD5"
	DataCompareMethodDatabaseCheckCRC32 = "DATABASE_CRC32"
	DataCompareMethodDatabaseCheckRows  = "DATABASE_ROWS"
	DataCompareMethodProgramCheckCRC32  = "PROGRAM_CRC32"

	DataCompareMethodCheckMD5ValueLength = 32

	DataCompareMYSQLCompatibleDatabaseSupportDecimalSubtypes = []string{
		BuildInMySQLDatatypeTinyint,
		BuildInMySQLDatatypeSmallint,
		BuildInMySQLDatatypeInt,
		BuildInMySQLDatatypeInteger,
		BuildInMySQLDatatypeMediumint,
		BuildInMySQLDatatypeBigint,
		BuildInMySQLDatatypeFloat,
		BuildInMySQLDatatypeDecimal,
		BuildInMySQLDatatypeDoublePrecision,
		BuildInMySQLDatatypeNumeric,
		BuildInMySQLDatatypeReal,
		BuildInMySQLDatatypeDouble,
	}
	DataCompareMYSQLCompatibleDatabaseColumnTimeSubtypes = []string{
		BuildInMySQLDatatypeDatetime,
		BuildInMySQLDatatypeTimestamp,
		BuildInMySQLDatatypeDate,
	}
	DataCompareMYSQLCompatibleDatabaseColumnDatatypeSupportCollation = []string{
		BuildInMySQLDatatypeChar,
		BuildInMySQLDatatypeVarchar,
		BuildInMySQLDatatypeText,
		BuildInMySQLDatatypeTinyText,
		BuildInMySQLDatatypeMediumText,
		BuildInMySQLDatatypeLongText}

	DataComparePostgresCompatibleDatabaseSupportDecimalSubtypes = []string{
		BuildInPostgresDatatypeInteger,
		BuildInPostgresDatatypeSmallInt,
		BuildInPostgresDatatypeBigInt,
		BuildInPostgresDatatypeNumeric,
		BuildInPostgresDatatypeDecimal,
		BuildInPostgresDatatypeMoney,
		BuildInPostgresDatatypeReal,
		BuildInPostgresDatatypeDoublePrecision,
	}
	DataComparePostgresCompatibleDatabaseColumnTimeSubtypes = []string{
		BuildInPostgresDatatypeTimeWithoutTimeZone,
		BuildInPostgresDatatypeTimestampWithoutTimeZone,
		BuildInPostgresDatatypeInterval,
	}
	DataComparePostgresCompatibleDatabaseColumnDateSubtypes = []string{
		BuildInPostgresDatatypeDate,
	}
	DataComparePostgresCompatibleDatabaseColumnDatatypeSupportCollation = []string{
		BuildInPostgresDatatypeText,
		BuildInPostgresDatatypeCharacterVarying,
		BuildInPostgresDatatypeCharacter,
	}
)

const (
	DataCompareFixStmtTypeDelete = "D"
	DataCompareFixStmtTypeInsert = "I"
	DataCompareFixStmtTypeRows   = "R" // row counts compare

	DataCompareRepairStmtFlowUpstream   = "UPSTREAM"
	DataCompareRepairStmtFlowDownstream = "DOWNSTREAM"

	// DataCompareDisabledCollationSettingFillEmptyString represent disabled collation setting and the column name is not support collation setting
	DataCompareDisabledCollationSettingFillEmptyString = ""

	// Garbled characters and rare characters can cause inconsistent data verification, but no repair statements can be generated
	// verify scan command tool can be used to scan and locate specific problem lines
	DataCompareSeekAsciiColumnPrefix         = "ASCII"
	DataCompareSeekGarbledAbnormalData       = "GARBLED"
	DataCompareSeekUncommonWordsAbnormalData = "UNCOMMON_WORDS"
)
