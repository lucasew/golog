package golog_logger_filter

import (
    "github.com/lucasew/golog/logger"
    "fmt"
)

type LogLevel int
const (
    LvlPanic LogLevel = iota;
    LvlError
    LvlWarn
    LvlInfo
)
func LvlVerbose(level int) LogLevel {
    return LogLevel(int(LvlInfo) + 1 + level)
}

// Passa pra frente somente niveis de log que são iguais ou mais importantes que filter
type FilterLogger struct {
    logger golog_logger.Logger
    filter LogLevel
}

func NewLogger(logger golog_logger.Logger, filter LogLevel) FilterLogger {
    return FilterLogger{
        logger: logger,
        filter: filter,
    }
}

func (l FilterLogger) NewLogger(submodule string) golog_logger.Logger {
    return FilterLogger{
        logger: l.logger.NewLogger(submodule),
        filter: l.filter,
    }
}

func (l FilterLogger) checkLog(level LogLevel) bool {
    return level <= l.filter
}

func (l FilterLogger) Panic(str string, v ...interface{}) {
    if (l.checkLog(LvlPanic)) {
        l.logger.Panic(str, v...)
    }
}

func (l FilterLogger) Error(str string, v ...interface{}) error {
    if (l.checkLog(LvlError)) {
        return l.logger.Error(str, v...)
    }
    return fmt.Errorf(str, v...)
}

func (l FilterLogger) Warn(str string, v ...interface{}) {
    if (l.checkLog(LvlWarn)) {
        l.logger.Warn(str, v...)
    }
}

func (l FilterLogger) Info(str string, v ...interface{}) {
    if (l.checkLog(LvlInfo)) {
        l.logger.Info(str, v...)
    }
}

func (l FilterLogger) Verbose(level int, str string, v ...interface{}) {
    if (l.checkLog(LvlVerbose(level))) {
        l.logger.Verbose(level, str, v...)
    }
}
