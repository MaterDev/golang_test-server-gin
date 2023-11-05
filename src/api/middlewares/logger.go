package middlewares

import (
	"io"
	"log"
	"os"
)


type LoggerStruct struct {
	info    *log.Logger
	warning *log.Logger
	error   *log.Logger
}

// New creates a new LoggerStruct instance.
func New(infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) *LoggerStruct {
	// Create custom loggers for Info, Warning, and Error
	return &LoggerStruct{
		info:    log.New(infoHandle, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile),
		warning: log.New(warningHandle, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile),
		error:   log.New(errorHandle, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

// Info logs an informational message.
func (l *LoggerStruct) Info(v ...interface{}) {
	l.info.Println(v...)
}

// Warning logs a warning message.
func (l *LoggerStruct) Warning(v ...interface{}) {
	l.warning.Println(v...)
}

// Error logs an error message.
func (l *LoggerStruct) Error(v ...interface{}) {
	l.error.Println(v...)
}

// Initialize the default logger with the common setup.
var defaultLogger = New(os.Stdout, os.Stdout, os.Stderr)

/*
Logger is a wrapper around the defaultLogger instance. It provides convenience methods for logging messages at different levels.
*/
var Logger = defaultLogger

// Info is a convenience method for defaultLogger.Info
func Info(v ...interface{}) {
	defaultLogger.Info(v...)
}

// Warning is a convenience method for defaultLogger.Warning
func Warning(v ...interface{}) {
	defaultLogger.Warning(v...)
}

// Error is a convenience method for defaultLogger.Error
func Error(v ...interface{}) {
	defaultLogger.Error(v...)
}
