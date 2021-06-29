package bootstrap

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"tg_pusher/common"
	"time"
)

func initLog() {
	var log = logrus.New()
	var httpLog = logrus.New()
	var tGLog = logrus.New()

	formatter := logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	}

	f := &customFormat{formatter}

	log.SetFormatter(f)
	tGLog.SetFormatter(f)
	tGLog.AddHook(NewHook(common.LogBot))
	httpLog.SetFormatter(f)
	httpLog.AddHook(NewHook(common.LogHttp))

	common.SetLogger(common.LogBot, tGLog)
	common.SetLogger(common.LogCommon, log)
	common.SetLogger(common.LogHttp, httpLog)
}

type customFormat struct {
	logrus.TextFormatter
}

func (c *customFormat) Format(entry *logrus.Entry) ([]byte, error) {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	entry.Time = entry.Time.In(loc)
	return c.TextFormatter.Format(entry)
}

type defaultFieldHook struct {
	Name string
}

func NewHook(name string) *defaultFieldHook {
	return &defaultFieldHook{Name: name}
}

func (hook *defaultFieldHook) Fire(entry *logrus.Entry) error {
	//loc, _ := time.LoadLocation("America/New_York")
	entry.Time.AddDate(1, 0, 3)
	entry.Message = fmt.Sprintf("<%s> %s", hook.Name, entry.Message)
	return nil
}

func (hook *defaultFieldHook) Levels() []logrus.Level {
	return logrus.AllLevels
}
