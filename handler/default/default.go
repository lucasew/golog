package golog_handler_default

import (
    "io"
    "fmt"
    "github.com/lucasew/golog/handler"
)

type DefaultLogHandler struct {
    component string;
    writter io.Writer;
    father *DefaultLogHandler;
}

func NewLogger(writter io.Writer, component string) DefaultLogHandler{
    return DefaultLogHandler{
        component: component,
        writter: writter,
        father: nil,
    }
}

func (h DefaultLogHandler) HandleLine(str string) error {
    _, err := fmt.Fprintln(h.writter, str)
    return err
}

func (h DefaultLogHandler) GetComponentName() string {
    if (h.father != nil) {
        if (h.father.GetComponentName() != "") {
            return fmt.Sprintf("%s.%s", 
               h.father.GetComponentName(),
                h.component,
            )
        }
    }
    return h.component
}

func (h DefaultLogHandler) NewLogger(subcomponent string) golog_handler.LoggerHandler {
    return DefaultLogHandler{
        component: subcomponent,
        writter: h.writter,
        father: &h,
    }
}
