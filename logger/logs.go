package logger

import (
	"github.com/sirupsen/logrus"
	"time"
)

var Logs *logrus.Logger

func init() {
	Logs = NewLogger(&Config{
		Debug:        true,
		Name:         "./log/run.log",
		Level:        5,
		Caller:       true,
		Color:        true,
		MaxAge:       10 * time.Second,
		RotationTime: 10 * time.Second,
	})
}
