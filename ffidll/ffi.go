package ffidll

// #include <stdlib.h>
// #include <dlfcn.h>
// #include <ffi.h>
import "C"
import "unsafe"
// import "fmt"


// Memory is a wrapper type around a block of memory allocated for 
// this program using the C allocator library functions.  
type Memory struct {
  ptr   unsafe.Pointer 
  size  int
}
  
  
  
// Returns the size of the memory block.  
func (mem * Memory) Size() (int) {
  if (mem == nil) { return -1 }
  if (mem.ptr == nil) { return -1 }
  return mem.size
}

// Returns the unsafe.Pointer that points to this memory block. 
func (mem * Memory) Ptr() (unsafe.Pointer) {
  if (mem == nil)     { return nil }
  if (mem.ptr == nil) { return nil }
  return mem.ptr
}

// Frees any memory allocated. Protected from double deallocation.
func (mem * Memory) Free() {
  if (mem == nil)     { return }
  if (mem.ptr == nil) { return }
  C.free(mem.ptr)
  return  
}

// Allocates a pointer to new block of memorty of the given size.
// May return nil if no memory could be allocated.
func Allocate(size int) (* Memory) {
  mem         := &Memory{};
  mem.size     = size
  mem.ptr      = C.malloc(C.size_t(mem.size));
  if (mem.ptr == nil) { return nil; }
  return mem;
}



type Kind uint16 

const (
  TYPE_VOID   = Kind(1 + iota)
  TYPE_UINT8  = Kind(1 + iota)
  TYPE_SINT8  = Kind(1 + iota)
  TYPE_UINT16 = Kind(1 + iota)
  TYPE_SINT16 = Kind(1 + iota)
  TYPE_UINT32 = Kind(1 + iota)
  TYPE_SINT32 = Kind(1 + iota)
  TYPE_UINT64 = Kind(1 + iota)
  TYPE_SINT64 = Kind(1 + iota)
  TYPE_FLOAT  = Kind(1 + iota)
  TYPE_DOUBLE = Kind(1 + iota)
  TYPE_UCHAR  = Kind(1 + iota)
  TYPE_SCHAR  = Kind(1 + iota)
  TYPE_USHORT = Kind(1 + iota)
  TYPE_SSHORT = Kind(1 + iota)
  TYPE_UINT   = Kind(1 + iota)
  TYPE_SINT   = Kind(1 + iota)
  TYPE_ULONG  = Kind(1 + iota)
  TYPE_SLONG  = Kind(1 + iota)
  TYPE_LONGD  = Kind(1 + iota)
  TYPE_POINTER= Kind(1 + iota)
  TYPE_STRUCT = Kind(1 + iota)
)

type Type struct {

}




func (kind Kind) c() (* C.ffi_type) { 
  switch (kind) { 
  case TYPE_VOID: 	return &C.ffi_type_void
  case TYPE_UINT8:	return &C.ffi_type_uint8
  case TYPE_SINT8:	return &C.ffi_type_sint8
  case TYPE_UINT16:	return &C.ffi_type_uint16
  case TYPE_SINT16:	return &C.ffi_type_sint16
  case TYPE_UINT32:	return &C.ffi_type_uint32
  case TYPE_SINT32:	return &C.ffi_type_sint32
  case TYPE_UINT64:	return &C.ffi_type_uint64
  case TYPE_SINT64:	return &C.ffi_type_sint64
  case TYPE_DOUBLE:	return &C.ffi_type_double
  case TYPE_FLOAT:	return &C.ffi_type_float
  case TYPE_UCHAR:	return &C.ffi_type_uchar
  case TYPE_SCHAR:	return &C.ffi_type_schar
  case TYPE_USHORT:	return &C.ffi_type_ushort
  case TYPE_SSHORT:	return &C.ffi_type_sshort
  case TYPE_UINT:	return &C.ffi_type_uint
  case TYPE_SINT:	return &C.ffi_type_sint  
  case TYPE_ULONG:	return &C.ffi_type_ulong
  case TYPE_SLONG:	return &C.ffi_type_slong
  case TYPE_POINTER:	return &C.ffi_type_pointer
  default:
    return nil
  }
  return nil
}

type Callable struct {  
  cif     * C.ffi_cif
  fun     * Function
  cifmem  * Memory
  argimem * Memory 
  cifs C.ffi_cif;
}



type ffi_func *[0]uint8;

func Prepare(fun * Function, nargs uint) (* Callable, int) {
  var ffi_type_void C.ffi_type = C.ffi_type { C.size_t(1), C.ushort(1), C.FFI_TYPE_VOID, nil }

  var cif C.ffi_cif
  var typp * C.ffi_type
  // recover from panic here 
  defer func() {
    if x := recover(); x != nil {
      println("panicking with value", x)      
    }
    println("function returns normally") // executes only when hideErrors==true
  }()

  
  call 		 := &Callable{}
  println("cif size: ", unsafe.Sizeof(cif));
  call.cifmem = Allocate(unsafe.Sizeof(cif) + 100)
  call.argimem= Allocate(unsafe.Sizeof(typp) * 10)
  call.cif    = &call.cifs 
  
  // (*C.ffi_cif)(call.cifmem.Ptr())
  // fmt.Println("cif ABI, nargs:", int(call.cif.abi), int(call.cif.nargs))
  call.cif.abi   = C.FFI_DEFAULT_ABI
  call.cif.nargs = C.uint(nargs)
  call.fun       = fun
  call.cif.rtype = &C.ffi_type_void; 
  //stat          := 777
  println("cif address:", (unsafe.Pointer)(call.cif));
  println("void type address:", (unsafe.Pointer)(&ffi_type_void));
  println("cgo void type address:", (unsafe.Pointer)(&C.ffi_type_void));

  stat := C.ffi_prep_cif(call.cif, C.FFI_DEFAULT_ABI, C.uint(nargs),
          &ffi_type_void, nil);
          
		   /*&C.ffi_type_void, (**C.ffi_type)(call.argimem.ptr))*/
  return call, int(stat);
} 

func (call * Callable) Call() {
  C.ffi_call(call.cif, ffi_func(call.fun.ptr), nil, nil);
}


/*
ffi_status ffi_prep_cif (ffi_cif *CIF, ffi_abi ABI,
          unsigned int NARGS, ffi_type *RTYPE, ffi_type **ARGTYPES)
     This initializes CIF according to the given parameters.

     ABI is the ABI to use; normally `FFI_DEFAULT_ABI' is what you
     want.  *note Multiple ABIs:: for more information.

     NARGS is the number of arguments that this function accepts.
     `libffi' does not yet handle varargs functions; see *note Missing
     Features:: for more information.

     RTYPE is a pointer to an `ffi_type' structure that describes the
     return type of the function.  *Note Types::.

     ARGTYPES is a vector of `ffi_type' pointers.  ARGTYPES must have
     NARGS elements.  If NARGS is 0, this argument is ignored.

     `ffi_prep_cif' returns a `libffi' status code, of type
     `ffi_status'.  This will be either `FFI_OK' if everything worked
     properly; `FFI_BAD_TYPEDEF' if one of the `ffi_type' objects is
     incorrect; or `FFI_BAD_ABI' if the ABI parameter is invalid.

   To call a function using an initialized `ffi_cif', use the
`ffi_call' function:

 -- Function: void ffi_call (ffi_cif *CIF, void *FN, void *RVALUE, void
          **AVALUES)
	  
ffi_type_void'
     The type `void'.  This cannot be used for argument types, only for
     return values.

`ffi_type_uint8'
     An unsigned, 8-bit integer type.

`ffi_type_sint8'
     A signed, 8-bit integer type.

`ffi_type_uint16'
     An unsigned, 16-bit integer type.

`ffi_type_sint16'
     A signed, 16-bit integer type.

`ffi_type_uint32'
     An unsigned, 32-bit integer type.

`ffi_type_sint32'
     A signed, 32-bit integer type.

`ffi_type_uint64'
     An unsigned, 64-bit integer type.

`ffi_type_sint64'
     A signed, 64-bit integer type.

`ffi_type_float'
     The C `float' type.

`ffi_type_double'
     The C `double' type.

`ffi_type_uchar'
     The C `unsigned char' type.

`ffi_type_schar'
     The C `signed char' type.  (Note that there is not an exact
     equivalent to the C `char' type in `libffi'; ordinarily you should
     either use `ffi_type_schar' or `ffi_type_uchar' depending on
     whether `char' is signed.)

`ffi_type_ushort'
     The C `unsigned short' type.

`ffi_type_sshort'
     The C `short' type.

`ffi_type_uint'
     The C `unsigned int' type.

`ffi_type_sint'
     The C `int' type.

`ffi_type_ulong'
     The C `unsigned long' type.

`ffi_type_slong'
     The C `long' type.

`ffi_type_longdouble'
     On platforms that have a C `long double' type, this is defined.
     On other platforms, it is not.

`ffi_type_pointer'
     A generic `void *' pointer.  You should use this for all pointers,
     regardless of their real type.
 -- ffi_type:
     The `ffi_type' has the following members:
    `size_t size'
          This is set by `libffi'; you should initialize it to zero.

    `unsigned short alignment'
          This is set by `libffi'; you should initialize it to zero.

    `unsigned short type'
          For a structure, this should be set to `FFI_TYPE_STRUCT'.

    `ffi_type **elements'
          This is a `NULL'-terminated array of pointers to `ffi_type'
          objects.  There is one element per field of the struct.


 -- Function: void ffi_closure_free (void *WRITABLE)
     Free memory allocated using `ffi_closure_alloc'.  The argument is
     the writable address that was returned.

   Once you have allocated the memory for a closure, you must construct
a `ffi_cif' describing the function call.  Finally you can prepare the
closure function:

 -- Function: ffi_status ffi_prep_closure_loc (ffi_closure *CLOSURE,
          ffi_cif *CIF, void (*FUN) (ffi_cif *CIF, void *RET, void
          **ARGS, void *USER_DATA), void *USER_DATA, void *CODELOC)
     Prepare a closure function.

	  
*/




