package main

import (
	"log"
	"net/http"
	
	"github.com/labstack/echo/v4"
	"backend/model"
	"backend/controller"
)

func connect(c echo.Context) error {
	db, err := model.DB.DB()
	if err != nil {
        log.Printf("Failed to get database connection: %v", err)
        return c.String(http.StatusInternalServerError, "Failed to connect to database")
    }
	err = db.Ping()
	if err != nil {
        log.Printf("Failed to ping database: %v", err)
        return c.String(http.StatusInternalServerError, "Failed to connect to database")
    }
    return c.String(http.StatusOK, "Connected to database")
}



func main() {
	e := echo.New()
	e.GET("/", connect)
	e.GET("/users", controller.GetUsers)
	e.GET("/users/:id", controller.GetUser)
	e.POST("/users", controller.CreateUser)
	e.PUT("/users/:id", controller.UpdateUser)
	e.DELETE("/users/:id", controller.DeleteUser)
	e.Logger.Fatal(e.Start(":8080"))
}

