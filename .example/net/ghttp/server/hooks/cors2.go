package main

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/frame/gmvc"
	"github.com/snail007/gf/net/ghttp"
)

type Order2 struct {
	gmvc.Controller
}

func (o *Order2) Get() {
	o.Response.Write("GET")
}

func main() {
	s := g.Server()
	s.BindHookHandlerByMap("/api.v2/*any", map[string]ghttp.HandlerFunc{
		"BeforeServe": func(r *ghttp.Request) {
			r.Response.CORSDefault()
		},
	})
	s.BindControllerRest("/api.v2/{.struct}", new(Order2))
	s.SetPort(8199)
	s.Run()
}
