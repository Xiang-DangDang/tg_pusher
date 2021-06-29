# TG_PUSHER

## 概述

通过 TG 机器人向 会话 或者 群组发送消息，封装了 TG API，更简洁，部署在 heroku 或 google 应用引擎 以解决 国内无法直连 TG API 的问题。

#### API

* /send

  * 请求方式 `POST 或者 GET`

  * 参数

    **POST 参数必须为 JSON 格式**

    ```json
    {
      "sendkey":"638782213:425a5df9f6e94ea9c71c074cbd0f9b", // 消息发送key
      "text":"hello server 酱", //消息内容
      "parse_mode":"" //可选 HTML 或 MarkdownV2 默认为普通文本，决定了 $text 的解析方式
    }
    ```

    `parse_mode` 参考 [Telegram Bot API](https://core.telegram.org/bots/api#formatting-options)

* 环境变量

  * **BOT_API_TOKEN**  (必须)

    机器人 API TOKEN

  * **SECRET_ID ** (必须)

    保证接口 KEY 安全的随机字符串

  * **OWNER_ID** (必须)

    机器人所有者 TG ID，只有此用户才可以获取 `sendkey`
    不知道可以先随便写一个，应用启动后 通过 /myid 命令获取并重新配置

  * **SITE_URL**

    API 访问域名 如: `https://tg-pusher.herokuapp.com`

* 机器人可用命令

  * `/sendkey`

    生成当前会话的 `sendkey` 参数，在群聊或者单聊中发送此命令，用于向此会话发送消息

  * `/myid`

    获取自己的 Tg ID

  * `/chatid`

    获取当前会话的 ID

## 部署

[![Deploy](https://www.herokucdn.com/deploy/button.png)](https://dashboard.heroku.com/new?template=https://github.com/Xiang-DangDang/tg_pusher)