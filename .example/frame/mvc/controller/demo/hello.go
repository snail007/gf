package demo

import "github.com/snail007/gf/net/ghttp"

func init() {
	ghttp.GetServer().BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("Hello World!")
	})
}
