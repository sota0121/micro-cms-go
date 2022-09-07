package mcmslib

import (
	"net/http"

	"github.com/labstack/echo"
)

func RootHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hi, this is micro-cms-go!")
}
