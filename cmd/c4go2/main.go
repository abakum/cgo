//go:generate cmd /c "set CGO_ENABLED=1 && PATH E:/msys64/mingw64/bin;%PATH% && go install"
//go:generate cmd /c "PATH E:/msys64/mingw64/bin;%PATH% && go run ."

package main

/*
#cgo windows CFLAGS: -IE:/msys64/mingw64/include
#cgo windows LDFLAGS: -LE:/msys64/mingw64/lib

#cgo CFLAGS: -I${SRCDIR}
#include "h.h"
*/
import "C"
import "unsafe"

func main() {
	name := C.CString("World")
	defer C.free(unsafe.Pointer(name)) // Более правильное освобождение памяти
	C.say_hello_from_c(name)
}
