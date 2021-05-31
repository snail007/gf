// Copyright 2017 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

// Package gsha1 provides useful API for SHA1 encryption algorithms.
package gsha1

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"

	"github.com/snail007/gf/errors/gerror"
	"github.com/snail007/gf/util/gconv"
)

// Encrypt encrypts any type of variable using SHA1 algorithms.
// It uses gconv package to convert <v> to its bytes type.
func Encrypt(v interface{}) string {
	r := sha1.Sum(gconv.Bytes(v))
	return hex.EncodeToString(r[:])
}

// EncryptFile encrypts file content of <path> using SHA1 algorithms.
func EncryptFile(path string) (encrypt string, err error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer func() {
		err = gerror.Wrap(f.Close(), "file closing error")
	}()
	h := sha1.New()
	_, err = io.Copy(h, f)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
