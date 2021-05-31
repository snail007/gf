package main

import (
	"fmt"
	"time"

	"github.com/snail007/gf/os/gfcache"
	"github.com/snail007/gf/os/gfile"
)

func main() {
	s := 0
	r := ""
	path := gfile.TempDir() + gfile.Separator + "temp"
	gfile.PutContents(path, "hello")

	s = gfcache.GetSize()
	r = gfcache.GetContents(path)
	fmt.Println(s, r)

	gfile.PutContentsAppend(path, " john")

	// 等待1秒以便gfsnotify回调能处理完成
	time.Sleep(time.Second)

	s = gfcache.GetSize()
	r = gfcache.GetContents(path)
	fmt.Println(s, r)
}
