package message

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"strconv"
	"strings"
	"tg_pusher/common"
	"tg_pusher/pkg/utils"
)

type req struct {
	SendKey   string `json:"sendkey"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode"`
}

func Send(ctx echo.Context) error {
	secretID := os.Getenv("SECRET_ID")

	var req req

	if ctx.Request().Method == http.MethodPost {
		ctx.Bind(&req)
	} else {
		req.Text = ctx.QueryParam("text")
		req.SendKey = ctx.QueryParam("sendkey")
		req.ParseMode = ctx.QueryParam("")
	}

	token := strings.Split(req.SendKey, ":")
	if req.Text == "" || req.SendKey == "" || len(token) != 2 {
		return ctx.JSON(422, map[string]interface{}{
			"code":    422,
			"message": "参数有误",
		})
	}

	if token[1] != utils.Md5Sum(fmt.Sprintf("%s:%s", token[0], secretID)) {
		return ctx.JSON(422, map[string]interface{}{
			"code":    401,
			"message": "未授权访问",
		})
	}

	req.Text = strings.ReplaceAll(req.Text, "\\n", "\n")

	// 发送消息, 后面写到 service 中
	chatId, _ := strconv.ParseInt(token[0], 10, 64)

	message := tgbotapi.NewMessage(chatId, req.Text)
	if req.ParseMode != "" {
		message.ParseMode = req.ParseMode
	}

	res, _ := common.Bot.Send(message)

	return ctx.JSON(200, map[string]interface{}{
		"code":      0,
		"message":   "ok",
		"tg_result": res,
	})
}
