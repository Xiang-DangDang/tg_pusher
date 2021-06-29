package bootstrap

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"tg_pusher/common"
	bot2 "tg_pusher/route/bot"
)

var bot *tgbotapi.BotAPI

func initBot() {
	var err error
	token := os.Getenv("TG_BOT_TOKEN")
	bot, err = tgbotapi.NewBotAPI(token)

	if err != nil {
		common.GetBotLog().Fatalf("初始化 Bot 失败, TOKEN [%s]。", token)
	}

	bot.Debug = false

	common.Bot = bot
}

func RunBot() {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := common.Bot.GetUpdatesChan(u)

	// 开始从 tg 去拿消息
	for update := range updates {
		if update.Message == nil {
			continue
		} else {
			go func() {
				common.GetBotLog().Printf("(USERID:%d|CHAT_ID:%d)[%s -- %s] %s", update.Message.From.ID, update.Message.Chat.ID, update.Message.Chat.Title, update.Message.From.String(), update.Message.Text)
				// 创建并执行路由内容
				bot2.NewRoute(bot).Exec(update.Message)
			}()
		}
	}
}
