package main

import (
	"github.com/PranavDev97/echo-middleware/app"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	app.Routes(e)

}
