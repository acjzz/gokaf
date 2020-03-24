package gokaf

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

type logWrapper interface {
	Tracef(format string, args ...interface{})
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
}

//NewLogger returns an instance of Logger
func NewLogger(ctx context.Context) logWrapper {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(getLogLevelKeyFromCtx(ctx))
	return logrus.WithFields(getLogFields(ctx))
}
