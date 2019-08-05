package main

import (
	"github.com/lucasew/golog/handler/default"
	"github.com/lucasew/golog/logger/default"
	"github.com/lucasew/golog/logger/filter"
	"os"
)

func main() {
	h := hdefault.NewHandler(os.Stderr, "")
	l := ldefault.NewLogger(h)
	f := lfilter.NewLogger(l, lfilter.LvlInfo)
	f.Info("Isto tem que aparecer")
	f.Error("Isto também")
	f.Verbose(0, "Já isto não")
}
