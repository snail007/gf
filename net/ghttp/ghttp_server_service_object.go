// Copyright 2018 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package ghttp

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/snail007/gf/os/gfile"
	"github.com/snail007/gf/os/glog"
	"github.com/snail007/gf/text/gregex"
	"github.com/snail007/gf/text/gstr"
)

// 绑定对象到URI请求处理中，会自动识别方法名称，并附加到对应的URI地址后面
// 第三个参数methods用以指定需要注册的方法，支持多个方法名称，多个方法以英文“,”号分隔，区分大小写
func (s *Server) BindObject(pattern string, obj interface{}, methods ...string) {
	// 当pattern中的method为all时，去掉该method，以便于后续方法判断
	domain, method, path, err := s.parsePattern(pattern)
	if err != nil {
		glog.Error(err)
		return
	}
	if strings.EqualFold(method, gDEFAULT_METHOD) {
		pattern = s.serveHandlerKey("", path, domain)
	}

	methodMap := (map[string]bool)(nil)
	if len(methods) > 0 {
		methodMap = make(map[string]bool)
		for _, v := range strings.Split(methods[0], ",") {
			methodMap[strings.TrimSpace(v)] = true
		}
	}
	m := make(handlerMap)
	v := reflect.ValueOf(obj)
	t := v.Type()
	sname := t.Elem().Name()
	initFunc := (func(*Request))(nil)
	shutFunc := (func(*Request))(nil)
	if v.MethodByName("Init").IsValid() {
		initFunc = v.MethodByName("Init").Interface().(func(*Request))
	}
	if v.MethodByName("Shut").IsValid() {
		shutFunc = v.MethodByName("Shut").Interface().(func(*Request))
	}
	pkgPath := t.Elem().PkgPath()
	pkgName := gfile.Basename(pkgPath)
	for i := 0; i < v.NumMethod(); i++ {
		mname := t.Method(i).Name
		if methodMap != nil && !methodMap[mname] {
			continue
		}
		if mname == "Init" || mname == "Shut" {
			continue
		}
		objName := gstr.Replace(t.String(), fmt.Sprintf(`%s.`, pkgName), "")
		if objName[0] == '*' {
			objName = fmt.Sprintf(`(%s)`, objName)
		}
		itemFunc, ok := v.Method(i).Interface().(func(*Request))
		if !ok {
			if len(methodMap) > 0 {
				// 指定的方法名称注册，那么需要使用错误提示
				glog.Errorf(`invalid route method: %s.%s.%s defined as "%s", but "func(*ghttp.Request)" is required for object registry`,
					pkgPath, objName, mname, v.Method(i).Type().String())
			} else {
				// 否则只是Debug提示
				glog.Debugf(`ignore route method: %s.%s.%s defined as "%s", no match "func(*ghttp.Request)"`,
					pkgPath, objName, mname, v.Method(i).Type().String())
			}
			continue
		}
		key := s.mergeBuildInNameToPattern(pattern, sname, mname, true)
		m[key] = &handlerItem{
			itemName: fmt.Sprintf(`%s.%s.%s`, pkgPath, objName, mname),
			itemType: gHANDLER_TYPE_OBJECT,
			itemFunc: itemFunc,
			initFunc: initFunc,
			shutFunc: shutFunc,
		}
		// 如果方法中带有Index方法，那么额外自动增加一个路由规则匹配主URI。
		// 注意，当pattern带有内置变量时，不会自动加该路由。
		if strings.EqualFold(mname, "Index") && !gregex.IsMatchString(`\{\.\w+\}`, pattern) {
			p := gstr.PosRI(key, "/index")
			k := key[0:p] + key[p+6:]
			if len(k) == 0 || k[0] == '@' {
				k = "/" + k
			}
			m[k] = &handlerItem{
				itemName: fmt.Sprintf(`%s.%s.%s`, pkgPath, objName, mname),
				itemType: gHANDLER_TYPE_OBJECT,
				itemFunc: itemFunc,
				initFunc: initFunc,
				shutFunc: shutFunc,
			}
		}
	}
	s.bindHandlerByMap(m)
}

// 绑定对象到URI请求处理中，会自动识别方法名称，并附加到对应的URI地址后面，
// 第三个参数method仅支持一个方法注册，不支持多个，并且区分大小写。
func (s *Server) BindObjectMethod(pattern string, obj interface{}, method string) {
	m := make(handlerMap)
	v := reflect.ValueOf(obj)
	t := v.Type()
	sname := t.Elem().Name()
	mname := strings.TrimSpace(method)
	fval := v.MethodByName(mname)
	if !fval.IsValid() {
		glog.Error("invalid method name:" + mname)
		return
	}
	initFunc := (func(*Request))(nil)
	shutFunc := (func(*Request))(nil)
	if v.MethodByName("Init").IsValid() {
		initFunc = v.MethodByName("Init").Interface().(func(*Request))
	}
	if v.MethodByName("Shut").IsValid() {
		shutFunc = v.MethodByName("Shut").Interface().(func(*Request))
	}
	pkgPath := t.Elem().PkgPath()
	pkgName := gfile.Basename(pkgPath)
	objName := gstr.Replace(t.String(), fmt.Sprintf(`%s.`, pkgName), "")
	if objName[0] == '*' {
		objName = fmt.Sprintf(`(%s)`, objName)
	}
	itemFunc, ok := fval.Interface().(func(*Request))
	if !ok {
		glog.Errorf(`invalid route method: %s.%s.%s defined as "%s", but "func(*ghttp.Request)" is required for object registry`,
			pkgPath, objName, mname, fval.Type().String())
		return
	}
	key := s.mergeBuildInNameToPattern(pattern, sname, mname, false)
	m[key] = &handlerItem{
		itemName: fmt.Sprintf(`%s.%s.%s`, pkgPath, objName, mname),
		itemType: gHANDLER_TYPE_OBJECT,
		itemFunc: itemFunc,
		initFunc: initFunc,
		shutFunc: shutFunc,
	}

	s.bindHandlerByMap(m)
}

// 绑定对象到URI请求处理中，会自动识别方法名称，并附加到对应的URI地址后面,
// 需要注意对象方法的定义必须按照 ghttp.HandlerFunc 来定义
func (s *Server) BindObjectRest(pattern string, obj interface{}) {
	m := make(handlerMap)
	v := reflect.ValueOf(obj)
	t := v.Type()
	sname := t.Elem().Name()
	initFunc := (func(*Request))(nil)
	shutFunc := (func(*Request))(nil)
	if v.MethodByName("Init").IsValid() {
		initFunc = v.MethodByName("Init").Interface().(func(*Request))
	}
	if v.MethodByName("Shut").IsValid() {
		shutFunc = v.MethodByName("Shut").Interface().(func(*Request))
	}
	pkgPath := t.Elem().PkgPath()
	for i := 0; i < v.NumMethod(); i++ {
		mname := t.Method(i).Name
		method := strings.ToUpper(mname)
		if _, ok := methodsMap[method]; !ok {
			continue
		}
		pkgName := gfile.Basename(pkgPath)
		objName := gstr.Replace(t.String(), fmt.Sprintf(`%s.`, pkgName), "")
		if objName[0] == '*' {
			objName = fmt.Sprintf(`(%s)`, objName)
		}
		itemFunc, ok := v.Method(i).Interface().(func(*Request))
		if !ok {
			glog.Errorf(`invalid route method: %s.%s.%s defined as "%s", but "func(*ghttp.Request)" is required for object registry`,
				pkgPath, objName, mname, v.Method(i).Type().String())
			continue
		}
		key := s.mergeBuildInNameToPattern(mname+":"+pattern, sname, mname, false)
		m[key] = &handlerItem{
			itemName: fmt.Sprintf(`%s.%s.%s`, pkgPath, objName, mname),
			itemType: gHANDLER_TYPE_OBJECT,
			itemFunc: itemFunc,
			initFunc: initFunc,
			shutFunc: shutFunc,
		}
	}
	s.bindHandlerByMap(m)
}
