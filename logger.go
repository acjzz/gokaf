package gokaf

import (
	"context"
	"github.com/sirupsen/logrus"
	"os"
)

//NewLogger returns an instance of Logger
func NewLogger(ctx context.Context) *logrus.Entry {
	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(getLogLevelKey(ctx))
	return logrus.WithFields(getLogFields(ctx))
}
