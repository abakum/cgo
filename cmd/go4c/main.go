//go:build cgo

package main

/*
#include <stdio.h>

// Экспортируем функцию SayHelloFromGo в C
extern void SayHelloFromGo(const char *name);

// Функция, которая будет вызвана из Go
void say_hello_from_c_wrapper(const char *name) {
    SayHelloFromGo(name);
}
*/
import "C"
import "fmt"

// Го функция, которую нужно вызвать из C
//export SayHelloFromGo
func SayHelloFromGo(name *C.char) {
	fmt.Printf("Hello from Go, %s!\n", C.GoString(name))
}

func main() {
	// Задаем имя
	name := C.CString("World")
	// Завершаем программу
	defer C.free(C.malloc(C.size_t(len(C.GoString(name)))))
	// Вызываем обертку
	C.say_hello_from_c_wrapper(name)
}
