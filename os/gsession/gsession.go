// Copyright 2019 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

// Package gsession implements manager and storage features for sessions.
package gsession

import (
	"strconv"
	"strings"

	"github.com/snail007/gf/os/gtime"
	"github.com/snail007/gf/util/grand"
)

// NewSessionId creates and returns a new and unique session id string,
// the length of which is 18 bytes.
func NewSessionId() string {
	return strings.ToUpper(strconv.FormatInt(gtime.Nanosecond(), 36) + grand.Str(6))
}
