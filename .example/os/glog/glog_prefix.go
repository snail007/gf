package main

import (
	"github.com/snail007/gf/os/glog"
)

func main() {
	l := glog.New()
	l.SetPrefix("[API]")
	l.Println("hello world")
	l.Error("error occurred")
}
