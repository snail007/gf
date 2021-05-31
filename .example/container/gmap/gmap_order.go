package main

import (
	"fmt"

	"github.com/snail007/gf/container/gmap"
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/util/gutil"
)

func main() {
	array := g.Slice{2, 3, 1, 5, 4, 6, 8, 7, 9}
	hashMap := gmap.New(true)
	linkMap := gmap.NewLinkMap(true)
	treeMap := gmap.NewTreeMap(gutil.ComparatorInt, true)
	for _, v := range array {
		hashMap.Set(v, v)
	}
	for _, v := range array {
		linkMap.Set(v, v)
	}
	for _, v := range array {
		treeMap.Set(v, v)
	}
	fmt.Println("HashMap   Keys:", hashMap.Keys())
	fmt.Println("HashMap Values:", hashMap.Values())
	fmt.Println("LinkMap   Keys:", linkMap.Keys())
	fmt.Println("LinkMap Values:", linkMap.Values())
	fmt.Println("TreeMap   Keys:", treeMap.Keys())
	fmt.Println("TreeMap Values:", treeMap.Values())
}
