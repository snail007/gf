package demo

import (
	"github.com/snail007/gf/frame/gins"
	"github.com/snail007/gf/net/ghttp"
)

func init() {
	ghttp.GetServer().BindHandler("/config", func(r *ghttp.Request) {
		r.Response.Write(gins.Config().GetString("database.default.0.host"))
	})
}
