package utils

import "log"

// APILogger is responsible for the main logging of the API
var APILogger *Logger

func init() {
	APILogger = &Logger{}
}

// Logger is an abstract for logging relevant information. It can be
// a standard STDOUT or write into a file
type Logger struct {
}

// Infof records useful information for reference
func (l *Logger) Infof(format string, v ...interface{}) {
	log.Printf(format, v...)
}

// Fatal logs errors and exit the application
func (l *Logger) Fatal(v ...interface{}) {
	log.Fatal(v...)
}

// Fatalf logs errors and exit the application
func (l *Logger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
