package main

import "C"
import "unsafe"

func main() {
	cstr := C.CString("Hello, world")
	C.puts(cstr)
	C.free(unsafe.Pointer(cstr))
}
