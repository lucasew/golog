package lfilter

import (
	"fmt"
	"github.com/lucasew/golog/logger"
)

// LogLevel represents a tier of logging. The less this value is the more important is the log
type LogLevel int

const (
	// LvlPanic or more important
	LvlPanic LogLevel = iota
	// LvlError Error or more important
	LvlError
	// LvlWarn Warn or more important
	LvlWarn
	// LvlInfo Info or more important
	LvlInfo
)

// LvlVerbose VerboseN or more important
func LvlVerbose(level int) LogLevel {
	return LogLevel(int(LvlInfo) + 1 + level)
}

// FilterLogger Forward only calls of log levels that are the same or more imporant than filter
type FilterLogger struct {
	logger logger.Logger
	filter LogLevel
}

// NewLogger creates a new object to filter calls of another logger
func NewLogger(logger logger.Logger, filter LogLevel) FilterLogger {
	return FilterLogger{
		logger: logger,
		filter: filter,
	}
}

// NewLogger creates a new object now pointing to a logging subnamespace, all the characteristics are inherited
func (l FilterLogger) NewLogger(submodule string) logger.Logger {
	return FilterLogger{
		logger: l.logger.NewLogger(submodule),
		filter: l.filter,
	}
}

// checkLog Do I need to forward this log call?
func (l FilterLogger) checkLog(level LogLevel) bool {
	return level <= l.filter
}

// Panic same meaning as in other structs that implement this
func (l FilterLogger) Panic(str string, v ...interface{}) {
	if l.checkLog(LvlPanic) {
		l.logger.Panic(str, v...)
	}
}

// Error same meaning as in other structs that implement this. The difference is that I return the error even if I dont show it on the screen
func (l FilterLogger) Error(str string, v ...interface{}) error {
	if l.checkLog(LvlError) {
		return l.logger.Error(str, v...)
	}
	return fmt.Errorf(str, v...)
}

// Warn same meaning as in other structs that implement this
func (l FilterLogger) Warn(str string, v ...interface{}) {
	if l.checkLog(LvlWarn) {
		l.logger.Warn(str, v...)
	}
}

// Info same meaning as in other structs that implement this
func (l FilterLogger) Info(str string, v ...interface{}) {
	if l.checkLog(LvlInfo) {
		l.logger.Info(str, v...)
	}
}

// Verbose same meaning as in other structs that implement this
func (l FilterLogger) Verbose(level int, str string, v ...interface{}) {
	if l.checkLog(LvlVerbose(level)) {
		l.logger.Verbose(level, str, v...)
	}
}
