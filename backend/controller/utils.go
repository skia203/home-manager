package controller

import (
	"log"
	"net/http"

	"backend/model"

	"github.com/labstack/echo/v4"
)

func Connect(c echo.Context) error {
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
