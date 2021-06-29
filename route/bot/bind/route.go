package bind

import (
	"tg_pusher/service"
	_ "tg_pusher/service"
)

type Route struct {
	Command   string
	handle    service.CommandHandler
	IsRegular bool
	IsCommand bool
}

func NewRoute(command string, handleFun service.CommandHandler, isCommand bool, isRegular bool) *Route {
	return &Route{
		Command:   command,
		handle:    handleFun,
		IsRegular: isRegular,
		IsCommand: isCommand,
	}
}
