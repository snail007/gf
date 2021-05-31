package main

import (
	"fmt"

	"github.com/snail007/gf/net/ghttp"
	"github.com/snail007/gf/os/glog"
)

func main() {
	path := "/home/john/Workspace/Go/github.com/snail007/gf/version.go"
	r, e := ghttp.Post("http://127.0.0.1:8199/upload", "name=john&age=18&upload-file=@file:"+path)
	if e != nil {
		glog.Error(e)
	} else {
		fmt.Println(string(r.ReadAll()))
		r.Close()
	}
}
