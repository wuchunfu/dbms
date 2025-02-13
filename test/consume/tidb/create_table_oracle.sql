CREATE TABLE c_int (
    id          NUMBER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    c_tinyint   NUMBER(3),
    c_smallint  NUMBER(5),
    c_mediumint NUMBER(7),
    c_int       NUMBER(10),
    c_bigint    NUMBER(19)
);

CREATE TABLE c_unsigned_int (
    id                  NUMBER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    c_unsigned_tinyint  NUMBER(3) CHECK (c_unsigned_tinyint >= 0),
    c_unsigned_smallint NUMBER(5) CHECK (c_unsigned_smallint >= 0),
    c_unsigned_mediumint NUMBER(7) CHECK (c_unsigned_mediumint >= 0),
    c_unsigned_int      NUMBER(10) CHECK (c_unsigned_int >= 0),
    c_unsigned_bigint   NUMBER(20) CHECK (c_unsigned_bigint >= 0)
);

CREATE TABLE c_text (
    id           NUMBER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    c_tinytext   VARCHAR2(255),
    c_text       CLOB,
    c_mediumtext CLOB,
    c_longtext   CLOB
);

CREATE TABLE c_char_binary (
    id           NUMBER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    c_char       CHAR(16),
    c_varchar    VARCHAR2(16),
    c_binary     RAW(16),
    c_varbinary  RAW(16)
);

CREATE TABLE c_blob (
    id           NUMBER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    c_tinyblob   BLOB,
    c_blob       BLOB,
    c_mediumblob BLOB,
    c_longblob   BLOB
);

CREATE TABLE c_time (
    id          NUMBER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    c_date      DATE,
    c_datetime  TIMESTAMP,
    c_timestamp TIMESTAMP,
    c_time      INTERVAL DAY TO SECOND,
    c_year      NUMBER(4)
);

CREATE TABLE c_real (
    id            NUMBER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    c_float       BINARY_FLOAT,
    c_double      BINARY_DOUBLE,
    c_decimal     NUMBER,
    c_decimal_2   NUMBER(10, 4)
);

CREATE TABLE c_unsigned_real (
    id                   NUMBER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    c_unsigned_float     BINARY_FLOAT CHECK (c_unsigned_float >= 0),
    c_unsigned_double    BINARY_DOUBLE CHECK (c_unsigned_double >= 0),
    c_unsigned_decimal   NUMBER CHECK (c_unsigned_decimal >= 0),
    c_unsigned_decimal_2 NUMBER(10, 4) CHECK (c_unsigned_decimal_2 >= 0)
);

CREATE TABLE c_other_datatype (
    id         NUMBER GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    c_enum     VARCHAR2(1) CHECK (c_enum IN ('a', 'b', 'c')),
    c_set      VARCHAR2(10),
    c_bit      RAW(8),
    c_json     CLOB
);

CREATE TABLE c_update00 (
	id NUMBER primary key,
	val VARCHAR2(16)
);

CREATE TABLE c_update01 (
	a NUMBER PRIMARY KEY,
	b NUMBER
);

CREATE TABLE c_update02 (
	a NUMBER PRIMARY KEY,
	b NUMBER
);


CREATE TABLE c_partition_hash (
    a INT,
    CONSTRAINT pk PRIMARY KEY (a)
) PARTITION BY HASH (a)
PARTITIONS 5;


CREATE TABLE c_partition_range (
    a INT PRIMARY KEY
) PARTITION BY RANGE (a) (
    PARTITION p0 VALUES LESS THAN (6),
    PARTITION p1 VALUES LESS THAN (11),
    PARTITION p2 VALUES LESS THAN (21)
);

CREATE TABLE c_clustered_t0 (
    a INT PRIMARY KEY,
    b INT
);

CREATE TABLE c_clustered_t1 (
    a INT PRIMARY KEY,
    b INT,
    CONSTRAINT unique_b UNIQUE (b)
);

CREATE TABLE c_clustered_t2 (
    a CHAR(10) PRIMARY KEY,
    b INT
);

CREATE TABLE c_clustered_t3 (
    a INT CONSTRAINT unique_a UNIQUE NOT NULL
);

CREATE TABLE c_nonclustered_t0 (
    a INT PRIMARY KEY,
    b INT
);

CREATE TABLE c_nonclustered_t1 (
    a INT PRIMARY KEY,
    b INT CONSTRAINT unique_k UNIQUE
);

CREATE TABLE c_nonclustered_t2 (
    a VARCHAR2(255) PRIMARY KEY
);

CREATE TABLE c_nonclustered_t3 (
    a INT CONSTRAINT unique_q UNIQUE NOT NULL
);

CREATE TABLE c_store_gen_col (
    col1 INT NOT NULL,
    col2 VARCHAR2(255) NOT NULL,
    col3 INT GENERATED ALWAYS AS (col1 * 2) NOT NULL,
    CONSTRAINT UK_C UNIQUE (col1, col2, col3)
);

CREATE TABLE c_vitual_gen_col (
    col1 INT NOT NULL,
    col2 VARCHAR2(255) NOT NULL,
    col3 INT GENERATED ALWAYS AS (col1 * 2) VIRTUAL NOT NULL,
    CONSTRAINT UK UNIQUE (col1, col2, col3)
);

CREATE TABLE c_compression_t (
    id                  NUMBER PRIMARY KEY,    
    c_tinyint           NUMBER(3) NULL,
    c_smallint          NUMBER(5) NULL,
    c_mediumint         NUMBER(8) NULL,
    c_int               NUMBER NULL,
    c_bigint            NUMBER(19) NULL,
    c_unsigned_tinyint  NUMBER(3) CHECK (c_unsigned_tinyint >= 0) NULL,
    c_unsigned_smallint NUMBER(5) CHECK (c_unsigned_smallint >= 0) NULL,
    c_unsigned_mediumint NUMBER(8) CHECK (c_unsigned_mediumint >= 0) NULL,
    c_unsigned_int      NUMBER CHECK (c_unsigned_int >= 0) NULL,
    c_unsigned_bigint   NUMBER(19) CHECK (c_unsigned_bigint >= 0) NULL,
    c_float             BINARY_FLOAT NULL,
    c_double            BINARY_DOUBLE NULL,
    c_decimal           NUMBER NULL,
    c_decimal_2         NUMBER(10, 4) NULL,
    c_unsigned_float    BINARY_FLOAT CHECK (c_unsigned_float >= 0) NULL,
    c_unsigned_double   BINARY_DOUBLE CHECK (c_unsigned_double >= 0) NULL,
    c_unsigned_decimal  NUMBER CHECK (c_unsigned_decimal >= 0) NULL,
    c_unsigned_decimal_2 NUMBER(10, 4) CHECK (c_unsigned_decimal_2 >= 0) NULL,
    c_date              DATE NULL,
    c_datetime          TIMESTAMP NULL,
    c_timestamp         TIMESTAMP NULL,
    c_time              INTERVAL DAY TO SECOND NULL,
    c_year              NUMBER(4) NULL,
    c_tinytext          CLOB NULL,
    c_text              CLOB NULL,
    c_mediumtext        CLOB NULL,
    c_longtext          CLOB NULL,
    c_tinyblob          BLOB NULL,
    c_blob              BLOB NULL,
    c_mediumblob        BLOB NULL,
    c_longblob          BLOB NULL,
    c_char              CHAR(16) NULL,
    c_varchar           VARCHAR2(16) NULL,
    c_binary            RAW(16) NULL,
    c_varbinary         RAW(16) NULL,
    c_bit               RAW(8) NULL,  
    c_json              CLOB NULL
);

CREATE TABLE c_gen_store_t (
    a NUMBER,
    b NUMBER GENERATED ALWAYS AS (a + 1),
    CONSTRAINT pk_a PRIMARY KEY (b)
);

CREATE TABLE c_gen_virtual_t (
    a INT,
    b INT GENERATED ALWAYS AS (a + 1) VIRTUAL NOT NULL,
    c INT NOT NULL,
    CONSTRAINT idx1 UNIQUE (b),
    CONSTRAINT idx2 UNIQUE (c)
);

CREATE TABLE c_random (
    id NUMBER(19) PRIMARY KEY,
    data NUMBER
);

CREATE TABLE c_savepoint_t (
    id NUMBER PRIMARY KEY,
    a NUMBER
);
create index idx_a on c_savepoint_t(a);

CREATE TABLE c_multi_data_type (
    id           NUMBER PRIMARY KEY,
    t_boolean    NUMBER(1) CHECK (t_boolean IN (0, 1)),
    t_bigint     NUMBER(19),
    t_double     BINARY_DOUBLE,
    t_decimal    NUMBER(38, 19),
    t_bit        RAW(8),
    t_date       DATE,
    t_datetime   TIMESTAMP,
    t_timestamp  TIMESTAMP NULL,
    t_time       INTERVAL DAY TO SECOND,
    t_year       NUMBER(4),
    t_char       CHAR(3),
    t_varchar    VARCHAR2(10),
    t_blob       BLOB,
    t_text       CLOB,
    t_json       CLOB -- Use CLOB for JSON in older versions of Oracle
);