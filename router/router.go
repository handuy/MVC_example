package router

import (
	"git.hocngay.com/techmaster-example/controller"
	"github.com/kataras/iris"
)

func DemoRoutes(c *controller.Controller, api iris.Party) {
	api.Get("/", c.About)
	api.Post("/create", c.Create)
	api.Post("/update", c.Update)
	api.Post("/delete", c.Delete)

	api.Post("/add-student", c.AddStudent)
}
