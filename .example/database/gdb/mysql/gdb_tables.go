package main

import (
	"github.com/snail007/gf/frame/g"
)

func main() {
	db := g.DB()
	db.SetDebug(true)

	tables, err := db.Tables()
	if err != nil {
		panic(err)
	}
	if tables != nil {
		g.Dump(tables)
	}
}
