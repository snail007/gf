// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gvalid_test

import (
	"strings"
	"testing"

	"github.com/snail007/gf/test/gtest"
	"github.com/snail007/gf/util/gvalid"
)

func Test_Map(t *testing.T) {
	rule := "ipv4"
	val := "0.0.0"
	msg := map[string]string{
		"ipv4": "IPv4地址格式不正确",
	}

	err := gvalid.Check(val, rule, nil)
	gtest.Assert(err.Map(), msg)
}

func Test_FirstString(t *testing.T) {
	rule := "ipv4"
	val := "0.0.0"

	err := gvalid.Check(val, rule, nil)
	n := err.FirstString()
	gtest.Assert(n, "IPv4地址格式不正确")
}

func Test_SetDefaultErrorMsgs(t *testing.T) {
	rule := "integer|length:6,16"
	msgs := map[string]string{
		"integer": "请输入一个整数",
		"length":  "参数长度不对啊老铁",
	}
	gvalid.SetDefaultErrorMsgs(msgs)
	e := gvalid.Check("6.66", rule, nil)
	if e == nil || len(e.Map()) != 2 {
		t.Error("规则校验失败")
	} else {
		if v, ok := e.Map()["integer"]; ok {
			if strings.Compare(v, msgs["integer"]) != 0 {
				t.Error("错误信息不匹配")
			}
		}
		if v, ok := e.Map()["length"]; ok {
			if strings.Compare(v, msgs["length"]) != 0 {
				t.Error("错误信息不匹配")
			}
		}
	}
}

func Test_CustomError1(t *testing.T) {
	rule := "integer|length:6,16"
	msgs := map[string]string{
		"integer": "请输入一个整数",
		"length":  "参数长度不对啊老铁",
	}
	e := gvalid.Check("6.66", rule, msgs)
	if e == nil || len(e.Map()) != 2 {
		t.Error("规则校验失败")
	} else {
		if v, ok := e.Map()["integer"]; ok {
			if strings.Compare(v, msgs["integer"]) != 0 {
				t.Error("错误信息不匹配")
			}
		}
		if v, ok := e.Map()["length"]; ok {
			if strings.Compare(v, msgs["length"]) != 0 {
				t.Error("错误信息不匹配")
			}
		}
	}
}

func Test_CustomError2(t *testing.T) {
	rule := "integer|length:6,16"
	msgs := "请输入一个整数|参数长度不对啊老铁"
	e := gvalid.Check("6.66", rule, msgs)
	if e == nil || len(e.Map()) != 2 {
		t.Error("规则校验失败")
	} else {
		if v, ok := e.Map()["integer"]; ok {
			if strings.Compare(v, "请输入一个整数") != 0 {
				t.Error("错误信息不匹配")
			}
		}
		if v, ok := e.Map()["length"]; ok {
			if strings.Compare(v, "参数长度不对啊老铁") != 0 {
				t.Error("错误信息不匹配")
			}
		}
	}
}
