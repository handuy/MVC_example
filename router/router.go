package router

import (
	"git.hocngay.com/techmaster-example/controller"
	"github.com/kataras/iris"
)

func DemoRoutes(c *controller.Controller, api iris.Party) {
	api.Get("/about", c.About)
	api.Post("/create", c.Create)
}
