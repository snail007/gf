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
	"github.com/snail007/gf/frame/gmvc"
	"github.com/snail007/gf/net/ghttp"
	"github.com/snail007/gf/test/gtest"
)

type DomainControllerRest struct {
	gmvc.Controller
}

func (c *DomainControllerRest) Init(r *ghttp.Request) {
	c.Controller.Init(r)
	c.Response.Write("1")
}

func (c *DomainControllerRest) Shut() {
	c.Response.Write("2")
}

func (c *DomainControllerRest) Get() {
	c.Response.Write("Controller Get")
}

func (c *DomainControllerRest) Put() {
	c.Response.Write("Controller Put")
}

func (c *DomainControllerRest) Post() {
	c.Response.Write("Controller Post")
}

func (c *DomainControllerRest) Delete() {
	c.Response.Write("Controller Delete")
}

func (c *DomainControllerRest) Patch() {
	c.Response.Write("Controller Patch")
}

func (c *DomainControllerRest) Options() {
	c.Response.Write("Controller Options")
}

func (c *DomainControllerRest) Head() {
	c.Response.Header().Set("head-ok", "1")
}

// 控制器注册测试
func Test_Router_DomainControllerRest(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	d := s.Domain("localhost, local")
	d.BindControllerRest("/", new(DomainControllerRest))
	s.SetPort(p)
	s.SetDumpRouteMap(false)
	s.Start()
	defer s.Shutdown()

	// 等待启动完成
	time.Sleep(200 * time.Millisecond)
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		gtest.Assert(client.GetContent("/"), "Not Found")
		gtest.Assert(client.PutContent("/"), "Not Found")
		gtest.Assert(client.PostContent("/"), "Not Found")
		gtest.Assert(client.DeleteContent("/"), "Not Found")
		gtest.Assert(client.PatchContent("/"), "Not Found")
		gtest.Assert(client.OptionsContent("/"), "Not Found")
		resp1, err := client.Head("/")
		if err == nil {
			defer resp1.Close()
		}
		gtest.Assert(err, nil)
		gtest.Assert(resp1.Header.Get("head-ok"), "")
		gtest.Assert(client.GetContent("/none-exist"), "Not Found")
	})
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://localhost:%d", p))

		gtest.Assert(client.GetContent("/"), "1Controller Get2")
		gtest.Assert(client.PutContent("/"), "1Controller Put2")
		gtest.Assert(client.PostContent("/"), "1Controller Post2")
		gtest.Assert(client.DeleteContent("/"), "1Controller Delete2")
		gtest.Assert(client.PatchContent("/"), "1Controller Patch2")
		gtest.Assert(client.OptionsContent("/"), "1Controller Options2")
		resp1, err := client.Head("/")
		if err == nil {
			defer resp1.Close()
		}
		gtest.Assert(err, nil)
		gtest.Assert(resp1.Header.Get("head-ok"), "1")
		gtest.Assert(client.GetContent("/none-exist"), "Not Found")
	})
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://local:%d", p))

		gtest.Assert(client.GetContent("/"), "1Controller Get2")
		gtest.Assert(client.PutContent("/"), "1Controller Put2")
		gtest.Assert(client.PostContent("/"), "1Controller Post2")
		gtest.Assert(client.DeleteContent("/"), "1Controller Delete2")
		gtest.Assert(client.PatchContent("/"), "1Controller Patch2")
		gtest.Assert(client.OptionsContent("/"), "1Controller Options2")
		resp1, err := client.Head("/")
		if err == nil {
			defer resp1.Close()
		}
		gtest.Assert(err, nil)
		gtest.Assert(resp1.Header.Get("head-ok"), "1")
		gtest.Assert(client.GetContent("/none-exist"), "Not Found")
	})
}
