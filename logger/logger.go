package logger

import (
	"io"
	"os"
	"time"

	rotate "github.com/lestrrat-go/file-rotatelogs"
	"github.com/shiena/ansicolor"
	"github.com/sirupsen/logrus"
)

// getWriter
func getWriter(debug bool, logName string, maxAge, rotationTime time.Duration) []io.Writer {
	if w, err := rotate.New(
		logName+".%Y%m%d",
		rotate.WithLinkName(logName),          // 生成软链，指向最新日志文件
		rotate.WithMaxAge(maxAge),             // 文件最大保存时间
		rotate.WithRotationTime(rotationTime), // 日志切割时间间隔
	); err != nil {
		panic(err)
	} else if debug {
		return []io.Writer{w, ansicolor.NewAnsiColorWriter(os.Stdout)}
	} else {
		return []io.Writer{w}
	}
}

// NewLogger
// @title NewLogger
// @description 创建新的日志对象
// @param conf *Config 日志配置信息
// @return logger *logrus.Logger 返回日志对象
func NewLogger(conf *Config) *logrus.Logger {
	logger := logrus.New()
	output := io.MultiWriter(getWriter(conf.Debug, conf.Name, conf.MaxAge, conf.RotationTime)...)
	logger.SetOutput(output)
	logger.AddHook(&Hook{})
	logger.SetLevel(logrus.Level(conf.Level))
	logger.SetReportCaller(conf.Caller)
	logger.SetFormatter(&Formatter{Color: conf.Color})
	logger.ExitFunc = func(i int) {}
	return logger
}
