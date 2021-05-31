package main

import (
	"fmt"
	"time"

	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/os/gtime"
)

func main() {
	v := g.View()
	v.SetPath(`D:\Workspace\Go\GOPATH\src\gitee.com\johng\gf\geg\os\gview`)
	gtime.SetInterval(time.Second, func() bool {
		b, _ := v.Parse("gview.tpl", nil)
		fmt.Println(string(b))
		return true
	})
	select {}
}
