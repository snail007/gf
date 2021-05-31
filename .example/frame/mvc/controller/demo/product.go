package demo

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/net/ghttp"
)

type Product struct {
	total int
}

func init() {
	p := &Product{}
	g.Server().BindHandler("/product/total", p.Total)
	g.Server().BindHandler("/product/list/{page}.html", p.List)
}

func (p *Product) Total(r *ghttp.Request) {
	p.total++
	r.Response.Write("total: ", p.total)
}

func (p *Product) List(r *ghttp.Request) {
	r.Response.Write("page: ", r.Get("page"))
}
