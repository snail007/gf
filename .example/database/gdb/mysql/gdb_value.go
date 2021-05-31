package main

import (
	"fmt"

	"github.com/snail007/gf/frame/g"
)

func main() {
	db := g.DB()
	// 开启调试模式，以便于记录所有执行的SQL
	db.SetDebug(true)

	r, e := db.Table("test").Where("id IN (?)", []interface{}{1, 2}).All()
	if e != nil {
		panic(e)
	}
	if r != nil {
		fmt.Println(r.ToList())
	}
}
