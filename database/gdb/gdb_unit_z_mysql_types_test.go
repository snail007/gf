// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gdb_test

import (
	"fmt"
	"testing"

	"github.com/snail007/gf/frame/g"

	"github.com/snail007/gf/test/gtest"
)

func Test_Types(t *testing.T) {

	gtest.Case(t, func() {
		if _, err := db.Exec(fmt.Sprintf(`
    CREATE TABLE IF NOT EXISTS types (
        id int(10) unsigned NOT NULL AUTO_INCREMENT,
        %s blob NOT NULL,
        %s binary(8) NOT NULL,
        %s date NOT NULL,
        %s decimal(5,2) NOT NULL,
        %s double NOT NULL,
        %s bit(2) NOT NULL,
        %s tinyint(1) NOT NULL,
        %s bool NOT NULL,
        PRIMARY KEY (id)
    ) ENGINE=InnoDB DEFAULT CHARSET=utf8;
    `, "`blob`", "`binary`", "`date`", "`decimal`", "`double`", "`bit`", "`tinyint`", "`bool`")); err != nil {
			gtest.Error(err)
		}
		defer dropTable("types")
		data := g.Map{
			"id":      1,
			"blob":    "i love gf",
			"binary":  []byte("abcdefgh"),
			"date":    "2018-10-24",
			"decimal": 123.456,
			"double":  123.456,
			"bit":     2,
			"tinyint": true,
			"bool":    false,
		}
		r, err := db.Table("types").Data(data).Insert()
		gtest.Assert(err, nil)
		n, _ := r.RowsAffected()
		gtest.Assert(n, 1)

		one, err := db.Table("types").One()
		gtest.Assert(err, nil)
		gtest.Assert(one["id"].Int(), 1)
		gtest.Assert(one["blob"].String(), data["blob"])
		gtest.Assert(one["binary"].String(), data["binary"])
		gtest.Assert(one["date"].String(), data["date"])
		gtest.Assert(one["decimal"].String(), 123.46)
		gtest.Assert(one["double"].String(), data["double"])
		gtest.Assert(one["bit"].Int(), data["bit"])
		gtest.Assert(one["tinyint"].Bool(), data["tinyint"])
		gtest.Assert(one["tinyint"].Bool(), data["tinyint"])
	})
}
