// Copyright 2021 The Atomika Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package command

import (
	"errors"
	"fmt"
	"os/exec"
	"time"

	"github.com/metacatdud/reaktor/cli/arch"
	"github.com/metacatdud/reaktor/cli/dependecies"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

const (
	CommandNameCheck = "check"
)

var (
	ErrCheckCommandOSUnsupported = errors.New("detected unsupported os")
)

type checkCommand struct {
	use      string
	shorDesc string
	longDesc string
}

func NewCheckCommand() Command {
	c := &checkCommand{
		use:      "check",
		shorDesc: "Test reaktor dependencies",
		longDesc: "Test reaktor dependencies",
	}

	return c
}

func (cmd *checkCommand) Name() string {
	return CommandNameCheck
}

func (cmd *checkCommand) Use() string {
	return cmd.use
}

func (cmd *checkCommand) ShortDesc() string {
	return cmd.shorDesc
}

func (cmd *checkCommand) LongDesc() string {
	return cmd.longDesc
}

func (cmd *checkCommand) RunE() RunE {
	return cmd.executeWithError
}

func (cmd *checkCommand) Run() Run {
	return nil
}

func (cmd *checkCommand) Flags() []Flag {
	return []Flag{}
}

func (cmd *checkCommand) Sub() []Command {
	return []Command{}
}

func (cmd *checkCommand) executeWithError(c *cobra.Command, args []string) error {
	Header()

	// Check OS
	pterm.DefaultSection.Println("1. Checking OS ...")

	os := arch.NewOS()
	if os.Name() == arch.UnsupportedOS {
		pterm.Warning.Printf("%s:%s", ErrCheckCommandOSUnsupported, os.Name())
		pterm.Println()
		pterm.Println()
		return errors.New("error")
	}
	pterm.Success.Println("OS:", os.Name())
	pterm.Println()

	// 	Check dependencies
	pterm.DefaultSection.Println("2. Checking dependencies ...")

	p, _ := pterm.DefaultProgressbar.WithTotal(len(dependecies.DepsList)).WithTitle("Checking...").Start()
	for i := 0; i < p.Total; i++ {
		depStr := dependecies.DepsList[i].String()
		p.Title = fmt.Sprintf("Check [%s]", depStr)

		path, err := exec.LookPath(depStr)
		if err != nil {
			pterm.Warning.Println("Not found", depStr)

			FooterFail()
			return errors.New("error")
		}

		pterm.Success.Println("Found:", depStr, "Path:", path)

		p.Increment()
		time.Sleep(1000 * time.Millisecond)
	}
	p.Stop()

	FooterSuccess()
	return nil
}
