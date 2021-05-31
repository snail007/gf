package demo

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/net/ghttp"
)

type Object struct{}

func init() {
	g.Server().BindObject("/object", new(Object))
}

func (o *Object) Index(r *ghttp.Request) {
	r.Response.Write("object index")
}

func (o *Object) Show(r *ghttp.Request) {
	r.Response.Write("object show")
}
