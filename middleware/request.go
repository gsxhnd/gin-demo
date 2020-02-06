package middleware

import (
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"os"
	"time"
)

var loggerReq *zap.Logger

func ReqZapLogger(filename string) gin.HandlerFunc {
	proEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "ts",
		CallerKey:   "caller",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000000-07:00"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})

	debugEncoder := zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		TimeKey:     "ts",
		CallerKey:   "caller",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000000-07:00"))
		},
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
	})
	var core zapcore.Core
	if gin.Mode() == "debug" {
		core = zapcore.NewTee(
			zapcore.NewCore(debugEncoder, os.Stdout, zap.InfoLevel),
			zapcore.NewCore(proEncoder, zapcore.AddSync(getWriter(filename)), zap.InfoLevel),
		)
	} else {
		core = zapcore.NewCore(proEncoder, zapcore.AddSync(getWriter(filename)), zap.InfoLevel)
	}

	loggerReq = zap.New(core)
	return func(c *gin.Context) {
		start := time.Now()
		// some evil middlewares modify this values
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		end := time.Now()
		latency := end.Sub(start)

		if len(c.Errors) > 0 {
			// Append error field if this is an erroneous request.
			for _, e := range c.Errors.Errors() {
				loggerReq.Error(e)
			}
		} else {
			loggerReq.Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user_agent", c.Request.UserAgent()),
				zap.String("start_time", start.UTC().Format("2006-01-02T15:04:05.000000-07:00")),
				zap.String("end_time", end.UTC().Format("2006-01-02T15:04:05.000000-07:00")),
				zap.Duration("latency", latency),
			)
		}
	}
}

func getWriter(filename string) io.Writer {
	hook, err := rotatelogs.New(
		"./log/"+filename+"_request.%Y%m%d.log",
		//rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24),
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if err != nil {
		panic(err)
	}
	return hook
}
