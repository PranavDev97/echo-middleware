package app

import (
	"github.com/PranavDev97/echo-middleware/delivery/api/handler"
	mw "github.com/PranavDev97/echo-middleware/delivery/api/middleware"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	// preMiddleware := []echo.MiddlewareFunc{
	// 	middleware.Logger(),
	// 	middleware.Recover(),
	// }

	eGroup := e.Group("/api", mw.SimpleMiddleware, mw.SimpleMiddleware2)
	eGroup.GET("/middle", handler.SimpleHandler)
	e.Start(":8000")
}
