package main

import (
	"encoding/json"

	"github.com/snail007/gf/.example/net/gtcp/pkg_operations/monitor/types"
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/net/gtcp"
	"github.com/snail007/gf/os/glog"
	"github.com/snail007/gf/os/gtime"
)

func main() {
	// 数据上报客户端
	conn, err := gtcp.NewConn("127.0.0.1:8999")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	// 使用JSON格式化数据字段
	info, err := json.Marshal(types.NodeInfo{
		Cpu:  float32(66.66),
		Host: "localhost",
		Ip: g.Map{
			"etho": "192.168.1.100",
			"eth1": "114.114.10.11",
		},
		MemUsed:  15560320,
		MemTotal: 16333788,
		Time:     int(gtime.Second()),
	})
	if err != nil {
		panic(err)
	}
	// 使用 SendRecvPkg 发送消息包并接受返回
	if result, err := conn.SendRecvPkg(info); err != nil {
		if err.Error() == "EOF" {
			glog.Println("server closed")
		}
	} else {
		glog.Println(string(result))
	}
}
