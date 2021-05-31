package main

import (
	"fmt"
	"time"

	"github.com/snail007/gf/os/gtimer"
)

func main() {
	now := time.Now()
	interval := 510 * time.Millisecond
	gtimer.Add(interval, func() {
		fmt.Println(time.Now(), time.Duration(time.Now().UnixNano()-now.UnixNano()))
		now = time.Now()
	})
	time.Sleep(time.Hour)
	select {}
}
