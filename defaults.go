package golog

import (
    "github.com/lucasew/golog/handler/default"
    "github.com/lucasew/golog/logger/default"
    "os"
)

var Default = golog_logger_default.NewLogger(
    golog_handler_default.NewHandler(os.Stderr, ""),
)
