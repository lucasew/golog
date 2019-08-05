package hdefault

import (
	"fmt"
	"github.com/lucasew/golog/handler"
	"io"
)

// DefaultLogHandler is a handler that just writes the line to a io.Writer interface, plus controlling the log namespace
type DefaultLogHandler struct {
	component string
	writter   io.Writer
	father    *DefaultLogHandler
}

// NewHandler creates a handler
func NewHandler(writter io.Writer, component string) DefaultLogHandler {
	return DefaultLogHandler{
		component: component,
		writter:   writter,
		father:    nil,
	}
}

// HandleLine handles a line from a logger
func (h DefaultLogHandler) HandleLine(str string) error {
	_, err := fmt.Fprintln(h.writter, str)
	return err
}

// GetComponentName gets the log namespace of current handler
func (h DefaultLogHandler) GetComponentName() string {
	if h.father != nil {
		if h.father.GetComponentName() != "" {
			return fmt.Sprintf("%s.%s",
				h.father.GetComponentName(),
				h.component,
			)
		}
	}
	return h.component
}

// NewHandler Creates a logger using a subnamespace of the father
func (h DefaultLogHandler) NewHandler(subcomponent string) handler.LoggerHandler {
	return DefaultLogHandler{
		component: subcomponent,
		writter:   h.writter,
		father:    &h,
	}
}
