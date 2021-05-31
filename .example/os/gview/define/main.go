package main

import (
	"fmt"

	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/os/gfile"
)

func main() {
	v := g.View()
	v.AddPath(gfile.MainPkgPath() + gfile.Separator + "template")
	b, err := v.Parse("index.html", map[string]interface{}{
		"k": "v",
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

}
