package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// ResponseMiddleware struct untuk standardisasi respons
func ResponseMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Jalankan handler berikutnya
		err := next(c)
		if err != nil {
			c.Error(err)
		}

		// Jika respons belum dihasilkan
		if !c.Response().Committed {
			// Ambil status dari context, default adalah 200 OK
			status := c.Get("status")
			if status == nil {
				status = http.StatusOK
			}

			// Ambil pesan dan data dari context, atau gunakan default jika tidak disetel
			message := c.Get("message")
			if message == nil {
				message = "Request was successful"
			}

			data := c.Get("data")

			// Kirim respons success
			return c.JSON(status.(int), map[string]interface{}{
				"status":  "success",
				"message": message,
				"data":    data,
			})
		}

		return nil
	}
}
