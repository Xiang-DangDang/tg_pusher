package api

import (
	"github.com/labstack/echo/v4"
	"tg_pusher/controller/message"
)

func BindRoute(e *echo.Echo) {
	e.GET("/", func(context echo.Context) error {
		return context.JSON(200, map[string]interface{}{
			"code":    200,
			"message": "welcome tg pusher ^_^#",
		})
	})

	e.POST("/api/send", message.Send)
	e.GET("/api/send", message.Send)
}
