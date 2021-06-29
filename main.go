package main

import (
	"context"
	"encoding/json"
	"os"
	"os/signal"
	"runtime"
	"tg_pusher/bootstrap"
	"tg_pusher/common"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU()) //协程并发跑在多核cpu
}

func main() {
	bootstrap.Init()
	log := common.GetLog()
	go bootstrap.RunBot()
	go bootstrap.RunHttpServer()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// delete bot api
	if resp, err := common.Bot.MakeRequest("deleteWebhook", nil); err != nil {
		log.Errorf("清除 bot webhook 失败 -> %s", err)
	} else {
		res, _ := json.Marshal(resp)
		log.Infof("清除 bot webhook -> %s", res)
	}

	log.Info("关闭 api server")
	if err := common.Echo.Shutdown(ctx); err != nil {
		common.Echo.Logger.Fatal(err)
	}

	log.Info("Bye Bye !!!")
}
