package logger

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
)

// Hook 日志写入钩子函数
type Hook struct{}

func (Hook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (Hook) Fire(entry *logrus.Entry) error {
	if entry.HasCaller() {
		entry.Message = fmt.Sprintf(`[%s] %s`, findCaller(6), entry.Message)
	}
	return nil
}

// findCaller
func findCaller(skip int) string {
	file, line := "", 0
	for i := 0; i < 10; i++ {
		file, line = getCaller(skip + i)
		if !strings.HasPrefix(file, `logrus`) {
			break
		}
	}
	return fmt.Sprintf("%s:%d", file, line)
}

// getCaller
func getCaller(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip)
	if !ok {
		return "", 0
	}
	return path.Join(path.Base(path.Dir(file)), path.Base(file)), line
}
