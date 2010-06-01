// Created by cgo - DO NOT EDIT
package dll

import "unsafe"

type _ unsafe.Pointer

type _C_size_t _C_uint
type _C_uint uint32
type _C_int int32
type _C_char int8
type _C_void [0]byte
const RTLD_LAZY = 0x00001
const RTLD_LOCAL = 0

func _C_GoString(*_C_char) string
func _C_calloc(_C_size_t, _C_size_t) unsafe.Pointer
func _C_CString(string) *_C_char
func _C_dlerror() *_C_char
func _C_dlsym(unsafe.Pointer, *_C_char) unsafe.Pointer
func _C_free(unsafe.Pointer)
func _C_dlclose(unsafe.Pointer) _C_int
func _C_dlopen(*_C_char, _C_int) unsafe.Pointer
