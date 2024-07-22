package code

import (
	"fmt"
	"strings"

	"github.com/starter-go/afs"
	"github.com/starter-go/afs/files"
)

// LocalFileSystemDriver ...
type LocalFileSystemDriver struct {

	//starter:component

	_as func(afs.Registry) //starter:as(".")

	fs afs.FS
}

func (inst *LocalFileSystemDriver) _impl() (afs.Driver, afs.Registry) {
	return inst, inst
}

// Registration ...
func (inst *LocalFileSystemDriver) Registration() *afs.Registration {
	return &afs.Registration{
		Name:     "local_fs",
		Enabled:  true,
		Priority: 100,
		Driver:   inst,
	}
}

// Support ...
func (inst *LocalFileSystemDriver) Support(uri afs.URI) bool {
	str := uri.String()
	return strings.HasPrefix(str, "file:")
}

// Fetch ...
func (inst *LocalFileSystemDriver) Fetch(uri afs.URI) (afs.Path, error) {
	const (
		prefix = "file:"
	)
	str := uri.String()
	if strings.HasPrefix(str, prefix) {
		str = str[len(prefix):]
	} else {
		return nil, fmt.Errorf("unsupported AFS URI [%s]", str)
	}
	fs := inst.getFS()
	path := fs.NewPath(str)
	return path, nil
}

func (inst *LocalFileSystemDriver) getFS() afs.FS {
	fs := inst.fs
	if fs != nil {
		return fs
	}
	fs = files.FS()
	inst.fs = fs
	return fs
}
