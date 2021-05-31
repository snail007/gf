// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gres_test

import (
	"strings"
	"testing"

	"github.com/snail007/gf/os/gfile"

	"github.com/snail007/gf/debug/gdebug"
	"github.com/snail007/gf/os/gres"
	"github.com/snail007/gf/test/gtest"
)

func Test_Pack(t *testing.T) {
	gtest.Case(t, func() {
		srcPath := gdebug.CallerDirectory() + "/testdata/files"
		goFilePath := gdebug.CallerDirectory() + "/testdata/testdata.go"
		pkgName := "testdata"
		err := gres.PackToGoFile(srcPath, goFilePath, pkgName)
		gtest.Assert(err, nil)
	})
}

func Test_PackMulti(t *testing.T) {
	gtest.Case(t, func() {
		srcPath := gdebug.CallerDirectory() + "/testdata/files"
		goFilePath := gdebug.CallerDirectory() + "/testdata/data/data.go"
		pkgName := "data"
		array, err := gfile.ScanDir(srcPath, "*", false)
		gtest.Assert(err, nil)
		err = gres.PackToGoFile(strings.Join(array, ","), goFilePath, pkgName)
		gtest.Assert(err, nil)
	})
}

func Test_PackWithPrefix1(t *testing.T) {
	gtest.Case(t, func() {
		srcPath := gdebug.CallerDirectory() + "/testdata/files"
		goFilePath := gdebug.CallerDirectory() + "/testdata/testdata.go"
		pkgName := "testdata"
		err := gres.PackToGoFile(srcPath, goFilePath, pkgName, "www/gf-site/test")
		gtest.Assert(err, nil)
	})
}

func Test_PackWithPrefix2(t *testing.T) {
	gtest.Case(t, func() {
		srcPath := gdebug.CallerDirectory() + "/testdata/files"
		goFilePath := gdebug.CallerDirectory() + "/testdata/testdata.go"
		pkgName := "testdata"
		err := gres.PackToGoFile(srcPath, goFilePath, pkgName, "/var/www/gf-site/test")
		gtest.Assert(err, nil)
	})
}
