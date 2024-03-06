package utils

import (
	"log"
	"os"
)

type Logger interface{
	Info(v ...interface{})
	Error(v ...interface{})
	Debug(v ...interface{})
}

// StdLogger is a standard logger implementation using the log package.
type StdLogger struct {
	infoLogger  *log.Logger
	errorLogger *log.Logger
	debugLogger *log.Logger
}

// NewStdLogger creates a new instance of StdLogger.
func NewStdLogger() *StdLogger {
	return &StdLogger{
		infoLogger:  log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
		debugLogger: log.New(os.Stdout, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile), // Adding debugLogger initialization
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

// Debug logs a debug message.
func (l *StdLogger) Debug(v ...interface{}) {
	l.debugLogger.Println(v...)
}