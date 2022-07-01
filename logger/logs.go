package logger

import (
	"github.com/sirupsen/logrus"
)

var Logs *logrus.Logger

func init() {
	Logs = NewLogger(&Config{
		Debug: true,
		//Name:         "./log/run.log",
		Level:        5,
		Caller:       true,
		Color:        true,
		MaxAge:       0,
		RotationTime: 0,
	})
}
