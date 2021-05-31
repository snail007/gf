// Copyright 2017 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package ghttp

import (
	"time"

	"github.com/snail007/gf/os/glog"
)

// 设置http server参数 - CookieMaxAge
func (s *Server) SetCookieMaxAge(ttl time.Duration) {
	if s.Status() == SERVER_STATUS_RUNNING {
		glog.Error(gCHANGE_CONFIG_WHILE_RUNNING_ERROR)
		return
	}
	s.config.CookieMaxAge = ttl
}

// 设置http server参数 - CookiePath
func (s *Server) SetCookiePath(path string) {
	if s.Status() == SERVER_STATUS_RUNNING {
		glog.Error(gCHANGE_CONFIG_WHILE_RUNNING_ERROR)
		return
	}
	s.config.CookiePath = path
}

// 设置http server参数 - CookieDomain
func (s *Server) SetCookieDomain(domain string) {
	if s.Status() == SERVER_STATUS_RUNNING {
		glog.Error(gCHANGE_CONFIG_WHILE_RUNNING_ERROR)
		return
	}
	s.config.CookieDomain = domain
}

// 获取http server参数 - CookieMaxAge
func (s *Server) GetCookieMaxAge() time.Duration {
	return s.config.CookieMaxAge
}

// 获取http server参数 - CookiePath
func (s *Server) GetCookiePath() string {
	return s.config.CookiePath
}

// 获取http server参数 - CookieDomain
func (s *Server) GetCookieDomain() string {
	return s.config.CookieDomain
}
