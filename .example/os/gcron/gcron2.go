package main

import (
	"time"

	"github.com/snail007/gf/os/gcron"
	"github.com/snail007/gf/os/glog"
)

func test() {
	glog.Println(111)
}

func main() {
	_, err := gcron.AddOnce("@every 2s", test)
	if err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Second)
}
