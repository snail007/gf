package main

import (
	"fmt"

	"github.com/snail007/gf/text/gregex"
)

func main() {
	match, _ := gregex.MatchString(`(\w+).+\-\-\s*(.+)`, `GF is best! -- John`)
	fmt.Printf(`%s says "%s" is the one he loves!`, match[2], match[1])
}
