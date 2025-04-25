package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ResponseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err != nil {
			c.Error(err)
		}

		if !c.Response().Committed {
			status := c.Get("status")
			if status == nil {
				status = http.StatusOK
			}

			message := c.Get("message")
			if message == nil {
				message = "Request was successful"
			}

			data := c.Get("data")

			return c.JSON(status.(int), map[string]interface{}{
				"status":  "success",
				"message": message,
				"data":    data,
			})
		}

		return nil
	}
}
