package users

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func Handler(c *echo.Echo) {
	c.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
}
