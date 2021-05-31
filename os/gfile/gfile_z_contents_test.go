// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gfile_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/snail007/gf/text/gstr"

	"github.com/snail007/gf/os/gfile"
	"github.com/snail007/gf/test/gtest"
)

// 创建测试文件
func createTestFile(filename, content string) error {
	TempDir := testpath()
	err := ioutil.WriteFile(TempDir+filename, []byte(content), 0666)
	return err
}

// 测试完删除文件或目录
func delTestFiles(filenames string) {
	os.RemoveAll(testpath() + filenames)
}

// 创建目录
func createDir(paths string) {
	TempDir := testpath()
	os.Mkdir(TempDir+paths, 0777)
}

// 统一格式化文件目录为"/"
func formatpaths(paths []string) []string {
	for k, v := range paths {
		paths[k] = filepath.ToSlash(v)
		paths[k] = strings.Replace(paths[k], "./", "/", 1)
	}

	return paths
}

// 统一格式化文件目录为"/"
func formatpath(paths string) string {
	paths = filepath.ToSlash(paths)
	paths = strings.Replace(paths, "./", "/", 1)
	return paths
}

// 指定返回要测试的目录
func testpath() string {
	return gstr.TrimRight(os.TempDir(), "\\/")
}

func Test_GetContents(t *testing.T) {
	gtest.Case(t, func() {

		var (
			filepaths string = "/testfile_t1.txt"
		)
		createTestFile(filepaths, "my name is jroam")
		defer delTestFiles(filepaths)

		gtest.Assert(gfile.GetContents(testpath()+filepaths), "my name is jroam")
		gtest.Assert(gfile.GetContents(""), "")

	})
}

func Test_GetBinContents(t *testing.T) {
	gtest.Case(t, func() {
		var (
			filepaths1  string = "/testfile_t1.txt"                 // 文件存在时
			filepaths2  string = testpath() + "/testfile_t1_no.txt" // 文件不存在时
			readcontent []byte
			str1        string = "my name is jroam"
		)
		createTestFile(filepaths1, str1)
		defer delTestFiles(filepaths1)
		readcontent = gfile.GetBytes(testpath() + filepaths1)
		gtest.Assert(readcontent, []byte(str1))

		readcontent = gfile.GetBytes(filepaths2)
		gtest.Assert(string(readcontent), "")

		gtest.Assert(string(gfile.GetBytes(filepaths2)), "")

	})
}

// 截断文件为指定的大小
func Test_Truncate(t *testing.T) {
	gtest.Case(t, func() {
		var (
			filepaths1 string = "/testfile_GetContentsyyui.txt" //文件存在时
			err        error
			files      *os.File
		)
		createTestFile(filepaths1, "abcdefghijkmln")
		defer delTestFiles(filepaths1)
		err = gfile.Truncate(testpath()+filepaths1, 10)
		gtest.Assert(err, nil)

		//=========================检查修改文后的大小，是否与期望一致
		files, err = os.Open(testpath() + filepaths1)
		defer files.Close()
		gtest.Assert(err, nil)
		fileinfo, err2 := files.Stat()
		gtest.Assert(err2, nil)
		gtest.Assert(fileinfo.Size(), 10)

		//====测试当为空时，是否报错
		err = gfile.Truncate("", 10)
		gtest.AssertNE(err, nil)

	})
}

func Test_PutContents(t *testing.T) {
	gtest.Case(t, func() {
		var (
			filepaths   string = "/testfile_PutContents.txt"
			err         error
			readcontent []byte
		)
		createTestFile(filepaths, "a")
		defer delTestFiles(filepaths)

		err = gfile.PutContents(testpath()+filepaths, "test!")
		gtest.Assert(err, nil)

		//==================判断是否真正写入
		readcontent, err = ioutil.ReadFile(testpath() + filepaths)
		gtest.Assert(err, nil)
		gtest.Assert(string(readcontent), "test!")

		err = gfile.PutContents("", "test!")
		gtest.AssertNE(err, nil)

	})
}

func Test_PutContentsAppend(t *testing.T) {
	gtest.Case(t, func() {
		var (
			filepaths   string = "/testfile_PutContents.txt"
			err         error
			readcontent []byte
		)

		createTestFile(filepaths, "a")
		defer delTestFiles(filepaths)
		err = gfile.PutContentsAppend(testpath()+filepaths, "hello")
		gtest.Assert(err, nil)

		//==================判断是否真正写入
		readcontent, err = ioutil.ReadFile(testpath() + filepaths)
		gtest.Assert(err, nil)
		gtest.Assert(string(readcontent), "ahello")

		err = gfile.PutContentsAppend("", "hello")
		gtest.AssertNE(err, nil)

	})

}

func Test_PutBinContents(t *testing.T) {
	gtest.Case(t, func() {
		var (
			filepaths   string = "/testfile_PutContents.txt"
			err         error
			readcontent []byte
		)
		createTestFile(filepaths, "a")
		defer delTestFiles(filepaths)

		err = gfile.PutBytes(testpath()+filepaths, []byte("test!!"))
		gtest.Assert(err, nil)

		// 判断是否真正写入
		readcontent, err = ioutil.ReadFile(testpath() + filepaths)
		gtest.Assert(err, nil)
		gtest.Assert(string(readcontent), "test!!")

		err = gfile.PutBytes("", []byte("test!!"))
		gtest.AssertNE(err, nil)

	})
}

func Test_PutBinContentsAppend(t *testing.T) {
	gtest.Case(t, func() {
		var (
			filepaths   string = "/testfile_PutContents.txt" //原文件内容: yy
			err         error
			readcontent []byte
		)
		createTestFile(filepaths, "test!!")
		defer delTestFiles(filepaths)
		err = gfile.PutBytesAppend(testpath()+filepaths, []byte("word"))
		gtest.Assert(err, nil)

		// 判断是否真正写入
		readcontent, err = ioutil.ReadFile(testpath() + filepaths)
		gtest.Assert(err, nil)
		gtest.Assert(string(readcontent), "test!!word")

		err = gfile.PutBytesAppend("", []byte("word"))
		gtest.AssertNE(err, nil)

	})
}

func Test_GetBinContentsByTwoOffsetsByPath(t *testing.T) {
	gtest.Case(t, func() {
		var (
			filepaths   string = "/testfile_GetContents.txt" // 文件内容: abcdefghijk
			readcontent []byte
		)

		createTestFile(filepaths, "abcdefghijk")
		defer delTestFiles(filepaths)
		readcontent = gfile.GetBytesByTwoOffsetsByPath(testpath()+filepaths, 2, 5)

		gtest.Assert(string(readcontent), "cde")

		readcontent = gfile.GetBytesByTwoOffsetsByPath("", 2, 5)
		gtest.Assert(len(readcontent), 0)

	})

}

func Test_GetNextCharOffsetByPath(t *testing.T) {
	gtest.Case(t, func() {
		var (
			filepaths  string = "/testfile_GetContents.txt" // 文件内容: abcdefghijk
			localindex int64
		)
		createTestFile(filepaths, "abcdefghijk")
		defer delTestFiles(filepaths)
		localindex = gfile.GetNextCharOffsetByPath(testpath()+filepaths, 'd', 1)
		gtest.Assert(localindex, 3)

		localindex = gfile.GetNextCharOffsetByPath("", 'd', 1)
		gtest.Assert(localindex, -1)

	})
}

func Test_GetNextCharOffset(t *testing.T) {
	gtest.Case(t, func() {
		var (
			localindex int64
		)
		reader := strings.NewReader("helloword")

		localindex = gfile.GetNextCharOffset(reader, 'w', 1)
		gtest.Assert(localindex, 5)

		localindex = gfile.GetNextCharOffset(reader, 'j', 1)
		gtest.Assert(localindex, -1)

	})
}

func Test_GetBinContentsByTwoOffsets(t *testing.T) {
	gtest.Case(t, func() {
		var (
			reads []byte
		)
		reader := strings.NewReader("helloword")

		reads = gfile.GetBytesByTwoOffsets(reader, 1, 3)
		gtest.Assert(string(reads), "el")

		reads = gfile.GetBytesByTwoOffsets(reader, 10, 30)
		gtest.Assert(string(reads), "")

	})
}

func Test_GetBinContentsTilChar(t *testing.T) {
	gtest.Case(t, func() {
		var (
			reads  []byte
			indexs int64
		)
		reader := strings.NewReader("helloword")

		reads, _ = gfile.GetBytesTilChar(reader, 'w', 2)
		gtest.Assert(string(reads), "llow")

		_, indexs = gfile.GetBytesTilChar(reader, 'w', 20)
		gtest.Assert(indexs, -1)

	})
}

func Test_GetBinContentsTilCharByPath(t *testing.T) {
	gtest.Case(t, func() {
		var (
			reads     []byte
			indexs    int64
			filepaths string = "/testfile_GetContents.txt"
		)

		createTestFile(filepaths, "abcdefghijklmn")
		defer delTestFiles(filepaths)

		reads, _ = gfile.GetBytesTilCharByPath(testpath()+filepaths, 'c', 2)
		gtest.Assert(string(reads), "c")

		reads, _ = gfile.GetBytesTilCharByPath(testpath()+filepaths, 'y', 1)
		gtest.Assert(string(reads), "")

		_, indexs = gfile.GetBytesTilCharByPath(testpath()+filepaths, 'x', 1)
		gtest.Assert(indexs, -1)

	})
}

func Test_Home(t *testing.T) {
	gtest.Case(t, func() {
		var (
			reads string
			err   error
		)

		reads, err = gfile.Home()
		gtest.Assert(err, nil)
		gtest.AssertNE(reads, "")

	})
}
