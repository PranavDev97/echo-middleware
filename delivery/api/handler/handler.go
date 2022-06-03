package handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func SimpleHandler(c echo.Context) error {
	fmt.Println("This is a simple handler")
	return nil
}
