package main

import (
	"fmt"
	"time"

	"github.com/snail007/gf/net/gtcp"
	"github.com/snail007/gf/os/glog"
	"github.com/snail007/gf/os/gtime"
)

func main() {
	// Server
	go gtcp.NewServer("127.0.0.1:8999", func(conn *gtcp.Conn) {
		defer conn.Close()
		for {
			data, err := conn.Recv(-1)
			if len(data) > 0 {
				if err := conn.Send(append([]byte("> "), data...)); err != nil {
					fmt.Println(err)
				}
			}
			if err != nil {
				break
			}
			return
		}
	}).Run()

	time.Sleep(time.Second)

	// Client
	for {
		if conn, err := gtcp.NewPoolConn("127.0.0.1:8999"); err == nil {
			if b, err := conn.SendRecv([]byte(gtime.Datetime()), -1); err == nil {
				fmt.Println(string(b), conn.LocalAddr(), conn.RemoteAddr())
			} else {
				fmt.Println(err)
			}
			conn.Close()
		} else {
			glog.Error(err)
		}
		time.Sleep(time.Second)
	}
}
