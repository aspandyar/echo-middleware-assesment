package app

import (
	"github.com/aspandyar/echo-middleware-assessment/internal/app/endpoint"
	"github.com/aspandyar/echo-middleware-assessment/internal/app/mw"
	"github.com/aspandyar/echo-middleware-assessment/internal/app/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type App struct {
	endpoint *endpoint.Endpoint
	service  *service.Service
	echo     *echo.Echo
}

func New() (*App, error) {
	app := &App{}

	app.service = service.New()

	app.endpoint = endpoint.New(app.service)

	app.echo = echo.New()

	app.echo.Use(middleware.Logger())
	app.echo.Use(middleware.Recover())
	app.echo.Use(mw.RoleCheckMW)

	app.echo.GET("/days", app.endpoint.Status)

	return app, nil
}

func (app *App) Run() {
	app.echo.Logger.Fatal(app.echo.Start(":1323"))
}
