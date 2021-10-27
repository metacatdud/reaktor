// Copyright 2021 The Atomika Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package command

type Flag interface {
	Name() string
	Description() string
}

// BoolFlag flag type boolean
type BoolFlag interface {
	Flag
	RefVar() *bool
	DefaultVal() bool
}

type BoolFlagConfig struct {
	Name         string
	RefVar       *bool
	DefaultValue bool
	Description  string
}

type boolFlag struct {
	name       string
	refVar     *bool
	defaultVal bool
	desc       string
}

func NewBoolFlag(c BoolFlagConfig) BoolFlag {
	return &boolFlag{
		name:       c.Name,
		refVar:     c.RefVar,
		defaultVal: c.DefaultValue,
		desc:       c.Description,
	}
}

func (flag *boolFlag) Name() string {
	return flag.name
}

func (flag *boolFlag) Description() string {
	return flag.desc
}

func (flag *boolFlag) DefaultVal() bool {
	return flag.defaultVal
}

func (flag *boolFlag) RefVar() *bool {
	return flag.refVar
}

// StringFlag flag type string
type StringFlag interface {
	Flag
	RefVar() *string
	DefaultVal() string
}

type StringFlagConfig struct {
	Name         string
	RefVar       *string
	DefaultValue string
	Description  string
}

type stringFlag struct {
	name       string
	refVar     *string
	defaultVal string
	desc       string
}

func NewStringFlag(c StringFlagConfig) StringFlag {
	return &stringFlag{
		name:       c.Name,
		refVar:     c.RefVar,
		defaultVal: c.DefaultValue,
		desc:       c.Description,
	}
}

func (flag *stringFlag) Name() string {
	return flag.name
}

func (flag *stringFlag) Description() string {
	return flag.desc
}

func (flag *stringFlag) DefaultVal() string {
	return flag.defaultVal
}

func (flag *stringFlag) RefVar() *string {
	return flag.refVar
}
