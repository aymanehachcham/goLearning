package linearmodels

import (
	"C"
	"unsafe"
)

func sayHello(name string) {
	
	cData := C.CString(name)
	defer C.free(unsafe.Pointer(cData))
	C.sayHello(cData)
}