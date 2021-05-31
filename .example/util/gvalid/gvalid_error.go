package main

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/util/gvalid"
)

// 返回结果方法示例
func main() {
	type User struct {
		Password        string `gvalid:"password@password"`
		ConfiemPassword string `gvalid:"confirm_password@password|same:password#|密码与确认密码不一致"`
	}

	user := &User{
		Password:        "123456",
		ConfiemPassword: "",
	}

	e := gvalid.CheckStruct(user, nil)
	g.Dump(e.Map())
	g.Dump(e.Maps())
	g.Dump(e.String())
	g.Dump(e.Strings())
	g.Dump(e.FirstItem())
	g.Dump(e.FirstRule())
	g.Dump(e.FirstString())
}
