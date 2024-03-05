package utils

import (
	"log"
	"os"
)

type Logger interface{
	Info(v ...interface{})
	Error(v ...interface{})
}

// StdLogger is a standard logger implementation using the log package.
type StdLogger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

// NewStdLogger creates a new instance of StdLogger.
func NewStdLogger() *StdLogger {
	return &StdLogger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info logs an informational message.
func (l *StdLogger) Info(v ...interface{}) {
	l.infoLogger.Println(v...)
}

// Error logs an error message.
func (l *StdLogger) Error(v ...interface{}) {
	l.errorLogger.Println(v...)
}