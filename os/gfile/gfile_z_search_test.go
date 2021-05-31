// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gfile_test

import (
	"path/filepath"
	"testing"

	"github.com/snail007/gf/os/gfile"
	"github.com/snail007/gf/test/gtest"
)

func Test_Search(t *testing.T) {
	gtest.Case(t, func() {
		var (
			paths1  string = "/testfiless"
			paths2  string = "./testfile/dirfiles_no"
			tpath   string
			tpath2  string
			tempstr string
			ypaths1 string
			err     error
		)

		createDir(paths1)
		defer delTestFiles(paths1)
		ypaths1 = paths1

		tpath, err = gfile.Search(testpath() + paths1)
		gtest.Assert(err, nil)

		tpath = filepath.ToSlash(tpath)

		// 自定义优先路径
		tpath2, err = gfile.Search(testpath() + paths1)
		gtest.Assert(err, nil)
		tpath2 = filepath.ToSlash(tpath2)

		tempstr = testpath()
		paths1 = tempstr + paths1
		paths1 = filepath.ToSlash(paths1)

		gtest.Assert(tpath, paths1)

		gtest.Assert(tpath2, tpath)

		// 测试给定目录
		tpath2, err = gfile.Search(paths1, "testfiless")
		tpath2 = filepath.ToSlash(tpath2)
		tempss := filepath.ToSlash(paths1)
		gtest.Assert(tpath2, tempss)

		// 测试当前目录
		tempstr, _ = filepath.Abs("./")
		tempstr = testpath()
		paths1 = tempstr + ypaths1
		paths1 = filepath.ToSlash(paths1)

		gtest.Assert(tpath2, paths1)

		// 测试目录不存在时
		_, err = gfile.Search(paths2)
		gtest.AssertNE(err, nil)

	})
}
