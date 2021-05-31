// Copyright 2017 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.
package gurl_test

import (
	"net/url"
	"testing"

	"github.com/snail007/gf/encoding/gurl"
	"github.com/snail007/gf/test/gtest"
)

var urlStr string = `https://golang.org/x/crypto?go-get=1 +`
var urlEncode string = `https%3A%2F%2Fgolang.org%2Fx%2Fcrypto%3Fgo-get%3D1+%2B`
var rawUrlEncode string = `https%3A%2F%2Fgolang.org%2Fx%2Fcrypto%3Fgo-get%3D1%20%2B`

func TestEncodeAndDecode(t *testing.T) {
	gtest.Case(t, func() {
		gtest.Assert(gurl.Encode(urlStr), urlEncode)

		res, err := gurl.Decode(urlEncode)
		if err != nil {
			t.Errorf("decode failed. %v", err)
			return
		}
		gtest.Assert(res, urlStr)
	})
}

func TestRowEncodeAndDecode(t *testing.T) {
	gtest.Case(t, func() {
		gtest.Assert(gurl.RawEncode(urlStr), rawUrlEncode)

		res, err := gurl.RawDecode(rawUrlEncode)
		if err != nil {
			t.Errorf("decode failed. %v", err)
			return
		}
		gtest.Assert(res, urlStr)
	})
}

func TestBuildQuery(t *testing.T) {
	src := url.Values{
		"a": {"a2", "a1"},
		"b": {"b2", "b1"},
		"c": {"c1", "c2"},
	}
	expect := "a=a2&a=a1&b=b2&b=b1&c=c1&c=c2"

	gtest.Assert(gurl.BuildQuery(src), expect)
}

func TestParseURL(t *testing.T) {
	src := `http://username:password@hostname:9090/path?arg=value#anchor`
	expect := map[string]string{
		"scheme":   "http",
		"host":     "hostname",
		"port":     "9090",
		"user":     "username",
		"pass":     "password",
		"path":     "/path",
		"query":    "arg=value",
		"fragment": "anchor",
	}

	gtest.Case(t, func() {
		component := 0
		for k, v := range []string{"all", "scheme", "host", "port", "user", "pass", "path", "query", "fragment"} {
			if v == "all" {
				component = -1
			} else {
				component = 1 << (uint(k - 1))
			}

			res, err := gurl.ParseURL(src, component)
			if err != nil {
				t.Errorf("ParseURL failed. component:%v, err:%v", component, err)
				return
			}

			if v == "all" {
				gtest.Assert(res, expect)
			} else {
				gtest.Assert(res[v], expect[v])
			}

		}
	})
}
