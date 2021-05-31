package main

import (
	_ "github.com/snail007/gf/.example/frame/mvc/controller/demo"
	_ "github.com/snail007/gf/.example/frame/mvc/controller/stats"
	"github.com/snail007/gf/frame/g"
)

func main() {

	//g.Server().SetDumpRouteMap(false)
	g.Server().SetPort(8199)
	g.Server().Run()

}
