// Created by cgo - DO NOT EDIT
package ffidll

import "unsafe"

type _ unsafe.Pointer

type _Cstruct___0 struct {
	abi		_C_ffi_abi
	nargs		_C_uint
	arg_types	**_C_ffi_type
	rtype		*_C_ffi_type
	bytes		_C_uint
	flags		_C_uint
}
type _C_ushort uint16
type _Cstruct__ffi_type struct {
	size		_C_size_t
	alignment	_C_ushort
	_type		_C_ushort
	elements	**_Cstruct__ffi_type
}
type _C_size_t _C_uint
type _C_uint uint32
type _C_int int32
type _C_ffi_type _Cstruct__ffi_type
type _C_char int8
type _C_ffi_status uint32
type _C_ffi_abi uint32
type _C_ffi_cif _Cstruct___0
type _C_void [0]byte
var _C_ffi_type_void *_C_ffi_type
const FFI_TYPE_FLOAT = 2
const FFI_TYPE_UINT32 = 9
const FFI_TYPE_VOID = 0
const FFI_TYPE_UINT16 = 7
const FFI_TYPE_UINT8 = 5
const FFI_TYPE_SINT16 = 8
const RTLD_LAZY = 0x00001
const RTLD_LOCAL = 0
const FFI_TYPE_STRUCT = 13
const FFI_TYPE_SINT32 = 10
const FFI_TYPE_SINT64 = 12
const FFI_TYPE_SINT8 = 6
const FFI_TYPE_UINT64 = 11
const FFI_TYPE_POINTER = 14
const FFI_TYPE_DOUBLE = 3
const FFI_DEFAULT_ABI = 1

func _C_GoString(*_C_char) string
func _C_calloc(_C_size_t, _C_size_t) unsafe.Pointer
func _C_CString(string) *_C_char
func _C_dlerror() *_C_char
func _C_dlsym(unsafe.Pointer, *_C_char) unsafe.Pointer
func _C_malloc(_C_size_t) unsafe.Pointer
func _C_free(unsafe.Pointer)
func _C_dlclose(unsafe.Pointer) _C_int
func _C_ffi_prep_cif(*_C_ffi_cif, _C_ffi_abi, _C_uint, *_C_ffi_type, **_C_ffi_type) _C_ffi_status
func _C_dlopen(*_C_char, _C_int) unsafe.Pointer
func _C_ffi_call(*_C_ffi_cif, *[0]byte, unsafe.Pointer, *unsafe.Pointer)
