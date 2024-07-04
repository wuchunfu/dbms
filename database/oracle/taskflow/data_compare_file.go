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
package taskflow

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/shopspring/decimal"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"github.com/jedib0t/go-pretty/v6/table"

	"github.com/wentaojin/dbms/utils/stringutil"

	"github.com/wentaojin/dbms/model"
	"github.com/wentaojin/dbms/model/task"
	"github.com/wentaojin/dbms/utils/constant"
)

type DataCompareFile struct {
	Ctx        context.Context `json:"-"`
	Mutex      *sync.Mutex     `json:"-"`
	CompFile   *os.File        `json:"-"`
	CompWriter *bufio.Writer   `json:"-"`
	TaskName   string          `json:"taskName"`
	TaskFlow   string          `json:"taskFlow"`
	OutputDir  string          `json:"outputDir"`
}

func NewDataCompareFile(ctx context.Context,
	taskName, taskFlow, outputDir string) *DataCompareFile {
	return &DataCompareFile{
		Ctx:       ctx,
		TaskName:  taskName,
		TaskFlow:  taskFlow,
		OutputDir: outputDir,
		Mutex:     &sync.Mutex{},
	}
}

func (s *DataCompareFile) InitFile() error {
	err := s.initOutputCompareFile()
	if err != nil {
		return err
	}
	return nil
}

func (s *DataCompareFile) SyncFile() error {
	var (
		err          error
		migrateTasks []*task.DataCompareTask
	)
	// get migrate task tables
	migrateTasks, err = model.GetIDataCompareTaskRW().QueryDataCompareTaskByTaskStatus(s.Ctx, &task.DataCompareTask{
		TaskName: s.TaskName, TaskStatus: constant.TaskDatabaseStatusNotEqual})
	if err != nil {
		return err
	}

	migrateTaskResults, err := model.GetIDataCompareResultRW().FindDataCompareResult(s.Ctx, &task.DataCompareResult{TaskName: s.TaskName})
	if err != nil {
		return err
	}

	if len(migrateTasks) == 0 && len(migrateTaskResults) == 0 {
		// fmt.Printf("the data compare task all of the table records are equal, current not exist not equal table records.\n")
		return errors.New(constant.TaskDatabaseStatusEqual)
	}

	tableTaskM := make(map[string]string)
	compareMethodM := make(map[string]string)
	for _, mt := range migrateTasks {
		tableTaskM[stringutil.StringBuilder(mt.SchemaNameS, constant.StringSeparatorAite, mt.TableNameS)] = stringutil.StringBuilder(mt.SchemaNameT, constant.StringSeparatorAite, mt.TableNameT)
		compareMethodM[stringutil.StringBuilder(mt.SchemaNameS, constant.StringSeparatorAite, mt.TableNameS)] = mt.CompareMethod
	}

	for k, v := range tableTaskM {
		keySli := stringutil.StringSplit(k, constant.StringSeparatorAite)
		valSli := stringutil.StringSplit(v, constant.StringSeparatorAite)

		if strings.EqualFold(compareMethodM[k], constant.DataCompareMethodCheckRows) {
			var sqlComp strings.Builder

			sqlComp.WriteString("/*\n")
			sqlComp.WriteString(" database schema table compare row counts\n")
			wt := table.NewWriter()
			wt.SetStyle(table.StyleLight)
			wt.AppendHeader(table.Row{"#", "TASK_NAME", "TASK_FLOW", "TABLE_NAME_S", "TABLE_NAME_T", "COMPARE_METHOD", "SUGGEST"})

			var rowsCountsS, rowsCountsT decimal.Decimal

			for _, res := range migrateTaskResults {
				if strings.EqualFold(keySli[0], res.SchemaNameS) && strings.EqualFold(keySli[1], res.TableNameS) &&
					strings.EqualFold(res.FixStmtType, constant.DataCompareFixStmtTypeRows) {
					//sample: rowCountsS:500 rowCountsT:600
					rows := stringutil.StringSplit(res.FixDetailT, " ")
					rowsS := rows[0]
					rowsT := rows[1]

					resultS, err := decimal.NewFromString(stringutil.StringSplit(rowsS, ":")[1])
					if err != nil {
						return fmt.Errorf("parse the database source rowcounts failed: %v", err)
					}

					resultT, err := decimal.NewFromString(stringutil.StringSplit(rowsT, ":")[1])
					if err != nil {
						return fmt.Errorf("parse the database target rowcounts failed: %v", err)
					}

					rowsCountsS.Add(resultS)
					rowsCountsT.Add(resultT)
				}
			}

			wt.AppendRows([]table.Row{
				{"Schema", s.TaskName, s.TaskFlow,
					fmt.Sprintf("%s.%s: %v", keySli[0], keySli[1], rowsCountsS.String()),
					fmt.Sprintf("%s.%s: %v", valSli[0], valSli[1], rowsCountsT.String()),
					compareMethodM[k], "Row Counts Difference, Please Verify Rows"},
			})
			sqlComp.WriteString(wt.Render() + "\n")
			sqlComp.WriteString("*/\n")

			if !strings.EqualFold(sqlComp.String(), "") {
				_, err = s.writeCompareFile(sqlComp.String())
				if err != nil {
					return err
				}
			}

			continue
		}

		var sqlComp strings.Builder

		sqlComp.WriteString("/*\n")
		sqlComp.WriteString(" database schema table compare row verify\n")
		wt := table.NewWriter()
		wt.SetStyle(table.StyleLight)
		wt.AppendHeader(table.Row{"#", "TASK_NAME", "TASK_FLOW", "TABLE_NAME_S", "TABLE_NAME_T", "COMPARE_METHOD", "SUGGEST"})
		wt.AppendRows([]table.Row{
			{"Schema", s.TaskName, s.TaskFlow,
				fmt.Sprintf("%s.%s", keySli[0], keySli[1]),
				fmt.Sprintf("%s.%s", valSli[0], valSli[1]),
				compareMethodM[k], "Row Verify Difference, Please Fixed SQL"},
		})
		sqlComp.WriteString(wt.Render() + "\n")
		sqlComp.WriteString("*/\n")

		var delSqls, addSqls []string
		for _, res := range migrateTaskResults {
			if strings.EqualFold(keySli[0], res.SchemaNameS) && strings.EqualFold(keySli[1], res.TableNameS) {
				desDetails, err := stringutil.Decrypt(res.FixDetailT, []byte(constant.DefaultDataEncryptDecryptKey))
				if err != nil {
					return err
				}
				if !strings.EqualFold(desDetails, "") {
					switch res.FixStmtType {
					case constant.DataCompareFixStmtTypeDelete:
						delSqls = append(delSqls, desDetails)
					case constant.DataCompareFixStmtTypeInsert:
						addSqls = append(addSqls, desDetails)
					default:
						return fmt.Errorf("the data compare statement type is invalid: %v", res.FixStmtType)
					}
				}
			}
		}

		if len(delSqls) > 0 {
			sqlComp.WriteString(stringutil.StringJoin(delSqls, "\n") + "\n")
		}

		if len(addSqls) > 0 {
			sqlComp.WriteString(stringutil.StringJoin(addSqls, "\n") + "\n")
		}

		if !strings.EqualFold(sqlComp.String(), "") {
			_, err = s.writeCompareFile(sqlComp.String())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (s *DataCompareFile) writeCompareFile(str string) (int, error) {
	s.Mutex.Lock()
	defer s.Mutex.Unlock()
	return s.CompWriter.WriteString(str)
}

func (s *DataCompareFile) initOutputCompareFile() error {
	outCompFile, err := os.OpenFile(filepath.Join(s.OutputDir, fmt.Sprintf("data_compare_%s.sql", stringutil.StringLower(s.TaskName))), os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	s.CompWriter, s.CompFile = bufio.NewWriter(outCompFile), outCompFile
	return nil
}

func (s *DataCompareFile) Close() error {
	if s.CompFile != nil {
		err := s.CompWriter.Flush()
		if err != nil {
			return err
		}
		err = s.CompFile.Close()
		if err != nil {
			return err
		}
	}
	return nil
}
