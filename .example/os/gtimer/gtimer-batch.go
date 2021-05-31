package main

import (
	"fmt"
	"time"

	"github.com/snail007/gf/os/gtimer"
)

func main() {
	for i := 0; i < 100000; i++ {
		gtimer.Add(time.Second, func() {

		})
	}
	fmt.Println("start")
	time.Sleep(48 * time.Hour)
}
