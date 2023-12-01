package code

import (
	"os"

	"github.com/starter-go/afs"
	"github.com/starter-go/units"
	"github.com/starter-go/vlog"
)

// ExampleController ...
type ExampleController struct {

	//starter:component

	FS afs.FS //starter:inject("#")

}

func (inst *ExampleController) _impl() units.Units {
	return inst
}

// Units ...
func (inst *ExampleController) Units(list []*units.Registration) []*units.Registration {

	r1 := &units.Registration{
		Name:     "test-1",
		Enabled:  true,
		Priority: 0,
		Test:     inst.test1,
	}

	list = append(list, r1)
	return list
}

// Units ...
func (inst *ExampleController) test1() error {

	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	path := inst.FS.NewPath(wd)
	info := path.GetInfo()
	vlog.Debug("wd = %s", path.GetPath())

	if info.IsDirectory() {
	}

	if info.IsFile() {
	}

	return nil
}
