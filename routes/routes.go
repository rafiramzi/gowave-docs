package routes

import (
	"golang_projects/controllers"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	controller := controllers.Controller{DB: db}

	r := e.Group("/")
		r.GET("", controller.GetHello)
}