package main

import (
	"github.com/snail007/gf/net/ghttp"
)

func main() {
	s := ghttp.GetServer()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Writeln("您可以同时通过HTTP和HTTPS方式看到该内容！")
	})
	s.EnableHTTPS("./server.crt", "./server.key")
	s.SetHTTPSPort(8100, 8200)
	s.SetPort(8300, 8400)
	s.EnableAdmin()
	s.Run()
}
