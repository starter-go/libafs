package main

import (
	"os"

	"github.com/starter-go/libafs/modules/libafs"
	"github.com/starter-go/units"
)

func main() {
	m := libafs.ModuleForTest()

	// i := starter.Init(os.Args)
	// i.MainModule(m)
	// i.WithPanic(true).Run()

	r := units.NewRunner()
	r.Dependencies(m)
	r.EnablePanic(true).Run(os.Args)
}
