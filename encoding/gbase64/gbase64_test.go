// Copyright 2017 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.
package gbase64_test

import (
	"testing"

	"github.com/snail007/gf/encoding/gbase64"
	"github.com/snail007/gf/test/gtest"
)

type testpair struct {
	decoded, encoded string
}

var pairs = []testpair{
	// RFC 3548 examples
	{"\x14\xfb\x9c\x03\xd9\x7e", "FPucA9l+"},
	{"\x14\xfb\x9c\x03\xd9", "FPucA9k="},
	{"\x14\xfb\x9c\x03", "FPucAw=="},

	// RFC 4648 examples
	{"", ""},
	{"f", "Zg=="},
	{"fo", "Zm8="},
	{"foo", "Zm9v"},
	{"foob", "Zm9vYg=="},
	{"fooba", "Zm9vYmE="},
	{"foobar", "Zm9vYmFy"},

	// Wikipedia examples
	{"sure.", "c3VyZS4="},
	{"sure", "c3VyZQ=="},
	{"sur", "c3Vy"},
	{"su", "c3U="},
	{"leasure.", "bGVhc3VyZS4="},
	{"easure.", "ZWFzdXJlLg=="},
	{"asure.", "YXN1cmUu"},
	{"sure.", "c3VyZS4="},
}

func TestBase64(t *testing.T) {
	gtest.Case(t, func() {
		for k := range pairs {
			// []byte
			gtest.Assert(gbase64.Encode([]byte(pairs[k].decoded)), []byte(pairs[k].encoded))
			e1, _ := gbase64.Decode([]byte(pairs[k].encoded))
			gtest.Assert(e1, []byte(pairs[k].decoded))

			// string
			gtest.Assert(gbase64.EncodeString([]byte(pairs[k].decoded)), pairs[k].encoded)
			e2, _ := gbase64.DecodeString(pairs[k].encoded)
			gtest.Assert(e2, []byte(pairs[k].decoded))
		}
	})
}
