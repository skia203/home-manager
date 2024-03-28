package main

import (
	"backend/controller"

	"github.com/labstack/echo/v4"
)

func newRouter() *echo.Echo {
	e := echo.New()
	e.GET("/", controller.Connect)
	e.GET("/users", controller.GetUsers)
	e.POST("/users", controller.CreateUser)
	e.GET("/users/:id", controller.GetUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)
	return e
}
