package golog

import (
	"os"

	hconsole "github.com/lucasew/golog/handler/console"
	lconsole "github.com/lucasew/golog/logger/console"
)

// Default logger already built for more out of the box experience.
var Default = lconsole.NewLogger(
	hconsole.NewHandler(os.Stderr, ""),
)
