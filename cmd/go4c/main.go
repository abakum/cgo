//go:build cgo

package main

/*
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

// Объявление функции, которая будет реализована в Go
extern void SayHelloFromGo(char* name);

// Функция-обертка для вызова из Go
static void say_hello_from_c_wrapper(char* name) {
    SayHelloFromGo(name);
}
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
