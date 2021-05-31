package stats

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/net/ghttp"
)

var (
	total int
)

func init() {
	g.Server().BindHandler("/stats/total", showTotal)
}

func showTotal(r *ghttp.Request) {
	total++
	r.Response.Write("total:", total)
}
