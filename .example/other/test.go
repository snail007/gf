package main

import (
	"fmt"
	"github.com/snail007/gf/text/gregex"
)

func main() {
	file := "xxx/github.com/hg-hh/ww/gf/.example/"
	fmt.Println(gregex.IsMatchString(`/github.com/[^/]+/gf/\.example/`, file))
}
