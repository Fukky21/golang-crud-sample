package infrastructure

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"example.com/m/interfaces/controllers"
)

var EchoRouter Router

type Router struct {
	Echo *echo.Echo 
}

func init() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	sqlHandler := NewSqlHandler()
	sqlHandler.AutoMigration()

	userController := controllers.NewUserController(sqlHandler)
	e.GET("/users", func(c echo.Context) error { return userController.Index(c) })
	e.GET("/users/:id", func(c echo.Context) error { return userController.Show(c) })
	e.POST("/users", func(c echo.Context) error { return userController.Create(c) })
	e.PATCH("/users/:id", func(c echo.Context) error { return userController.Update(c) })
	e.DELETE("/users/:id", func(c echo.Context) error { return userController.Destroy(c) })

	EchoRouter = Router{Echo: e}
}

func (router *Router) Run() {
	e := router.Echo
	e.Logger.Fatal(e.Start(":8080"))
}