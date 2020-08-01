package utils

import "log"

var ApiLogger *Logger

func init() {
	ApiLogger = &Logger{}
}

type Logger struct {
}

func (l *Logger) Infof(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func (l *Logger) Fatal(v ...interface{}) {
	log.Fatal(v...)
}

func (l *Logger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
