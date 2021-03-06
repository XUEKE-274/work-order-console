package logger
import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/sirupsen/logrus"
	"go.uber.org/fx"
)
var Model = fx.Options(
	sourceLogger,
	NewLoggerReg,
)



var sourceLogger = fx.Provide(func() *logrus.Logger {
	//now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = dir + "/logs/"
	}
	if err := os.MkdirAll(logFilePath, 0777); err != nil {
		fmt.Println(err.Error())
	}
	logFileName := "work-order-console" + ".log" // now.Format("2006_01_02_15") +
	fileName := path.Join(logFilePath, logFileName)
	src, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err", err)
	}
	logger := logrus.New()
	logger.Out = src
	logger.SetLevel(logrus.InfoLevel)
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	// add 多端输出
	writes := []io.Writer{
		src,
		os.Stdout,
	}
	fileAndStdoutWriter := io.MultiWriter(writes...)
	logger.SetOutput(fileAndStdoutWriter)
	return logger
})