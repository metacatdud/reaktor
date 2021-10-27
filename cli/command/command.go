// Copyright 2021 The Atomika Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package command

import (
	"os"
	"strings"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var (
	DefaultPath string
)

func init() {
	dirname, _ := os.UserHomeDir()
	DefaultPath = strings.Join([]string{dirname, ".reaktor"}, "/")
}

type RunE func(c *cobra.Command, args []string) error
type Run func(c *cobra.Command, args []string)

type Command interface {
	Name() string
	Use() string
	ShortDesc() string
	LongDesc() string
	RunE() RunE
	Run() Run

	Flags() []Flag
	Sub() []Command
}

func Header() {
	pterm.DefaultSection.
		WithStyle(pterm.NewStyle(pterm.FgGreen, pterm.Bold)).
		Println(
			"REAKTOR - CLI Tool",
		)
	pterm.Println()
}

func FooterSuccess() {
	pterm.Println()
	pterm.DefaultSection.
		WithStyle(pterm.NewStyle(pterm.FgGreen, pterm.Bold)).
		WithLevel(2).
		Println("Done!")
}

func FooterFail() {
	pterm.Println()
	pterm.DefaultSection.
		WithStyle(pterm.NewStyle(pterm.FgRed, pterm.Bold)).
		WithLevel(2).
		Println("Fail!")
}
