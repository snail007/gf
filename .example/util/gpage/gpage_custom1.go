package main

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/net/ghttp"
	"github.com/snail007/gf/os/gview"
	"github.com/snail007/gf/text/gstr"
	"github.com/snail007/gf/util/gpage"
)

// 分页标签使用li标签包裹
func wrapContent(page *gpage.Page) string {
	content := page.GetContent(4)
	content = gstr.ReplaceByMap(content, map[string]string{
		"<span":  "<li><span",
		"/span>": "/span></li>",
		"<a":     "<li><a",
		"/a>":    "/a></li>",
	})
	return "<ul>" + content + "</ul>"
}

func main() {
	s := ghttp.GetServer()
	s.BindHandler("/page/custom1/*page", func(r *ghttp.Request) {
		page := gpage.New(100, 10, r.Get("page"), r.URL.String(), r.Router)
		content := wrapContent(page)
		buffer, _ := gview.ParseContent(`
        <html>
            <head>
                <style>
                    a,span {padding:8px; font-size:16px;}
                    div{margin:5px 5px 20px 5px}
                </style>
            </head>
            <body>
                <div>{{.page}}</div>
            </body>
        </html>
        `, g.Map{
			"page": content,
		})
		r.Response.Write(buffer)
	})
	s.SetPort(10000)
	s.Run()
}
