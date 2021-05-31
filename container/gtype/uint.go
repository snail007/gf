// Copyright 2018 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gtype

import (
	"sync/atomic"
)

type Uint struct {
	value uint64
}

// NewUint returns a concurrent-safe object for uint type,
// with given initial value <value>.
func NewUint(value ...uint) *Uint {
	if len(value) > 0 {
		return &Uint{
			value: uint64(value[0]),
		}
	}
	return &Uint{}
}

// Clone clones and returns a new concurrent-safe object for uint type.
func (v *Uint) Clone() *Uint {
	return NewUint(v.Val())
}

// Set atomically stores <value> into t.value and returns the previous value of t.value.
func (v *Uint) Set(value uint) (old uint) {
	return uint(atomic.SwapUint64(&v.value, uint64(value)))
}

// Val atomically loads t.value.
func (v *Uint) Val() uint {
	return uint(atomic.LoadUint64(&v.value))
}

// Add atomically adds <delta> to t.value and returns the new value.
func (v *Uint) Add(delta uint) (new uint) {
	return uint(atomic.AddUint64(&v.value, uint64(delta)))
}

// Cas executes the compare-and-swap operation for value.
func (v *Uint) Cas(old, new uint) bool {
	return atomic.CompareAndSwapUint64(&v.value, uint64(old), uint64(new))
}
