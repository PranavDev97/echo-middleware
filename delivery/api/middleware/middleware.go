package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func SimpleMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("this is a simple middleware")
		return next(c)
	}
}

func SimpleMiddleware2(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("this is a simple middleware 2")
		return next(c)
	}
}
