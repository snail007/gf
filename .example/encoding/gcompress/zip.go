package main

import (
	"fmt"

	"github.com/snail007/gf/encoding/gcompress"
)

func main() {
	err := gcompress.ZipPath(
		`D:\Workspace\Go\GOPATH\src\github.com\gogf\gf\geg`,
		`D:\Workspace\Go\GOPATH\src\github.com\gogf\gf\geg\encoding\gcompress\data.zip`,
		"my-dir",
	)
	fmt.Println(err)
}
