package ipv4

// #include <stdlib.h>
// #include "ipv4.h"
import "C"
import "unsafe"

// Code to work with IP addresses
func PackIPv4(address *string) uint32 {
	cAddr := C.CString(*address)
	defer C.free(unsafe.Pointer(cAddr))
	return uint32(C.pack_ipv4_address(cAddr))
}

// Code to work with IP addresses
func UnpackIPv4(address uint32) string {
	buf := make([]byte, 32)
	val := C.fast_intoa(C.uint(address), (*C.char)(unsafe.Pointer(&buf[0])), C.int(len(buf)))
	return C.GoString(val)
}
