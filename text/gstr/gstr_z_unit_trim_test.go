// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

// go test *.go -bench=".*"

package gstr_test

import (
	"testing"

	"github.com/snail007/gf/test/gtest"
	"github.com/snail007/gf/text/gstr"
)

func Test_Trim(t *testing.T) {
	gtest.Case(t, func() {
		gtest.Assert(gstr.Trim(" 123456\n "), "123456")
		gtest.Assert(gstr.Trim("#123456#;", "#;"), "123456")
	})
}

func Test_TrimRight(t *testing.T) {
	gtest.Case(t, func() {
		gtest.Assert(gstr.TrimRight(" 123456\n "), " 123456")
		gtest.Assert(gstr.TrimRight("#123456#;", "#;"), "#123456")
	})
}

func Test_TrimRightStr(t *testing.T) {
	gtest.Case(t, func() {
		gtest.Assert(gstr.TrimRightStr("gogo我爱gogo", "go"), "gogo我爱")
	})
}

func Test_TrimLeft(t *testing.T) {
	gtest.Case(t, func() {
		gtest.Assert(gstr.TrimLeft(" \r123456\n "), "123456\n ")
		gtest.Assert(gstr.TrimLeft("#;123456#;", "#;"), "123456#;")
	})
}

func Test_TrimLeftStr(t *testing.T) {
	gtest.Case(t, func() {
		gtest.Assert(gstr.TrimLeftStr("gogo我爱gogo", "go"), "我爱gogo")
	})
}
