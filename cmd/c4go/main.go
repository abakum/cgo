//go:generate cmd /c "set CGO_ENABLED=1 && PATH E:/msys64/mingw64/bin;%PATH% && go install"

package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void say_hello_from_c(const char *name) {
    printf("Hello from C, %s!\n", name);
}
*/
import "C"
import "unsafe"

func main() {
	name := C.CString("World")
	defer C.free(unsafe.Pointer(name)) // Более правильное освобождение памяти
	C.say_hello_from_c(name)
}
