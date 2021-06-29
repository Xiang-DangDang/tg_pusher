package bootstrap

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"tg_pusher/common"
)

func initLog() {
	var log = logrus.New()
	var httpLog = logrus.New()
	var tGLog = logrus.New()

	formatter := &logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}

	log.SetFormatter(formatter)
	tGLog.SetFormatter(formatter)
	tGLog.AddHook(NewHook(common.LogBot))
	httpLog.SetFormatter(formatter)
	httpLog.AddHook(NewHook(common.LogHttp))

	common.SetLogger(common.LogBot, tGLog)
	common.SetLogger(common.LogCommon, log)
	common.SetLogger(common.LogHttp, httpLog)
}

type defaultFieldHook struct {
	Name string
}

func NewHook(name string) *defaultFieldHook {
	return &defaultFieldHook{Name: name}
}

func (hook *defaultFieldHook) Fire(entry *logrus.Entry) error {
	entry.Message = fmt.Sprintf("<%s> %s", hook.Name, entry.Message)
	return nil
}

func (hook *defaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
