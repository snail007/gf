package main

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.BindHandler("/:name", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router.Uri)
	})
	s.BindHandler("/:name/update", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router.Uri)
	})
	s.BindHandler("/:name/:action", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router.Uri)
	})
	s.BindHandler("/:name/*any", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router.Uri)
	})
	s.BindHandler("/user/list/{field}.html", func(r *ghttp.Request) {
		r.Response.Writeln(r.Router.Uri)
	})
	s.SetPort(8199)
	s.Run()
}
