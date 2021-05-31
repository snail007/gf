// Copyright 2018 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gtype

import (
	"sync/atomic"
)

type Int32 struct {
	value int32
}

// NewInt32 returns a concurrent-safe object for int32 type,
// with given initial value <value>.
func NewInt32(value ...int32) *Int32 {
	if len(value) > 0 {
		return &Int32{
			value: value[0],
		}
	}
	return &Int32{}
}

// Clone clones and returns a new concurrent-safe object for int32 type.
func (v *Int32) Clone() *Int32 {
	return NewInt32(v.Val())
}

// Set atomically stores <value> into t.value and returns the previous value of t.value.
func (v *Int32) Set(value int32) (old int32) {
	return atomic.SwapInt32(&v.value, value)
}

// Val atomically loads t.value.
func (v *Int32) Val() int32 {
	return atomic.LoadInt32(&v.value)
}

// Add atomically adds <delta> to t.value and returns the new value.
func (v *Int32) Add(delta int32) (new int32) {
	return atomic.AddInt32(&v.value, delta)
}

// Cas executes the compare-and-swap operation for value.
func (v *Int32) Cas(old, new int32) bool {
	return atomic.CompareAndSwapInt32(&v.value, old, new)
}
