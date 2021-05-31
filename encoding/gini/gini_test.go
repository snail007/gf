// Copyright 2017 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gini_test

import (
	"github.com/snail007/gf/encoding/gini"
	"github.com/snail007/gf/encoding/gjson"
	"github.com/snail007/gf/test/gtest"
	"testing"
)

var iniContent = `

;注释
aa=bb
[addr] 
#注释
ip = 127.0.0.1
port=9001
enable=true

	[DBINFO]
	type=mysql
	user=root
	password=password
[键]
呵呵=值

`

func TestDecode(t *testing.T) {
	gtest.Case(t, func() {
		res, err := gini.Decode([]byte(iniContent))
		if err != nil {
			gtest.Fatal(err)
		}
		gtest.Assert(res["addr"].(map[string]interface{})["ip"], "127.0.0.1")
		gtest.Assert(res["addr"].(map[string]interface{})["port"], "9001")
		gtest.Assert(res["DBINFO"].(map[string]interface{})["user"], "root")
		gtest.Assert(res["DBINFO"].(map[string]interface{})["type"], "mysql")
		gtest.Assert(res["键"].(map[string]interface{})["呵呵"], "值")
	})

	gtest.Case(t, func() {
		errContent := `
		a = b
`
		_, err := gini.Decode([]byte(errContent))
		if err == nil {
			gtest.Fatal(err)
		}
	})
}

func TestEncode(t *testing.T) {
	gtest.Case(t, func() {
		iniMap, err := gini.Decode([]byte(iniContent))
		if err != nil {
			gtest.Fatal(err)
		}

		iniStr, err := gini.Encode(iniMap)
		if err != nil {
			gtest.Fatal(err)
		}

		res, err := gini.Decode(iniStr)
		if err != nil {
			gtest.Fatal(err)
		}

		gtest.Assert(res["addr"].(map[string]interface{})["ip"], "127.0.0.1")
		gtest.Assert(res["addr"].(map[string]interface{})["port"], "9001")
		gtest.Assert(res["DBINFO"].(map[string]interface{})["user"], "root")
		gtest.Assert(res["DBINFO"].(map[string]interface{})["type"], "mysql")

	})
}

func TestToJson(t *testing.T) {
	gtest.Case(t, func() {
		jsonStr, err := gini.ToJson([]byte(iniContent))
		if err != nil {
			gtest.Fatal(err)
		}

		json, err := gjson.LoadContent(jsonStr)
		if err != nil {
			gtest.Fatal(err)
		}

		iniMap, err := gini.Decode([]byte(iniContent))
		gtest.Assert(err, nil)

		gtest.Assert(iniMap["addr"].(map[string]interface{})["ip"], json.GetString("addr.ip"))
		gtest.Assert(iniMap["addr"].(map[string]interface{})["port"], json.GetString("addr.port"))
		gtest.Assert(iniMap["DBINFO"].(map[string]interface{})["user"], json.GetString("DBINFO.user"))
		gtest.Assert(iniMap["DBINFO"].(map[string]interface{})["type"], json.GetString("DBINFO.type"))
	})
}
