package controller

import (
	"log"

	model_package "git.hocngay.com/model-package"
	"github.com/kataras/iris"
)

func (c *Controller) About(ctx iris.Context) {
	var class model_package.Class

	classes, err := class.GetClasses(c.DB)
	if err != nil {
		log.Println(err)
		return
	}

	ctx.ViewData("classes", classes)
	ctx.View("book/index.html")
}

func (c *Controller) Create(ctx iris.Context) {
	var class model_package.Class
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
	var class model_package.Class
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
	var class model_package.Class
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

func (c *Controller) AddStudent(ctx iris.Context) {
	tx, err := c.DB.Begin()
	if err != nil {
		tx.Rollback()
		return
	}

	// Tạo tài khoản sinh viên
	var student model_package.Student
	student.Name = "HTML"
	err = student.InsertStudent(tx)
	if err != nil {
		log.Println(err)
		tx.Rollback()
		return
	}

	// Lấy thông tin lớp
	var class model_package.Class
	lastClass, err := class.GetLastClass(tx)
	if err != nil {
		tx.Rollback()
		return
	}

	// Cập nhật số học viên trong lớp
	lastClass.Students++
	err = lastClass.UpdateClass(tx)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()
	ctx.Redirect("/")
}