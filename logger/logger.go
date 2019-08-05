package golog_logger

type Logger interface {
	NewLogger(subcomponent string) Logger
	Panic(str string, v ...interface{})
	Error(str string, v ...interface{}) error
	Warn(str string, v ...interface{})
	Info(str string, v ...interface{})
	Verbose(level int, str string, v ...interface{})
}
