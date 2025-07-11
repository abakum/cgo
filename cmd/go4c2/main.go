//go:buid windows
//go:generate cmd /c "set CGO_ENABLED=1 && PATH E:/msys64/mingw64/bin;%PATH% && go install"
package main

/*
#cgo windows CFLAGS: -IE:/msys64/mingw64/include
#cgo windows LDFLAGS: -LE:/msys64/mingw64/lib

#cgo CFLAGS: -I${SRCDIR}
#include "h.h"
*/

import "C"
import (
	"fmt"
	"unsafe"
)

// Реализация Go-функции для вызова из C
//
//export SayHelloFromGo
func SayHelloFromGo(name *C.char) {
	fmt.Printf("Hello from Go, %s!\n", C.GoString(name))
}

func main() {
	// Создаем C-строку
	name := C.CString("World")
	defer C.free(unsafe.Pointer(name)) // Освобождаем память

	// Вызываем C-функцию
	C.say_hello_from_c_wrapper(name)
}
