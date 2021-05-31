// Copyright 2018 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package ghttp_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/net/ghttp"
	"github.com/snail007/gf/test/gtest"
)

func Test_Session_Cookie(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/set", func(r *ghttp.Request) {
		r.Session.Set(r.Get("k"), r.Get("v"))
	})
	s.BindHandler("/get", func(r *ghttp.Request) {
		r.Response.Write(r.Session.Get(r.Get("k")))
	})
	s.BindHandler("/remove", func(r *ghttp.Request) {
		r.Session.Remove(r.Get("k"))
	})
	s.BindHandler("/clear", func(r *ghttp.Request) {
		r.Session.Clear()
	})
	s.SetPort(p)
	s.SetDumpRouteMap(false)
	s.Start()
	defer s.Shutdown()

	// 等待启动完成
	time.Sleep(200 * time.Millisecond)
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetBrowserMode(true)
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		r1, e1 := client.Get("/set?k=key1&v=100")
		if r1 != nil {
			defer r1.Close()
		}
		gtest.Assert(e1, nil)
		gtest.Assert(r1.ReadAllString(), "")

		gtest.Assert(client.GetContent("/set?k=key2&v=200"), "")

		gtest.Assert(client.GetContent("/get?k=key1"), "100")
		gtest.Assert(client.GetContent("/get?k=key2"), "200")
		gtest.Assert(client.GetContent("/get?k=key3"), "")
		gtest.Assert(client.GetContent("/remove?k=key1"), "")
		gtest.Assert(client.GetContent("/remove?k=key3"), "")
		gtest.Assert(client.GetContent("/remove?k=key4"), "")
		gtest.Assert(client.GetContent("/get?k=key1"), "")
		gtest.Assert(client.GetContent("/get?k=key2"), "200")
		gtest.Assert(client.GetContent("/clear"), "")
		gtest.Assert(client.GetContent("/get?k=key2"), "")
	})
}

func Test_Session_Header(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	s.BindHandler("/set", func(r *ghttp.Request) {
		r.Session.Set(r.Get("k"), r.Get("v"))
	})
	s.BindHandler("/get", func(r *ghttp.Request) {
		r.Response.Write(r.Session.Get(r.Get("k")))
	})
	s.BindHandler("/remove", func(r *ghttp.Request) {
		r.Session.Remove(r.Get("k"))
	})
	s.BindHandler("/clear", func(r *ghttp.Request) {
		r.Session.Clear()
	})
	s.SetPort(p)
	s.SetDumpRouteMap(false)
	s.Start()
	defer s.Shutdown()

	// 等待启动完成
	time.Sleep(200 * time.Millisecond)
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		response, e1 := client.Get("/set?k=key1&v=100")
		if response != nil {
			defer response.Close()
		}
		sessionId := response.GetCookie(s.GetSessionIdName())
		gtest.Assert(e1, nil)
		gtest.AssertNE(sessionId, nil)
		gtest.Assert(response.ReadAllString(), "")

		client.SetHeader(s.GetSessionIdName(), sessionId)

		gtest.Assert(client.GetContent("/set?k=key2&v=200"), "")

		gtest.Assert(client.GetContent("/get?k=key1"), "100")
		gtest.Assert(client.GetContent("/get?k=key2"), "200")
		gtest.Assert(client.GetContent("/get?k=key3"), "")
		gtest.Assert(client.GetContent("/remove?k=key1"), "")
		gtest.Assert(client.GetContent("/remove?k=key3"), "")
		gtest.Assert(client.GetContent("/remove?k=key4"), "")
		gtest.Assert(client.GetContent("/get?k=key1"), "")
		gtest.Assert(client.GetContent("/get?k=key2"), "200")
		gtest.Assert(client.GetContent("/clear"), "")
		gtest.Assert(client.GetContent("/get?k=key2"), "")
	})
}

func Test_Session_StorageFile(t *testing.T) {
	sessionId := ""
	gtest.Case(t, func() {
		p := ports.PopRand()
		s := g.Server(p)
		s.BindHandler("/set", func(r *ghttp.Request) {
			r.Session.Set(r.Get("k"), r.Get("v"))
			r.Response.Write(r.Get("k"), "=", r.Get("v"))
		})
		s.SetPort(p)
		s.SetDumpRouteMap(false)
		s.Start()
		defer s.Shutdown()
		time.Sleep(200 * time.Millisecond)

		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		response, e1 := client.Get("/set?k=key&v=100")
		if response != nil {
			defer response.Close()
		}
		sessionId = response.GetCookie(s.GetSessionIdName())
		gtest.Assert(e1, nil)
		gtest.AssertNE(sessionId, nil)
		gtest.Assert(response.ReadAllString(), "key=100")
	})
	time.Sleep(time.Second)
	gtest.Case(t, func() {
		p := ports.PopRand()
		s := g.Server(p)
		s.BindHandler("/get", func(r *ghttp.Request) {
			r.Response.Write(r.Session.Get(r.Get("k")))
		})
		s.SetPort(p)
		s.SetDumpRouteMap(false)
		s.Start()
		defer s.Shutdown()
		time.Sleep(200 * time.Millisecond)

		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))
		client.SetHeader(s.GetSessionIdName(), sessionId)
		gtest.Assert(client.GetContent("/get?k=key"), "100")
		gtest.Assert(client.GetContent("/get?k=key1"), "")
	})
}
