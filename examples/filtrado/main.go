package main

import (
	"os"

	hconsole "github.com/lucasew/golog/handler/console"
	lconsole "github.com/lucasew/golog/logger/console"
	lfilter "github.com/lucasew/golog/logger/filter"
)

func main() {
	h := hconsole.NewHandler(os.Stderr, "")
	l := lconsole.NewLogger(h)
	f := lfilter.NewLogger(l, lfilter.LvlInfo)
	f.Info("Isto tem que aparecer")
	f.Error("Isto também")
	f.Verbose(0, "Já isto não")
}
