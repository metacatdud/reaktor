// Copyright 2021 The Atomika Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dependecies

type Dep string

func (d Dep) String() string {
	return string(d)
}

const (
	DepIPFS Dep = "ipfs"
	Dep7Z   Dep = "7z"
	DepFake Dep = "fake"
)

var (
	DepsList []Dep = []Dep{DepIPFS, Dep7Z, DepFake}
)
