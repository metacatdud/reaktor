// Copyright 2021 The Atomika Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package repository

type InfoRepo struct {
	Files []string
	Size  int
}

type Repository interface {
	Info() *InfoRepo
	GetFile()
	AddFile()
}

type repository struct {
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) Info() *InfoRepo {
	panic("implement me")
}

func (r *repository) GetFile() {
	panic("implement me")
}

func (r *repository) AddFile() {
	panic("implement me")
}
