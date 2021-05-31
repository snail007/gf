// Copyright 2017 gf Author(https://github.com/snail007/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/snail007/gf.

package gtcp

import "time"

// 简单协议: (面向短链接)发送消息包
func SendPkg(addr string, data []byte, option ...PkgOption) error {
	conn, err := NewConn(addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.SendPkg(data, option...)
}

// 简单协议: (面向短链接)发送数据并等待接收返回数据
func SendRecvPkg(addr string, data []byte, option ...PkgOption) ([]byte, error) {
	conn, err := NewConn(addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.SendRecvPkg(data, option...)
}

// 简单协议: (面向短链接)带超时时间的数据发送
func SendPkgWithTimeout(addr string, data []byte, timeout time.Duration, option ...PkgOption) error {
	conn, err := NewConn(addr)
	if err != nil {
		return err
	}
	defer conn.Close()
	return conn.SendPkgWithTimeout(data, timeout, option...)
}

// 简单协议: (面向短链接)发送数据并等待接收返回数据(带返回超时等待时间)
func SendRecvPkgWithTimeout(addr string, data []byte, timeout time.Duration, option ...PkgOption) ([]byte, error) {
	conn, err := NewConn(addr)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	return conn.SendRecvPkgWithTimeout(data, timeout, option...)
}
