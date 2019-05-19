package controller

import (
	"git.hocngay.com/techmaster-example/config"
	"github.com/go-pg/pg"
)

type Controller struct {
	// DB instance
	DB *pg.DB

	// Configuration
	Config config.Config
}

func NewController() *Controller {
	var c Controller
	return &c
}
