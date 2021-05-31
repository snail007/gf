// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gsession

import (
	"testing"

	"github.com/snail007/gf/test/gtest"
)

func Test_NewSessionId(t *testing.T) {
	gtest.Case(t, func() {
		id1 := NewSessionId()
		id2 := NewSessionId()
		gtest.AssertNE(id1, id2)
		gtest.Assert(len(id1), 18)
	})
}
