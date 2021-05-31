package demo

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/net/ghttp"
)

type Order struct{}

func init() {
	g.Server().BindObject("/{.struct}-{.method}", new(Order))
}

func (o *Order) List(r *ghttp.Request) {
	r.Response.Write("List")
}
