package main

import (
	"github.com/snail007/gf/net/ghttp"
)

func main() {
	s := ghttp.GetServer()
	s.BindHandler("/log/error", func(r *ghttp.Request) {
		if j := r.GetJson(); j != nil {
			r.Response.Write(j.Get("test"))
		}
	})
	s.SetErrorLogEnabled(true)
	s.SetPort(8199)
	s.Run()
}
