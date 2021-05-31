package demo

import (
	"github.com/snail007/gf/frame/gmvc"
	"github.com/snail007/gf/net/ghttp"
)

type ControllerExit struct {
	gmvc.Controller
}

func (c *ControllerExit) Init(r *ghttp.Request) {
	c.Controller.Init(r)
	c.Response.Write("exit, it will not print \"show\"")
	c.Request.Exit()
}

func (c *ControllerExit) Show() {
	c.Response.Write("show")
}

func init() {
	ghttp.GetServer().BindController("/exit", &ControllerExit{})
}
