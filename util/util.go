package util

import "C"

import (
	"unsafe"
)

// @fix(judah): Use cgo.Handle or runtime.Pinner to avoid heap allocating every c-string
type CString = *byte

func ToGoString(str CString) string {
	if str == nil {
		panic("cannot convert nil CString to string")
	}

	return C.GoStringN((*C.char)(unsafe.Pointer(str)), C.int(CStringLen(str)))
}

func ToCString(str string) CString {
	return (CString)(unsafe.Pointer(C.CString(str)))
}

func CStringLen(str CString) int {
	ptr := uintptr(unsafe.Pointer(str))
	for *(*byte)((unsafe.Pointer)(ptr)) != 0 {
		ptr += 1
	}

	return int(ptr - uintptr(unsafe.Pointer(str)))
}
