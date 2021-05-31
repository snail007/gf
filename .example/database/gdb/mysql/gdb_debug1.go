package main

import (
	"fmt"

	"github.com/snail007/gf/database/gdb"
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/os/glog"
)

func main() {
	gdb.AddDefaultConfigNode(gdb.ConfigNode{
		Host:    "127.0.0.1",
		Port:    "3306",
		User:    "root",
		Pass:    "12345678",
		Name:    "test",
		Type:    "mysql",
		Role:    "master",
		Charset: "utf8",
	})
	db, err := gdb.New()
	if err != nil {
		panic(err)
	}
	db.SetDebug(true)

	glog.SetPath("/tmp")

	// 执行3条SQL查询
	for i := 1; i <= 3; i++ {
		db.Table("user").Where("uid=?", i).One()
	}
	// 构造一条错误查询
	db.Table("user").Where("no_such_field=?", "just_test").One()

	db.Table("user").Data(g.Map{"name": "smith"}).Where("uid=?", 1).Save()

	db.PrintQueriedSqls()

	fmt.Println(db.GetLastSql())
}
