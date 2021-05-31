package main

import (
	"fmt"
	"os"
	"time"

	"github.com/snail007/gf/os/gfpool"
)

func main() {
	for {
		time.Sleep(time.Second)
		if f, err := gfpool.Open("/home/john/temp/log.log", os.O_RDONLY, 0666, 60000000*1000); err == nil {
			fmt.Println(f.Name())
			f.Close()
		} else {
			fmt.Println(err)
		}
	}
}
