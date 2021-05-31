// Copyright 2017 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

// Package gyaml provides accessing and converting for YAML content.
package gyaml

import (
	"encoding/json"

	"github.com/snail007/gf/util/gconv"

	yaml3 "gopkg.in/yaml.v3"
)

func Encode(v interface{}) ([]byte, error) {
	return yaml3.Marshal(v)
}

func Decode(v []byte) (interface{}, error) {
	var result map[string]interface{}
	if err := yaml3.Unmarshal(v, &result); err != nil {
		return nil, err
	}
	return gconv.Map(result), nil
}

func DecodeTo(v []byte, result interface{}) error {
	return yaml3.Unmarshal(v, result)
}

func ToJson(v []byte) ([]byte, error) {
	if r, err := Decode(v); err != nil {
		return nil, err
	} else {
		return json.Marshal(r)
	}
}
