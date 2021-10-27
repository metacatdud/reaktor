// Copyright 2021 The Atomika Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arch

type Linux struct {
}

func NewLinuxOS() OS {
	return &Linux{}
}

func (arch *Linux) Name() OSName {
	return "linux"
}
