package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/snail007/gf/os/grpool"
	"github.com/snail007/gf/os/gtime"
)

func main() {
	start := gtime.Millisecond()
	wg := sync.WaitGroup{}
	for i := 0; i < 100000; i++ {
		wg.Add(1)
		grpool.Add(func() {
			time.Sleep(time.Second)
			wg.Done()
		})
	}
	wg.Wait()
	fmt.Println(grpool.Size())
	fmt.Println("time spent:", gtime.Millisecond()-start)
}
