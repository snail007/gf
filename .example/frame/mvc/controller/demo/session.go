package demo

import (
	"strconv"

	"github.com/snail007/gf/net/ghttp"
)

func init() {
	ghttp.GetServer().BindHandler("/session", Session)
}

func Session(r *ghttp.Request) {
	id := r.Session.GetInt("id")
	r.Session.Set("id", id+1)
	r.Response.Write("id:" + strconv.Itoa(id))
}
