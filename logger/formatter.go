package logger

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// Formatter 日志格式化对象
type Formatter struct {
	Color bool
}

func (s *Formatter) Format(entry *logrus.Entry) ([]byte, error) {
	Level := strings.ToUpper(entry.Level.String())
	if s.Color {
		Level = colors[entry.Level](Level)
	}
	msg := fmt.Sprintf(
		"%s [%s] %s",
		entry.Time.Local().Format(`2006/01/02 15:04:05.000`),
		Level,
		entry.Message,
	)
	for field, val := range entry.Data {
		msg += fmt.Sprintf(` %s=%s`, colors[entry.Level](field), fmt.Sprint(val))
	}
	return []byte(msg + "\n"), nil
}
