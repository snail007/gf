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

type ObjectRest struct{}

func (o *ObjectRest) Init(r *ghttp.Request) {
	r.Response.Write("1")
}

func (o *ObjectRest) Shut(r *ghttp.Request) {
	r.Response.Write("2")
}

func (o *ObjectRest) Get(r *ghttp.Request) {
	r.Response.Write("Object Get")
}

func (o *ObjectRest) Put(r *ghttp.Request) {
	r.Response.Write("Object Put")
}

func (o *ObjectRest) Post(r *ghttp.Request) {
	r.Response.Write("Object Post")
}

func (o *ObjectRest) Delete(r *ghttp.Request) {
	r.Response.Write("Object Delete")
}

func (o *ObjectRest) Patch(r *ghttp.Request) {
	r.Response.Write("Object Patch")
}

func (o *ObjectRest) Options(r *ghttp.Request) {
	r.Response.Write("Object Options")
}

func (o *ObjectRest) Head(r *ghttp.Request) {
	r.Response.Header().Set("head-ok", "1")
}

func Test_Router_ObjectRest(t *testing.T) {
	p := ports.PopRand()
	s := g.Server(p)
	s.BindObjectRest("/", new(ObjectRest))
	s.BindObjectRest("/{.struct}/{.method}", new(ObjectRest))
	s.SetPort(p)
	s.SetDumpRouteMap(false)
	s.Start()
	defer s.Shutdown()

	// 等待启动完成
	time.Sleep(200 * time.Millisecond)
	gtest.Case(t, func() {
		client := ghttp.NewClient()
		client.SetPrefix(fmt.Sprintf("http://127.0.0.1:%d", p))

		gtest.Assert(client.GetContent("/"), "1Object Get2")
		gtest.Assert(client.PutContent("/"), "1Object Put2")
		gtest.Assert(client.PostContent("/"), "1Object Post2")
		gtest.Assert(client.DeleteContent("/"), "1Object Delete2")
		gtest.Assert(client.PatchContent("/"), "1Object Patch2")
		gtest.Assert(client.OptionsContent("/"), "1Object Options2")
		resp1, err := client.Head("/")
		if err == nil {
			defer resp1.Close()
		}
		gtest.Assert(err, nil)
		gtest.Assert(resp1.Header.Get("head-ok"), "1")

		gtest.Assert(client.GetContent("/object-rest/get"), "1Object Get2")
		gtest.Assert(client.PutContent("/object-rest/put"), "1Object Put2")
		gtest.Assert(client.PostContent("/object-rest/post"), "1Object Post2")
		gtest.Assert(client.DeleteContent("/object-rest/delete"), "1Object Delete2")
		gtest.Assert(client.PatchContent("/object-rest/patch"), "1Object Patch2")
		gtest.Assert(client.OptionsContent("/object-rest/options"), "1Object Options2")
		resp2, err := client.Head("/object-rest/head")
		if err == nil {
			defer resp2.Close()
		}
		gtest.Assert(err, nil)
		gtest.Assert(resp2.Header.Get("head-ok"), "1")

		gtest.Assert(client.GetContent("/none-exist"), "Not Found")
	})
}
