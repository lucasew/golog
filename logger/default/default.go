package ldefault

import (
	"errors"
	"fmt"
	"github.com/lucasew/golog/handler"
	"github.com/lucasew/golog/logger"
)

// DefaultLogger A logger that just print the log event type as a char plus the payload
type DefaultLogger struct {
	handler handler.LoggerHandler
}

// NewLogger Creates a new logger from a given handler
func NewLogger(handler handler.LoggerHandler) DefaultLogger {
	return DefaultLogger{
		handler: handler,
	}
}

// NewLogger Creates a new logger from a given logger with given subnamespace
func (l DefaultLogger) NewLogger(subcomponent string) logger.Logger {
	return DefaultLogger{
		handler: l.handler.NewHandler(subcomponent),
	}
}

// processString receives the payload from the log functions and process generically
func (l DefaultLogger) processString(prefix, text string, v ...interface{}) string {
	str := fmt.Sprintf(text, v...)
	line := fmt.Sprintf("%s %s: %s", prefix, l.handler.GetComponentName(), str)
	l.handler.HandleLine(line)
	return str
}

// Panic log and panic
func (l DefaultLogger) Panic(text string, v ...interface{}) {
	panic(l.processString("P", text, v...))
}

// Error log and return the error with the given text
func (l DefaultLogger) Error(text string, v ...interface{}) error {
	return errors.New(
		l.processString("E", text, v...),
	)
}

// Warn log a warning to the console
func (l DefaultLogger) Warn(text string, v ...interface{}) {
	l.processString("W", text, v...)
}

// Info log some information to the console, it is not for every bit the program process.
func (l DefaultLogger) Info(text string, v ...interface{}) {
	l.processString("I", text, v...)
}

// Verbose log some detailed information about the program execution
func (l DefaultLogger) Verbose(level int, text string, v ...interface{}) {
	l.processString("V", text, v...)
}
