package main

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/os/glog"
)

func main() {
	g.TryCatch(func() {
		glog.Println("hello")
		g.Throw("exception")
		glog.Println("world")
	})

	g.TryCatch(func() {
		glog.Println("hello")
		g.Throw("exception")
		glog.Println("world")
	}, func(exception interface{}) {
		glog.Error(exception)
	})
}
