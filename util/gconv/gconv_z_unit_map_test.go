// Copyright 2018 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gconv_test

import (
	"testing"

	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/test/gtest"
	"github.com/snail007/gf/util/gconv"
)

func Test_Map_Basic(t *testing.T) {
	gtest.Case(t, func() {
		m1 := map[string]string{
			"k": "v",
		}
		m2 := map[int]string{
			3: "v",
		}
		m3 := map[float64]float32{
			1.22: 3.1,
		}
		gtest.Assert(gconv.Map(m1), g.Map{
			"k": "v",
		})
		gtest.Assert(gconv.Map(m2), g.Map{
			"3": "v",
		})
		gtest.Assert(gconv.Map(m3), g.Map{
			"1.22": "3.1",
		})
	})
}

func Test_Map_StructWithGconvTag(t *testing.T) {
	gtest.Case(t, func() {
		type User struct {
			Uid      int
			Name     string
			SiteUrl  string `gconv:"-"`
			NickName string `gconv:"nickname, omitempty"`
			Pass1    string `gconv:"password1"`
			Pass2    string `gconv:"password2"`
		}
		user1 := User{
			Uid:     100,
			Name:    "john",
			SiteUrl: "https://goframe.org",
			Pass1:   "123",
			Pass2:   "456",
		}
		user2 := &user1
		map1 := gconv.Map(user1)
		map2 := gconv.Map(user2)
		gtest.Assert(map1["Uid"], 100)
		gtest.Assert(map1["Name"], "john")
		gtest.Assert(map1["SiteUrl"], nil)
		gtest.Assert(map1["NickName"], nil)
		gtest.Assert(map1["nickname"], nil)
		gtest.Assert(map1["password1"], "123")
		gtest.Assert(map1["password2"], "456")

		gtest.Assert(map2["Uid"], 100)
		gtest.Assert(map2["Name"], "john")
		gtest.Assert(map2["SiteUrl"], nil)
		gtest.Assert(map2["NickName"], nil)
		gtest.Assert(map2["nickname"], nil)
		gtest.Assert(map2["password1"], "123")
		gtest.Assert(map2["password2"], "456")
	})
}

func Test_Map_StructWithJsonTag(t *testing.T) {
	gtest.Case(t, func() {
		type User struct {
			Uid      int
			Name     string
			SiteUrl  string `json:"-"`
			NickName string `json:"nickname, omitempty"`
			Pass1    string `json:"password1"`
			Pass2    string `json:"password2"`
		}
		user1 := User{
			Uid:     100,
			Name:    "john",
			SiteUrl: "https://goframe.org",
			Pass1:   "123",
			Pass2:   "456",
		}
		user2 := &user1
		map1 := gconv.Map(user1)
		map2 := gconv.Map(user2)
		gtest.Assert(map1["Uid"], 100)
		gtest.Assert(map1["Name"], "john")
		gtest.Assert(map1["SiteUrl"], nil)
		gtest.Assert(map1["NickName"], nil)
		gtest.Assert(map1["nickname"], nil)
		gtest.Assert(map1["password1"], "123")
		gtest.Assert(map1["password2"], "456")

		gtest.Assert(map2["Uid"], 100)
		gtest.Assert(map2["Name"], "john")
		gtest.Assert(map2["SiteUrl"], nil)
		gtest.Assert(map2["NickName"], nil)
		gtest.Assert(map2["nickname"], nil)
		gtest.Assert(map2["password1"], "123")
		gtest.Assert(map2["password2"], "456")
	})
}

func Test_Map_PrivateAttribute(t *testing.T) {
	type User struct {
		Id   int
		name string
	}
	gtest.Case(t, func() {
		user := &User{1, "john"}
		gtest.Assert(gconv.Map(user), g.Map{"Id": 1})
	})
}

func Test_Map_StructInherit(t *testing.T) {
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
	gtest.Case(t, func() {
		user := new(User)
		user.Id = 100
		user.Nickname = "john"
		user.CreateTime = "2019"
		m := gconv.Map(user)
		gtest.Assert(m["id"], "")
		gtest.Assert(m["nickname"], user.Nickname)
		gtest.Assert(m["create_time"], "")
	})
	gtest.Case(t, func() {
		user := new(User)
		user.Id = 100
		user.Nickname = "john"
		user.CreateTime = "2019"
		m := gconv.MapDeep(user)
		gtest.Assert(m["id"], user.Id)
		gtest.Assert(m["nickname"], user.Nickname)
		gtest.Assert(m["create_time"], user.CreateTime)
	})
}

func Test_MapStruct(t *testing.T) {
	type User struct {
		Id   int
		Name string
	}
	params := g.Map{
		"key": g.Map{
			"id":   1,
			"name": "john",
		},
	}
	gtest.Case(t, func() {
		m := make(map[string]User)
		err := gconv.MapStruct(params, &m)
		gtest.Assert(err, nil)
		gtest.Assert(len(m), 1)
		gtest.Assert(m["key"].Id, 1)
		gtest.Assert(m["key"].Name, "john")
	})
	gtest.Case(t, func() {
		m := (map[string]User)(nil)
		err := gconv.MapStruct(params, &m)
		gtest.Assert(err, nil)
		gtest.Assert(len(m), 1)
		gtest.Assert(m["key"].Id, 1)
		gtest.Assert(m["key"].Name, "john")
	})
	gtest.Case(t, func() {
		m := make(map[string]*User)
		err := gconv.MapStruct(params, &m)
		gtest.Assert(err, nil)
		gtest.Assert(len(m), 1)
		gtest.Assert(m["key"].Id, 1)
		gtest.Assert(m["key"].Name, "john")
	})
	gtest.Case(t, func() {
		m := (map[string]*User)(nil)
		err := gconv.MapStruct(params, &m)
		gtest.Assert(err, nil)
		gtest.Assert(len(m), 1)
		gtest.Assert(m["key"].Id, 1)
		gtest.Assert(m["key"].Name, "john")
	})
}

func Test_MapStructDeep(t *testing.T) {
	type Ids struct {
		Id  int
		Uid int
	}
	type Base struct {
		Ids
		Time string
	}
	type User struct {
		Base
		Name string
	}
	params := g.Map{
		"key": g.Map{
			"id":   1,
			"name": "john",
		},
	}
	gtest.Case(t, func() {
		m := (map[string]*User)(nil)
		err := gconv.MapStruct(params, &m)
		gtest.Assert(err, nil)
		gtest.Assert(len(m), 1)
		gtest.Assert(m["key"].Id, 0)
		gtest.Assert(m["key"].Name, "john")
	})
	gtest.Case(t, func() {
		m := (map[string]*User)(nil)
		err := gconv.MapStructDeep(params, &m)
		gtest.Assert(err, nil)
		gtest.Assert(len(m), 1)
		gtest.Assert(m["key"].Id, 1)
		gtest.Assert(m["key"].Name, "john")
	})
}

func Test_MapStructs1(t *testing.T) {
	type User struct {
		Id   int
		Name int
	}
	params := g.Map{
		"key1": g.Slice{
			g.Map{"id": 1, "name": "john"},
			g.Map{"id": 2, "name": "smith"},
		},
		"key2": g.Slice{
			g.Map{"id": 3, "name": "green"},
			g.Map{"id": 4, "name": "jim"},
		},
	}
	gtest.Case(t, func() {
		m := make(map[string][]User)
		err := gconv.MapStructs(params, &m)
		gtest.Assert(err, nil)
		gtest.Assert(len(m), 2)
		gtest.Assert(m["key1"][0].Id, 1)
		gtest.Assert(m["key1"][1].Id, 2)
		gtest.Assert(m["key2"][0].Id, 3)
		gtest.Assert(m["key2"][1].Id, 4)
	})
	gtest.Case(t, func() {
		m := (map[string][]User)(nil)
		err := gconv.MapStructs(params, &m)
		gtest.Assert(err, nil)
		gtest.Assert(len(m), 2)
		gtest.Assert(m["key1"][0].Id, 1)
		gtest.Assert(m["key1"][1].Id, 2)
		gtest.Assert(m["key2"][0].Id, 3)
		gtest.Assert(m["key2"][1].Id, 4)
	})
	gtest.Case(t, func() {
		m := make(map[string][]*User)
		err := gconv.MapStructs(params, &m)
		gtest.Assert(err, nil)
		gtest.Assert(len(m), 2)
		gtest.Assert(m["key1"][0].Id, 1)
		gtest.Assert(m["key1"][1].Id, 2)
		gtest.Assert(m["key2"][0].Id, 3)
		gtest.Assert(m["key2"][1].Id, 4)
	})
	gtest.Case(t, func() {
		m := (map[string][]*User)(nil)
		err := gconv.MapStructs(params, &m)
		gtest.Assert(err, nil)
		gtest.Assert(len(m), 2)
		gtest.Assert(m["key1"][0].Id, 1)
		gtest.Assert(m["key1"][1].Id, 2)
		gtest.Assert(m["key2"][0].Id, 3)
		gtest.Assert(m["key2"][1].Id, 4)
	})
}

func Test_MapStructs2(t *testing.T) {
	type User struct {
		Id   int
		Name int
	}
	params := g.MapIntAny{
		100: g.Slice{
			g.Map{"id": 1, "name": "john"},
			g.Map{"id": 2, "name": "smith"},
		},
		200: g.Slice{
			g.Map{"id": 3, "name": "green"},
			g.Map{"id": 4, "name": "jim"},
		},
	}
	gtest.Case(t, func() {
		m := make(map[int][]User)
		err := gconv.MapStructs(params, &m)
		gtest.Assert(err, nil)
		gtest.Assert(len(m), 2)
		gtest.Assert(m[100][0].Id, 1)
		gtest.Assert(m[100][1].Id, 2)
		gtest.Assert(m[200][0].Id, 3)
		gtest.Assert(m[200][1].Id, 4)
	})
	gtest.Case(t, func() {
		m := make(map[int][]*User)
		err := gconv.MapStructs(params, &m)
		gtest.Assert(err, nil)
		gtest.Assert(len(m), 2)
		gtest.Assert(m[100][0].Id, 1)
		gtest.Assert(m[100][1].Id, 2)
		gtest.Assert(m[200][0].Id, 3)
		gtest.Assert(m[200][1].Id, 4)
	})
	gtest.Case(t, func() {
		m := make(map[string][]*User)
		err := gconv.MapStructs(params, &m)
		gtest.Assert(err, nil)
		gtest.Assert(len(m), 2)
		gtest.Assert(m["100"][0].Id, 1)
		gtest.Assert(m["100"][1].Id, 2)
		gtest.Assert(m["200"][0].Id, 3)
		gtest.Assert(m["200"][1].Id, 4)
	})
}

func Test_MapStructsDeep(t *testing.T) {
	type Ids struct {
		Id  int
		Uid int
	}
	type Base struct {
		Ids
		Time string
	}
	type User struct {
		Base
		Name string
	}
	params := g.MapIntAny{
		100: g.Slice{
			g.Map{"id": 1, "name": "john"},
			g.Map{"id": 2, "name": "smith"},
		},
		200: g.Slice{
			g.Map{"id": 3, "name": "green"},
			g.Map{"id": 4, "name": "jim"},
		},
	}
	gtest.Case(t, func() {
		m := make(map[string][]*User)
		err := gconv.MapStructs(params, &m)
		gtest.Assert(err, nil)
		gtest.Assert(len(m), 2)
		gtest.Assert(m["100"][0].Id, 0)
		gtest.Assert(m["100"][1].Id, 0)
		gtest.Assert(m["100"][0].Name, "john")
		gtest.Assert(m["100"][1].Name, "smith")
		gtest.Assert(m["200"][0].Id, 0)
		gtest.Assert(m["200"][1].Id, 0)
		gtest.Assert(m["200"][0].Name, "green")
		gtest.Assert(m["200"][1].Name, "jim")
	})
	gtest.Case(t, func() {
		m := make(map[string][]*User)
		err := gconv.MapStructsDeep(params, &m)
		gtest.Assert(err, nil)
		gtest.Assert(len(m), 2)
		gtest.Assert(m["100"][0].Id, 1)
		gtest.Assert(m["100"][1].Id, 2)
		gtest.Assert(m["100"][0].Name, "john")
		gtest.Assert(m["100"][1].Name, "smith")
		gtest.Assert(m["200"][0].Id, 3)
		gtest.Assert(m["200"][1].Id, 4)
		gtest.Assert(m["200"][0].Name, "green")
		gtest.Assert(m["200"][1].Name, "jim")
	})
}
