// Copyright 2017 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package ghttp

import "github.com/snail007/gf/container/gvar"

// 设置请求流程共享变量
func (r *Request) SetParam(key string, value interface{}) {
	if r.params == nil {
		r.params = make(map[string]interface{})
	}
	r.params[key] = value
}

// 获取请求流程共享变量
func (r *Request) GetParam(key string, def ...interface{}) *gvar.Var {
	if r.params != nil {
		if v, ok := r.params[key]; ok {
			return gvar.New(v)
		}
	}
	if len(def) > 0 {
		return gvar.New(def[0])
	}
	return gvar.New(nil)
}
