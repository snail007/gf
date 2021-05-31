package main

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/net/ghttp"
	"github.com/snail007/gf/os/gview"
	"github.com/snail007/gf/util/gpage"
)

// 自定义分页名称
func pageContent(page *gpage.Page) string {
	page.NextPageTag = "NextPage"
	page.PrevPageTag = "PrevPage"
	page.FirstPageTag = "HomePage"
	page.LastPageTag = "LastPage"
	pageStr := page.FirstPage()
	pageStr += page.PrevPage()
	pageStr += page.PageBar("current-page")
	pageStr += page.NextPage()
	pageStr += page.LastPage()
	return pageStr
}

func main() {
	s := ghttp.GetServer()
	s.BindHandler("/page/custom2/*page", func(r *ghttp.Request) {
		page := gpage.New(100, 10, r.Get("page"), r.URL.String(), r.Router)
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
			"page": pageContent(page),
		})
		r.Response.Write(buffer)
	})
	s.SetPort(10000)
	s.Run()
}
