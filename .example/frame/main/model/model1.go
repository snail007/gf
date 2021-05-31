package main

import (
	"github.com/snail007/gf/.example/frame/mvc/model/test"
	"github.com/snail007/gf/frame/g"
)

func main() {
	g.DB().SetDebug(true)
	user, err := test.ModelUser().One()
	g.Dump(err)
	g.Dump(user)
}
