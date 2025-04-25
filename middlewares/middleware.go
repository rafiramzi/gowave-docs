package middlewares

import (
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

// JWTMiddleware returns middleware function for JWT authentication
func JWTMiddleware() echo.MiddlewareFunc {
	// Gunakan kunci rahasia yang sama dengan yang digunakan saat membuat token
	secretKey := []byte("your_secret_key")

	config := echojwt.Config{
		SigningKey: secretKey, // Ganti dengan secret key Anda
		ErrorHandler: func(c echo.Context, err error) error {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "Token tidak valid atau tidak ditemukan",
			})
		},
	}

	return echojwt.WithConfig(config)
}
