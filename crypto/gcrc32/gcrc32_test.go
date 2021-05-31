// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

// go test *.go -bench=".*"

package gcrc32_test

import (
	"testing"

	"github.com/snail007/gf/crypto/gcrc32"
	"github.com/snail007/gf/crypto/gmd5"
	"github.com/snail007/gf/test/gtest"
)

func TestEncrypt(t *testing.T) {
	gtest.Case(t, func() {
		s := "pibigstar"
		result := 693191136
		encrypt1 := gcrc32.Encrypt(s)
		encrypt2 := gcrc32.Encrypt([]byte(s))
		gtest.AssertEQ(int(encrypt1), result)
		gtest.AssertEQ(int(encrypt2), result)

		strmd5, _ := gmd5.Encrypt(s)
		test1 := gcrc32.Encrypt(strmd5)
		test2 := gcrc32.Encrypt([]byte(strmd5))
		gtest.AssertEQ(test2, test1)
	})
}
