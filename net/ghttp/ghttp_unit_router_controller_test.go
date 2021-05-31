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

// 控制器
type Controller struct {
	gmvc.Controller
}

func (c *Controller) Init(r *ghttp.Request) {
	c.Controller.Init(r)
	c.Response.Write("1")
}

func (c *Controller) Shut() {
	c.Response.Write("2")
}

func (c *Controller) Index() {
	c.Response.Write("Controller Index")
}

func (c *Controller) Show() {
	c.Response.Write("Controller Show")
}

func (c *Controller) Info() {
	c.Response.Write("Controller Info")
}

func Test_Router_Controller1(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	s.BindController("/", new(Controller))
	s.BindController("/{.struct}/{.method}", new(Controller))
	s.SetPort(p)
	s.SetDumpRouteMap(false)
	s.Start()
	defer s.Shutdown()

	// 等待启动完成
	time.Sleep(200 * time.Millisecond)
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		gtest.Assert(client.GetContent("/"), "1Controller Index2")
		gtest.Assert(client.GetContent("/init"), "Not Found")
		gtest.Assert(client.GetContent("/shut"), "Not Found")
		gtest.Assert(client.GetContent("/index"), "1Controller Index2")
		gtest.Assert(client.GetContent("/show"), "1Controller Show2")

		gtest.Assert(client.GetContent("/controller"), "Not Found")
		gtest.Assert(client.GetContent("/controller/init"), "Not Found")
		gtest.Assert(client.GetContent("/controller/shut"), "Not Found")
		gtest.Assert(client.GetContent("/controller/index"), "1Controller Index2")
		gtest.Assert(client.GetContent("/controller/show"), "1Controller Show2")

		gtest.Assert(client.GetContent("/none-exist"), "Not Found")
	})
}

func Test_Router_Controller2(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	s.BindController("/controller", new(Controller), "Show, Info")
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
		gtest.Assert(client.GetContent("/controller"), "Not Found")
		gtest.Assert(client.GetContent("/controller/init"), "Not Found")
		gtest.Assert(client.GetContent("/controller/shut"), "Not Found")
		gtest.Assert(client.GetContent("/controller/index"), "Not Found")
		gtest.Assert(client.GetContent("/controller/show"), "1Controller Show2")
		gtest.Assert(client.GetContent("/controller/info"), "1Controller Info2")

		gtest.Assert(client.GetContent("/none-exist"), "Not Found")
	})
}

func Test_Router_ControllerMethod(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	s.BindControllerMethod("/controller-info", new(Controller), "Info")
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
		gtest.Assert(client.GetContent("/controller"), "Not Found")
		gtest.Assert(client.GetContent("/controller/init"), "Not Found")
		gtest.Assert(client.GetContent("/controller/shut"), "Not Found")
		gtest.Assert(client.GetContent("/controller/index"), "Not Found")
		gtest.Assert(client.GetContent("/controller/show"), "Not Found")
		gtest.Assert(client.GetContent("/controller/info"), "Not Found")
		gtest.Assert(client.GetContent("/controller-info"), "1Controller Info2")

		gtest.Assert(client.GetContent("/none-exist"), "Not Found")
	})
}
