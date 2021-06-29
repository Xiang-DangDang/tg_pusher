package service

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

// CommandHandler 通用的 service handler method 签名
type CommandHandler func(Command string, message *tgbotapi.Message, parameters ...string)
