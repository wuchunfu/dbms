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
package oracle

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/wentaojin/dbms/utils/constant"

	"github.com/wentaojin/dbms/utils/stringutil"
)

func (d *Database) GetDatabasePartitionTable(schemaName string) ([]string, error) {
	_, res, err := d.GeneralQuery(fmt.Sprintf(`SELECT TABLE_NAME AS TABLE_NAME
	FROM DBA_TABLES
 WHERE PARTITIONED = 'YES'
   AND OWNER = '%s'`, schemaName))
	if err != nil {
		return []string{}, err
	}

	var tables []string
	for _, r := range res {
		tables = append(tables, r["TABLE_NAME"])
	}
	return tables, nil
}

func (d *Database) GetDatabaseTemporaryTable(schemaName string) ([]string, error) {
	_, res, err := d.GeneralQuery(fmt.Sprintf(`SELECT
	TABLE_NAME AS TABLE_NAME
  FROM DBA_TABLES
 WHERE TEMPORARY = 'Y'
   AND OWNER = '%s'`, schemaName))
	if err != nil {
		return []string{}, err
	}

	var tables []string
	for _, r := range res {
		tables = append(tables, r["TABLE_NAME"])
	}
	return tables, nil
}

func (d *Database) GetDatabaseClusteredTable(schemaName string) ([]string, error) {
	_, res, err := d.GeneralQuery(fmt.Sprintf(`SELECT 
    TABLE_NAME AS TABLE_NAME
  FROM dba_tables
 WHERE CLUSTER_NAME IS NOT NULL
   AND OWNER = '%s'`, schemaName))
	if err != nil {
		return []string{}, err
	}

	var tables []string
	for _, r := range res {
		tables = append(tables, r["TABLE_NAME"])
	}
	return tables, nil
}

func (d *Database) GetDatabaseMaterializedView(schemaName string) ([]string, error) {
	_, res, err := d.GeneralQuery(fmt.Sprintf(`SELECT
	OWNER,
	MVIEW_NAME
FROM
	DBA_MVIEWS
WHERE
	OWNER = '%s'`, schemaName))
	if err != nil {
		return []string{}, err
	}

	var tables []string
	for _, r := range res {
		tables = append(tables, r["MVIEW_NAME"])
	}
	return tables, nil
}

func (d *Database) GetDatabaseTableType(schemaName string) (map[string]string, error) {
	tableTypeMap := make(map[string]string)

	_, res, err := d.GeneralQuery(fmt.Sprintf(`SELECT
	F.TABLE_NAME,
	(
	CASE
		WHEN F.CLUSTER_NAME IS NOT NULL THEN 'CLUSTERED'
		WHEN (
		SELECT
			COUNT(1)
		FROM
			DBA_MVIEWS V
		WHERE
			V.OWNER = F.OWNER
			AND F.TABLE_NAME = V.MVIEW_NAME) = 1 THEN 'MATERIALIZED VIEW'
		ELSE
		CASE
			WHEN F.IOT_TYPE = 'IOT' THEN
			CASE
				WHEN T.IOT_TYPE != 'IOT' THEN T.IOT_TYPE
				ELSE 'IOT'
			END
			ELSE
				CASE	
						WHEN F.PARTITIONED = 'YES' THEN 'PARTITIONED'
				ELSE
					CASE
							WHEN F.TEMPORARY = 'Y' THEN 
							DECODE(F.DURATION, 'SYS$SESSION', 'SESSION TEMPORARY', 'SYS$TRANSACTION', 'TRANSACTION TEMPORARY')
					ELSE 'HEAP'
				END
			END
		END
	END ) TABLE_TYPE
FROM
	(
	SELECT
		TMP.OWNER,
		TMP.TABLE_NAME,
		TMP.CLUSTER_NAME,
		TMP.PARTITIONED,
		TMP.TEMPORARY,
		TMP.DURATION,
		TMP.IOT_TYPE
	FROM
		DBA_TABLES TMP,
		DBA_TABLES W
	WHERE
		TMP.OWNER = W.OWNER
		AND TMP.TABLE_NAME = W.TABLE_NAME
		AND TMP.OWNER = '%s'
		AND (W.IOT_TYPE IS NULL
			OR W.IOT_TYPE = 'IOT')) F
LEFT JOIN (
	SELECT
		OWNER,
		IOT_NAME,
		IOT_TYPE
	FROM
		DBA_TABLES
	WHERE
		OWNER = '%s')T 
ON
	F.OWNER = T.OWNER
	AND F.TABLE_NAME = T.IOT_NAME`, schemaName, schemaName))
	if err != nil {
		return tableTypeMap, err
	}
	if len(res) == 0 {
		return tableTypeMap, fmt.Errorf("oracle schema [%s] get table type result null, but can't be null", schemaName)
	}

	for _, r := range res {
		if len(r) > 2 || len(r) == 0 || len(r) == 1 {
			return tableTypeMap, fmt.Errorf("oracle schema [%s] table type values should be 2, result: %v", schemaName, r)
		}
		tableTypeMap[r["TABLE_NAME"]] = r["TABLE_TYPE"]
	}

	return tableTypeMap, nil
}

func (d *Database) GetDatabaseCharset() (string, error) {
	_, res, err := d.GeneralQuery(fmt.Sprintf(`SELECT
	USERENV('LANGUAGE') AS LANG
FROM
	DUAL`))
	if err != nil {
		return "", err
	}
	return stringutil.StringSplit(res[0]["LANG"], constant.StringSeparatorDot)[1], nil
}

func (d *Database) GetDatabaseCharsetCollation() (string, string, error) {
	_, res, err := d.GeneralQuery(fmt.Sprintf(`SELECT
	PARAMETER,
	VALUE
FROM
	NLS_DATABASE_PARAMETERS
WHERE
	PARAMETER IN ('NLS_SORT', 'NLS_COMP')`))
	if err != nil {
		return "", "", err
	}
	var (
		nlsComp string
		nlsSort string
	)
	for _, r := range res {
		if strings.EqualFold(r["PARAMETER"], "NLS_COMP") {
			nlsComp = r["VALUE"]
		} else if strings.EqualFold(r["PARAMETER"], "NLS_SORT") {
			nlsSort = r["VALUE"]
		} else {
			return "", "", fmt.Errorf("get database character set collation failed: [%v]", r)
		}
	}
	return nlsComp, nlsSort, nil
}

func (d *Database) GetDatabaseVersion() (string, error) {
	_, res, err := d.GeneralQuery(fmt.Sprintf(`SELECT
	VALUE
FROM
	NLS_DATABASE_PARAMETERS
WHERE
	PARAMETER = 'NLS_RDBMS_VERSION'`))
	if err != nil {
		return res[0]["VALUE"], err
	}
	return res[0]["VALUE"], nil
}

func (d *Database) GetDatabaseTableColumns(schemaName string, tableName string, collation bool) ([]map[string]string, error) {
	var queryStr string
	/*
			1、dataPrecision Accuracy range ORA-01727: numeric precision specifier is out of range (1 to 38)
			2、dataScale Accuracy range ORA-01728: numeric scale specifier is out of range (-84 to 127)
			3、oracle number datatype，view table struct by using command desc tableName
			- number(*,10) -> number(38,10)
			- number(*,0) -> number(38,0)
			- number(*) -> number
			- number -> number
		    - number(x,y) -> number(x,y)
			4、SQL Query
			- number(*,10) -> number(38,10)
			- number(*,0) -> number(38,0)
			- number(*) -> number(38,127)
			- number -> number(38,127)
			- number(x,y) -> number(x,y)
	*/
	if collation {
		queryStr = fmt.Sprintf(`SELECT 
		T.COLUMN_NAME,
	    T.DATA_TYPE,
		T.CHAR_LENGTH,
		NVL(T.CHAR_USED, 'UNKNOWN') CHAR_USED,
	    NVL(T.DATA_LENGTH, 0) AS DATA_LENGTH,
	    DECODE(NVL(TO_CHAR(T.DATA_PRECISION), '*'), '*', '38', TO_CHAR(T.DATA_PRECISION)) AS DATA_PRECISION,
	    DECODE(NVL(TO_CHAR(T.DATA_SCALE), '*'), '*', '127', TO_CHAR(T.DATA_SCALE)) AS DATA_SCALE,
		T.NULLABLE,
	    NVL(S.DATA_DEFAULT, 'NULLSTRING') DATA_DEFAULT,
		DECODE(T.COLLATION, 'USING_NLS_COMP',(SELECT VALUE FROM NLS_DATABASE_PARAMETERS WHERE PARAMETER = 'NLS_COMP'), T.COLLATION) COLLATION,
	    C.COMMENTS
FROM
	DBA_TAB_COLUMNS T,
	DBA_COL_COMMENTS C,
	(
	SELECT
		XS.OWNER,
		XS.TABLE_NAME,
		XS.COLUMN_NAME,
		XS.DATA_DEFAULT
	FROM
		XMLTABLE (
		'/ROWSET/ROW' PASSING (
		SELECT
			DBMS_XMLGEN.GETXMLTYPE (
				Q'[SELECT D.OWNER,D.TABLE_NAME,D.COLUMN_NAME,D.DATA_DEFAULT FROM DBA_TAB_COLUMNS D WHERE D.OWNER = '%s' AND D.TABLE_NAME = '%s']')
		FROM
			DUAL ) COLUMNS OWNER VARCHAR2 (300) PATH 'OWNER',
		TABLE_NAME VARCHAR2(300) PATH 'TABLE_NAME',
		COLUMN_NAME VARCHAR2(300) PATH 'COLUMN_NAME',
		DATA_DEFAULT VARCHAR2(4000) PATH 'DATA_DEFAULT') XS
		) S
WHERE
	T.TABLE_NAME = C.TABLE_NAME
	AND T.COLUMN_NAME = C.COLUMN_NAME
	AND T.OWNER = C.OWNER
	AND C.OWNER = S.OWNER
	AND C.TABLE_NAME = S.TABLE_NAME
	AND C.COLUMN_NAME = S.COLUMN_NAME
	AND T.OWNER = '%s'
	AND T.TABLE_NAME = '%s'
ORDER BY
	T.COLUMN_ID`, schemaName, tableName, schemaName, tableName)
	} else {
		queryStr = fmt.Sprintf(`SELECT 
		T.COLUMN_NAME,
	    T.DATA_TYPE,
		T.CHAR_LENGTH,
		NVL(T.CHAR_USED, 'UNKNOWN') CHAR_USED,
	    NVL(T.DATA_LENGTH, 0) AS DATA_LENGTH,
	    DECODE(NVL(TO_CHAR(T.DATA_PRECISION), '*'), '*', '38', TO_CHAR(T.DATA_PRECISION)) AS DATA_PRECISION,
	    DECODE(NVL(TO_CHAR(T.DATA_SCALE), '*'), '*', '127', TO_CHAR(T.DATA_SCALE)) AS DATA_SCALE,
		T.NULLABLE,
	    NVL(S.DATA_DEFAULT, 'NULLSTRING') DATA_DEFAULT,
	    C.COMMENTS
FROM
	DBA_TAB_COLUMNS T,
	DBA_COL_COMMENTS C,
	(
	SELECT
		XS.OWNER,
		XS.TABLE_NAME,
		XS.COLUMN_NAME,
		XS.DATA_DEFAULT
	FROM
		XMLTABLE (
		'/ROWSET/ROW' PASSING (
		SELECT
			DBMS_XMLGEN.GETXMLTYPE (
				Q'[SELECT D.OWNER,D.TABLE_NAME,D.COLUMN_NAME,D.DATA_DEFAULT FROM DBA_TAB_COLUMNS D WHERE D.OWNER = '%s' AND D.TABLE_NAME = '%s']')
		FROM
			DUAL ) COLUMNS OWNER VARCHAR2 (300) PATH 'OWNER',
		TABLE_NAME VARCHAR2(300) PATH 'TABLE_NAME',
		COLUMN_NAME VARCHAR2(300) PATH 'COLUMN_NAME',
		DATA_DEFAULT VARCHAR2(4000) PATH 'DATA_DEFAULT') XS
		) S
WHERE
	T.TABLE_NAME = C.TABLE_NAME
	AND T.COLUMN_NAME = C.COLUMN_NAME
	AND T.OWNER = C.OWNER
	AND C.OWNER = S.OWNER
	AND C.TABLE_NAME = S.TABLE_NAME
	AND C.COLUMN_NAME = S.COLUMN_NAME
	AND T.OWNER = '%s'
	AND T.TABLE_NAME = '%s'
ORDER BY
	T.COLUMN_ID`, schemaName, tableName, schemaName, tableName)
	}

	_, queryRes, err := d.GeneralQuery(queryStr)
	if err != nil {
		return queryRes, err
	}
	if len(queryRes) == 0 {
		return queryRes, fmt.Errorf("oracle table [%s.%s] column info cann't be null", schemaName, tableName)
	}

	// check constraints notnull
	// search_condition long datatype
	_, condRes, err := d.GeneralQuery(fmt.Sprintf(`SELECT
				COL.COLUMN_NAME,
				CONS.SEARCH_CONDITION
FROM
				DBA_CONS_COLUMNS COL,
				DBA_CONSTRAINTS CONS
WHERE
				COL.OWNER = CONS.OWNER
	AND COL.TABLE_NAME = CONS.TABLE_NAME
	AND COL.CONSTRAINT_NAME = CONS.CONSTRAINT_NAME
	AND CONS.CONSTRAINT_TYPE = 'C'
	AND COL.OWNER = '%s'
	AND COL.TABLE_NAME = '%s'`, schemaName, tableName))
	if err != nil {
		return queryRes, err
	}

	if len(condRes) == 0 {
		return queryRes, nil
	}

	rep, err := regexp.Compile(`(^.*)(?i:IS NOT NULL)`)
	if err != nil {
		return queryRes, fmt.Errorf("check notnull constraint regexp complile failed: %v", err)
	}
	for _, r := range queryRes {
		for _, c := range condRes {
			if r["COLUMN_NAME"] == c["COLUMN_NAME"] && r["NULLABLE"] == "Y" {
				// check constraint -> not null
				if rep.MatchString(c["SEARCH_CONDITION"]) {
					r["NULLABLE"] = "N"
				}
			}
		}
	}
	return queryRes, nil
}

func (d *Database) GetDatabaseTablePrimaryKey(schemaName string, tableName string) ([]map[string]string, error) {
	// for the primary key of an Engine table, you can use the following command to set whether the primary key takes effect.
	// disable the primary key: alter table tableName disable primary key;
	// enable the primary key: alter table tableName enable primary key;
	// primary key status Disabled will not do primary key processing
	queryStr := fmt.Sprintf(`SELECT
	CU.CONSTRAINT_NAME,
	LISTAGG(CU.COLUMN_NAME, ',') WITHIN GROUP(
	ORDER BY CU.POSITION) AS COLUMN_LIST
FROM
	DBA_CONS_COLUMNS CU,
	DBA_CONSTRAINTS AU
WHERE
	CU.CONSTRAINT_NAME = AU.CONSTRAINT_NAME
	AND AU.CONSTRAINT_TYPE = 'P'
	AND AU.STATUS = 'ENABLED'
	AND CU.OWNER = AU.OWNER
	AND CU.TABLE_NAME = AU.TABLE_NAME
	AND AU.TABLE_NAME = '%s'
	AND CU.OWNER = '%s'
GROUP BY
	CU.CONSTRAINT_NAME`,
		tableName,
		schemaName)
	_, res, err := d.GeneralQuery(queryStr)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (d *Database) GetDatabaseTableUniqueKey(schemaName string, tableName string) ([]map[string]string, error) {
	queryStr := fmt.Sprintf(`SELECT
	CU.CONSTRAINT_NAME,
	AU.INDEX_NAME,
	LISTAGG(CU.COLUMN_NAME, ',') WITHIN GROUP(
	ORDER BY CU.POSITION) AS COLUMN_LIST
FROM
	DBA_CONS_COLUMNS CU,
	DBA_CONSTRAINTS AU
WHERE
	CU.CONSTRAINT_NAME = AU.CONSTRAINT_NAME
	AND CU.OWNER = AU.OWNER
	AND CU.TABLE_NAME = AU.TABLE_NAME
	AND AU.CONSTRAINT_TYPE = 'U'
	AND AU.STATUS = 'ENABLED'
	AND AU.TABLE_NAME = '%s'
	AND CU.OWNER = '%s'
GROUP BY
	CU.CONSTRAINT_NAME,
	AU.INDEX_NAME`,
		tableName,
		schemaName)
	_, res, err := d.GeneralQuery(queryStr)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (d *Database) GetDatabaseTableForeignKey(schemaName string, tableName string) ([]map[string]string, error) {
	queryStr := fmt.Sprintf(`WITH TEMP1 AS
 (
SELECT
	T1.OWNER,
	T1.R_OWNER,
	T1.CONSTRAINT_NAME,
	T1.R_CONSTRAINT_NAME,
	T1.DELETE_RULE,
	LISTAGG(A1.COLUMN_NAME, ',') WITHIN GROUP(ORDER BY A1.POSITION) AS COLUMN_LIST
FROM
	DBA_CONSTRAINTS T1,
	DBA_CONS_COLUMNS A1
WHERE
	T1.CONSTRAINT_NAME = A1.CONSTRAINT_NAME
	AND T1.TABLE_NAME = '%s'
	AND T1.OWNER = '%s'
	AND T1.STATUS = 'ENABLED'
	AND T1.CONSTRAINT_TYPE = 'R'
GROUP BY
	T1.OWNER,
	T1.R_OWNER,
	T1.CONSTRAINT_NAME,
	T1.R_CONSTRAINT_NAME,
	T1.DELETE_RULE),
TEMP2 AS
 (
SELECT
	T1.OWNER,
	T1.TABLE_NAME,
	T1.CONSTRAINT_NAME,
	LISTAGG(A1.COLUMN_NAME, ',') WITHIN GROUP(ORDER BY A1.POSITION) AS COLUMN_LIST
FROM
	DBA_CONSTRAINTS T1,
	DBA_CONS_COLUMNS A1
WHERE
	T1.CONSTRAINT_NAME = A1.CONSTRAINT_NAME
	AND T1.OWNER = '%s'
		AND T1.STATUS = 'ENABLED'
		AND T1.CONSTRAINT_TYPE = 'P'
	GROUP BY
		T1.OWNER,
		T1.TABLE_NAME,
		T1.R_OWNER,
		T1.CONSTRAINT_NAME)
SELECT
	X.CONSTRAINT_NAME,
	X.COLUMN_LIST,
	X.R_OWNER,
	Y.TABLE_NAME AS RTABLE_NAME,
	Y.COLUMN_LIST AS RCOLUMN_LIST,
	X.DELETE_RULE
FROM
	TEMP1 X,
	TEMP2 Y
WHERE
	X.R_OWNER = Y.OWNER
	AND X.R_CONSTRAINT_NAME = Y.CONSTRAINT_NAME`,
		tableName,
		schemaName,
		schemaName)
	_, res, err := d.GeneralQuery(queryStr)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (d *Database) GetDatabaseTableCheckKey(schemaName string, tableName string) ([]map[string]string, error) {
	queryStr := fmt.Sprintf(`SELECT
	CU.CONSTRAINT_NAME,
	AU.SEARCH_CONDITION
FROM
	DBA_CONS_COLUMNS CU,
	DBA_CONSTRAINTS AU
WHERE
	CU.OWNER = AU.OWNER
	AND CU.TABLE_NAME = AU.TABLE_NAME
	AND CU.CONSTRAINT_NAME = AU.CONSTRAINT_NAME
	AND AU.CONSTRAINT_TYPE = 'C'
	AND AU.STATUS = 'ENABLED'
	AND AU.TABLE_NAME = '%s'
	AND AU.OWNER = '%s'`,
		tableName,
		schemaName)
	_, res, err := d.GeneralQuery(queryStr)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (d *Database) GetDatabaseTableNormalIndex(schemaName string, tableName string) ([]map[string]string, error) {
	queryStr := fmt.Sprintf(`SELECT
	TEMP.TABLE_NAME,
	TEMP.UNIQUENESS,
	TEMP.INDEX_NAME,
	TEMP.INDEX_TYPE,
	TEMP.ITYP_OWNER,
	TEMP.ITYP_NAME,
	TEMP.PARAMETERS,
	LISTAGG ( TEMP.COLUMN_NAME, ',' ) WITHIN GROUP (ORDER BY TEMP.COLUMN_POSITION) AS COLUMN_LIST
FROM
	(
	SELECT
		T.TABLE_OWNER,
		T.TABLE_NAME,
		I.UNIQUENESS,
		T.INDEX_NAME,
		I.INDEX_TYPE,
		NVL(I.ITYP_OWNER, '') ITYP_OWNER,
		NVL(I.ITYP_NAME, '') ITYP_NAME,
		NVL(I.PARAMETERS, '') PARAMETERS,
		DECODE((SELECT
	COUNT(1) 
FROM
	DBA_IND_EXPRESSIONS S
WHERE
	S.TABLE_OWNER = T.TABLE_OWNER
	AND S.TABLE_NAME = T.TABLE_NAME
	AND S.INDEX_OWNER = T.INDEX_OWNER
	AND S.INDEX_NAME = T.INDEX_NAME
	AND S.COLUMN_POSITION = T.COLUMN_POSITION), 0, T.COLUMN_NAME, (SELECT
	XS.COLUMN_EXPRESSION
FROM
	XMLTABLE (
		'/ROWSET/ROW' PASSING ( SELECT DBMS_XMLGEN.GETXMLTYPE ( 
				Q'[SELECT
	S.INDEX_OWNER,
	S.INDEX_NAME,
	S.COLUMN_EXPRESSION,
	S.COLUMN_POSITION
FROM
	DBA_IND_EXPRESSIONS S
WHERE
	S.TABLE_OWNER = '%s'
	AND S.TABLE_NAME = '%s']' ) FROM DUAL ) COLUMNS 
	INDEX_OWNER VARCHAR2 (30) PATH 'INDEX_OWNER',
	INDEX_NAME VARCHAR2 (30) PATH 'INDEX_NAME',
	COLUMN_POSITION VARCHAR2 (30) PATH 'COLUMN_POSITION',
	COLUMN_EXPRESSION VARCHAR2 (4000) 
	) XS 
WHERE
XS.INDEX_OWNER = T.INDEX_OWNER
	AND XS.INDEX_NAME = T.INDEX_NAME
AND XS.COLUMN_POSITION = T.COLUMN_POSITION)) COLUMN_NAME,
		T.COLUMN_POSITION
	FROM
		DBA_IND_COLUMNS T,
		DBA_INDEXES I
	WHERE
		T.TABLE_OWNER = I.TABLE_OWNER
		AND T.TABLE_NAME = I.TABLE_NAME
		AND T.INDEX_NAME = I.INDEX_NAME
		AND I.UNIQUENESS = 'NONUNIQUE'
		AND T.TABLE_OWNER = '%s'
		AND T.TABLE_NAME = '%s'
		-- exclude constraint index
		AND NOT EXISTS (
		SELECT
			1
		FROM
			DBA_CONSTRAINTS C
		WHERE
			C.OWNER = I.TABLE_OWNER
			AND C.TABLE_NAME = I.TABLE_NAME
			AND C.INDEX_NAME = I.INDEX_NAME)) TEMP
GROUP BY
		TEMP.TABLE_OWNER,
		TEMP.TABLE_NAME,
		TEMP.UNIQUENESS,
		TEMP.INDEX_NAME,
		TEMP.INDEX_TYPE,
		TEMP.ITYP_OWNER,
		TEMP.ITYP_NAME,
		TEMP.PARAMETERS`,
		schemaName,
		tableName,
		schemaName,
		tableName)
	_, res, err := d.GeneralQuery(queryStr)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (d *Database) GetDatabaseTableUniqueIndex(schemaName string, tableName string) ([]map[string]string, error) {
	queryStr := fmt.Sprintf(`SELECT
	TEMP.TABLE_NAME,
	TEMP.UNIQUENESS,
	TEMP.INDEX_NAME,
	TEMP.INDEX_TYPE,
	TEMP.ITYP_OWNER,
	TEMP.ITYP_NAME,
	TEMP.PARAMETERS,
	LISTAGG ( TEMP.COLUMN_NAME, ',' ) WITHIN GROUP ( ORDER BY TEMP.COLUMN_POSITION ) AS COLUMN_LIST
FROM
	(
	SELECT
		I.TABLE_OWNER,
		I.TABLE_NAME,
		I.UNIQUENESS,
		I.INDEX_NAME,
		I.INDEX_TYPE,
		NVL(I.ITYP_OWNER, '') ITYP_OWNER,
		NVL(I.ITYP_NAME, '') ITYP_NAME,
		NVL(I.PARAMETERS, '') PARAMETERS,
		DECODE((SELECT
	COUNT(1) 
FROM
	DBA_IND_EXPRESSIONS S
WHERE
	S.TABLE_OWNER = T.TABLE_OWNER
	AND S.TABLE_NAME = T.TABLE_NAME
	AND S.INDEX_OWNER = T.INDEX_OWNER
	AND S.INDEX_NAME = T.INDEX_NAME
	AND S.COLUMN_POSITION = T.COLUMN_POSITION), 0, T.COLUMN_NAME, (SELECT
	XS.COLUMN_EXPRESSION
FROM
	XMLTABLE (
		'/ROWSET/ROW' PASSING ( SELECT DBMS_XMLGEN.GETXMLTYPE ( 
				Q'[SELECT
	S.INDEX_OWNER,
	S.INDEX_NAME,
	S.COLUMN_EXPRESSION,
	S.COLUMN_POSITION
FROM
	DBA_IND_EXPRESSIONS S
WHERE
	S.TABLE_OWNER = '%s'
	AND S.TABLE_NAME = '%s']' ) FROM DUAL ) COLUMNS 
	INDEX_OWNER VARCHAR2 (30) PATH 'INDEX_OWNER',
	INDEX_NAME VARCHAR2 (30) PATH 'INDEX_NAME',
	COLUMN_POSITION VARCHAR2 (30) PATH 'COLUMN_POSITION',
	COLUMN_EXPRESSION VARCHAR2 (4000) 
	) XS 
WHERE
XS.INDEX_OWNER = T.INDEX_OWNER
	AND XS.INDEX_NAME = T.INDEX_NAME
AND XS.COLUMN_POSITION = T.COLUMN_POSITION)) COLUMN_NAME,
		T.COLUMN_POSITION
	FROM
		DBA_INDEXES I,
		DBA_IND_COLUMNS T
	WHERE
		I.INDEX_NAME = T.INDEX_NAME
		AND I.TABLE_OWNER = T.TABLE_OWNER
		AND I.TABLE_NAME = T.TABLE_NAME
		AND I.UNIQUENESS = 'UNIQUE'
		AND I.TABLE_OWNER = '%s'
		AND I.TABLE_NAME = '%s'
		-- exclude primary and unique 
		AND NOT EXISTS (
		SELECT
			1
		FROM
			DBA_CONSTRAINTS C
		WHERE
			I.INDEX_NAME = C.INDEX_NAME
			AND I.TABLE_OWNER = C.OWNER
			AND I.TABLE_NAME = C.TABLE_NAME
			AND C.CONSTRAINT_TYPE IN ('P', 'U')
	)) TEMP
GROUP BY
		TEMP.TABLE_OWNER,
		TEMP.TABLE_NAME,
		TEMP.UNIQUENESS,
		TEMP.INDEX_NAME,
		TEMP.INDEX_TYPE,
		TEMP.ITYP_OWNER,
		TEMP.ITYP_NAME,
		TEMP.PARAMETERS`,
		schemaName,
		tableName,
		schemaName,
		tableName)
	_, res, err := d.GeneralQuery(queryStr)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (d *Database) GetDatabaseTableComment(schemaName string, tableName string) ([]map[string]string, error) {
	var (
		comments []map[string]string
		err      error
	)
	queryStr := fmt.Sprintf(`SELECT
	TABLE_NAME,
	TABLE_TYPE,
	COMMENTS
FROM
	DBA_TAB_COMMENTS
WHERE
	TABLE_TYPE = 'TABLE'
	AND OWNER= '%s'
	AND TABLE_NAME = '%s'`, schemaName, tableName)
	_, res, err := d.GeneralQuery(queryStr)
	if err != nil {
		return comments, err
	}
	if len(res) > 1 {
		return res, fmt.Errorf("database schema [%s] table [%s] comments exist multiple values: [%v]", schemaName, tableName, res)
	}

	return res, nil
}

func (d *Database) GetDatabaseTableColumnComment(schemaName string, tableName string) ([]map[string]string, error) {
	var queryStr string

	queryStr = fmt.Sprintf(`SELECT
	T.COLUMN_NAME,
	    C.COMMENTS
FROM
	DBA_TAB_COLUMNS T,
	DBA_COL_COMMENTS C
WHERE
	T.TABLE_NAME = C.TABLE_NAME
	AND T.COLUMN_NAME = C.COLUMN_NAME
	AND T.OWNER = C.OWNER
	AND T.OWNER = '%s'
	AND T.TABLE_NAME = '%s'
ORDER BY
	T.COLUMN_ID`,
		schemaName,
		tableName)
	_, queryRes, err := d.GeneralQuery(queryStr)
	if err != nil {
		return queryRes, err
	}
	if len(queryRes) == 0 {
		return queryRes, fmt.Errorf("database table [%s.%s] column info cann't be null", schemaName, tableName)
	}
	return queryRes, nil
}

func (d *Database) GetDatabaseSchemaCollation(schemaName string) (string, error) {
	queryStr := fmt.Sprintf(`SELECT
	DECODE(DEFAULT_COLLATION,'USING_NLS_COMP',( SELECT VALUE FROM NLS_DATABASE_PARAMETERS WHERE PARAMETER = 'NLS_COMP'), DEFAULT_COLLATION) DEFAULT_COLLATION
FROM
	DBA_USERS
WHERE
	USERNAME = '%s'`, schemaName)
	_, res, err := d.GeneralQuery(queryStr)
	if err != nil {
		return "", err
	}
	return res[0]["DEFAULT_COLLATION"], nil
}

func (d *Database) GetDatabaseSchemaTableCollation(schemaName, schemaCollation string) (map[string]string, error) {
	var err error

	tablesMap := make(map[string]string)

	_, res, err := d.GeneralQuery(fmt.Sprintf(`SELECT
	TABLE_NAME,
	DEFAULT_COLLATION
FROM
	DBA_TABLES
WHERE
	OWNER = '%s'
	AND (IOT_TYPE IS NULL
		OR IOT_TYPE = 'IOT')`, schemaName))
	if err != nil {
		return tablesMap, err
	}
	for _, r := range res {
		if strings.ToUpper(r["DEFAULT_COLLATION"]) == constant.OracleDatabaseUserTableColumnDefaultCollation {
			tablesMap[strings.ToUpper(r["TABLE_NAME"])] = strings.ToUpper(schemaCollation)
		} else {
			tablesMap[strings.ToUpper(r["TABLE_NAME"])] = strings.ToUpper(r["DEFAULT_COLLATION"])
		}
	}

	return tablesMap, nil
}

func (d *Database) GetDatabaseTableOriginStruct(schemaName, tableName, tableType string) (string, error) {
	ddlSql := fmt.Sprintf(`SELECT REGEXP_REPLACE(REPLACE(DBMS_METADATA.GET_DDL('%s','%s','%s'),'"'), ',\s*''NLS_CALENDAR=GREGORIAN''', '') ORIGIN_DDL FROM DUAL`, tableType, tableName, schemaName)

	rows, err := d.QueryContext(d.Ctx, ddlSql)
	if err != nil {
		return "", err
	}
	defer rows.Close()

	var ddl string
	for rows.Next() {
		if err = rows.Scan(&ddl); err != nil {
			return ddl, fmt.Errorf("get database schema table origin ddl scan failed: %v", err)
		}
	}

	if err = rows.Close(); err != nil {
		return ddl, fmt.Errorf("get database schema table origin ddl close failed: %v", err)
	}
	if err = rows.Err(); err != nil {
		return ddl, fmt.Errorf("get database schema table origin ddl Err failed: %v", err)
	}

	return ddl, nil
}
