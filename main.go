package main

import (
	"fmt"
	"todo/config"
	"todo/database/mysql"
	factory "todo/faktory"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.GetConfig()
	db := mysql.InitDB(cfg)
	// db := prostgest.InitDB(cfg)
	e := echo.New()

	factory.InitFactory(e, db)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", cfg.SERVER_PORT)))
}
