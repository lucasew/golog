package main

import (
    "github.com/lucasew/golog/logger/default"
    "github.com/lucasew/golog/handler/default"
    "github.com/lucasew/golog/logger/filter"
    "os"
)

func main() {
    h := golog_handler_default.NewHandler(os.Stderr, "")
    l := golog_logger_default.NewLogger(h)
    f := golog_logger_filter.NewLogger(l, golog_logger_filter.LvlInfo)
    f.Info("Isto tem que aparecer")
    f.Error("Isto também")
    f.Verbose(0, "Já isto não")
}
