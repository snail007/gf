// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gdb_test

import (
	"testing"

	"github.com/snail007/gf/database/gdb"
	"github.com/snail007/gf/test/gtest"
)

func Test_Instance(t *testing.T) {
	gtest.Case(t, func() {
		_, err := gdb.Instance("none")
		gtest.AssertNE(err, nil)

		db, err := gdb.Instance()
		gtest.Assert(err, nil)

		err1 := db.PingMaster()
		err2 := db.PingSlave()
		gtest.Assert(err1, nil)
		gtest.Assert(err2, nil)
	})
}
