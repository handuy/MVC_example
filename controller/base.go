package controller

import (
	"git.hocngay.com/techmaster-example/config"
	"github.com/go-pg/pg"
)

type Controller struct {
	// Configuration
	Config config.Config
	// DB instance
	DB *pg.DB
}

func NewController() *Controller {
	var c Controller
	return &c
}
