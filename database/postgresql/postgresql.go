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
package postgresql

import (
	"context"
	"database/sql"
	"strings"
	"time"

	_ "github.com/lib/pq"

	"fmt"

	"github.com/wentaojin/dbms/model/datasource"
	"github.com/wentaojin/dbms/utils/stringutil"
)

type Database struct {
	Ctx         context.Context
	DBConn      *sql.DB
	CallTimeout int64 // unit: seconds, sql execute timeout
}

func NewDatabase(ctx context.Context, datasource *datasource.Datasource, callTimeout int64) (*Database, error) {
	var (
		connString string
		err        error
	)
	if strings.EqualFold(datasource.DbName, "") {
		connString = fmt.Sprintf("postgres://%s:%s@%s:%d/postgres", datasource.Username, datasource.Password, datasource.Host, datasource.Port)
	} else {
		connString = fmt.Sprintf("postgres://%s:%s@%s:%d/%s", datasource.Username, datasource.Password, datasource.Host, datasource.Port, datasource.DbName)
	}

	if strings.EqualFold(datasource.ConnectParams, "") {
		connString = fmt.Sprintf("%s?sslmode=disable&client_encoding=%s", connString, datasource.ConnectCharset)
	} else {
		connString = fmt.Sprintf("%s?sslmode=disable&client_encoding=%s&%s", connString, datasource.ConnectCharset, datasource.ConnectParams)
	}

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("error on open postgresql database connection: %v", err)
	}

	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("error on ping postgresql database connection: %v", err)
	}
	return &Database{Ctx: ctx, DBConn: db, CallTimeout: callTimeout}, nil
}

func (d *Database) PingDatabaseConnection() error {
	err := d.DBConn.Ping()
	if err != nil {
		return fmt.Errorf("error on ping postgresql database connection:%v", err)
	}
	return nil
}

func (d *Database) PrepareContext(ctx context.Context, sqlStr string) (*sql.Stmt, error) {
	return d.DBConn.PrepareContext(ctx, sqlStr)
}

func (d *Database) QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error) {
	return d.DBConn.QueryContext(ctx, query, args...)
}

func (d *Database) BeginTxn(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error) {
	return d.DBConn.BeginTx(ctx, opts)
}

func (d *Database) CommitTxn(txn *sql.Tx) error {
	return txn.Commit()
}

func (d *Database) Transaction(ctx context.Context, opts *sql.TxOptions, fns []func(ctx context.Context, tx *sql.Tx) error) error {
	tx, err := d.BeginTxn(ctx, opts)
	if err != nil {
		return err
	}
	for _, fn := range fns {
		if err = fn(ctx, tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	err = d.CommitTxn(tx)
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (d *Database) ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error) {
	return d.DBConn.ExecContext(ctx, query, args...)
}

func (d *Database) GeneralQuery(query string, args ...any) ([]string, []map[string]string, error) {
	var (
		columns []string
		results []map[string]string
	)

	deadline := time.Now().Add(time.Duration(d.CallTimeout) * time.Second)

	ctx, cancel := context.WithDeadline(d.Ctx, deadline)
	defer cancel()

	rows, err := d.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, nil, err
	}
	defer rows.Close()

	// general query, automatic get column name
	columns, err = rows.Columns()
	if err != nil {
		return columns, results, fmt.Errorf("query rows.Columns failed, sql: [%v], error: [%v]", query, err)
	}

	values := make([][]byte, len(columns))
	scans := make([]interface{}, len(columns))
	for i := range values {
		scans[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scans...)
		if err != nil {
			return columns, results, fmt.Errorf("query rows.Scan failed, sql: [%v], error: [%v]", query, err)
		}

		row := make(map[string]string)
		for k, v := range values {
			if v == nil {
				row[columns[k]] = "NULLABLE"
			} else {
				// Handling empty string and other values, the return value output string
				row[columns[k]] = stringutil.BytesToString(v)
			}
		}
		results = append(results, row)
	}

	if err = rows.Err(); err != nil {
		return columns, results, fmt.Errorf("query rows.Next failed, sql: [%v], error: [%v]", query, err.Error())
	}
	return columns, results, nil
}

func (d *Database) Close() error {
	return d.DBConn.Close()
}
