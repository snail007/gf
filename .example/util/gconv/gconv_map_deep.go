package main

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/util/gconv"
)

func main() {
	type Ids struct {
		Id  int `json:"id"`
		Uid int `json:"uid"`
	}
	type Base struct {
		Ids
		CreateTime string `json:"create_time"`
	}
	type User struct {
		Base
		Passport string `json:"passport"`
		Password string `json:"password"`
		Nickname string `json:"nickname"`
	}
	user := new(User)
	user.Id = 1
	user.Uid = 100
	user.Nickname = "John"
	user.Passport = "johng"
	user.Password = "123456"
	user.CreateTime = "2019"
	g.Dump(gconv.MapDeep(user))
}
