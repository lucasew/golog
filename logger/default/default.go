package golog_logger_default

import (
    "github.com/lucasew/golog/handler"
    "fmt"
    "errors"
)

type DefaultLogger struct {
    handler golog_handler.LoggerHandler;
}

func NewLogger(handler golog_handler.LoggerHandler) DefaultLogger {
    return DefaultLogger{
        handler: handler,
    }
}

func (l DefaultLogger) NewLogger(subcomponent string) DefaultLogger {
    return DefaultLogger{
        handler: l.handler.NewLogger(subcomponent),
    }
}

func (l DefaultLogger) processString(prefix, text string, v ...interface{}) string {
    str := fmt.Sprintf(text, v...)
    line := fmt.Sprintf("%s %s: %s", prefix, l.handler.GetComponentName(), str)
    l.handler.HandleLine(line)
    return str
}

func (l DefaultLogger) Panic(text string, v ...interface{}) {
    panic(l.processString("P", text, v...))
}

func (l DefaultLogger) Error(text string, v ...interface{}) error {
    return errors.New(
        l.processString("E", text, v...),
    )
}

func (l DefaultLogger) Warn(text string, v ...interface{}) {
    l.processString("W", text, v...)
}

func (l DefaultLogger) Info(text string, v ...interface{}) {
    l.processString("I", text, v...)
}

func (l DefaultLogger) Verbose(level int, text string, v ...interface{}) {
    l.processString("V", text, v...)
}
