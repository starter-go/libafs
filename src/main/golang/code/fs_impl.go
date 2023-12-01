package code

// ImplAFS ... use:github.com/starter-go/starter/common/fs/AFSProxy
type ImplAFS struct {

	//starter:component

	// _as func(afs.FS) //starter:as(".")
	// inner afs.FS
}

// func (inst *ImplAFS) _impl() afs.FS {
// 	return inst
// }

// func (inst *ImplAFS) getInner() afs.FS {
// 	i := inst.inner
// 	if i == nil {
// 		i = files.FS()
// 		inst.inner = i
// 	}
// 	return i
// }

// // NewPath ...
// func (inst *ImplAFS) NewPath(path string) afs.Path {
// 	return inst.getInner().NewPath(path)
// }

// // ListRoots ...
// func (inst *ImplAFS) ListRoots() []afs.Path {
// 	return inst.getInner().ListRoots()
// }

// // CreateTempFile ...
// func (inst *ImplAFS) CreateTempFile(prefix, suffix string, dir afs.Path) (afs.Path, error) {
// 	return inst.getInner().CreateTempFile(prefix, suffix, dir)
// }

// // PathSeparator return ';'(windows) | ':'(unix)
// func (inst *ImplAFS) PathSeparator() string {
// 	return inst.getInner().PathSeparator()
// }

// // Separator return '/'(unix) | '\'(windows)
// func (inst *ImplAFS) Separator() string {
// 	return inst.getInner().Separator()
// }

// // SetDefaultOptionsHandler 设置一个函数，用来处理默认的I/O选项
// func (inst *ImplAFS) SetDefaultOptionsHandler(fn afs.OptionsHandlerFunc) error {
// }
