// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gpool_test

import (
	"errors"
	"testing"
	"time"

	"github.com/snail007/gf/frame/g"

	"github.com/snail007/gf/container/gpool"
	"github.com/snail007/gf/test/gtest"
)

var nf gpool.NewFunc = func() (i interface{}, e error) {
	return "hello", nil
}

var assertIndex int = 0
var ef gpool.ExpireFunc = func(i interface{}) {
	assertIndex++
	gtest.Assert(i, assertIndex)
}

func Test_Gpool(t *testing.T) {
	gtest.Case(t, func() {
		//
		//expire = 0
		p1 := gpool.New(0, nf)
		p1.Put(1)
		p1.Put(2)
		time.Sleep(1 * time.Second)
		//test won't be timeout
		v1, err1 := p1.Get()
		gtest.Assert(err1, nil)
		gtest.AssertIN(v1, g.Slice{1, 2})
		//test clear
		p1.Clear()
		gtest.Assert(p1.Size(), 0)
		//test newFunc
		v1, err1 = p1.Get()
		gtest.Assert(err1, nil)
		gtest.Assert(v1, "hello")
		//put data again
		p1.Put(3)
		p1.Put(4)
		v1, err1 = p1.Get()
		gtest.Assert(err1, nil)
		gtest.AssertIN(v1, g.Slice{3, 4})
		//test close
		p1.Close()
		v1, err1 = p1.Get()
		gtest.Assert(err1, nil)
		gtest.Assert(v1, "hello")
	})

	gtest.Case(t, func() {
		//
		//expire > 0
		p2 := gpool.New(2000, nil, ef)
		for index := 0; index < 10; index++ {
			p2.Put(index)
		}
		gtest.Assert(p2.Size(), 10)
		v2, err2 := p2.Get()
		gtest.Assert(err2, nil)
		gtest.Assert(v2, 0)
		//test timeout expireFunc
		time.Sleep(3 * time.Second)
		v2, err2 = p2.Get()
		gtest.Assert(err2, errors.New("pool is empty"))
		gtest.Assert(v2, nil)
		//test close expireFunc
		for index := 0; index < 10; index++ {
			p2.Put(index)
		}
		gtest.Assert(p2.Size(), 10)
		v2, err2 = p2.Get()
		gtest.Assert(err2, nil)
		gtest.Assert(v2, 0)
		assertIndex = 0
		p2.Close()
		time.Sleep(3 * time.Second)
	})

	gtest.Case(t, func() {
		//
		//expire < 0
		p3 := gpool.New(-1, nil)
		v3, err3 := p3.Get()
		gtest.Assert(err3, errors.New("pool is empty"))
		gtest.Assert(v3, nil)
	})
}
