package main

import (
	"time"

	"github.com/snail007/gf/os/glog"
	"github.com/snail007/gf/os/gtime"
	"github.com/snail007/gf/os/gtimer"
)

func main() {
	gtimer.SetTimeout(3*time.Second, func() {
		glog.SetDebug(false)
	})
	for {
		glog.Debug(gtime.Datetime())
		time.Sleep(time.Second)
	}
}
