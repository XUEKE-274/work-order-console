package logger

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)


type NewLogger interface {
	NewInstance(context context.Context) *logrus.Entry

}

type newLogger struct {
	logger *logrus.Logger
}


var NewLoggerReg = fx.Provide(func(logger *logrus.Logger) NewLogger {
	return &newLogger{
		logger,
	}
})

func (mine *newLogger) NewInstance(context context.Context) *logrus.Entry   {
	requestId := context.Value("requestId").(string)
	logger := mine.logger
	newLog := logger.WithFields(logrus.Fields{
		"requestId": requestId,
	})
	return newLog
}