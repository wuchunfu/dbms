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
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/wentaojin/dbms/errconcurrent"

	"golang.org/x/sync/errgroup"

	"github.com/google/uuid"
	"github.com/wentaojin/dbms/database"
	"github.com/wentaojin/dbms/logger"
	"github.com/wentaojin/dbms/model"
	"github.com/wentaojin/dbms/model/datasource"
	"github.com/wentaojin/dbms/model/rule"
	"github.com/wentaojin/dbms/model/task"
	"github.com/wentaojin/dbms/proto/pb"
	"github.com/wentaojin/dbms/utils/constant"
	"github.com/wentaojin/dbms/utils/stringutil"
	"go.uber.org/zap"
)

type StmtMigrateTask struct {
	Ctx         context.Context
	Task        *task.Task
	DatasourceS *datasource.Datasource
	DatasourceT *datasource.Datasource
	TaskParams  *pb.StatementMigrateParam
}

func (dmt *StmtMigrateTask) Start() error {
	schemaTaskTime := time.Now()
	logger.Info("stmt migrate task get schema route",
		zap.String("task_name", dmt.Task.TaskName), zap.String("task_mode", dmt.Task.TaskMode), zap.String("task_flow", dmt.Task.TaskFlow))
	schemaRoute, err := model.GetIMigrateSchemaRouteRW().GetSchemaRouteRule(dmt.Ctx, &rule.SchemaRouteRule{TaskName: dmt.Task.TaskName})
	if err != nil {
		return err
	}

	logger.Info("stmt migrate task init database connection",
		zap.String("task_name", dmt.Task.TaskName), zap.String("task_mode", dmt.Task.TaskMode), zap.String("task_flow", dmt.Task.TaskFlow))

	sourceDatasource, err := model.GetIDatasourceRW().GetDatasource(dmt.Ctx, dmt.Task.DatasourceNameS)
	if err != nil {
		return err
	}
	databaseS, err := database.NewDatabase(dmt.Ctx, sourceDatasource, schemaRoute.SchemaNameS)
	if err != nil {
		return err
	}
	defer databaseS.Close()
	databaseT, err := database.NewDatabase(dmt.Ctx, dmt.DatasourceT, "")
	if err != nil {
		return err
	}
	defer databaseT.Close()

	logger.Info("stmt migrate task inspect migrate task",
		zap.String("task_name", dmt.Task.TaskName), zap.String("task_mode", dmt.Task.TaskMode), zap.String("task_flow", dmt.Task.TaskFlow))

	dbCollationS, err := InspectMigrateTask(databaseS, stringutil.StringUpper(dmt.DatasourceS.ConnectCharset), stringutil.StringUpper(dmt.DatasourceT.ConnectCharset))
	if err != nil {
		return err
	}

	logger.Info("stmt migrate task init task",
		zap.String("task_name", dmt.Task.TaskName), zap.String("task_mode", dmt.Task.TaskMode), zap.String("task_flow", dmt.Task.TaskFlow))

	err = dmt.InitDataMigrateTask(databaseS, dbCollationS, schemaRoute)
	if err != nil {
		return err
	}

	logger.Info("stmt migrate task get tables",
		zap.String("task_name", dmt.Task.TaskName), zap.String("task_mode", dmt.Task.TaskMode), zap.String("task_flow", dmt.Task.TaskFlow))

	summaries, err := model.GetIDataMigrateSummaryRW().FindDataMigrateSummary(dmt.Ctx, &task.DataMigrateSummary{
		TaskName:    dmt.Task.TaskName,
		SchemaNameS: schemaRoute.SchemaNameS,
	})
	if err != nil {
		return err
	}

	for _, s := range summaries {
		startTableTime := time.Now()
		logger.Info("stmt migrate task process table",
			zap.String("task_name", dmt.Task.TaskName),
			zap.String("task_mode", dmt.Task.TaskMode),
			zap.String("task_flow", dmt.Task.TaskFlow),
			zap.String("schema_name_s", s.SchemaNameS),
			zap.String("table_name_s", s.TableNameS))

		var migrateTasks []*task.DataMigrateTask
		err = model.Transaction(dmt.Ctx, func(txnCtx context.Context) error {
			// get migrate task tables
			migrateTasks, err = model.GetIDataMigrateTaskRW().FindDataMigrateTask(txnCtx,
				&task.DataMigrateTask{
					TaskName:    s.TaskName,
					SchemaNameS: s.SchemaNameS,
					TableNameS:  s.TableNameS,
					TaskStatus:  constant.TaskDatabaseStatusWaiting,
				})
			if err != nil {
				return err
			}
			migrateFailedTasks, err := model.GetIDataMigrateTaskRW().FindDataMigrateTask(txnCtx,
				&task.DataMigrateTask{
					TaskName:    s.TaskName,
					SchemaNameS: s.SchemaNameS,
					TableNameS:  s.TableNameS,
					TaskStatus:  constant.TaskDatabaseStatusFailed})
			if err != nil {
				return err
			}
			migrateRunningTasks, err := model.GetIDataMigrateTaskRW().FindDataMigrateTask(txnCtx,
				&task.DataMigrateTask{
					TaskName:    s.TaskName,
					SchemaNameS: s.SchemaNameS,
					TableNameS:  s.TableNameS,
					TaskStatus:  constant.TaskDatabaseStatusRunning})
			if err != nil {
				return err
			}
			migrateStopTasks, err := model.GetIDataMigrateTaskRW().FindDataMigrateTask(txnCtx,
				&task.DataMigrateTask{
					TaskName:    s.TaskName,
					SchemaNameS: s.SchemaNameS,
					TableNameS:  s.TableNameS,
					TaskStatus:  constant.TaskDatabaseStatusRunning})
			if err != nil {
				return err
			}
			migrateTasks = append(migrateTasks, migrateFailedTasks...)
			migrateTasks = append(migrateTasks, migrateRunningTasks...)
			migrateTasks = append(migrateTasks, migrateStopTasks...)
			return nil
		})
		if err != nil {
			return err
		}

		logger.Info("stmt migrate task process chunks",
			zap.String("task_name", dmt.Task.TaskName),
			zap.String("task_mode", dmt.Task.TaskMode),
			zap.String("task_flow", dmt.Task.TaskFlow),
			zap.String("schema_name_s", s.SchemaNameS),
			zap.String("table_name_s", s.TableNameS))

		var (
			sqlTSmt *sql.Stmt
		)
		switch {
		case strings.EqualFold(dmt.Task.TaskFlow, constant.TaskFlowOracleToTiDB) || strings.EqualFold(dmt.Task.TaskFlow, constant.TaskFlowOracleToMySQL):
			limitOne, err := model.GetIDataMigrateTaskRW().GetDataMigrateTask(dmt.Ctx, &task.DataMigrateTask{
				TaskName:    s.TaskName,
				SchemaNameS: s.SchemaNameS,
				TableNameS:  s.TableNameS})
			if err != nil {
				return err
			}
			sqlStr := GenMYSQLCompatibleDatabasePrepareStmt(s.SchemaNameT, s.TableNameT, dmt.TaskParams.SqlHintT, limitOne.ColumnDetailT, int(dmt.TaskParams.BatchSize), true)
			sqlTSmt, err = databaseT.PrepareContext(dmt.Ctx, sqlStr)
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("oracle current task [%s] schema [%s] task_mode [%s] task_flow [%s] prepare isn't support, please contact author", dmt.Task.TaskName, s.SchemaNameS, dmt.Task.TaskMode, dmt.Task.TaskFlow)
		}

		g := errconcurrent.NewGroup()
		g.SetLimit(int(dmt.TaskParams.SqlThreadS))
		for _, j := range migrateTasks {
			gTime := time.Now()
			g.Go(j, gTime, func(j interface{}) error {
				select {
				case <-dmt.Ctx.Done():
					return nil
				default:
					dt := j.(*task.DataMigrateTask)
					errW := model.Transaction(dmt.Ctx, func(txnCtx context.Context) error {
						_, err = model.GetIDataMigrateTaskRW().UpdateDataMigrateTask(txnCtx,
							&task.DataMigrateTask{TaskName: dt.TaskName, SchemaNameS: dt.SchemaNameS, TableNameS: dt.TableNameS, ChunkDetailS: dt.ChunkDetailS},
							map[string]interface{}{
								"TaskStatus": constant.TaskDatabaseStatusRunning,
							})
						if err != nil {
							return err
						}
						_, err = model.GetITaskLogRW().CreateLog(txnCtx, &task.Log{
							TaskName:    dt.TaskName,
							SchemaNameS: dt.SchemaNameS,
							TableNameS:  dt.TableNameS,
							LogDetail: fmt.Sprintf("%v [%v] stmt migrate task [%v] taskflow [%v] source table [%v.%v] chunk [%s] start",
								stringutil.CurrentTimeFormatString(),
								stringutil.StringLower(constant.TaskModeStmtMigrate),
								dt.TaskName,
								dmt.Task.TaskMode,
								dt.SchemaNameS,
								dt.TableNameS,
								dt.ChunkDetailS),
						})
						if err != nil {
							return err
						}
						return nil
					})
					if errW != nil {
						return errW
					}

					err = database.IDataMigrateProcess(&StmtMigrateRow{
						Ctx:           dmt.Ctx,
						TaskMode:      dmt.Task.TaskMode,
						TaskFlow:      dmt.Task.TaskFlow,
						Dmt:           dt,
						DatabaseS:     databaseS,
						DatabaseT:     databaseT,
						DatabaseTStmt: sqlTSmt,
						DBCharsetS:    constant.MigrateOracleCharsetStringConvertMapping[stringutil.StringUpper(dmt.DatasourceS.ConnectCharset)],
						DBCharsetT:    stringutil.StringUpper(dmt.DatasourceT.ConnectCharset),
						SqlThreadT:    int(dmt.TaskParams.SqlThreadT),
						BatchSize:     int(dmt.TaskParams.BatchSize),
						CallTimeout:   int(dmt.TaskParams.CallTimeout),
						SafeMode:      true,
						ReadChan:      make(chan []interface{}, constant.DefaultMigrateTaskQueueSize),
						WriteChan:     make(chan []interface{}, constant.DefaultMigrateTaskQueueSize),
					})
					if err != nil {
						return err
					}

					errW = model.Transaction(dmt.Ctx, func(txnCtx context.Context) error {
						_, err = model.GetIDataMigrateTaskRW().UpdateDataMigrateTask(txnCtx,
							&task.DataMigrateTask{TaskName: dt.TaskName, SchemaNameS: dt.SchemaNameS, TableNameS: dt.TableNameS, ChunkDetailS: dt.ChunkDetailS},
							map[string]interface{}{
								"TaskStatus": constant.TaskDatabaseStatusSuccess,
								"Duration":   fmt.Sprintf("%f", time.Now().Sub(gTime).Seconds()),
							})
						if err != nil {
							return err
						}
						_, err = model.GetITaskLogRW().CreateLog(txnCtx, &task.Log{
							TaskName:    dt.TaskName,
							SchemaNameS: dt.SchemaNameS,
							TableNameS:  dt.TableNameS,
							LogDetail: fmt.Sprintf("%v [%v] stmt migrate task [%v] taskflow [%v] source table [%v.%v] chunk [%s] success",
								stringutil.CurrentTimeFormatString(),
								stringutil.StringLower(constant.TaskModeStmtMigrate),
								dt.TaskName,
								dmt.Task.TaskMode,
								dt.SchemaNameS,
								dt.TableNameS,
								dt.ChunkDetailS),
						})
						if err != nil {
							return err
						}
						return nil
					})
					if errW != nil {
						return errW
					}
					return nil
				}
			})
		}

		for _, r := range g.Wait() {
			if r.Err != nil {
				smt := r.Task.(*task.DataMigrateTask)
				logger.Warn("stmt migrate task process tables",
					zap.String("task_name", dmt.Task.TaskName), zap.String("task_mode", dmt.Task.TaskMode), zap.String("task_flow", dmt.Task.TaskFlow),
					zap.String("schema_name_s", smt.SchemaNameS),
					zap.String("table_name_s", smt.TableNameS),
					zap.Error(r.Err))

				errW := model.Transaction(dmt.Ctx, func(txnCtx context.Context) error {
					_, err = model.GetIDataMigrateTaskRW().UpdateDataMigrateTask(txnCtx,
						&task.DataMigrateTask{TaskName: smt.TaskName, SchemaNameS: smt.SchemaNameS, TableNameS: smt.TableNameS, ChunkDetailS: smt.ChunkDetailS},
						map[string]interface{}{
							"TaskStatus":  constant.TaskDatabaseStatusFailed,
							"Duration":    fmt.Sprintf("%f", time.Now().Sub(r.Time).Seconds()),
							"ErrorDetail": r.Err.Error(),
						})
					if err != nil {
						return err
					}
					_, err = model.GetITaskLogRW().CreateLog(txnCtx, &task.Log{
						TaskName:    smt.TaskName,
						SchemaNameS: smt.SchemaNameS,
						TableNameS:  smt.TableNameS,
						LogDetail: fmt.Sprintf("%v [%v] stmt migrate task [%v] taskflow [%v] source table [%v.%v] failed, please see [data_migrate_task] detail",
							stringutil.CurrentTimeFormatString(),
							stringutil.StringLower(constant.TaskModeStmtMigrate),
							smt.TaskName,
							dmt.Task.TaskMode,
							smt.SchemaNameS,
							smt.TableNameS),
					})
					if err != nil {
						return err
					}
					return nil
				})
				if errW != nil {
					return errW
				}
			}
		}

		err = sqlTSmt.Close()
		if err != nil {
			return err
		}

		endTableTime := time.Now()
		err = model.Transaction(dmt.Ctx, func(txnCtx context.Context) error {
			tableStatusRecs, err := model.GetIDataMigrateTaskRW().FindDataMigrateTaskBySchemaTableChunkStatus(txnCtx, &task.DataMigrateTask{
				TaskName:    s.TaskName,
				SchemaNameS: s.SchemaNameS,
				TableNameS:  s.TableNameS,
			})
			if err != nil {
				return err
			}
			for _, rec := range tableStatusRecs {
				switch rec.TaskStatus {
				case constant.TaskDatabaseStatusSuccess:
					_, err = model.GetIDataMigrateSummaryRW().UpdateDataMigrateSummary(txnCtx, &task.DataMigrateSummary{
						TaskName:    rec.TaskName,
						SchemaNameS: rec.SchemaNameS,
						TableNameS:  rec.TableNameS,
					}, map[string]interface{}{
						"ChunkSuccess": rec.StatusTotals,
					})
					if err != nil {
						return err
					}
				case constant.TaskDatabaseStatusFailed:
					_, err = model.GetIDataMigrateSummaryRW().UpdateDataMigrateSummary(txnCtx, &task.DataMigrateSummary{
						TaskName:    rec.TaskName,
						SchemaNameS: rec.SchemaNameS,
						TableNameS:  rec.TableNameS,
					}, map[string]interface{}{
						"ChunkFails": rec.StatusTotals,
					})
					if err != nil {
						return err
					}
				case constant.TaskDatabaseStatusWaiting:
					_, err = model.GetIDataMigrateSummaryRW().UpdateDataMigrateSummary(txnCtx, &task.DataMigrateSummary{
						TaskName:    rec.TaskName,
						SchemaNameS: rec.SchemaNameS,
						TableNameS:  rec.TableNameS,
					}, map[string]interface{}{
						"ChunkWaits": rec.StatusTotals,
					})
					if err != nil {
						return err
					}
				case constant.TaskDatabaseStatusRunning:
					_, err = model.GetIDataMigrateSummaryRW().UpdateDataMigrateSummary(txnCtx, &task.DataMigrateSummary{
						TaskName:    rec.TaskName,
						SchemaNameS: rec.SchemaNameS,
						TableNameS:  rec.TableNameS,
					}, map[string]interface{}{
						"ChunkRuns": rec.StatusTotals,
					})
					if err != nil {
						return err
					}
				case constant.TaskDatabaseStatusStopped:
					_, err = model.GetIDataMigrateSummaryRW().UpdateDataMigrateSummary(txnCtx, &task.DataMigrateSummary{
						TaskName:    rec.TaskName,
						SchemaNameS: rec.SchemaNameS,
						TableNameS:  rec.TableNameS,
					}, map[string]interface{}{
						"ChunkStops": rec.StatusTotals,
					})
					if err != nil {
						return err
					}
				default:
					return fmt.Errorf("the task [%v] task_mode [%s] task_flow [%v] schema_name_s [%v] table_name_s [%v] task_status [%v] panic, please contact auhtor or reselect", s.TaskName, dmt.Task.TaskMode, dmt.Task.TaskFlow, rec.SchemaNameS, rec.TableNameS, rec.TaskStatus)
				}
			}

			_, err = model.GetIDataMigrateSummaryRW().UpdateDataMigrateSummary(txnCtx, &task.DataMigrateSummary{
				TaskName:    s.TaskName,
				SchemaNameS: s.SchemaNameS,
				TableNameS:  s.TableNameS,
			}, map[string]interface{}{
				"Duration": fmt.Sprintf("%f", time.Now().Sub(startTableTime).Seconds()),
			})
			if err != nil {
				return err
			}
			return nil
		})
		if err != nil {
			return err
		}

		logger.Info("stmt migrate task process table",
			zap.String("task_name", dmt.Task.TaskName),
			zap.String("task_mode", dmt.Task.TaskMode),
			zap.String("task_flow", dmt.Task.TaskFlow),
			zap.String("schema_name_s", s.SchemaNameS),
			zap.String("table_name_s", s.TableNameS),
			zap.String("cost", endTableTime.Sub(startTableTime).String()))
	}
	logger.Info("stmt migrate task",
		zap.String("task_name", dmt.Task.TaskName), zap.String("task_mode", dmt.Task.TaskMode), zap.String("task_flow", dmt.Task.TaskFlow),
		zap.String("cost", time.Now().Sub(schemaTaskTime).String()))
	return nil
}

func (dmt *StmtMigrateTask) InitDataMigrateTask(databaseS database.IDatabase, dbCollationS bool, schemaRoute *rule.SchemaRouteRule) error {
	// filter database table
	schemaTaskTables, err := model.GetIMigrateTaskTableRW().FindMigrateTaskTable(dmt.Ctx, &rule.MigrateTaskTable{
		TaskName:    schemaRoute.TaskName,
		SchemaNameS: schemaRoute.SchemaNameS,
	})
	if err != nil {
		return err
	}
	var (
		includeTables  []string
		excludeTables  []string
		databaseTables []string // task tables
		globalScn      uint64
	)
	databaseTableTypeMap := make(map[string]string)

	for _, t := range schemaTaskTables {
		if strings.EqualFold(t.IsExclude, constant.MigrateTaskTableIsExclude) {
			excludeTables = append(excludeTables, t.TableNameS)
		}
		if strings.EqualFold(t.IsExclude, constant.MigrateTaskTableIsNotExclude) {
			includeTables = append(includeTables, t.TableNameS)
		}
	}

	databaseTables, err = databaseS.FilterDatabaseTable(schemaRoute.SchemaNameS, includeTables, excludeTables)
	if err != nil {
		return err
	}

	// clear the stmt migrate task table
	// repeatInitTableMap used for store the struct_migrate_task table name has be finished, avoid repeated initialization
	migrateGroupTasks, err := model.GetIDataMigrateTaskRW().FindDataMigrateTaskGroupByTaskSchemaTable(dmt.Ctx, dmt.Task.TaskName)
	if err != nil {
		return err
	}
	repeatInitTableMap := make(map[string]struct{})

	if len(migrateGroupTasks) > 0 {
		taskTablesMap := make(map[string]struct{})
		for _, t := range databaseTables {
			taskTablesMap[t] = struct{}{}
		}
		for _, smt := range migrateGroupTasks {
			if smt.SchemaNameS == schemaRoute.SchemaNameS {
				if _, ok := taskTablesMap[smt.TableNameS]; !ok {
					err = model.Transaction(dmt.Ctx, func(txnCtx context.Context) error {
						err = model.GetIDataMigrateSummaryRW().DeleteDataMigrateSummary(txnCtx, &task.DataMigrateSummary{
							TaskName:    smt.TaskName,
							SchemaNameS: smt.SchemaNameS,
							TableNameS:  smt.TableNameS,
						})
						if err != nil {
							return err
						}
						err = model.GetIDataMigrateTaskRW().DeleteDataMigrateTask(txnCtx, &task.DataMigrateTask{
							TaskName:    smt.TaskName,
							SchemaNameS: smt.SchemaNameS,
							TableNameS:  smt.TableNameS,
						})
						if err != nil {
							return err
						}
						return nil
					})
					if err != nil {
						return err
					}

					continue
				}
				var summary *task.DataMigrateSummary

				summary, err = model.GetIDataMigrateSummaryRW().GetDataMigrateSummary(dmt.Ctx, &task.DataMigrateSummary{
					TaskName:    smt.TaskName,
					SchemaNameS: smt.SchemaNameS,
					TableNameS:  smt.TableNameS,
				})
				if err != nil {
					return err
				}

				if int64(summary.ChunkTotals) != smt.ChunkTotals {
					err = model.Transaction(dmt.Ctx, func(txnCtx context.Context) error {
						err = model.GetIDataMigrateSummaryRW().DeleteDataMigrateSummary(txnCtx, &task.DataMigrateSummary{
							TaskName:    smt.TaskName,
							SchemaNameS: smt.SchemaNameS,
							TableNameS:  smt.TableNameS,
						})
						if err != nil {
							return err
						}
						err = model.GetIDataMigrateTaskRW().DeleteDataMigrateTask(txnCtx, &task.DataMigrateTask{
							TaskName:    smt.TaskName,
							SchemaNameS: smt.SchemaNameS,
							TableNameS:  smt.TableNameS,
						})
						if err != nil {
							return err
						}
						return nil
					})
					if err != nil {
						return err
					}

					continue
				}

				repeatInitTableMap[smt.TableNameS] = struct{}{}
			}
		}
	}

	databaseTableTypeMap, err = databaseS.GetDatabaseTableType(schemaRoute.SchemaNameS)
	if err != nil {
		return err
	}

	globalScn, err = databaseS.GetDatabaseConsistentPos()
	if err != nil {
		return err
	}

	// database tables
	// init database table
	logger.Info("stmt migrate task init",
		zap.String("task_name", dmt.Task.TaskName), zap.String("task_mode", dmt.Task.TaskMode), zap.String("task_flow", dmt.Task.TaskFlow))

	g, gCtx := errgroup.WithContext(dmt.Ctx)
	g.SetLimit(int(dmt.TaskParams.TableThread))

	for _, taskJob := range databaseTables {
		sourceTable := taskJob
		g.Go(func() error {
			select {
			case <-gCtx.Done():
				return nil
			default:
				startTime := time.Now()
				if _, ok := repeatInitTableMap[sourceTable]; ok {
					// skip
					return nil
				}

				tableRows, err := databaseS.GetDatabaseTableRows(schemaRoute.SchemaNameS, sourceTable)
				if err != nil {
					return err
				}
				tableSize, err := databaseS.GetDatabaseTableSize(schemaRoute.SchemaNameS, sourceTable)
				if err != nil {
					return err
				}

				dataRule := &DataMigrateRule{
					Ctx:            gCtx,
					TaskMode:       dmt.Task.TaskMode,
					TaskName:       dmt.Task.TaskName,
					TaskFlow:       dmt.Task.TaskFlow,
					DatabaseS:      databaseS,
					DBCollationS:   dbCollationS,
					SchemaNameS:    schemaRoute.SchemaNameS,
					TableNameS:     sourceTable,
					TableTypeS:     databaseTableTypeMap,
					DBCharsetS:     dmt.DatasourceS.ConnectCharset,
					CaseFieldRuleS: dmt.Task.CaseFieldRuleS,
					CaseFieldRuleT: dmt.Task.CaseFieldRuleT,
					GlobalSqlHintS: dmt.TaskParams.SqlHintS,
				}

				attsRule, err := database.IDataMigrateAttributesRule(dataRule)
				if err != nil {
					return err
				}

				// only where range
				if !attsRule.EnableChunkStrategy && !strings.EqualFold(attsRule.WhereRange, "") {
					err = model.Transaction(gCtx, func(txnCtx context.Context) error {
						_, err = model.GetIDataMigrateTaskRW().CreateDataMigrateTask(txnCtx, &task.DataMigrateTask{
							TaskName:        dmt.Task.TaskName,
							SchemaNameS:     attsRule.SchemaNameS,
							TableNameS:      attsRule.TableNameS,
							SchemaNameT:     attsRule.SchemaNameT,
							TableNameT:      attsRule.TableNameT,
							TableTypeS:      attsRule.TableTypeS,
							GlobalScnS:      globalScn,
							ColumnDetailO:   attsRule.ColumnDetailO,
							ColumnDetailS:   attsRule.ColumnDetailS,
							ColumnDetailT:   attsRule.ColumnDetailT,
							SqlHintS:        attsRule.SqlHintS,
							SqlHintT:        dmt.TaskParams.SqlHintT,
							ChunkDetailS:    attsRule.WhereRange,
							ConsistentReadS: strconv.FormatBool(dmt.TaskParams.EnableConsistentRead),
							TaskStatus:      constant.TaskDatabaseStatusWaiting,
						})
						if err != nil {
							return err
						}
						_, err = model.GetIDataMigrateSummaryRW().CreateDataMigrateSummary(txnCtx, &task.DataMigrateSummary{
							TaskName:    dmt.Task.TaskName,
							SchemaNameS: attsRule.SchemaNameS,
							TableNameS:  attsRule.TableNameS,
							SchemaNameT: attsRule.SchemaNameT,
							TableNameT:  attsRule.TableNameT,
							GlobalScnS:  globalScn,
							TableRowsS:  tableRows,
							TableSizeS:  tableSize,
							ChunkTotals: 1,
						})
						if err != nil {
							return err
						}
						return nil
					})
					if err != nil {
						return err
					}
					return nil
				}

				chunkTask := uuid.New().String()

				chunks, err := databaseS.GetDatabaseTableChunkTask(chunkTask, schemaRoute.SchemaNameS, sourceTable, dmt.TaskParams.ChunkSize, dmt.TaskParams.CallTimeout)
				if err != nil {
					return err
				}

				var whereRange string

				if len(chunks) == 0 {
					switch {
					case attsRule.EnableChunkStrategy && !strings.EqualFold(attsRule.WhereRange, ""):
						whereRange = stringutil.StringBuilder(`1 = 1 AND `, attsRule.WhereRange)
					default:
						whereRange = `1 = 1`
					}

					err = model.Transaction(gCtx, func(txnCtx context.Context) error {
						_, err = model.GetIDataMigrateTaskRW().CreateDataMigrateTask(txnCtx, &task.DataMigrateTask{
							TaskName:        dmt.Task.TaskName,
							SchemaNameS:     attsRule.SchemaNameS,
							TableNameS:      attsRule.TableNameS,
							SchemaNameT:     attsRule.SchemaNameT,
							TableNameT:      attsRule.TableNameT,
							TableTypeS:      attsRule.TableTypeS,
							GlobalScnS:      globalScn,
							ColumnDetailO:   attsRule.ColumnDetailO,
							ColumnDetailS:   attsRule.ColumnDetailS,
							ColumnDetailT:   attsRule.ColumnDetailT,
							SqlHintS:        attsRule.SqlHintS,
							SqlHintT:        dmt.TaskParams.SqlHintT,
							ChunkDetailS:    whereRange,
							ConsistentReadS: strconv.FormatBool(dmt.TaskParams.EnableConsistentRead),
							TaskStatus:      constant.TaskDatabaseStatusWaiting,
						})
						if err != nil {
							return err
						}
						_, err = model.GetIDataMigrateSummaryRW().CreateDataMigrateSummary(txnCtx, &task.DataMigrateSummary{
							TaskName:    dmt.Task.TaskName,
							SchemaNameS: attsRule.SchemaNameS,
							TableNameS:  attsRule.TableNameS,
							SchemaNameT: attsRule.SchemaNameT,
							TableNameT:  attsRule.TableNameT,
							GlobalScnS:  globalScn,
							TableRowsS:  tableRows,
							TableSizeS:  tableSize,
							ChunkTotals: 1,
						})
						if err != nil {
							return err
						}
						return nil
					})
					if err != nil {
						return err
					}

					return nil
				}

				var metas []*task.DataMigrateTask
				for _, r := range chunks {
					switch {
					case attsRule.EnableChunkStrategy && !strings.EqualFold(attsRule.WhereRange, ""):
						whereRange = stringutil.StringBuilder(r["CMD"], ` AND `, attsRule.WhereRange)
					default:
						whereRange = r["CMD"]
					}

					metas = append(metas, &task.DataMigrateTask{
						TaskName:        dmt.Task.TaskName,
						SchemaNameS:     attsRule.SchemaNameS,
						TableNameS:      attsRule.TableNameS,
						SchemaNameT:     attsRule.SchemaNameT,
						TableNameT:      attsRule.TableNameT,
						TableTypeS:      attsRule.TableTypeS,
						GlobalScnS:      globalScn,
						ColumnDetailO:   attsRule.ColumnDetailO,
						ColumnDetailS:   attsRule.ColumnDetailS,
						ColumnDetailT:   attsRule.ColumnDetailT,
						SqlHintS:        attsRule.SqlHintS,
						SqlHintT:        dmt.TaskParams.SqlHintT,
						ChunkDetailS:    whereRange,
						ConsistentReadS: strconv.FormatBool(dmt.TaskParams.EnableConsistentRead),
						TaskStatus:      constant.TaskDatabaseStatusWaiting,
					})
				}

				err = model.Transaction(gCtx, func(txnCtx context.Context) error {
					err = model.GetIDataMigrateTaskRW().CreateInBatchDataMigrateTask(txnCtx, metas, int(dmt.TaskParams.BatchSize))
					if err != nil {
						return err
					}
					_, err = model.GetIDataMigrateSummaryRW().CreateDataMigrateSummary(txnCtx, &task.DataMigrateSummary{
						TaskName:    dmt.Task.TaskName,
						SchemaNameS: attsRule.SchemaNameS,
						TableNameS:  attsRule.TableNameS,
						SchemaNameT: attsRule.SchemaNameT,
						TableNameT:  attsRule.TableNameT,
						GlobalScnS:  globalScn,
						TableRowsS:  tableRows,
						TableSizeS:  tableSize,
						ChunkTotals: uint64(len(chunks)),
					})
					if err != nil {
						return err
					}
					return nil
				})
				if err != nil {
					return err
				}

				logger.Info("stmt migrate task init",
					zap.String("task_name", dmt.Task.TaskName),
					zap.String("task_mode", dmt.Task.TaskMode),
					zap.String("task_flow", dmt.Task.TaskFlow),
					zap.String("schema_name_s", attsRule.SchemaNameS),
					zap.String("table_name_s", attsRule.TableNameS),
					zap.String("cost", time.Now().Sub(startTime).String()))
				return nil
			}
		})
	}

	// ignore context canceled error
	if err = g.Wait(); !errors.Is(err, context.Canceled) {
		logger.Warn("stmt migrate task init",
			zap.String("task_name", dmt.Task.TaskName), zap.String("task_mode", dmt.Task.TaskMode), zap.String("task_flow", dmt.Task.TaskFlow),
			zap.String("schema_name_s", schemaRoute.SchemaNameS),
			zap.Error(err))
		return err
	}
	return nil
}
