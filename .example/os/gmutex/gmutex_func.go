package main

import (
	"time"

	"github.com/snail007/gf/os/glog"

	"github.com/snail007/gf/os/gmutex"
)

func main() {
	mu := gmutex.New()
	go mu.LockFunc(func() {
		glog.Println("lock func1")
		time.Sleep(1 * time.Second)
	})
	time.Sleep(time.Millisecond)
	go mu.LockFunc(func() {
		glog.Println("lock func2")
	})
	time.Sleep(2 * time.Second)
}
