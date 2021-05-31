// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gdb_test

import (
	"fmt"
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/os/gtime"
	"github.com/snail007/gf/test/gtest"
	"strings"
	"testing"
)

func Test_TX_Query_Oracle(t *testing.T) {
	tx, err := db.Begin()
	if err != nil {
		gtest.Fatal(err)
	}
	if rows, err := tx.Query("SELECT ? FROM DUAL", 1); err != nil {
		gtest.Fatal(err)
	} else {
		rows.Close()
	}
	if rows, err := tx.Query("SELECT ?+? FROM DUAL", 1, 2); err != nil {
		gtest.Fatal(err)
	} else {
		rows.Close()
	}
	if rows, err := tx.Query("SELECT ?+? FROM DUAL", g.Slice{1, 2}); err != nil {
		gtest.Fatal(err)
	} else {
		rows.Close()
	}
	if _, err := tx.Query("ERROR"); err == nil {
		gtest.Fatal("FAIL")
	}
	if err := tx.Commit(); err != nil {
		gtest.Fatal(err)
	}
}

func Test_TX_Exec_Oracle(t *testing.T) {
	tx, err := db.Begin()
	if err != nil {
		gtest.Fatal(err)
	}
	if _, err := tx.Exec("SELECT ? FROM DUAL", 1); err != nil {
		gtest.Fatal(err)
	}
	if _, err := tx.Exec("SELECT ?+? FROM DUAL", 1, 2); err != nil {
		gtest.Fatal(err)
	}
	if _, err := tx.Exec("SELECT ?+? FROM DUAL", g.Slice{1, 2}); err != nil {
		gtest.Fatal(err)
	}
	if _, err := tx.Exec("ERROR"); err == nil {
		gtest.Fatal("FAIL")
	}
	if err := tx.Commit(); err != nil {
		gtest.Fatal(err)
	}
}

func Test_TX_Commit_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}
	tx, err := oradb.Begin()
	if err != nil {
		gtest.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		gtest.Fatal(err)
	}
}

func Test_TX_Rollback_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}
	tx, err := oradb.Begin()
	if err != nil {
		gtest.Fatal(err)
	}
	if err := tx.Rollback(); err != nil {
		gtest.Fatal(err)
	}
}

func Test_TX_Prepare_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}
	tx, err := oradb.Begin()
	if err != nil {
		gtest.Fatal(err)
	}
	st, err := tx.Prepare("SELECT 100 FROM DUAL")
	if err != nil {
		gtest.Fatal(err)
	}
	rows, err := st.Query()
	if err != nil {
		gtest.Fatal(err)
	}
	array, err := rows.Columns()
	if err != nil {
		gtest.Fatal(err)
	}
	gtest.Assert(array[0], "100")
	if err := rows.Close(); err != nil {
		gtest.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		gtest.Fatal(err)
	}
}

func Test_TX_Insert_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}
	table := createTableOracle()
	defer dropTableOracle(table)

	tx, err := oradb.Begin()
	if err != nil {
		gtest.Fatal(err)
	}
	if _, err := tx.Insert(table, g.Map{
		"id":          1,
		"passport":    "t1",
		"password":    "25d55ad283aa400af464c76d713c07ad",
		"nickname":    "T1",
		"create_time": gtime.Now().String(),
	}); err != nil {
		gtest.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		gtest.Fatal(err)
	}
	if n, err := oradb.Table(table).Count(); err != nil {
		gtest.Fatal(err)
	} else {
		gtest.Assert(n, 1)
	}
}

func Test_TX_BatchInsert_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}

	table := createTableOracle()
	defer dropTableOracle(table)

	tx, err := oradb.Begin()
	if err != nil {
		gtest.Fatal(err)
	}
	if _, err := tx.BatchInsert(table, g.List{
		{
			"id":          2,
			"passport":    "t",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T2",
			"create_time": gtime.Now().String(),
		},
		{
			"id":          3,
			"passport":    "t3",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T3",
			"create_time": gtime.Now().String(),
		},
	}, 10); err != nil {
		gtest.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		gtest.Fatal(err)
	}
	if n, err := oradb.Table(table).Count(); err != nil {
		gtest.Fatal(err)
	} else {
		gtest.Assert(n, 2)
	}
}

func Test_TX_BatchReplace_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}

	table := createTableOracle()
	defer dropTableOracle(table)

	tx, err := oradb.Begin()
	if err != nil {
		gtest.Fatal(err)
	}
	if _, err := tx.BatchReplace(table, g.List{
		{
			"id":          2,
			"passport":    "t2",
			"password":    "p2",
			"nickname":    "T2",
			"create_time": gtime.Now().String(),
		},
		{
			"id":          4,
			"passport":    "t4",
			"password":    "25d55ad283aa400af464c76d713c07ad",
			"nickname":    "T4",
			"create_time": gtime.Now().String(),
		},
	}, 10); err != nil {
		gtest.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		gtest.Fatal(err)
	}
	// 数据数量
	if n, err := oradb.Table(table).Count(); err != nil {
		gtest.Fatal(err)
	} else {
		gtest.Assert(n, 2)
	}
	// 检查replace后的数值
	if value, err := oradb.Table(table).Fields("password").Where("id", 2).Value(); err != nil {
		gtest.Fatal(err)
	} else {
		gtest.Assert(strings.TrimSpace(value.String()), "p2")
	}
}

func Test_TX_BatchSave_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}
	table := createTableOracle()
	defer dropTableOracle(table)

	tx, err := oradb.Begin()
	if err != nil {
		gtest.Fatal(err)
	}
	if _, err := tx.BatchSave(table, g.List{
		{
			"id":          4,
			"passport":    "t4",
			"password":    "p4",
			"nickname":    "T4",
			"create_time": gtime.Now().String(),
		},
	}, 10); err != nil {
		gtest.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		gtest.Fatal(err)
	}
	// 数据数量
	if n, err := oradb.Table(table).Count(); err != nil {
		gtest.Fatal(err)
	} else {
		gtest.Assert(n, 1)
	}
	// 检查replace后的数值
	if value, err := oradb.Table(table).Fields("password").Where("id", 4).Value(); err != nil {
		gtest.Fatal(err)
	} else {
		gtest.Assert(strings.TrimSpace(value.String()), "p4")
	}
}

func Test_TX_Replace_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}
	table := createInitTableOracle()
	defer dropTableOracle(table)

	tx, err := oradb.Begin()
	if err != nil {
		gtest.Fatal(err)
	}
	if _, err := tx.Replace(table, g.Map{
		"id":          1,
		"passport":    "t11",
		"password":    "25d55ad283aa400af464c76d713c07ad",
		"nickname":    "T11",
		"create_time": gtime.Now().String(),
	}); err != nil {
		gtest.Fatal(err)
	}
	if err := tx.Rollback(); err != nil {
		gtest.Fatal(err)
	}
	if value, err := oradb.Table(table).Fields("nickname").Where("id", 1).Value(); err != nil {
		gtest.Fatal(err)
	} else {
		gtest.Assert(value.String(), "T1")
	}
}

func Test_TX_Save_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}
	table := createTableOracle()
	defer dropTableOracle(table)

	tx, err := oradb.Begin()
	if err != nil {
		gtest.Fatal(err)
	}
	if _, err := tx.Save(table, g.Map{
		"id":          1,
		"passport":    "t11",
		"password":    "25d55ad283aa400af464c76d713c07ad",
		"nickname":    "T11",
		"create_time": gtime.Now().String(),
	}); err != nil {
		gtest.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		gtest.Fatal(err)
	}
	if value, err := oradb.Table(table).Fields("nickname").Where("id", 1).Value(); err != nil {
		gtest.Fatal(err)
	} else {
		gtest.Assert(value.String(), "T11")
	}
}

func Test_TX_Update_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}
	table := createInitTableOracle()
	defer dropTableOracle(table)

	gtest.Case(t, func() {
		tx, err := oradb.Begin()
		if err != nil {
			gtest.Fatal(err)
		}
		if result, err := oradb.Update(table, "create_time='2010-10-10 00:00:01'", "id=3"); err != nil {
			gtest.Fatal(err)
		} else {
			n, _ := result.RowsAffected()
			gtest.Assert(n, 1)
		}
		if err := tx.Commit(); err != nil {
			gtest.Fatal(err)
		}
		if value, err := oradb.Table(table).Fields("create_time").Where("id", 3).Value(); err != nil {
			gtest.Fatal(err)
		} else {
			gtest.Assert(value.String(), "2010-10-10 00:00:01")
		}
	})
}

func Test_TX_GetAll_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}

	table := createInitTableOracle()
	defer dropTableOracle(table)

	tx, err := oradb.Begin()
	if err != nil {
		gtest.Fatal(err)
	}
	if result, err := tx.GetAll(fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 1); err != nil {
		gtest.Fatal(err)
	} else {
		gtest.Assert(len(result), 1)
	}
	if err := tx.Commit(); err != nil {
		gtest.Fatal(err)
	}
}

func Test_TX_GetOne_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}

	table := createInitTableOracle()
	defer dropTableOracle(table)

	tx, err := oradb.Begin()
	if err != nil {
		gtest.Fatal(err)
	}
	if record, err := tx.GetOne(fmt.Sprintf("SELECT * FROM %s WHERE passport=?", table), "t2"); err != nil {
		gtest.Fatal(err)
	} else {
		if record == nil {
			gtest.Fatal("FAIL")
		}
		gtest.Assert(record["NICKNAME"].String(), "T2")
	}
	if err := tx.Commit(); err != nil {
		gtest.Fatal(err)
	}
}

func Test_TX_GetValue_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}

	table := createInitTableOracle()
	defer dropTableOracle(table)

	tx, err := oradb.Begin()
	if err != nil {
		gtest.Fatal(err)
	}
	if value, err := tx.GetValue(fmt.Sprintf("SELECT id FROM %s WHERE passport=?", table), "t3"); err != nil {
		gtest.Fatal(err)
	} else {
		gtest.Assert(value.Int(), 3)
	}
	if err := tx.Commit(); err != nil {
		gtest.Fatal(err)
	}
}

func Test_TX_GetCount_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}

	table := createInitTableOracle()
	defer dropTableOracle(table)

	tx, err := oradb.Begin()
	if err != nil {
		gtest.Fatal(err)
	}
	if count, err := tx.GetCount(fmt.Sprintf("SELECT * FROM %s", table)); err != nil {
		gtest.Fatal(err)
	} else {
		gtest.Assert(count, INIT_DATA_SIZE)
	}
	if err := tx.Commit(); err != nil {
		gtest.Fatal(err)
	}
}

func Test_TX_GetStruct_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}
	table := createInitTableOracle()
	defer dropTableOracle(table)
	_, err := oradb.Table(table).Data("create_time", "2010-10-10 00:00:01").Where("id = ?", 3).Update()
	gtest.Assert(err, nil)

	gtest.Case(t, func() {
		tx, err := oradb.Begin()
		if err != nil {
			gtest.Fatal(err)
		}
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		user := new(User)
		if err := tx.GetStruct(user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3); err != nil {
			gtest.Fatal(err)
		}
		gtest.Assert(user.NickName, "T3")
		gtest.Assert(user.CreateTime.String(), "2010-10-10 00:00:01")
		if err := tx.Commit(); err != nil {
			gtest.Fatal(err)
		}
	})
	gtest.Case(t, func() {
		tx, err := oradb.Begin()
		if err != nil {
			gtest.Fatal(err)
		}
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		user := new(User)
		if err := tx.GetStruct(user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3); err != nil {
			gtest.Fatal(err)
		}
		gtest.Assert(user.NickName, "T3")
		gtest.Assert(user.CreateTime.String(), "2010-10-10 00:00:01")
		if err := tx.Commit(); err != nil {
			gtest.Fatal(err)
		}
	})
}

func Test_TX_GetStructs_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}
	table := createInitTableOracle()
	defer dropTableOracle(table)
	_, err := oradb.Table(table).Data("create_time", "2010-10-10 00:00:01").Where("id = ?", 1).Update()
	gtest.Assert(err, nil)

	gtest.Case(t, func() {
		tx, err := oradb.Begin()
		if err != nil {
			gtest.Fatal(err)
		}
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		var users []User
		if err := tx.GetStructs(&users, fmt.Sprintf("SELECT * FROM %s WHERE id>=? order by id", table), 1); err != nil {
			gtest.Fatal(err)
		}
		gtest.Assert(len(users), INIT_DATA_SIZE)
		gtest.Assert(users[0].Id, 1)
		gtest.Assert(users[1].Id, 2)
		gtest.Assert(users[2].Id, 3)
		gtest.Assert(users[0].NickName, "T1")
		gtest.Assert(users[1].NickName, "T2")
		gtest.Assert(users[2].NickName, "T3")
		gtest.Assert(users[0].CreateTime.String(), "2010-10-10 00:00:01")
		if err := tx.Commit(); err != nil {
			gtest.Fatal(err)
		}
	})

	gtest.Case(t, func() {
		tx, err := oradb.Begin()
		if err != nil {
			gtest.Fatal(err)
		}
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var users []User
		if err := tx.GetStructs(&users, fmt.Sprintf("SELECT * FROM %s WHERE id>=? order by id", table), 1); err != nil {
			gtest.Fatal(err)
		}
		gtest.Assert(len(users), INIT_DATA_SIZE)
		gtest.Assert(users[0].Id, 1)
		gtest.Assert(users[1].Id, 2)
		gtest.Assert(users[2].Id, 3)
		gtest.Assert(users[0].NickName, "T1")
		gtest.Assert(users[1].NickName, "T2")
		gtest.Assert(users[2].NickName, "T3")
		gtest.Assert(users[0].CreateTime.String(), "2010-10-10 00:00:01")
		if err := tx.Commit(); err != nil {
			gtest.Fatal(err)
		}
	})
}

func Test_TX_GetScan_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}
	table := createInitTableOracle()
	defer dropTableOracle(table)
	_, err := oradb.Table(table).Data("create_time", "2010-10-10 00:00:01").Where("id = ?", 3).Update()
	gtest.Assert(err, nil)

	gtest.Case(t, func() {
		tx, err := oradb.Begin()
		if err != nil {
			gtest.Fatal(err)
		}
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		user := new(User)
		if err := tx.GetScan(user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3); err != nil {
			gtest.Fatal(err)
		}
		gtest.Assert(user.NickName, "T3")
		gtest.Assert(user.CreateTime.String(), "2010-10-10 00:00:01")
		if err := tx.Commit(); err != nil {
			gtest.Fatal(err)
		}
	})
	gtest.Case(t, func() {
		tx, err := oradb.Begin()
		if err != nil {
			gtest.Fatal(err)
		}
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		user := new(User)
		if err := tx.GetScan(user, fmt.Sprintf("SELECT * FROM %s WHERE id=?", table), 3); err != nil {
			gtest.Fatal(err)
		}
		gtest.Assert(user.NickName, "T3")
		gtest.Assert(user.CreateTime.String(), "2010-10-10 00:00:01")
		if err := tx.Commit(); err != nil {
			gtest.Fatal(err)
		}
	})

	gtest.Case(t, func() {
		tx, err := oradb.Begin()
		if err != nil {
			gtest.Fatal(err)
		}
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime gtime.Time
		}
		var users []User
		if err := tx.GetScan(&users, fmt.Sprintf("SELECT * FROM %s WHERE id>=? order by id", table), 1); err != nil {
			gtest.Fatal(err)
		}
		gtest.Assert(len(users), INIT_DATA_SIZE)
		gtest.Assert(users[0].Id, 1)
		gtest.Assert(users[1].Id, 2)
		gtest.Assert(users[2].Id, 3)
		gtest.Assert(users[0].NickName, "T1")
		gtest.Assert(users[1].NickName, "T2")
		gtest.Assert(users[2].NickName, "T3")
		gtest.Assert(users[2].CreateTime.String(), "2010-10-10 00:00:01")
		if err := tx.Commit(); err != nil {
			gtest.Fatal(err)
		}
	})

	gtest.Case(t, func() {
		tx, err := oradb.Begin()
		if err != nil {
			gtest.Fatal(err)
		}
		type User struct {
			Id         int
			Passport   string
			Password   string
			NickName   string
			CreateTime *gtime.Time
		}
		var users []User
		if err := tx.GetScan(&users, fmt.Sprintf("SELECT * FROM %s WHERE id>=? order by id", table), 1); err != nil {
			gtest.Fatal(err)
		}
		gtest.Assert(len(users), INIT_DATA_SIZE)
		gtest.Assert(users[0].Id, 1)
		gtest.Assert(users[1].Id, 2)
		gtest.Assert(users[2].Id, 3)
		gtest.Assert(users[0].NickName, "T1")
		gtest.Assert(users[1].NickName, "T2")
		gtest.Assert(users[2].NickName, "T3")
		gtest.Assert(users[2].CreateTime.String(), "2010-10-10 00:00:01")
		if err := tx.Commit(); err != nil {
			gtest.Fatal(err)
		}
	})
}

func Test_TX_Delete_Oracle(t *testing.T) {
	if oradb == nil {
		return
	}
	table := createInitTableOracle()
	defer dropTableOracle(table)

	tx, err := oradb.Begin()
	if err != nil {
		gtest.Fatal(err)
	}
	if _, err := tx.Delete(table, nil); err != nil {
		gtest.Fatal(err)
	}
	if err := tx.Commit(); err != nil {
		gtest.Fatal(err)
	}
	if n, err := oradb.Table(table).Count(); err != nil {
		gtest.Fatal(err)
	} else {
		gtest.Assert(n, 0)
	}
}
