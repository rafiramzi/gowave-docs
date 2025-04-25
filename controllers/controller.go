package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

func (pc *Controller) GetHello(c echo.Context) error {

	var hello_world = "Hello World!"

	return c.String(http.StatusOK, hello_world)
}