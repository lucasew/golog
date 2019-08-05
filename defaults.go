package golog

import (
	"github.com/lucasew/golog/handler/default"
	"github.com/lucasew/golog/logger/default"
	"os"
)

// Default Default logger already built for more out of the box experience
var Default = ldefault.NewLogger(
	hdefault.NewHandler(os.Stderr, ""),
)
