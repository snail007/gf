package demo

import (
	"github.com/snail007/gf/net/ghttp"
	"github.com/snail007/gf/os/gtime"
)

func init() {
	ghttp.GetServer().BindHandler("/cookie", Cookie)
}

func Cookie(r *ghttp.Request) {
	datetime := r.Cookie.Get("datetime")
	r.Cookie.Set("datetime", gtime.Datetime())
	r.Response.Write("datetime:" + datetime)
}
