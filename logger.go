package gokaf

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

type LogWrapper interface {
	Trace(args ...interface{})
	Tracef(format string, args ...interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
}

//NewLogrusLogger returns an instance of Logger
func NewLogrusLogger(ctx context.Context) LogWrapper {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(getLogLevelKeyFromCtx(ctx))
	return logrus.WithFields(getLogFields(ctx))
}
