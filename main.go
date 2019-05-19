package main

import (
	"os"

	"git.hocngay.com/techmaster-example/config"
	"git.hocngay.com/techmaster-example/controller"
	"git.hocngay.com/techmaster-example/model"
	"git.hocngay.com/techmaster-example/router"
	"github.com/go-pg/pg"
	"github.com/kataras/iris"

	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

func main() {

	config := config.SetupConfig()

	// Khởi tạo controller
	c := controller.NewController()
	c.Config = config

	// Kết nối CSDL
	dbConfig := config.Database
	db := model.ConnectDb(dbConfig.User, dbConfig.Password, dbConfig.Database, dbConfig.Address)
	defer db.Close()
	c.DB = db
	setupDatabase(db, config)

	//Khởi tạo app iris
	app := iris.New()
	app.Logger().SetLevel("debug")

	// Logging in terminal
	app.Use(recover.New())
	app.Use(logger.New())

	// // Đăng ký thư mục chứa HTML
	tmpl := iris.HTML("./view", ".html").Reload(true)

	app.RegisterView(tmpl)

	router.DemoRoutes(c, app)

	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
}

func setupDatabase(db *pg.DB, config config.Config) {

	argsWithProg := os.Args
	if len(argsWithProg) > 1 && os.Args[1] == "release" {
	} else {
		model.LogQueryToConsole(db)
	}

	err := model.MigrationDb(db, config)
	if err != nil {
		panic(err)
	}
}
