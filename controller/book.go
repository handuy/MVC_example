package controller

import (
	"log"

	"git.hocngay.com/techmaster-example/model"
	"github.com/kataras/iris"
)

func (c *Controller) About(ctx iris.Context) {
	var books []model.Book

	books = c.BookService.GetBooks()

	log.Println(books)
	ctx.ViewData("books", books)
	ctx.View("book/index.html")
}

func (c *Controller) Create(ctx iris.Context) {
	lastBookId := c.BookService.GetLastBookId()
	
	var book model.Book
	book.Id = lastBookId + 1
	book.Name = "HTML"
	book.Author = "Long"
	book.Category = "Web"

	err := c.BookService.CreateBook(&book)
	if err != nil {
		log.Println(err)
		return
	}
	
	ctx.Redirect("/about")
}
