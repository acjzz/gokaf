package gokaf

import "fmt"

type Logger interface {
	Printf(format string, v ...interface{})
}

type SimpleLogger struct{}

func (l *SimpleLogger) Printf(format string, v ...interface{}) {
	fmt.Printf(format+"\n", v...)
}
