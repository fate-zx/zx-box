package logger

import "github.com/sirupsen/logrus"

var colors = map[logrus.Level]brush{
	logrus.PanicLevel: newBrush("1;36"),
	logrus.FatalLevel: newBrush("1;35"),
	logrus.ErrorLevel: newBrush("1;31"),
	logrus.WarnLevel:  newBrush("1;33"),
	logrus.InfoLevel:  newBrush("1;37"),
	logrus.DebugLevel: newBrush("1;32"),
	logrus.TraceLevel: newBrush("1;34"),
}

type brush func(string) string

// newBrush return a fix color Brush
func newBrush(color string) brush {
	pre := "\033["
	reset := "\033[0m"
	return func(text string) string {
		return pre + color + "m" + text + reset
	}
}
