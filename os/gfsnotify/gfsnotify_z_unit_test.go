// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

// go test *.go -bench=".*" -benchmem

package gfsnotify_test

import (
	"testing"
	"time"

	"github.com/snail007/gf/container/gtype"
	"github.com/snail007/gf/os/gfile"
	"github.com/snail007/gf/os/gfsnotify"
	"github.com/snail007/gf/os/gtime"
	"github.com/snail007/gf/test/gtest"
	"github.com/snail007/gf/util/gconv"
)

func TestWatcher_AddRemove(t *testing.T) {
	gtest.Case(t, func() {
		path1 := gfile.TempDir() + gfile.Separator + gconv.String(gtime.Nanosecond())
		path2 := gfile.TempDir() + gfile.Separator + gconv.String(gtime.Nanosecond()) + "2"
		gfile.PutContents(path1, "1")
		defer func() {
			gfile.Remove(path1)
			gfile.Remove(path2)
		}()
		v := gtype.NewInt(1)
		callback, err := gfsnotify.Add(path1, func(event *gfsnotify.Event) {
			if event.IsWrite() {
				v.Set(2)
				return
			}
			if event.IsRename() {
				v.Set(3)
				gfsnotify.Exit()
				return
			}
		})
		gtest.Assert(err, nil)
		gtest.AssertNE(callback, nil)

		gfile.PutContents(path1, "2")
		time.Sleep(100 * time.Millisecond)
		gtest.Assert(v.Val(), 2)

		gfile.Rename(path1, path2)
		time.Sleep(100 * time.Millisecond)
		gtest.Assert(v.Val(), 3)
	})

	gtest.Case(t, func() {
		path1 := gfile.TempDir() + gfile.Separator + gconv.String(gtime.Nanosecond())
		gfile.PutContents(path1, "1")
		defer func() {
			gfile.Remove(path1)
		}()
		v := gtype.NewInt(1)
		callback, err := gfsnotify.Add(path1, func(event *gfsnotify.Event) {
			if event.IsWrite() {
				v.Set(2)
				return
			}
			if event.IsRemove() {
				v.Set(4)
				return
			}
		})
		gtest.Assert(err, nil)
		gtest.AssertNE(callback, nil)

		gfile.PutContents(path1, "2")
		time.Sleep(100 * time.Millisecond)
		gtest.Assert(v.Val(), 2)

		gfile.Remove(path1)
		time.Sleep(100 * time.Millisecond)
		gtest.Assert(v.Val(), 4)

		gfile.PutContents(path1, "1")
		time.Sleep(100 * time.Millisecond)
		gtest.Assert(v.Val(), 4)
	})
}

func TestWatcher_Callback(t *testing.T) {
	gtest.Case(t, func() {
		path1 := gfile.TempDir() + gfile.Separator + gconv.String(gtime.Nanosecond())
		gfile.PutContents(path1, "1")
		defer func() {
			gfile.Remove(path1)
		}()
		v := gtype.NewInt(1)
		callback, err := gfsnotify.Add(path1, func(event *gfsnotify.Event) {
			if event.IsWrite() {
				v.Set(2)
				return
			}
		})
		gtest.Assert(err, nil)
		gtest.AssertNE(callback, nil)

		gfile.PutContents(path1, "2")
		time.Sleep(100 * time.Millisecond)
		gtest.Assert(v.Val(), 2)

		v.Set(3)
		gfsnotify.RemoveCallback(callback.Id)
		gfile.PutContents(path1, "3")
		time.Sleep(100 * time.Millisecond)
		gtest.Assert(v.Val(), 3)
	})
	// multiple callbacks
	gtest.Case(t, func() {
		path1 := gfile.TempDir() + gfile.Separator + gconv.String(gtime.Nanosecond())
		gfile.PutContents(path1, "1")
		defer func() {
			gfile.Remove(path1)
		}()
		v1 := gtype.NewInt(1)
		v2 := gtype.NewInt(1)
		callback1, err1 := gfsnotify.Add(path1, func(event *gfsnotify.Event) {
			if event.IsWrite() {
				v1.Set(2)
				return
			}
		})
		callback2, err2 := gfsnotify.Add(path1, func(event *gfsnotify.Event) {
			if event.IsWrite() {
				v2.Set(2)
				return
			}
		})
		gtest.Assert(err1, nil)
		gtest.Assert(err2, nil)
		gtest.AssertNE(callback1, nil)
		gtest.AssertNE(callback2, nil)

		gfile.PutContents(path1, "2")
		time.Sleep(100 * time.Millisecond)
		gtest.Assert(v1.Val(), 2)
		gtest.Assert(v2.Val(), 2)

		v1.Set(3)
		v2.Set(3)
		gfsnotify.RemoveCallback(callback1.Id)
		gfile.PutContents(path1, "3")
		time.Sleep(100 * time.Millisecond)
		gtest.Assert(v1.Val(), 3)
		gtest.Assert(v2.Val(), 2)
	})
}
