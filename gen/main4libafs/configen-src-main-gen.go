package main4libafs
import (
    p26b4ed446 "github.com/starter-go/libafs/src/main/golang/code"
     "github.com/starter-go/application"
)

// type p26b4ed446.ImplAFS in package:github.com/starter-go/libafs/src/main/golang/code
//
// id:com-26b4ed4469428f65-code-ImplAFS
// class:
// alias:
// scope:singleton
//
type p26b4ed4469_code_ImplAFS struct {
}

func (inst* p26b4ed4469_code_ImplAFS) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-26b4ed4469428f65-code-ImplAFS"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p26b4ed4469_code_ImplAFS) new() any {
    return &p26b4ed446.ImplAFS{}
}

func (inst* p26b4ed4469_code_ImplAFS) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p26b4ed446.ImplAFS)
	nop(ie, com)

	


    return nil
}


