package main

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/frame/gins"
	"github.com/snail007/gf/net/ghttp"
)

func main() {
	s := g.Server()
	s.SetServerRoot("public")
	s.SetNameToUriType(ghttp.NAME_TO_URI_TYPE_ALLLOWER)
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)
	s.SetPort(2333)

	s.BindHandler("/", func(r *ghttp.Request) {
		content, _ := gins.View().Parse("test.html", nil)
		r.Response.Write(content)
	})

	s.Run()
}
