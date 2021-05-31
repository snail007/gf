// Copyright 2017 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

// go test *.go -bench=".*" -benchmem

package gfcache_test

import (
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/snail007/gf/os/gfcache"
	"github.com/snail007/gf/os/gfile"
	"github.com/snail007/gf/test/gtest"
)

func TestGetContents(t *testing.T) {
	gtest.Case(t, func() {

		var f *os.File
		var err error
		fileName := "test"
		strTest := "123"

		if !gfile.Exists(fileName) {
			f, err = ioutil.TempFile("", fileName)
			if err != nil {
				t.Error("create file fail")
			}
		}

		defer f.Close()
		defer os.Remove(f.Name())

		if gfile.Exists(f.Name()) {

			f, err = gfile.OpenFile(f.Name(), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				t.Error("file open fail", err)
			}

			err = gfile.PutContents(f.Name(), strTest)
			if err != nil {
				t.Error("write error", err)
			}

			cache := gfcache.GetContents(f.Name(), 1)
			gtest.Assert(cache, strTest)
		}
	})

	gtest.Case(t, func() {

		var f *os.File
		var err error
		fileName := "test2"
		strTest := "123"

		if !gfile.Exists(fileName) {
			f, err = ioutil.TempFile("", fileName)
			if err != nil {
				t.Error("create file fail")
			}
		}

		defer f.Close()
		defer os.Remove(f.Name())

		if gfile.Exists(f.Name()) {
			cache := gfcache.GetContents(f.Name())

			f, err = gfile.OpenFile(f.Name(), os.O_APPEND|os.O_WRONLY, os.ModeAppend)
			if err != nil {
				t.Error("file open fail", err)
			}

			err = gfile.PutContents(f.Name(), strTest)
			if err != nil {
				t.Error("write error", err)
			}

			gtest.Assert(cache, "")

			time.Sleep(100 * time.Millisecond)
		}
	})
}
