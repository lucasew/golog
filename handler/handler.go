package golog_handler

type LoggerHandler interface {
    HandleLine(line string) error;
    GetComponentName() string;
    NewLogger(subcomponent string) LoggerHandler;
}


