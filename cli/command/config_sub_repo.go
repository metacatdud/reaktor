// Copyright 2021 The Atomika Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package command

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

const (
	CommandNameSubRepo = "repo [no options!]"
)

type configSubRepo struct {
	use      string
	shorDesc string
	longDesc string
}

func NewConfigSubRepo() Command {
	return &configSubRepo{
		use:      CommandNameSubRepo,
		shorDesc: "Set a path for the repository",
		longDesc: "",
	}
}

func (cmd *configSubRepo) Name() string {
	return CommandNameSubRepo
}

func (cmd *configSubRepo) Use() string {
	return cmd.use
}

func (cmd *configSubRepo) ShortDesc() string {
	return cmd.shorDesc
}

func (cmd *configSubRepo) LongDesc() string {
	return cmd.longDesc
}

func (cmd *configSubRepo) RunE() RunE {
	return nil
}

func (cmd *configSubRepo) Run() Run {
	return cmd.run
}

func (cmd *configSubRepo) Flags() []Flag {
	testBool := NewStringFlag(StringFlagConfig{
		Name:         "set-repo-path",
		RefVar:       &RepoPath,
		DefaultValue: "",
		Description:  "Some flag description",
	})
	return []Flag{
		testBool,
	}
}

func (cmd *configSubRepo) Sub() []Command {
	return []Command{}
}

func (cmd *configSubRepo) run(c *cobra.Command, args []string) {
	pterm.DefaultSection.Println("Setting repo path ...", args, RepoPath)
	pterm.Warning.Println("== NOT IMPLEMENTED")
}
