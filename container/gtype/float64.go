// Copyright 2018 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gtype

import (
	"math"
	"sync/atomic"
	"unsafe"
)

type Float64 struct {
	value uint64
}

// NewFloat64 returns a concurrent-safe object for float64 type,
// with given initial value <value>.
func NewFloat64(value ...float64) *Float64 {
	if len(value) > 0 {
		return &Float64{
			value: math.Float64bits(value[0]),
		}
	}
	return &Float64{}
}

// Clone clones and returns a new concurrent-safe object for float64 type.
func (v *Float64) Clone() *Float64 {
	return NewFloat64(v.Val())
}

// Set atomically stores <value> into t.value and returns the previous value of t.value.
func (v *Float64) Set(value float64) (old float64) {
	return math.Float64frombits(atomic.SwapUint64(&v.value, math.Float64bits(value)))
}

// Val atomically loads t.value.
func (v *Float64) Val() float64 {
	return math.Float64frombits(atomic.LoadUint64(&v.value))
}

// Add atomically adds <delta> to t.value and returns the new value.
func (v *Float64) Add(delta float64) (new float64) {
	for {
		old := math.Float64frombits(v.value)
		new = old + delta
		if atomic.CompareAndSwapUint64(
			(*uint64)(unsafe.Pointer(&v.value)),
			math.Float64bits(old),
			math.Float64bits(new),
		) {
			break
		}
	}
	return
}

// Cas executes the compare-and-swap operation for value.
func (v *Float64) Cas(old, new float64) bool {
	return atomic.CompareAndSwapUint64(&v.value, uint64(old), uint64(new))
}
