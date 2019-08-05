package main

import (
	"github.com/lucasew/golog"
)

func main() {
	log := golog.Default
	log.Info("Eoq")
	log.Error("Trabzeraaa")
	// log.Panic("Nem queria funcionar %s", "mesmo")
	u := log.NewLogger("main")
	u.Info("De outro modulo")
	v := u.NewLogger("LOGGER")
	v.Warn("Meu nome é %s", "lucas")
}
