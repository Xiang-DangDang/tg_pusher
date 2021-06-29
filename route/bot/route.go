package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	bind2 "tg_pusher/route/bot/bind"
	apikey3 "tg_pusher/service/apikey"
	"tg_pusher/service/utils"
)

var botApi *tgbotapi.BotAPI

func NewRoute(api *tgbotapi.BotAPI) *bind2.Map {
	botApi = api
	return bind2.NewRouteMap(routes)
}

func routes(routeMap *bind2.Map) {
	apikey := apikey3.ApiKey{}
	utils := utils.Utils{}
	//
	routeMap.AddCommandRoute("sendkey", apikey.SendKey)
	routeMap.AddCommandRoute("myid", utils.MyId)
	routeMap.AddCommandRoute("chatid", utils.ChatId)
	//routeMap.AddRegularExpression(`^#(\d+)`, topicService.ReplyToTopic)
	//routeMap.AddCommandRoute("reply", topicService.ReplyToTopic)
	//routeMap.AddCommandRoute("show", topicService.ShowTopic)
}
