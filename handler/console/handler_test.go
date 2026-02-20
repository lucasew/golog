package console

import (
	"bytes"
	"strings"
	"testing"
)

func TestHandler_GetComponentName(t *testing.T) {
	buf := &bytes.Buffer{}
	h := NewHandler(buf, "root")

	if name := h.GetComponentName(); name != "root" {
		t.Errorf("Expected root, got %s", name)
	}

	h2 := h.NewHandler("sub")
	if name := h2.GetComponentName(); name != "root.sub" {
		t.Errorf("Expected root.sub, got %s", name)
	}

	h3 := h2.NewHandler("leaf")
	if name := h3.GetComponentName(); name != "root.sub.leaf" {
		t.Errorf("Expected root.sub.leaf, got %s", name)
	}
}

func TestHandler_HandleLine(t *testing.T) {
	buf := &bytes.Buffer{}
	h := NewHandler(buf, "test")

	err := h.HandleLine("hello")
	if err != nil {
		t.Errorf("HandleLine failed: %v", err)
	}

	if !strings.Contains(buf.String(), "hello") {
		t.Errorf("Expected 'hello' in output, got %q", buf.String())
	}
}
