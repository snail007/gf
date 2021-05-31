// Copyright 2017 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.
package ghtml_test

import (
	"testing"

	"github.com/snail007/gf/encoding/ghtml"
	"github.com/snail007/gf/test/gtest"
)

func TestStripTags(t *testing.T) {
	src := `<p>Test paragraph.</p><!-- Comment -->  <a href="#fragment">Other text</a>`
	dst := `Test paragraph.  Other text`
	gtest.Assert(ghtml.StripTags(src), dst)
}

func TestEntities(t *testing.T) {
	src := `A 'quote' "is" <b>bold</b>`
	dst := `A &#39;quote&#39; &#34;is&#34; &lt;b&gt;bold&lt;/b&gt;`
	gtest.Assert(ghtml.Entities(src), dst)
	gtest.Assert(ghtml.EntitiesDecode(dst), src)
}

func TestSpecialChars(t *testing.T) {
	src := `A 'quote' "is" <b>bold</b>`
	dst := `A &#39;quote&#39; &#34;is&#34; &lt;b&gt;bold&lt;/b&gt;`
	gtest.Assert(ghtml.SpecialChars(src), dst)
	gtest.Assert(ghtml.SpecialCharsDecode(dst), src)
}
