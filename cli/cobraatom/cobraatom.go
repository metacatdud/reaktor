// Copyright 2021 The Atomika Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cobraatom

import (
	"github.com/metacatdud/reaktor/cli/command"
	"github.com/spf13/cobra"
)

type Cobra interface {
	Name() string
	AddCommand(commands ...command.Command)
	Execute()
}

type cobraAtom struct {
	name   string
	client *cobra.Command
}

func NewCobra(name string) Cobra {
	return &cobraAtom{
		name: name,
		client: &cobra.Command{
			Use: name,
		},
	}
}

func (c *cobraAtom) Name() string {
	return c.name
}

func (c *cobraAtom) AddCommand(commands ...command.Command) {
	for _, command := range commands {
		cobraComm, _ := parseToCobra(command)
		c.client.AddCommand(cobraComm)
	}
}

func (c *cobraAtom) Execute() {
	c.client.Execute()
}

func parseToCobra(c command.Command) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   c.Use(),
		Long:  c.LongDesc(),
		Short: c.ShortDesc(),
	}

	if c.Run() != nil {
		cmd.Run = c.Run()
	}

	if c.RunE() != nil {
		cmd.RunE = c.RunE()
	}

	if len(c.Sub()) > 0 {
		cmd.TraverseChildren = true

		for _, subComm := range c.Sub() {
			subCmd, _ := parseToCobra(subComm)
			cmd.AddCommand(subCmd)
		}
	}

	for _, flag := range c.Flags() {
		switch f := flag.(type) {
		case command.StringFlag:
			cmd.Flags().StringVarP(f.RefVar(), f.Name(), "", f.DefaultVal(), f.Description())
		case command.BoolFlag:
			cmd.Flags().BoolVarP(f.RefVar(), f.Name(), "", f.DefaultVal(), f.Description())
		}
	}

	return cmd, nil
}
