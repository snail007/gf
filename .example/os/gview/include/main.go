package main

import (
	"fmt"

	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/os/gfile"
)

func main() {
	v := g.View()
	// 设置模板目录为当前main.go所在目录下的template目录
	v.AddPath(gfile.MainPkgPath() + gfile.Separator + "template2")
	b, err := v.Parse("index.html", map[string]interface{}{
		"k": "v",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))
}
