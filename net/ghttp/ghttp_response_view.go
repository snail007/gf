// Copyright 2017 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.
//

package ghttp

import (
	"github.com/snail007/gf/frame/gins"
	"github.com/snail007/gf/os/gview"
)

// 展示模板，可以给定模板参数，及临时的自定义模板函数
func (r *Response) WriteTpl(tpl string, params ...gview.Params) {
	if b, err := r.ParseTpl(tpl, params...); err != nil {
		r.Write("Template Parsing Error: " + err.Error())
	} else {
		r.Write(b)
	}
}

// 展示模板内容，可以给定模板参数，及临时的自定义模板函数
func (r *Response) WriteTplContent(content string, params ...gview.Params) {
	if b, err := r.ParseTplContent(content, params...); err != nil {
		r.Write("Template Parsing Error: " + err.Error())
	} else {
		r.Write(b)
	}
}

// 解析模板文件，并返回模板内容
func (r *Response) ParseTpl(tpl string, params ...gview.Params) (string, error) {
	if r.Server.config.View != nil {
		return r.Server.config.View.Parse(tpl, r.buildInVars(params...))
	}
	return gview.Instance().Parse(tpl, r.buildInVars(params...))
}

// 解析并返回模板内容
func (r *Response) ParseTplContent(content string, params ...gview.Params) (string, error) {
	if r.Server.config.View != nil {
		return r.Server.config.View.ParseContent(content, r.buildInVars(params...))
	}
	return gview.Instance().ParseContent(content, r.buildInVars(params...))
}

// 内置变量/对象
func (r *Response) buildInVars(params ...map[string]interface{}) map[string]interface{} {
	vars := map[string]interface{}(nil)
	if len(params) > 0 {
		vars = params[0]
	} else {
		vars = make(map[string]interface{})
	}
	// 当配置文件不存在时就不赋值该模板变量，不然会报错
	if c := gins.Config(); c.FilePath() != "" {
		vars["Config"] = c.GetMap("")
	}
	vars["Cookie"] = r.request.Cookie.Map()
	vars["Session"] = r.request.Session.Map()
	vars["Get"] = r.request.GetQueryMap()
	vars["Post"] = r.request.GetPostMap()
	return vars
}
