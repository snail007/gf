// Copyright 2017-2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with gm file,
// You can obtain one at https://github.com/snail007/gf.

package gtree_test

import (
	"fmt"
	"testing"

	"github.com/snail007/gf/container/gtree"
	"github.com/snail007/gf/container/gvar"
	"github.com/snail007/gf/test/gtest"
	"github.com/snail007/gf/util/gutil"
)

func Test_BTree_Basic(t *testing.T) {
	gtest.Case(t, func() {
		m := gtree.NewBTree(3, gutil.ComparatorString)
		m.Set("key1", "val1")

		gtest.Assert(m.Height(), 1)

		gtest.Assert(m.Keys(), []interface{}{"key1"})

		gtest.Assert(m.Get("key1"), "val1")
		gtest.Assert(m.Size(), 1)
		gtest.Assert(m.IsEmpty(), false)

		gtest.Assert(m.GetOrSet("key2", "val2"), "val2")
		gtest.Assert(m.SetIfNotExist("key2", "val2"), false)

		gtest.Assert(m.SetIfNotExist("key3", "val3"), true)

		gtest.Assert(m.Remove("key2"), "val2")
		gtest.Assert(m.Contains("key2"), false)

		gtest.AssertIN("key3", m.Keys())
		gtest.AssertIN("key1", m.Keys())
		gtest.AssertIN("val3", m.Values())
		gtest.AssertIN("val1", m.Values())

		m.Clear()
		gtest.Assert(m.Size(), 0)
		gtest.Assert(m.IsEmpty(), true)

		m2 := gtree.NewBTreeFrom(3, gutil.ComparatorString, map[interface{}]interface{}{1: 1, "key1": "val1"})
		gtest.Assert(m2.Map(), map[interface{}]interface{}{1: 1, "key1": "val1"})
	})
}

func Test_BTree_Set_Fun(t *testing.T) {
	//GetOrSetFunc lock or unlock
	gtest.Case(t, func() {
		m := gtree.NewBTree(3, gutil.ComparatorString)
		gtest.Assert(m.GetOrSetFunc("fun", getValue), 3)
		gtest.Assert(m.GetOrSetFunc("fun", getValue), 3)
		gtest.Assert(m.GetOrSetFuncLock("funlock", getValue), 3)
		gtest.Assert(m.GetOrSetFuncLock("funlock", getValue), 3)
		gtest.Assert(m.Get("funlock"), 3)
		gtest.Assert(m.Get("fun"), 3)
	})
	//SetIfNotExistFunc lock or unlock
	gtest.Case(t, func() {
		m := gtree.NewBTree(3, gutil.ComparatorString)
		gtest.Assert(m.SetIfNotExistFunc("fun", getValue), true)
		gtest.Assert(m.SetIfNotExistFunc("fun", getValue), false)
		gtest.Assert(m.SetIfNotExistFuncLock("funlock", getValue), true)
		gtest.Assert(m.SetIfNotExistFuncLock("funlock", getValue), false)
		gtest.Assert(m.Get("funlock"), 3)
		gtest.Assert(m.Get("fun"), 3)
	})

}

func Test_BTree_Get_Set_Var(t *testing.T) {
	gtest.Case(t, func() {
		m := gtree.NewBTree(3, gutil.ComparatorString)
		gtest.AssertEQ(m.SetIfNotExist("key1", "val1"), true)
		gtest.AssertEQ(m.SetIfNotExist("key1", "val1"), false)
		gtest.AssertEQ(m.GetVarOrSet("key1", "val1"), gvar.New("val1", true))
		gtest.AssertEQ(m.GetVar("key1"), gvar.New("val1", true))
	})

	gtest.Case(t, func() {
		m := gtree.NewBTree(3, gutil.ComparatorString)
		gtest.AssertEQ(m.GetVarOrSetFunc("fun", getValue), gvar.New(3, true))
		gtest.AssertEQ(m.GetVarOrSetFunc("fun", getValue), gvar.New(3, true))
		gtest.AssertEQ(m.GetVarOrSetFuncLock("funlock", getValue), gvar.New(3, true))
		gtest.AssertEQ(m.GetVarOrSetFuncLock("funlock", getValue), gvar.New(3, true))
	})
}

func Test_BTree_Batch(t *testing.T) {
	m := gtree.NewBTree(3, gutil.ComparatorString)
	m.Sets(map[interface{}]interface{}{1: 1, "key1": "val1", "key2": "val2", "key3": "val3"})
	gtest.Assert(m.Map(), map[interface{}]interface{}{1: 1, "key1": "val1", "key2": "val2", "key3": "val3"})
	m.Removes([]interface{}{"key1", 1})
	gtest.Assert(m.Map(), map[interface{}]interface{}{"key2": "val2", "key3": "val3"})
}

func Test_BTree_Iterator(t *testing.T) {
	keys := []string{"1", "key1", "key2", "key3", "key4"}
	keyLen := len(keys)
	index := 0

	expect := map[interface{}]interface{}{"key4": "val4", 1: 1, "key1": "val1", "key2": "val2", "key3": "val3"}

	m := gtree.NewBTreeFrom(3, gutil.ComparatorString, expect)
	m.Iterator(func(k interface{}, v interface{}) bool {
		gtest.Assert(k, keys[index])
		index++
		gtest.Assert(expect[k], v)
		return true
	})

	m.IteratorDesc(func(k interface{}, v interface{}) bool {
		index--
		gtest.Assert(k, keys[index])
		gtest.Assert(expect[k], v)
		return true
	})

	m.Print()
	// 断言返回值对遍历控制
	gtest.Case(t, func() {
		i := 0
		j := 0
		m.Iterator(func(k interface{}, v interface{}) bool {
			i++
			return true
		})
		m.Iterator(func(k interface{}, v interface{}) bool {
			j++
			return false
		})
		gtest.Assert(i, keyLen)
		gtest.Assert(j, 1)
	})

	gtest.Case(t, func() {
		i := 0
		j := 0
		m.IteratorDesc(func(k interface{}, v interface{}) bool {
			i++
			return true
		})
		m.IteratorDesc(func(k interface{}, v interface{}) bool {
			j++
			return false
		})
		gtest.Assert(i, keyLen)
		gtest.Assert(j, 1)
	})
}

func Test_BTree_IteratorFrom(t *testing.T) {
	m := make(map[interface{}]interface{})
	for i := 1; i <= 10; i++ {
		m[i] = i * 10
	}
	tree := gtree.NewBTreeFrom(3, gutil.ComparatorInt, m)

	gtest.Case(t, func() {
		n := 5
		tree.IteratorFrom(5, true, func(key, value interface{}) bool {
			gtest.Assert(n, key)
			gtest.Assert(n*10, value)
			n++
			return true
		})

		i := 5
		tree.IteratorAscFrom(5, true, func(key, value interface{}) bool {
			gtest.Assert(i, key)
			gtest.Assert(i*10, value)
			i++
			return true
		})

		j := 5
		tree.IteratorDescFrom(5, true, func(key, value interface{}) bool {
			gtest.Assert(j, key)
			gtest.Assert(j*10, value)
			j--
			return true
		})
	})
}

func Test_BTree_Clone(t *testing.T) {
	//clone 方法是深克隆
	m := gtree.NewBTreeFrom(3, gutil.ComparatorString, map[interface{}]interface{}{1: 1, "key1": "val1"})
	m_clone := m.Clone()
	m.Remove(1)
	//修改原 map,clone 后的 map 不影响
	gtest.AssertIN(1, m_clone.Keys())

	m_clone.Remove("key1")
	//修改clone map,原 map 不影响
	gtest.AssertIN("key1", m.Keys())
}

func Test_BTree_LRNode(t *testing.T) {
	expect := map[interface{}]interface{}{"key4": "val4", "key1": "val1", "key2": "val2", "key3": "val3"}
	//safe
	gtest.Case(t, func() {
		m := gtree.NewBTreeFrom(3, gutil.ComparatorString, expect)
		gtest.Assert(m.Left().Key, "key1")
		gtest.Assert(m.Right().Key, "key4")
	})
	//unsafe
	gtest.Case(t, func() {
		m := gtree.NewBTreeFrom(3, gutil.ComparatorString, expect, true)
		gtest.Assert(m.Left().Key, "key1")
		gtest.Assert(m.Right().Key, "key4")
	})
}

func Test_BTree_Remove(t *testing.T) {
	m := gtree.NewBTree(3, gutil.ComparatorInt)
	for i := 1; i <= 100; i++ {
		m.Set(i, fmt.Sprintf("val%d", i))
	}
	expect := m.Map()
	gtest.Case(t, func() {
		for k, v := range expect {
			m1 := m.Clone()
			gtest.Assert(m1.Remove(k), v)
			gtest.Assert(m1.Remove(k), nil)
		}
	})
}
