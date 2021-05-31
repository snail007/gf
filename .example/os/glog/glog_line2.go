package main

import (
	"github.com/snail007/gf/os/glog"
)

func PrintLog(content string) {
	glog.Skip(0).Line().Println("line number with skip:", content)
	glog.Line(true).Println("line number without skip:", content)
}

func main() {
	PrintLog("just test")
}
