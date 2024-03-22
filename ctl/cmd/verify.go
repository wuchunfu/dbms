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
package cmd

import (
	"context"
	"fmt"
	"strings"

	"github.com/wentaojin/dbms/service"

	"github.com/spf13/cobra"
	"github.com/wentaojin/dbms/ctl/migrate"
)

type AppVerify struct {
	*App
}

func (a *App) AppVerify() Cmder {
	return &AppVerify{App: a}
}

func (a *AppVerify) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "verify",
		Short:            "Operator cluster data compare",
		Long:             `Operator cluster data compare`,
		RunE:             a.RunE,
		TraverseChildren: true,
		SilenceUsage:     true,
	}
	return cmd
}

func (a *AppVerify) RunE(cmd *cobra.Command, args []string) error {
	if err := cmd.Help(); err != nil {
		return err
	}
	return nil
}

type AppVerifyUpsert struct {
	*AppVerify
	config string
}

func (a *AppVerify) AppVerifyUpsert() Cmder {
	return &AppVerifyUpsert{AppVerify: a}
}

func (a *AppVerifyUpsert) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "upsert",
		Short:            "upsert cluster data compare task",
		Long:             `upsert cluster data compare task`,
		RunE:             a.RunE,
		TraverseChildren: true,
		SilenceUsage:     true,
	}
	cmd.Flags().StringVarP(&a.config, "config", "c", "config.toml", "config")
	return cmd
}

func (a *AppVerifyUpsert) RunE(cmd *cobra.Command, args []string) error {
	if strings.EqualFold(a.config, "") {
		return fmt.Errorf("flag parameter [config] is requirement, can not null")
	}

	if len(args) > 0 {
		if err := cmd.Help(); err != nil {
			return err
		}
	}

	err := migrate.UpsertDataCompare(a.Server, a.config)
	if err != nil {
		return err
	}
	return nil
}

type AppVerifyDelete struct {
	*AppVerify
	task string
}

func (a *AppVerify) AppVerifyDelete() Cmder {
	return &AppVerifyDelete{AppVerify: a}
}

func (a *AppVerifyDelete) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "delete",
		Short:            "delete cluster data compare task",
		Long:             `delete cluster data compare task`,
		RunE:             a.RunE,
		TraverseChildren: true,
		SilenceUsage:     true,
	}
	cmd.Flags().StringVarP(&a.task, "task", "t", "xxx", "delete data compare task")
	return cmd
}

func (a *AppVerifyDelete) RunE(cmd *cobra.Command, args []string) error {
	if strings.EqualFold(a.task, "") {
		return fmt.Errorf("flag parameter [task] is requirement, can not null")
	}

	if len(args) > 0 {
		if err := cmd.Help(); err != nil {
			return err
		}
	}

	err := migrate.DeleteDataCompare(a.Server, a.task)
	if err != nil {
		return err
	}
	fmt.Printf("Success Delete data compare Task [%v]！！！\n", a.task)
	return nil
}

type AppVerifyGet struct {
	*AppVerify
	task string
}

func (a *AppVerify) AppVerifyGet() Cmder {
	return &AppVerifyGet{AppVerify: a}
}

func (a *AppVerifyGet) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "get",
		Short:            "get cluster data compare task",
		Long:             `get cluster data compare task`,
		RunE:             a.RunE,
		TraverseChildren: true,
		SilenceUsage:     true,
	}
	return cmd
}

func (a *AppVerifyGet) RunE(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		if err := cmd.Help(); err != nil {
			return err
		}
	}

	err := migrate.GetDataCompare(a.Server, a.task)
	if err != nil {
		return err
	}

	return nil
}

type AppVerifyGen struct {
	*AppVerify
	task      string
	outputDir string
}

func (a *AppVerify) AppVerifyGen() Cmder {
	return &AppVerifyGen{AppVerify: a}
}

func (a *AppVerifyGen) Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:              "gen",
		Short:            "gen cluster data compare task detail",
		Long:             `gen cluster data compare task detail`,
		RunE:             a.RunE,
		TraverseChildren: true,
		SilenceUsage:     true,
	}
	cmd.Flags().StringVarP(&a.task, "task", "t", "", "the data compare task")
	cmd.Flags().StringVarP(&a.outputDir, "outputDir", "o", "/tmp", "the data compare task output file dir")
	return cmd
}

func (a *AppVerifyGen) RunE(cmd *cobra.Command, args []string) error {
	if len(args) > 0 {
		if err := cmd.Help(); err != nil {
			return err
		}
	}
	if strings.EqualFold(a.task, "") || strings.EqualFold(a.outputDir, "") {
		return fmt.Errorf("flag parameter [task] and [outputDir] are requirement, can not null")
	}

	err := service.GenDataCompareTask(context.Background(), a.Server, a.task, a.outputDir)
	if err != nil {
		return err
	}
	return nil
}