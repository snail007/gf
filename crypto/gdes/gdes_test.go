// Copyright 2018 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gdes_test

import (
	"encoding/hex"
	"testing"

	"github.com/snail007/gf/crypto/gdes"
	"github.com/snail007/gf/test/gtest"
)

var (
	errKey     = []byte("1111111111111234123456789")
	errIv      = []byte("123456789")
	errPadding = 5
)

func TestDesECB(t *testing.T) {
	gtest.Case(t, func() {
		key := []byte("11111111")
		text := []byte("12345678")
		padding := gdes.NOPADDING
		result := "858b176da8b12503"
		// encrypt test
		cipherText, err := gdes.EncryptECB(text, key, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptECB(cipherText, key, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(string(clearText), "12345678")

		// encrypt err test. when throw exception,the err is not equal nil and the string is nil
		errEncrypt, err := gdes.EncryptECB(text, key, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		errEncrypt, err = gdes.EncryptECB(text, errKey, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// err decrypt test.
		errDecrypt, err := gdes.DecryptECB(cipherText, errKey, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
		errDecrypt, err = gdes.DecryptECB(cipherText, key, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
	})

	gtest.Case(t, func() {
		key := []byte("11111111")
		text := []byte("12345678")
		padding := gdes.PKCS5PADDING
		errPadding := 5
		result := "858b176da8b12503ad6a88b4fa37833d"
		cipherText, err := gdes.EncryptECB(text, key, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptECB(cipherText, key, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(string(clearText), "12345678")

		// err test
		errEncrypt, err := gdes.EncryptECB(text, key, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		errDecrypt, err := gdes.DecryptECB(cipherText, errKey, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
	})
}

func Test3DesECB(t *testing.T) {
	gtest.Case(t, func() {
		key := []byte("1111111111111234")
		text := []byte("1234567812345678")
		padding := gdes.NOPADDING
		result := "a23ee24b98c26263a23ee24b98c26263"
		// encrypt test
		cipherText, err := gdes.EncryptECBTriple(text, key, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptECBTriple(cipherText, key, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(string(clearText), "1234567812345678")
		// err test
		errEncrypt, err := gdes.EncryptECB(text, key, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
	})

	gtest.Case(t, func() {
		key := []byte("111111111111123412345678")
		text := []byte("123456789")
		padding := gdes.PKCS5PADDING
		errPadding := 5
		result := "37989b1effc07a6d00ff89a7d052e79f"
		// encrypt test
		cipherText, err := gdes.EncryptECBTriple(text, key, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptECBTriple(cipherText, key, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(string(clearText), "123456789")
		// err test, when key is err, but text and padding is right
		errEncrypt, err := gdes.EncryptECBTriple(text, errKey, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// when padding is err,but key and text is right
		errEncrypt, err = gdes.EncryptECBTriple(text, key, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// decrypt err test,when key is err
		errEncrypt, err = gdes.DecryptECBTriple(text, errKey, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
	})
}

func TestDesCBC(t *testing.T) {
	gtest.Case(t, func() {
		key := []byte("11111111")
		text := []byte("1234567812345678")
		padding := gdes.NOPADDING
		iv := []byte("12345678")
		result := "40826a5800608c87585ca7c9efabee47"
		// encrypt test
		cipherText, err := gdes.EncryptCBC(text, key, iv, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptCBC(cipherText, key, iv, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(string(clearText), "1234567812345678")
		// encrypt err test.
		errEncrypt, err := gdes.EncryptCBC(text, errKey, iv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// the iv is err
		errEncrypt, err = gdes.EncryptCBC(text, key, errIv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// the padding is err
		errEncrypt, err = gdes.EncryptCBC(text, key, iv, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// decrypt err test. the key is err
		errDecrypt, err := gdes.DecryptCBC(cipherText, errKey, iv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
		// the iv is err
		errDecrypt, err = gdes.DecryptCBC(cipherText, key, errIv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
		// the padding is err
		errDecrypt, err = gdes.DecryptCBC(cipherText, key, iv, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
	})

	gtest.Case(t, func() {
		key := []byte("11111111")
		text := []byte("12345678")
		padding := gdes.PKCS5PADDING
		iv := []byte("12345678")
		result := "40826a5800608c87100a25d86ac7c52c"
		// encrypt test
		cipherText, err := gdes.EncryptCBC(text, key, iv, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptCBC(cipherText, key, iv, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(string(clearText), "12345678")
		// err test
		errEncrypt, err := gdes.EncryptCBC(text, key, errIv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
	})
}

func Test3DesCBC(t *testing.T) {
	gtest.Case(t, func() {
		key := []byte("1111111112345678")
		text := []byte("1234567812345678")
		padding := gdes.NOPADDING
		iv := []byte("12345678")
		result := "bfde1394e265d5f738d5cab170c77c88"
		// encrypt test
		cipherText, err := gdes.EncryptCBCTriple(text, key, iv, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptCBCTriple(cipherText, key, iv, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(string(clearText), "1234567812345678")
		// encrypt err test
		errEncrypt, err := gdes.EncryptCBCTriple(text, errKey, iv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// the iv is err
		errEncrypt, err = gdes.EncryptCBCTriple(text, key, errIv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// the padding is err
		errEncrypt, err = gdes.EncryptCBCTriple(text, key, iv, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errEncrypt, nil)
		// decrypt err test
		errDecrypt, err := gdes.DecryptCBCTriple(cipherText, errKey, iv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
		// the iv is err
		errDecrypt, err = gdes.DecryptCBCTriple(cipherText, key, errIv, padding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
		// the padding is err
		errDecrypt, err = gdes.DecryptCBCTriple(cipherText, key, iv, errPadding)
		gtest.AssertNE(err, nil)
		gtest.AssertEQ(errDecrypt, nil)
	})
	gtest.Case(t, func() {
		key := []byte("111111111234567812345678")
		text := []byte("12345678")
		padding := gdes.PKCS5PADDING
		iv := []byte("12345678")
		result := "40826a5800608c87100a25d86ac7c52c"
		// encrypt test
		cipherText, err := gdes.EncryptCBCTriple(text, key, iv, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(hex.EncodeToString(cipherText), result)
		// decrypt test
		clearText, err := gdes.DecryptCBCTriple(cipherText, key, iv, padding)
		gtest.AssertEQ(err, nil)
		gtest.AssertEQ(string(clearText), "12345678")
	})

}
