package main

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/util/gconv"
)

func main() {
	type User struct {
		Uid  int
		Name string
	}
	user := (*User)(nil)
	params := g.Map{
		"uid":  1,
		"name": "john",
	}
	err := gconv.Struct(params, &user)
	if err != nil {
		panic(err)
	}
	g.Dump(user)
}
