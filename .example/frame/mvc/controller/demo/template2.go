package demo

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/net/ghttp"
)

func init() {
	ghttp.GetServer().BindHandler("/template2", func(r *ghttp.Request) {
		content, _ := g.View().Parse("index.tpl", map[string]interface{}{
			"id":   123,
			"name": "john",
		})
		r.Response.Write(content)
	})
}
