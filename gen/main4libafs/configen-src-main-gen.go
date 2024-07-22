package main4libafs
import (
    p0d2a11d16 "github.com/starter-go/afs"
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



// type p26b4ed446.LocalFileSystemDriver in package:github.com/starter-go/libafs/src/main/golang/code
//
// id:com-26b4ed4469428f65-code-LocalFileSystemDriver
// class:class-0d2a11d163e349503a64168a1cdf48a2-Registry
// alias:
// scope:singleton
//
type p26b4ed4469_code_LocalFileSystemDriver struct {
}

func (inst* p26b4ed4469_code_LocalFileSystemDriver) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-26b4ed4469428f65-code-LocalFileSystemDriver"
	r.Classes = "class-0d2a11d163e349503a64168a1cdf48a2-Registry"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p26b4ed4469_code_LocalFileSystemDriver) new() any {
    return &p26b4ed446.LocalFileSystemDriver{}
}

func (inst* p26b4ed4469_code_LocalFileSystemDriver) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p26b4ed446.LocalFileSystemDriver)
	nop(ie, com)

	


    return nil
}



// type p26b4ed446.AFSystemService in package:github.com/starter-go/libafs/src/main/golang/code
//
// id:com-26b4ed4469428f65-code-AFSystemService
// class:
// alias:alias-0d2a11d163e349503a64168a1cdf48a2-Service
// scope:singleton
//
type p26b4ed4469_code_AFSystemService struct {
}

func (inst* p26b4ed4469_code_AFSystemService) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-26b4ed4469428f65-code-AFSystemService"
	r.Classes = ""
	r.Aliases = "alias-0d2a11d163e349503a64168a1cdf48a2-Service"
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p26b4ed4469_code_AFSystemService) new() any {
    return &p26b4ed446.AFSystemService{}
}

func (inst* p26b4ed4469_code_AFSystemService) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p26b4ed446.AFSystemService)
	nop(ie, com)

	
    com.RegList = inst.getRegList(ie)


    return nil
}


func (inst*p26b4ed4469_code_AFSystemService) getRegList(ie application.InjectionExt)[]p0d2a11d16.Registry{
    dst := make([]p0d2a11d16.Registry, 0)
    src := ie.ListComponents(".class-0d2a11d163e349503a64168a1cdf48a2-Registry")
    for _, item1 := range src {
        item2 := item1.(p0d2a11d16.Registry)
        dst = append(dst, item2)
    }
    return dst
}


