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

type ControllerRest struct {
	gmvc.Controller
}

func (c *ControllerRest) Init(r *ghttp.Request) {
	c.Controller.Init(r)
	c.Response.Write("1")
}

func (c *ControllerRest) Shut() {
	c.Response.Write("2")
}

func (c *ControllerRest) Get() {
	c.Response.Write("Controller Get")
}

func (c *ControllerRest) Put() {
	c.Response.Write("Controller Put")
}

func (c *ControllerRest) Post() {
	c.Response.Write("Controller Post")
}

func (c *ControllerRest) Delete() {
	c.Response.Write("Controller Delete")
}

func (c *ControllerRest) Patch() {
	c.Response.Write("Controller Patch")
}

func (c *ControllerRest) Options() {
	c.Response.Write("Controller Options")
}

func (c *ControllerRest) Head() {
	c.Response.Header().Set("head-ok", "1")
}

// 控制器注册测试
func Test_Router_ControllerRest(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	s.BindControllerRest("/", new(ControllerRest))
	s.BindControllerRest("/{.struct}/{.method}", new(ControllerRest))
	s.SetPort(p)
	s.SetDumpRouteMap(false)
	s.Start()
	defer s.Shutdown()

	// 等待启动完成
	time.Sleep(200 * time.Millisecond)
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

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

		gtest.Assert(client.GetContent("/controller-rest/get"), "1Controller Get2")
		gtest.Assert(client.PutContent("/controller-rest/put"), "1Controller Put2")
		gtest.Assert(client.PostContent("/controller-rest/post"), "1Controller Post2")
		gtest.Assert(client.DeleteContent("/controller-rest/delete"), "1Controller Delete2")
		gtest.Assert(client.PatchContent("/controller-rest/patch"), "1Controller Patch2")
		gtest.Assert(client.OptionsContent("/controller-rest/options"), "1Controller Options2")
		resp2, err := client.Head("/controller-rest/head")
		if err == nil {
			defer resp2.Close()
		}
		gtest.Assert(err, nil)
		gtest.Assert(resp2.Header.Get("head-ok"), "1")

		gtest.Assert(client.GetContent("/none-exist"), "Not Found")
	})
}
