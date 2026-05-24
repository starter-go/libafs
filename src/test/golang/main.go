package main

import (
	"os"

	"github.com/starter-go/libafs/modules/libafs"
	"github.com/starter-go/units"
)

func main() {
	m := libafs.ModuleForTest()

	ctx := units.NewContext()

	ctx.Arguments = os.Args
	ctx.Module = m
	ctx.UsePanic = true

	units.Run(ctx)

}
