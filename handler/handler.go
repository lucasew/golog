package handler

// LoggerHandler A object that receives a line of log from a logger and do something with it
type LoggerHandler interface {
	HandleLine(line string) error
	GetComponentName() string
	NewHandler(subcomponent string) LoggerHandler
}
