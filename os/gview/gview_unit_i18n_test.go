// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gview_test

import (
	"testing"

	"github.com/snail007/gf/debug/gdebug"
	"github.com/snail007/gf/os/gfile"

	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/test/gtest"
)

func Test_I18n(t *testing.T) {
	gtest.Case(t, func() {
		content := `{{.name}} says "{#hello}{#world}!"`
		expect1 := `john says "你好世界!"`
		expect2 := `john says "こんにちは世界!"`
		expect3 := `john says "{#hello}{#world}!"`

		g.I18n().SetPath(gdebug.CallerDirectory() + gfile.Separator + "testdata" + gfile.Separator + "i18n")

		g.I18n().SetLanguage("zh-CN")
		result1, err := g.View().ParseContent(content, g.Map{
			"name": "john",
		})
		gtest.Assert(err, nil)
		gtest.Assert(result1, expect1)

		g.I18n().SetLanguage("ja")
		result2, err := g.View().ParseContent(content, g.Map{
			"name": "john",
		})
		gtest.Assert(err, nil)
		gtest.Assert(result2, expect2)

		g.I18n().SetLanguage("none")
		result3, err := g.View().ParseContent(content, g.Map{
			"name": "john",
		})
		gtest.Assert(err, nil)
		gtest.Assert(result3, expect3)
	})
}
