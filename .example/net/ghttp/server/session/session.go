package main

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/net/ghttp"
	"github.com/snail007/gf/os/gtime"
	"time"
)

func main() {
	s := g.Server()
	s.SetSessionMaxAge(2 * time.Second)
	s.BindHandler("/set", func(r *ghttp.Request) {
		r.Session.Set("time", gtime.Second())
		r.Response.Write("ok")
	})
	s.BindHandler("/get", func(r *ghttp.Request) {
		r.Response.WriteJson(r.Session.Map())
	})
	s.BindHandler("/clear", func(r *ghttp.Request) {
		r.Session.Clear()
	})
	s.SetPort(8199)
	s.Run()
}
