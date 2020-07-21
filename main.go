package main

import (
	"github.com/danilovalente/project-api/config"
_ "github.com/danilovalente/project-api/usecase"
_ "github.com/danilovalente/project-api/gateway/mongodb"
	"github.com/danilovalente/project-api/controller"
	_ "github.com/danilovalente/project-api/gateway/customlog"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	controller.MapRoutes(e)

	e.Logger.Fatal(e.Start(":" + config.Values.Port))
}
