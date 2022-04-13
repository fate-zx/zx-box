package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

var Logger = NewLog()

type Log struct {
	log *logrus.Logger
}

func NewLog() *Log {
	mLog := logrus.New()               //新建一个实例
	mLog.SetOutput(os.Stderr)          //设置输出类型
	mLog.SetReportCaller(true)         //开启返回函数名和行号
	mLog.SetFormatter(&LogFormatter{}) //设置自己定义的Formatter
	mLog.SetLevel(logrus.DebugLevel)   //设置最低的Level
	file := fmt.Sprintf("./%v/%d-%d-%d.log", "log", time.Now().Year(), time.Now().Month(), time.Now().Day())
	logFile, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
		return nil
	}
	out := io.MultiWriter(os.Stdout, logFile)
	mLog.SetOutput(out)
	return &Log{
		log: mLog,
	}
}

//封装一些会用到的方法
func (l *Log) Debug(args ...interface{}) {
	l.log.Debugln(args...)
}
func (l *Log) Debugf(format string, args ...interface{}) {
	l.log.Debugf(format, args...)
}
func (l *Log) Info(args ...interface{}) {
	l.log.Infoln(args...)
}
func (l *Log) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}
func (l *Log) Error(args ...interface{}) {
	l.log.Errorln(args...)
}
func (l *Log) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}
func (l *Log) Trace(args ...interface{}) {
	l.log.Traceln()
}
func (l *Log) Tracef(format string, args ...interface{}) {
	l.log.Tracef(format, args...)
}
func (l *Log) Panic(args ...interface{}) {
	l.log.Panicln()
}
func (l *Log) Panicf(format string, args ...interface{}) {
	l.log.Panicf(format, args...)
}

func (l *Log) Print(args ...interface{}) {
	l.log.Println()
}
func (l *Log) Printf(format string, args ...interface{}) {
	l.log.Printf(format, args...)
}
