// Copyright 2021 The Atomika Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/metacatdud/reaktor/cli/cobraatom"
	"github.com/metacatdud/reaktor/cli/command"
)

//go:generate cp -r ../templates ./
//go:embed templates/*
var tpls embed.FS

const (
	appName = "reaktor-cli"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run() error {
	var err error

	cob := cobraatom.NewCobra(appName)
	cob.AddCommand(
		command.NewConfigCommand(tpls),
		command.NewCheckCommand(),
	)
	cob.Execute()

	return err
}
