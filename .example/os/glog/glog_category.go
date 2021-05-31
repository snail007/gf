package main

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/os/gfile"
	"github.com/snail007/gf/os/glog"
)

func main() {
	path := "/tmp/glog-cat"
	glog.SetPath(path)
	glog.Stdout(false).Cat("cat1").Cat("cat2").Println("test")
	list, err := gfile.ScanDir(path, "*", true)
	g.Dump(err)
	g.Dump(list)
}
