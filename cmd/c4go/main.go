//go:build cgo

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

func main() {
	name := C.CString("World")
	defer C.free(C.malloc(C.size_t(len(C.GoString(name))))) // освобождаем память
	C.say_hello_from_c(name)
}
