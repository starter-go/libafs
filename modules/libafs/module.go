package libafs

import (
	"github.com/starter-go/application"
	"github.com/starter-go/libafs"
	"github.com/starter-go/libafs/gen/main4libafs"
	"github.com/starter-go/libafs/gen/test4libafs"
)

// Module ...
func Module() application.Module {
	mb := libafs.NewMainModule()
	mb.Components(main4libafs.ExportComponents)
	return mb.Create()
}

// ModuleForTest ...
func ModuleForTest() application.Module {

	parent := Module()

	mb := libafs.NewTestModule()
	mb.Components(test4libafs.ExportComponents)
	mb.Depend(parent)
	return mb.Create()
}
