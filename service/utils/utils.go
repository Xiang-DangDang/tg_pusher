package utils

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"tg_pusher/common"
)

type Utils struct {
}

func (u Utils) MyId(Command string, message *tgbotapi.Message, parameters ...string) {
	if message.Chat.Title == "" {
		return
	}
	username := message.From.FirstName + message.From.LastName
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Hi, %s \n\n你的 Telegram ID 为 %d", username, message.From.ID))

	msg.ReplyToMessageID = message.MessageID
	common.Bot.Send(msg)
}

func (u Utils) ChatId(Command string, message *tgbotapi.Message, parameters ...string) {
	if message.Chat.Title == "" {
		return
	}

	username := message.From.FirstName + message.From.LastName
	msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Hi, %s \n\n 群组「%s」的会话 ID 为 %d", username, message.Chat.Title, message.Chat.ID))

	msg.ReplyToMessageID = message.MessageID
	common.Bot.Send(msg)
}
