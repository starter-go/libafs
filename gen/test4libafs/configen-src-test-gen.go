package test4libafs
import (
    p0d2a11d16 "github.com/starter-go/afs"
    p574cc1a10 "github.com/starter-go/libafs/src/test/golang/code"
     "github.com/starter-go/application"
)

// type p574cc1a10.ExampleController in package:github.com/starter-go/libafs/src/test/golang/code
//
// id:com-574cc1a10fd6f931-code-ExampleController
// class:
// alias:
// scope:singleton
//
type p574cc1a10f_code_ExampleController struct {
}

func (inst* p574cc1a10f_code_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-574cc1a10fd6f931-code-ExampleController"
	r.Classes = ""
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* p574cc1a10f_code_ExampleController) new() any {
    return &p574cc1a10.ExampleController{}
}

func (inst* p574cc1a10f_code_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*p574cc1a10.ExampleController)
	nop(ie, com)

	
    com.FS = inst.getFS(ie)


    return nil
}


func (inst*p574cc1a10f_code_ExampleController) getFS(ie application.InjectionExt)p0d2a11d16.FS{
    return ie.GetComponent("#alias-0d2a11d163e349503a64168a1cdf48a2-FS").(p0d2a11d16.FS)
}


