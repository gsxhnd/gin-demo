package logger

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var handlerLogger = logrus.New()

func HandlerLogger() *logrus.Logger {
	if gin.Mode() == "debug" {
		handlerLogger.SetFormatter(&logrus.TextFormatter{})
	} else {
		handlerLogger.SetFormatter(&logrus.JSONFormatter{})
	}
	handlerLogger.SetReportCaller(true)
	return handlerLogger
}
