package main

import (
	"github.com/snail007/gf/frame/g"
	"github.com/snail007/gf/frame/gmvc"
)

type Controller struct {
	gmvc.Controller
}

func (c *Controller) Test() {
	c.View.Display("layout.html")
}
func main() {
	s := g.Server()
	s.BindControllerMethod("/", new(Controller), "Test")
	s.SetPort(8199)
	s.Run()
}
