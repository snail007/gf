// Copyright 2018 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package ghttp_test

import (
	"github.com/snail007/gf/container/garray"
	"github.com/snail007/gf/os/genv"
)

var (
	// 用于测试的端口数组，随机获取
	ports = garray.NewIntArray(true)
)

func init() {
	genv.Set("UNDER_TEST", "1")
	for i := 8000; i <= 9000; i++ {
		ports.Append(i)
	}
}
