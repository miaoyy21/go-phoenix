package jsvm

import "github.com/sirupsen/logrus"

type Log struct{}

func NewLog() *Log {
	return &Log{}
}

func (l *Log) Info(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

func (l *Log) Debug(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

func (l *Log) Error(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}
