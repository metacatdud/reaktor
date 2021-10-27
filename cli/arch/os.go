// Copyright 2021 The Atomika Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arch

import (
	"runtime"
)

type (
	OSName string
)

const (
	LinuxOS       OSName = "linux"
	MacOS         OSName = "mac"
	WindowsOS     OSName = "windows"
	UnsupportedOS OSName = "unsupported"
)

type OS interface {
	Name() OSName
}

func NewOS() OS {

	var os OS

	runtimeOS := runtime.GOOS
	switch runtimeOS {
	case "windows":
		os = NewWindowsOS()
	case "darwin":
		os = NewMacOS()
	case "linux":
		os = NewLinuxOS()
	default:
		os = NewUnsupportedOS()
	}

	return os
}
