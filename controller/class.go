package controller

import (
	"log"

	"git.hocngay.com/techmaster-example/model"
	"github.com/kataras/iris"
)

func (c *Controller) About(ctx iris.Context) {
	var class model.Class

	classes, err := class.GetClasses(c.DB)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.ViewData("classes", classes)
	ctx.View("book/index.html")
}

func (c *Controller) Create(ctx iris.Context) {
	var class model.Class
	class.Name = "HTML"
	class.CourseId = "1"
	class.Type = 1

	err := class.InsertClass(c.DB)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.Redirect("/")
}

func (c *Controller) Update(ctx iris.Context) {
	var class model.Class
	lastClass, err := class.GetLastClass(c.DB)
	if err != nil {
		log.Println(err)
		return
	}
	lastClass.Name = "CSS"

	err = lastClass.UpdateClass(c.DB)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.Redirect("/")
}

func (c *Controller) Delete(ctx iris.Context) {
	var class model.Class
	lastClass, err := class.GetLastClass(c.DB)
	if err != nil {
		log.Println(err)
		return
	}

	err = lastClass.DeleteClass(c.DB)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.Redirect("/")
}
