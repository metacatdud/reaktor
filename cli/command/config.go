// Copyright 2021 The Atomika Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package command

import (
	"embed"
	"errors"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

const (
	CommandNameConfig = "config"

	TitleCheck = "Checklist"

	MsgCheckDone        = "export REAKTOR_WORKSPACE_PATH=%s"
	MsgCheckPermissions = "Make sure you have permission to write to that location"
	MsgCheckExist       = "Make sure config file exist"
)

var (
	WorkspacePath string
	RepoPath      string
)

var (
	ErrConfigSetup                    = errors.New("workspace already defined")
	ErrConfigUnableToCreateWorkspace  = errors.New("cannot create workspace")
	ErrConfigUnableToWriteToWorkspace = errors.New("cannot write to workspace")
	ErrConfigUnableToGenerateConfig   = errors.New("unable to generate config")
)

type configCommand struct {
	use      string
	shorDesc string
	longDesc string

	sub []Command

	tpl embed.FS
}

func NewConfigCommand(tpl embed.FS) Command {
	c := &configCommand{
		use:      CommandNameConfig,
		shorDesc: "Initialize reaktor work environment",
		longDesc: "Initialize reaktor work environment such as the local repo, push api",

		tpl: tpl,
	}

	return c
}

func (cmd *configCommand) Name() string {
	return CommandNameConfig
}

func (cmd *configCommand) Use() string {
	return cmd.use
}

func (cmd *configCommand) ShortDesc() string {
	return cmd.shorDesc
}

func (cmd *configCommand) LongDesc() string {
	return cmd.longDesc
}

func (cmd *configCommand) RunE() RunE {
	return cmd.runWithError
}

func (cmd *configCommand) Run() Run {
	return nil
}

func (cmd *configCommand) Flags() []Flag {
	return cmd.setFlags()
}

func (cmd *configCommand) Sub() []Command {
	return []Command{
		NewConfigSubRepo(),
	}
}

func (cmd *configCommand) setFlags() []Flag {
	pathFlag := NewStringFlag(StringFlagConfig{
		Name:         "path",
		RefVar:       &WorkspacePath,
		DefaultValue: DefaultPath,
		Description:  "Set custom location for workspace",
	})
	return []Flag{
		pathFlag,
	}
}

func (cmd *configCommand) runWithError(c *cobra.Command, args []string) error {
	Header()

	var workspacePath string

	workspacePath = DefaultPath
	if WorkspacePath != "" {
		workspacePath = WorkspacePath
	}

	workspaceDir, err := os.Open(workspacePath)
	if err == nil {
		pterm.Error.Printf("%s:[%s]\n", ErrConfigSetup, workspacePath)

		FooterFail()
		return errors.New("error")
	}
	workspaceDir.Close()

	err = os.MkdirAll(workspacePath, os.ModePerm)
	if err != nil {
		pterm.Error.Printf("%s:[%s]\n", ErrConfigUnableToCreateWorkspace, workspacePath)

		FooterFail()
		return errors.New("error")
	}

	f, err := os.Create(strings.Join([]string{workspacePath, "config.json"}, "/"))
	if err != nil {
		pterm.Error.Printf("%s:[%s]\n", ErrConfigUnableToGenerateConfig, workspacePath)
		pterm.Println()
		pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
			{Level: 0, Text: TitleCheck, TextStyle: pterm.NewStyle(pterm.FgCyan), BulletStyle: pterm.NewStyle(pterm.FgRed)},
			{Level: 1, Text: MsgCheckPermissions, TextStyle: pterm.NewStyle(pterm.FgGreen), Bullet: "-", BulletStyle: pterm.NewStyle(pterm.FgLightWhite)},
		}).Render()

		FooterFail()
		return errors.New("error")
	}

	err = os.Mkdir(strings.Join([]string{workspacePath, "repo"}, "/"), os.ModePerm)
	if err != nil {
		pterm.Error.Printf("%s:[%s]\n", ErrConfigUnableToCreateWorkspace, workspacePath)

		FooterFail()
		return errors.New("error")
	}

	t, err := template.ParseFS(cmd.tpl, "templates/reaktor-config.tpl")
	if err != nil {
		pterm.Error.Printf("%s:[%s]\n", ErrConfigUnableToWriteToWorkspace, workspacePath)
		pterm.Println()
		pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
			{Level: 0, Text: TitleCheck, TextStyle: pterm.NewStyle(pterm.FgCyan), BulletStyle: pterm.NewStyle(pterm.FgRed)},
			{Level: 1, Text: MsgCheckPermissions, TextStyle: pterm.NewStyle(pterm.FgGreen), Bullet: "-", BulletStyle: pterm.NewStyle(pterm.FgLightWhite)},
		}).Render()

		FooterFail()
		return errors.New("error")
	}

	err = t.Execute(f, map[string]string{
		"Name": "test",
	})

	if err != nil {
		pterm.Error.Printf("%s:[%s]\n", ErrConfigUnableToWriteToWorkspace, workspacePath)
		pterm.Println()
		pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
			{Level: 0, Text: TitleCheck, TextStyle: pterm.NewStyle(pterm.FgCyan), BulletStyle: pterm.NewStyle(pterm.FgRed)},
			{Level: 1, Text: MsgCheckPermissions, TextStyle: pterm.NewStyle(pterm.FgGreen), Bullet: "-", BulletStyle: pterm.NewStyle(pterm.FgLightWhite)},
			{Level: 1, Text: MsgCheckExist, TextStyle: pterm.NewStyle(pterm.FgGreen), Bullet: "-", BulletStyle: pterm.NewStyle(pterm.FgLightWhite)},
		}).Render()

		FooterFail()
		return errors.New("error")
	}
	f.Close()

	pterm.Success.Printf("Reaktor workspace set:[%s]\n", workspacePath)

	if DefaultPath != workspacePath {
		pterm.Println()
		pterm.DefaultBulletList.WithItems([]pterm.BulletListItem{
			{Level: 0, Text: TitleCheck, TextStyle: pterm.NewStyle(pterm.FgCyan), BulletStyle: pterm.NewStyle(pterm.FgRed)},
			{Level: 1, Text: fmt.Sprintf(MsgCheckDone, workspacePath), TextStyle: pterm.NewStyle(pterm.FgYellow), Bullet: "-", BulletStyle: pterm.NewStyle(pterm.FgLightWhite)},
		}).Render()

	}

	FooterSuccess()
	return nil
}
