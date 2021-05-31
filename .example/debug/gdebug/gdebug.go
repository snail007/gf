package main

import (
	"fmt"
	"github.com/snail007/gf/debug/gdebug"
)

func main() {
	fmt.Println(gdebug.CallerPackage())
	fmt.Println(gdebug.CallerFunction())
}
