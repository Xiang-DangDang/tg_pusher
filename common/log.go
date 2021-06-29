package common

import "github.com/sirupsen/logrus"

var logs = map[string]*logrus.Logger{}

const (
	LogCommon = "common"
	LogBot    = "bot"
	LogHttp   = "http"
)

func SetLogger(name string, logger *logrus.Logger) {
	if logger != nil {
		logs[name] = logger
	}
}

func GetLogger(name string) *logrus.Logger {
	if logger, ok := logs[name]; ok {
		return logger
	}

	return nil
}

func GetBotLog() *logrus.Logger {
	return GetLogger(LogBot)
}

func GetHttpLog() *logrus.Logger {
	return GetLogger(LogHttp)
}

func GetLog() *logrus.Logger {
	return GetLogger(LogCommon)
}
