package main

import (
	"github.com/snail007/gf/container/garray"
	"github.com/snail007/gf/frame/g"
)

func main() {
	array := garray.NewSortedStringArray()
	array.Add("1")
	array.Add("2")
	array.Add("3")
	array.Add("4")
	array.Add("5")
	array.Add("6")
	array.Add("7")
	array.Add("8")
	array.Add("9")
	g.Dump(array.Slice())
}
