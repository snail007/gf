package main

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/net/ghttp"
	"github.com/snail007/gf/util/gconv"
)

func main() {
	s := g.Server()
	s.BindHandler("/session", func(r *ghttp.Request) {
		id := r.Session.GetInt("id")
		r.Session.Set("id", id+1)
		r.Response.Write("id:" + gconv.String(id))
	})
	s.SetPort(8199)
	s.Run()
}
