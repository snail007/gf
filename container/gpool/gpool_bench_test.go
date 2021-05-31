// Copyright 2018 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

// go test *.go -bench=".*"

package gpool_test

import (
	"sync"
	"testing"

	"github.com/snail007/gf/container/gpool"
)

var pool = gpool.New(99999999, nil)
var syncp = sync.Pool{}

func BenchmarkGPoolPut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.Put(i)
	}
}

func BenchmarkGPoolGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pool.Get()
	}
}

func BenchmarkSyncPoolPut(b *testing.B) {
	for i := 0; i < b.N; i++ {
		syncp.Put(i)
	}
}

func BenchmarkGpoolGet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		syncp.Get()
	}
}
