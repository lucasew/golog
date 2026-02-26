package console

import (
	"errors"
	"fmt"

	"github.com/lucasew/golog/handler"
	gologErrors "github.com/lucasew/golog/internal/errors"
	"github.com/lucasew/golog/logger"
)

// Logger implements logger.Logger using a simple console output.
type Logger struct {
	handler handler.LoggerHandler
}

// NewLogger creates a new logger from a given handler.
func NewLogger(handler handler.LoggerHandler) logger.Logger {
	return Logger{
		handler: handler,
	}
}

// NewLogger creates a new logger from a given logger with given subnamespace.
func (l Logger) NewLogger(subcomponent string) logger.Logger {
	return Logger{
		handler: l.handler.NewHandler(subcomponent),
	}
}

// processString receives the payload from the log functions and process generically.
func (l Logger) processString(prefix, text string, v ...interface{}) string {
	str := fmt.Sprintf(text, v...)
	line := fmt.Sprintf("%s %s: %s", prefix, l.handler.GetComponentName(), str)
	err := l.handler.HandleLine(line)
	if err != nil {
		gologErrors.ReportError(err)
	}
	return str
}

// Panic log and panic.
func (l Logger) Panic(text string, v ...interface{}) {
	panic(l.processString("P", text, v...))
}

// Error log and return the error with the given text.
func (l Logger) Error(text string, v ...interface{}) error {
	return errors.New(
		l.processString("E", text, v...),
	)
}

// Warn log a warning to the console.
func (l Logger) Warn(text string, v ...interface{}) {
	l.processString("W", text, v...)
}

// Info log some information to the console.
func (l Logger) Info(text string, v ...interface{}) {
	l.processString("I", text, v...)
}

// Verbose log some detailed information about the program execution.
func (l Logger) Verbose(level int, text string, v ...interface{}) {
	l.processString("V", text, v...)
}
