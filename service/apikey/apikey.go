package apikey

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"os"
	"tg_pusher/common"
	"tg_pusher/pkg/utils"
)

type ApiKey struct {
}

func (a *ApiKey) SendKey(Command string, message *tgbotapi.Message, parameters ...string) {
	secretID := os.Getenv("SECRET_ID")
	ownerId := os.Getenv("OWNER_ID")

	if ownerId == "" || ownerId != fmt.Sprintf("%d", message.From.ID) {
		return
	}

	// 生成一个 api token chat_id  + md5(chat_idTSecretID) 秘密id
	sendKey := fmt.Sprintf("%d:%s", message.Chat.ID, utils.Md5Sum(fmt.Sprintf("%d:%s", message.Chat.ID, secretID)))
	userName := message.From.FirstName + message.From.LastName

	siteUrl := os.Getenv("SITE_URL")

	if siteUrl == "" {
		siteUrl = "https://tg-pusher.herokuapp.com"
	}

	msg := tgbotapi.NewMessage(message.From.ID,
		fmt.Sprintf(" Hi, %s , 已为你生成 sendkey \n\n 🔑 %s \n \n 🚀 使用该（测试） URL 可将消息发送到会话 [%s(%d)] 中 : \n \n %s/api/send?sendkey=%s&text=pusher_SayHi",
			userName,
			sendKey,
			message.Chat.Title,
			message.Chat.ID,
			siteUrl,
			sendKey,
		))

	common.Bot.Send(msg)
}
