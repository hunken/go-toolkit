package log

import (
	"github.com/sirupsen/logrus"
	"time"
)

type Logger struct {
}

func (logger Logger) Info(message string, data ...interface{}) {
	logrus.Info(message, data)
}

func (logger Logger) Debug(message string) {
	logrus.Debug(message)
}

func (logger Logger) Error(message string) {
	logrus.Error(message)
}

func (logger Logger) Fatal(message string) {
	logrus.Fatal(message)
}

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func functionTimer(time_start time.Time) {
	logrus.Info("FUNC TIME")
}
