package main

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/net/ghttp"
	"github.com/snail007/gf/os/glog"
)

func main() {
	s := g.Server()
	s.BindHandler("/priority/show", func(r *ghttp.Request) {
		r.Response.Write("priority test")
	})

	s.BindHookHandlerByMap("/priority/:name", map[string]ghttp.HandlerFunc{
		"BeforeServe": func(r *ghttp.Request) {
			glog.Println(r.Router.Uri)
		},
	})
	s.BindHookHandlerByMap("/priority/*any", map[string]ghttp.HandlerFunc{
		"BeforeServe": func(r *ghttp.Request) {
			glog.Println(r.Router.Uri)
		},
	})
	s.BindHookHandlerByMap("/priority/show", map[string]ghttp.HandlerFunc{
		"BeforeServe": func(r *ghttp.Request) {
			glog.Println(r.Router.Uri)
		},
	})
	s.SetPort(8199)
	s.Run()
}
