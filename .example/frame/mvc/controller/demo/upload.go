package demo

import (
	"github.com/snail007/gf/net/ghttp"
	"github.com/snail007/gf/os/gfile"
)

func Upload(r *ghttp.Request) {
	if f, h, e := r.FormFile("upload-file"); e == nil {
		defer f.Close()
		fname := gfile.Basename(h.Filename)
		buffer := make([]byte, h.Size)
		f.Read(buffer)
		gfile.PutBytes("/tmp/"+fname, buffer)
		r.Response.Write(fname + " uploaded successly")
	} else {
		r.Response.Write(e.Error())
	}
}

func UploadShow(r *ghttp.Request) {
	r.Response.Write(`
<html>
<head>
    <title>上传文件</title>
</head>
    <body>
        <form enctype="multipart/form-data" action="/upload" method="post">
            <input type="file" name="upload-file" />
            <input type="submit" value="upload" />
        </form>
    </body>
</html>
`)
}

func init() {
	ghttp.GetServer().BindHandler("/upload", Upload)
	ghttp.GetServer().BindHandler("/upload/show", UploadShow)
}
