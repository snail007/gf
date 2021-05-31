package main

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/text/gregex"
)

func main() {
	s := `-abc`
	m, err := gregex.MatchString(`^\-{1,2}a={0,1}(.*)`, s)
	g.Dump(err)
	g.Dump(m)
}
