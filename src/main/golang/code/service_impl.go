package code

import (
	"fmt"
	"sort"

	"github.com/starter-go/afs"
	"github.com/starter-go/vlog"
)

// AFSystemService ...
type AFSystemService struct {

	//starter:component

	_as func(afs.Service) //starter:as("#")

	RegList []afs.Registry //starter:inject(".")

	regListCache []*afs.Registration
}

func (inst *AFSystemService) _impl() afs.Service {
	return inst
}

// Fetch ...
func (inst *AFSystemService) Fetch(uri afs.URI) (afs.Path, error) {
	all := inst.getAll()
	for _, reg := range all {
		dr := reg.Driver
		if dr == nil {
			continue
		}
		if dr.Support(uri) {
			p, err := dr.Fetch(uri)
			if err == nil {
				return p, nil
			}
			vlog.Debug("afs_driver_error: %s", err.Error())
		}
	}
	return nil, fmt.Errorf("no driver support afs.URI [%s]", uri)
}

func (inst *AFSystemService) getAll() []*afs.Registration {
	dst := inst.regListCache
	if dst != nil {
		return dst
	}
	src := inst.RegList
	for _, r1 := range src {
		if r1 == nil {
			continue
		}
		r2 := r1.Registration()
		if r2 == nil {
			continue
		}
		if r2.Driver == nil || !r2.Enabled {
			continue
		}
		dst = append(dst, r2)
	}
	sort.Sort(&afsDriverListSorter{items: dst})
	inst.regListCache = dst
	return dst
}

////////////////////////////////////////////////////////////////////////////////

type afsDriverListSorter struct {
	items []*afs.Registration
}

func (inst *afsDriverListSorter) Len() int {
	return len(inst.items)
}

func (inst *afsDriverListSorter) Less(i1, i2 int) bool {
	o1 := inst.items[i1]
	o2 := inst.items[i2]
	return o1.Priority < o2.Priority
}

func (inst *afsDriverListSorter) Swap(i1, i2 int) {
	l := inst.items
	l[i1], l[i2] = l[i2], l[i1]
}
