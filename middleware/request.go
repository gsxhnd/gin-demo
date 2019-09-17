package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

var requestLogger = logrus.New()

// RequestLogger print log data when request in
func RequestLogger() gin.HandlerFunc {

	if gin.Mode() == "debug" {
		requestLogger.SetFormatter(&logrus.TextFormatter{})
	} else {
		requestLogger.SetFormatter(&logrus.JSONFormatter{})
	}

	return func(c *gin.Context) {
		startTime := time.Now()
		// 处理请求
		c.Next()

		// 结束时间
		endTime := time.Now()

		// 执行时间
		// latencyTime / 1000 = ms
		latencyTime := endTime.Sub(startTime)

		// 请求方式
		reqMethod := c.Request.Method

		// 请求路由
		reqUri := c.Request.RequestURI

		// 请求数据
		//reqData,_ := ioutil.ReadAll(c.Request.Body)
		//reqBody := string(reqData)

		// 状态码
		statusCode := c.Writer.Status()

		// 请求IP
		clientIP := c.ClientIP()

		requestLogger.WithFields(
			logrus.Fields{
				"status_code":  statusCode,
				"latency_time": latencyTime,
				"client_ip":    clientIP,
				"req_method":   reqMethod,
				"req_uri":      reqUri,
			}).Info()
	}
}
