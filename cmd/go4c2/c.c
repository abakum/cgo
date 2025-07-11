#include "h.h"

static void say_hello_from_c_wrapper(char* name) {
    SayHelloFromGo(name);
}
