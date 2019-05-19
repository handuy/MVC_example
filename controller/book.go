package controller

import (
	"fmt"

	"git.hocngay.com/techmaster-example/model"
	"github.com/kataras/iris"
)

func (c *Controller) About(ctx iris.Context) {
	var books []model.Book
	err := c.DB.Model(&books).Select()
	if err != nil {
		return
	}
	fmt.Println(books)
	ctx.ViewData("books", books)
	ctx.View("book/index.html")
}
