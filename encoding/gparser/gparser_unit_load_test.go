// Copyright 2017 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gparser_test

import (
	"io/ioutil"
	"testing"

	"github.com/snail007/gf/encoding/gparser"
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/os/gfile"
	"github.com/snail007/gf/test/gtest"
)

func Test_Load_JSON(t *testing.T) {
	data := []byte(`{"n":123456789, "m":{"k":"v"}, "a":[1,2,3]}`)
	// JSON
	gtest.Case(t, func() {
		j, err := gparser.LoadContent(data)
		gtest.Assert(err, nil)
		gtest.Assert(j.Get("n"), "123456789")
		gtest.Assert(j.Get("m"), g.Map{"k": "v"})
		gtest.Assert(j.Get("m.k"), "v")
		gtest.Assert(j.Get("a"), g.Slice{1, 2, 3})
		gtest.Assert(j.Get("a.1"), 2)
	})
	// JSON
	gtest.Case(t, func() {
		path := "test.json"
		gfile.PutBytes(path, data)
		defer gfile.Remove(path)
		j, err := gparser.Load(path)
		gtest.Assert(err, nil)
		gtest.Assert(j.Get("n"), "123456789")
		gtest.Assert(j.Get("m"), g.Map{"k": "v"})
		gtest.Assert(j.Get("m.k"), "v")
		gtest.Assert(j.Get("a"), g.Slice{1, 2, 3})
		gtest.Assert(j.Get("a.1"), 2)
	})
}

func Test_Load_XML(t *testing.T) {
	data := []byte(`<doc><a>1</a><a>2</a><a>3</a><m><k>v</k></m><n>123456789</n></doc>`)
	// XML
	gtest.Case(t, func() {
		j, err := gparser.LoadContent(data)
		gtest.Assert(err, nil)
		gtest.Assert(j.Get("doc.n"), "123456789")
		gtest.Assert(j.Get("doc.m"), g.Map{"k": "v"})
		gtest.Assert(j.Get("doc.m.k"), "v")
		gtest.Assert(j.Get("doc.a"), g.Slice{"1", "2", "3"})
		gtest.Assert(j.Get("doc.a.1"), 2)
	})
	// XML
	gtest.Case(t, func() {
		path := "test.xml"
		gfile.PutBytes(path, data)
		defer gfile.Remove(path)
		j, err := gparser.Load(path)
		gtest.Assert(err, nil)
		gtest.Assert(j.Get("doc.n"), "123456789")
		gtest.Assert(j.Get("doc.m"), g.Map{"k": "v"})
		gtest.Assert(j.Get("doc.m.k"), "v")
		gtest.Assert(j.Get("doc.a"), g.Slice{"1", "2", "3"})
		gtest.Assert(j.Get("doc.a.1"), 2)
	})

	// XML
	gtest.Case(t, func() {
		xml := `<?xml version="1.0"?>

	<Output type="o">
	<itotalSize>0</itotalSize>
	<ipageSize>1</ipageSize>
	<ipageIndex>2</ipageIndex>
	<itotalRecords>GF框架</itotalRecords>
	<nworkOrderDtos/>
	<nworkOrderFrontXML/>
	</Output>`
		j, err := gparser.LoadContent(xml)
		gtest.Assert(err, nil)
		gtest.Assert(j.Get("Output.ipageIndex"), "2")
		gtest.Assert(j.Get("Output.itotalRecords"), "GF框架")
	})
}

func Test_Load_YAML1(t *testing.T) {
	data := []byte(`
a:
- 1
- 2
- 3
m:
 k: v
"n": 123456789
    `)
	// YAML
	gtest.Case(t, func() {
		j, err := gparser.LoadContent(data)
		gtest.Assert(err, nil)
		gtest.Assert(j.Get("n"), "123456789")
		gtest.Assert(j.Get("m"), g.Map{"k": "v"})
		gtest.Assert(j.Get("m.k"), "v")
		gtest.Assert(j.Get("a"), g.Slice{1, 2, 3})
		gtest.Assert(j.Get("a.1"), 2)
	})
	// YAML
	gtest.Case(t, func() {
		path := "test.yaml"
		gfile.PutBytes(path, data)
		defer gfile.Remove(path)
		j, err := gparser.Load(path)
		gtest.Assert(err, nil)
		gtest.Assert(j.Get("n"), "123456789")
		gtest.Assert(j.Get("m"), g.Map{"k": "v"})
		gtest.Assert(j.Get("m.k"), "v")
		gtest.Assert(j.Get("a"), g.Slice{1, 2, 3})
		gtest.Assert(j.Get("a.1"), 2)
	})
}

func Test_Load_YAML2(t *testing.T) {
	data := []byte("i : 123456789")
	gtest.Case(t, func() {
		j, err := gparser.LoadContent(data)
		gtest.Assert(err, nil)
		gtest.Assert(j.Get("i"), "123456789")
	})
}

func Test_Load_TOML1(t *testing.T) {
	data := []byte(`
a = ["1", "2", "3"]
n = "123456789"

[m]
  k = "v"
`)
	// TOML
	gtest.Case(t, func() {
		j, err := gparser.LoadContent(data)
		gtest.Assert(err, nil)
		gtest.Assert(j.Get("n"), "123456789")
		gtest.Assert(j.Get("m"), g.Map{"k": "v"})
		gtest.Assert(j.Get("m.k"), "v")
		gtest.Assert(j.Get("a"), g.Slice{"1", "2", "3"})
		gtest.Assert(j.Get("a.1"), 2)
	})
	// TOML
	gtest.Case(t, func() {
		path := "test.toml"
		gfile.PutBytes(path, data)
		defer gfile.Remove(path)
		j, err := gparser.Load(path)
		gtest.Assert(err, nil)
		gtest.Assert(j.Get("n"), "123456789")
		gtest.Assert(j.Get("m"), g.Map{"k": "v"})
		gtest.Assert(j.Get("m.k"), "v")
		gtest.Assert(j.Get("a"), g.Slice{"1", "2", "3"})
		gtest.Assert(j.Get("a.1"), 2)
	})
}

func Test_Load_TOML2(t *testing.T) {
	data := []byte("i=123456789")
	gtest.Case(t, func() {
		j, err := gparser.LoadContent(data)
		gtest.Assert(err, nil)
		gtest.Assert(j.Get("i"), "123456789")
	})
}

func Test_Load_Nil(t *testing.T) {
	gtest.Case(t, func() {
		p := gparser.New(nil)
		gtest.Assert(p.Value(), nil)
		file := "test22222.json"
		filePath := gfile.Pwd() + gfile.Separator + file
		ioutil.WriteFile(filePath, []byte("{"), 0644)
		defer gfile.Remove(filePath)
		_, err := gparser.Load(file)
		gtest.AssertNE(err, nil)
		_, err = gparser.LoadContent("{")
		gtest.AssertNE(err, nil)
	})
}
