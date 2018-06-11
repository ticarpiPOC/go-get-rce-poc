package main
/*
#cgo CFLAGS: -fplugin=./rce.so
#include <stdio.h>
#include <stdlib.h>
void goputs(char* s) {
	printf("%s\n", s);
}
*/
import "C"
import "unsafe"
func main() {
  cs := C.CString("go got rced ;)\\n")
  C.goputs(cs)
  C.free(unsafe.Pointer(cs))
}
