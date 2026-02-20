package console

import (
	"fmt"
	"io"

	"github.com/lucasew/golog/handler"
)

// Handler writes the log line to an io.Writer, plus controlling the log namespace.
type Handler struct {
	component string
	writer    io.Writer
	parent    *Handler
}

// NewHandler creates a handler.
func NewHandler(writer io.Writer, component string) *Handler {
	return &Handler{
		component: component,
		writer:    writer,
		parent:    nil,
	}
}

// HandleLine handles a line from a logger.
func (h *Handler) HandleLine(str string) error {
	_, err := fmt.Fprintln(h.writer, str)
	return err
}

// GetComponentName gets the log namespace of current handler.
func (h *Handler) GetComponentName() string {
	if h.parent != nil {
		pName := h.parent.GetComponentName()
		if pName != "" {
			return fmt.Sprintf("%s.%s", pName, h.component)
		}
	}
	return h.component
}

// NewHandler creates a logger using a subnamespace of the parent.
func (h *Handler) NewHandler(subcomponent string) handler.LoggerHandler {
	return &Handler{
		component: subcomponent,
		writer:    h.writer,
		parent:    h,
	}
}
