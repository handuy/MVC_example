package controller

import (
	"git.hocngay.com/techmaster-example/config"
	"git.hocngay.com/techmaster-example/model"
)

type Controller struct {
	// Configuration
	Config config.Config
	BookService        *model.BookService
}

func NewController() *Controller {
	var c Controller
	return &c
}
