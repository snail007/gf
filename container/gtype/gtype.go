// Copyright 2018 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

// Package gtype provides kinds of high performance and concurrent-safe basic variable types.
package gtype

type Type = Interface

// See NewInterface.
func New(value ...interface{}) *Type {
	return NewInterface(value...)
}
