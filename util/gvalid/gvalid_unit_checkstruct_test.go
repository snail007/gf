// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gvalid_test

import (
	"testing"

	"github.com/snail007/gf/frame/g"

	"github.com/snail007/gf/test/gtest"
	"github.com/snail007/gf/util/gvalid"
)

func Test_CheckStruct(t *testing.T) {
	gtest.Case(t, func() {
		type Object struct {
			Name string
			Age  int
		}
		rules := []string{
			"@required|length:6,16",
			"@between:18,30",
		}
		msgs := map[string]interface{}{
			"Name": map[string]string{
				"required": "名称不能为空",
				"length":   "名称长度为:min到:max个字符",
			},
			"Age": "年龄为18到30周岁",
		}
		obj := &Object{"john", 16}
		err := gvalid.CheckStruct(obj, rules, msgs)
		gtest.Assert(err, nil)
	})

	gtest.Case(t, func() {
		type Object struct {
			Name string
			Age  int
		}
		rules := []string{
			"Name@required|length:6,16#名称不能为空",
			"Age@between:18,30",
		}
		msgs := map[string]interface{}{
			"Name": map[string]string{
				"required": "名称不能为空",
				"length":   "名称长度为:min到:max个字符",
			},
			"Age": "年龄为18到30周岁",
		}
		obj := &Object{"john", 16}
		err := gvalid.CheckStruct(obj, rules, msgs)
		gtest.AssertNE(err, nil)
		gtest.Assert(len(err.Maps()), 2)
		gtest.Assert(err.Maps()["Name"]["required"], "")
		gtest.Assert(err.Maps()["Name"]["length"], "名称长度为6到16个字符")
		gtest.Assert(err.Maps()["Age"]["between"], "年龄为18到30周岁")
	})

	gtest.Case(t, func() {
		type Object struct {
			Name string
			Age  int
		}
		rules := []string{
			"Name@required|length:6,16#名称不能为空|",
			"Age@between:18,30",
		}
		msgs := map[string]interface{}{
			"Name": map[string]string{
				"required": "名称不能为空",
				"length":   "名称长度为:min到:max个字符",
			},
			"Age": "年龄为18到30周岁",
		}
		obj := &Object{"john", 16}
		err := gvalid.CheckStruct(obj, rules, msgs)
		gtest.AssertNE(err, nil)
		gtest.Assert(len(err.Maps()), 2)
		gtest.Assert(err.Maps()["Name"]["required"], "")
		gtest.Assert(err.Maps()["Name"]["length"], "名称长度为6到16个字符")
		gtest.Assert(err.Maps()["Age"]["between"], "年龄为18到30周岁")
	})

	gtest.Case(t, func() {
		type Object struct {
			Name string
			Age  int
		}
		rules := map[string]string{
			"Name": "required|length:6,16",
			"Age":  "between:18,30",
		}
		msgs := map[string]interface{}{
			"Name": map[string]string{
				"required": "名称不能为空",
				"length":   "名称长度为:min到:max个字符",
			},
			"Age": "年龄为18到30周岁",
		}
		obj := &Object{"john", 16}
		err := gvalid.CheckStruct(obj, rules, msgs)
		gtest.AssertNE(err, nil)
		gtest.Assert(len(err.Maps()), 2)
		gtest.Assert(err.Maps()["Name"]["required"], "")
		gtest.Assert(err.Maps()["Name"]["length"], "名称长度为6到16个字符")
		gtest.Assert(err.Maps()["Age"]["between"], "年龄为18到30周岁")
	})

	gtest.Case(t, func() {
		type LoginRequest struct {
			Username string `json:"username" gvalid:"username@required#用户名不能为空"`
			Password string `json:"password" gvalid:"password@required#登录密码不能为空"`
		}
		var login LoginRequest
		err := gvalid.CheckStruct(login, nil)
		gtest.AssertNE(err, nil)
		gtest.Assert(len(err.Maps()), 2)
		gtest.Assert(err.Maps()["username"]["required"], "用户名不能为空")
		gtest.Assert(err.Maps()["password"]["required"], "登录密码不能为空")
	})

	gtest.Case(t, func() {
		type LoginRequest struct {
			Username string `json:"username" gvalid:"@required#用户名不能为空"`
			Password string `json:"password" gvalid:"@required#登录密码不能为空"`
		}
		var login LoginRequest
		err := gvalid.CheckStruct(login, nil)
		gtest.Assert(err, nil)
	})

	gtest.Case(t, func() {
		type LoginRequest struct {
			username string `json:"username" gvalid:"username@required#用户名不能为空"`
			Password string `json:"password" gvalid:"password@required#登录密码不能为空"`
		}
		var login LoginRequest
		err := gvalid.CheckStruct(login, nil)
		gtest.AssertNE(err, nil)
		gtest.Assert(err.Maps()["password"]["required"], "登录密码不能为空")
	})

	// gvalid tag
	gtest.Case(t, func() {
		type User struct {
			Id       int    `gvalid:"uid@required|min:10#|ID不能为空"`
			Age      int    `gvalid:"age@required#年龄不能为空"`
			Username string `json:"username" gvalid:"username@required#用户名不能为空"`
			Password string `json:"password" gvalid:"password@required#登录密码不能为空"`
		}
		user := &User{
			Id:       1,
			Username: "john",
			Password: "123456",
		}
		err := gvalid.CheckStruct(user, nil)
		gtest.AssertNE(err, nil)
		gtest.Assert(len(err.Maps()), 1)
		gtest.Assert(err.Maps()["uid"]["min"], "ID不能为空")
	})

	gtest.Case(t, func() {
		type User struct {
			Id       int    `gvalid:"uid@required|min:10#|ID不能为空"`
			Age      int    `gvalid:"age@required#年龄不能为空"`
			Username string `json:"username" gvalid:"username@required#用户名不能为空"`
			Password string `json:"password" gvalid:"password@required#登录密码不能为空"`
		}
		user := &User{
			Id:       1,
			Username: "john",
			Password: "123456",
		}

		rules := []string{
			"username@required#用户名不能为空",
		}

		err := gvalid.CheckStruct(user, rules)
		gtest.AssertNE(err, nil)
		gtest.Assert(len(err.Maps()), 1)
		gtest.Assert(err.Maps()["uid"]["min"], "ID不能为空")
	})

	gtest.Case(t, func() {
		type User struct {
			Id       int    `gvalid:"uid@required|min:10#ID不能为空"`
			Age      int    `gvalid:"age@required#年龄不能为空"`
			Username string `json:"username" gvalid:"username@required#用户名不能为空"`
			Password string `json:"password" gvalid:"password@required#登录密码不能为空"`
		}
		user := &User{
			Id:       1,
			Username: "john",
			Password: "123456",
		}
		err := gvalid.CheckStruct(user, nil)
		gtest.AssertNE(err, nil)
		gtest.Assert(len(err.Maps()), 1)
	})

	// valid tag
	gtest.Case(t, func() {
		type User struct {
			Id       int    `valid:"uid@required|min:10#|ID不能为空"`
			Age      int    `valid:"age@required#年龄不能为空"`
			Username string `json:"username" gvalid:"username@required#用户名不能为空"`
			Password string `json:"password" gvalid:"password@required#登录密码不能为空"`
		}
		user := &User{
			Id:       1,
			Username: "john",
			Password: "123456",
		}
		err := gvalid.CheckStruct(user, nil)
		gtest.AssertNE(err, nil)
		gtest.Assert(len(err.Maps()), 1)
		gtest.Assert(err.Maps()["uid"]["min"], "ID不能为空")
	})
}

func Test_CheckStruct_With_Inherit(t *testing.T) {
	gtest.Case(t, func() {
		type Pass struct {
			Pass1 string `valid:"password1@required|same:password2#请输入您的密码|您两次输入的密码不一致"`
			Pass2 string `valid:"password2@required|same:password1#请再次输入您的密码|您两次输入的密码不一致"`
		}
		type User struct {
			Id   int
			Name string `valid:"name@required#请输入您的姓名"`
			Pass Pass
		}
		user := &User{
			Name: "",
			Pass: Pass{
				Pass1: "1",
				Pass2: "2",
			},
		}
		err := gvalid.CheckStruct(user, nil)
		gtest.AssertNE(err, nil)
		gtest.Assert(err.Maps()["name"], g.Map{"required": "请输入您的姓名"})
		gtest.Assert(err.Maps()["password1"], g.Map{"same": "您两次输入的密码不一致"})
		gtest.Assert(err.Maps()["password2"], g.Map{"same": "您两次输入的密码不一致"})
	})
}
