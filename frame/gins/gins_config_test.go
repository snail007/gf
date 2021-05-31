// Copyright 2017 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gins_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/snail007/gf/os/gcfg"

	"github.com/snail007/gf/frame/gins"
	"github.com/snail007/gf/os/gfile"
	"github.com/snail007/gf/os/gtime"
	"github.com/snail007/gf/test/gtest"
)

func Test_Config(t *testing.T) {
	config := `
# 模板引擎目录
viewpath = "/home/www/templates/"
test = "v=1"
# MySQL数据库配置
[database]
    [[database.default]]
        host     = "127.0.0.1"
        port     = "3306"
        user     = "root"
        pass     = ""
        name     = "test"
        type     = "mysql"
        role     = "master"
        charset  = "utf8"
        priority = "1"
    [[database.default]]
        host     = "127.0.0.1"
        port     = "3306"
        user     = "root"
        pass     = "8692651"
        name     = "test"
        type     = "mysql"
        role     = "master"
        charset  = "utf8"
        priority = "1"
# Redis数据库配置
[redis]
    disk  = "127.0.0.1:6379,0"
    cache = "127.0.0.1:6379,1"
`
	gtest.Case(t, func() {
		gtest.AssertNE(gins.Config(), nil)
	})

	// relative path
	gtest.Case(t, func() {
		path := "config.toml"
		err := gfile.PutContents(path, config)
		gtest.Assert(err, nil)
		defer gfile.Remove(path)
		defer gins.Config().Clear()
		gtest.Assert(gins.Config().Get("test"), "v=1")
		gtest.Assert(gins.Config().Get("database.default.1.host"), "127.0.0.1")
		gtest.Assert(gins.Config().Get("redis.disk"), "127.0.0.1:6379,0")
	})
	// for gfsnotify callbacks to refresh cache of config file
	time.Sleep(500 * time.Millisecond)

	// relative path, config folder
	gtest.Case(t, func() {
		path := "config/config.toml"
		err := gfile.PutContents(path, config)
		gtest.Assert(err, nil)
		defer gfile.Remove(path)
		defer gins.Config().Clear()
		gtest.Assert(gins.Config().Get("test"), "v=1")
		gtest.Assert(gins.Config().Get("database.default.1.host"), "127.0.0.1")
		gtest.Assert(gins.Config().Get("redis.disk"), "127.0.0.1:6379,0")
	})
	// for gfsnotify callbacks to refresh cache of config file
	time.Sleep(500 * time.Millisecond)

	gtest.Case(t, func() {
		path := "test.toml"
		err := gfile.PutContents(path, config)
		gtest.Assert(err, nil)
		defer gfile.Remove(path)
		defer gins.Config("test").Clear()
		gins.Config("test").SetFileName("test.toml")
		gtest.Assert(gins.Config("test").Get("test"), "v=1")
		gtest.Assert(gins.Config("test").Get("database.default.1.host"), "127.0.0.1")
		gtest.Assert(gins.Config("test").Get("redis.disk"), "127.0.0.1:6379,0")
	})
	// for gfsnotify callbacks to refresh cache of config file
	time.Sleep(500 * time.Millisecond)

	gtest.Case(t, func() {
		path := "config/test.toml"
		err := gfile.PutContents(path, config)
		gtest.Assert(err, nil)
		defer gfile.Remove(path)
		defer gins.Config("test").Clear()
		gins.Config("test").SetFileName("test.toml")
		gtest.Assert(gins.Config("test").Get("test"), "v=1")
		gtest.Assert(gins.Config("test").Get("database.default.1.host"), "127.0.0.1")
		gtest.Assert(gins.Config("test").Get("redis.disk"), "127.0.0.1:6379,0")
	})
	// for gfsnotify callbacks to refresh cache of config file
	time.Sleep(500 * time.Millisecond)

	// absolute path
	gtest.Case(t, func() {
		path := fmt.Sprintf(`%s/%d`, gfile.TempDir(), gtime.Nanosecond())
		file := fmt.Sprintf(`%s/%s`, path, "config.toml")
		err := gfile.PutContents(file, config)
		gtest.Assert(err, nil)
		defer gfile.Remove(file)
		defer gins.Config().Clear()
		gtest.Assert(gins.Config().AddPath(path), nil)
		gtest.Assert(gins.Config().Get("test"), "v=1")
		gtest.Assert(gins.Config().Get("database.default.1.host"), "127.0.0.1")
		gtest.Assert(gins.Config().Get("redis.disk"), "127.0.0.1:6379,0")
	})
	time.Sleep(500 * time.Millisecond)

	gtest.Case(t, func() {
		path := fmt.Sprintf(`%s/%d/config`, gfile.TempDir(), gtime.Nanosecond())
		file := fmt.Sprintf(`%s/%s`, path, "config.toml")
		err := gfile.PutContents(file, config)
		gtest.Assert(err, nil)
		defer gfile.Remove(file)
		defer gins.Config().Clear()
		gtest.Assert(gins.Config().AddPath(path), nil)
		gtest.Assert(gins.Config().Get("test"), "v=1")
		gtest.Assert(gins.Config().Get("database.default.1.host"), "127.0.0.1")
		gtest.Assert(gins.Config().Get("redis.disk"), "127.0.0.1:6379,0")
	})
	time.Sleep(500 * time.Millisecond)

	gtest.Case(t, func() {
		path := fmt.Sprintf(`%s/%d`, gfile.TempDir(), gtime.Nanosecond())
		file := fmt.Sprintf(`%s/%s`, path, "test.toml")
		err := gfile.PutContents(file, config)
		gtest.Assert(err, nil)
		defer gfile.Remove(file)
		defer gins.Config("test").Clear()
		gins.Config("test").SetFileName("test.toml")
		gtest.Assert(gins.Config("test").AddPath(path), nil)
		gtest.Assert(gins.Config("test").Get("test"), "v=1")
		gtest.Assert(gins.Config("test").Get("database.default.1.host"), "127.0.0.1")
		gtest.Assert(gins.Config("test").Get("redis.disk"), "127.0.0.1:6379,0")
	})
	time.Sleep(500 * time.Millisecond)

	gtest.Case(t, func() {
		path := fmt.Sprintf(`%s/%d/config`, gfile.TempDir(), gtime.Nanosecond())
		file := fmt.Sprintf(`%s/%s`, path, "test.toml")
		err := gfile.PutContents(file, config)
		gtest.Assert(err, nil)
		defer gfile.Remove(file)
		defer gins.Config().Clear()
		gins.Config("test").SetFileName("test.toml")
		gtest.Assert(gins.Config("test").AddPath(path), nil)
		gtest.Assert(gins.Config("test").Get("test"), "v=1")
		gtest.Assert(gins.Config("test").Get("database.default.1.host"), "127.0.0.1")
		gtest.Assert(gins.Config("test").Get("redis.disk"), "127.0.0.1:6379,0")
	})
}

func Test_Basic2(t *testing.T) {
	config := `log-path = "logs"`
	gtest.Case(t, func() {
		path := gcfg.DEFAULT_CONFIG_FILE
		err := gfile.PutContents(path, config)
		gtest.Assert(err, nil)
		defer func() {
			_ = gfile.Remove(path)
		}()

		gtest.Assert(gins.Config().Get("log-path"), "logs")
	})
}
