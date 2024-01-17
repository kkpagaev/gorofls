package web

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func ValidateRequest(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}

	if err := c.Validate(i); err != nil {
		return err
	}

	return nil
}

func conflict(c echo.Context, messsage string) error {
	return c.JSON(http.StatusConflict, map[string]string{"message": messsage})
}
