package bootstrap

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"tg_pusher/common"
	"tg_pusher/route/api"
)

var echoServer *echo.Echo

func InitHttpServer() {
	echoServer = echo.New()
	common.Echo = echoServer
}

func RunHttpServer() {

	api.BindRoute(echoServer)

	listenPort := os.Getenv("PORT")

	if listenPort == "" {
		listenPort = "3000"
	}

	if err := echoServer.Start(":" + listenPort); err != nil && err != http.ErrServerClosed {
		echoServer.Logger.Fatal("shutting down the server:" + err.Error())
	}
}
