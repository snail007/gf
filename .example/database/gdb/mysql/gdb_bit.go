package main

import (
	"fmt"

	"github.com/snail007/gf/frame/g"
)

func main() {
	db := g.DB()
	db.SetDebug(true)

	r, e := db.Table("test").All()
	if e != nil {
		panic(e)
	}
	if r != nil {
		fmt.Println(r.ToList())
	}
}
