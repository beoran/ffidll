package ffidll
// #include <stdlib.h>
// #include <dlfcn.h>
import "C"
import "unsafe" 

type Library struct {
  name    string
  ptr     unsafe.Pointer
}

type Function struct { 
  name 	  string;
  ptr     unsafe.Pointer;
}



/*
const DLL_LAZY   int = C.RTLD_LAZY
// Relocations are performed at an implementation-defined time.
const DLL_NOW    int = C.RTLD_NOW
// Relocations are performed when the object is loaded.	
const DLL_GLOBAL int = C.RTLD_GLOBAL
// All symbols are available for relocation processing of other modules.	
const DLL_LOCAL  int = C.RTLD_LOCAL
// All symbols are not made available for relocation processing by other modules.
*/

/*
func ptr(pointer * C.void) (unsafe.Pointer) {
	return unsafe.Pointer(pointer)
}


func vptr(gopointer unsafe.Pointer) (* C.void)  {
	return (* C.void)((gopointer))
}
*/

type ptr unsafe.Pointer

// Helper functions 
// Calls C malloc
func malloc(size int) (unsafe.Pointer) { 
  return (unsafe.Pointer(C.calloc(C.size_t(size), C.size_t(1))))
} 

// Calls C free
func free(ptr unsafe.Pointer) { 
  C.free(ptr)
} 

// Allocates a string with the given byte length
// don't forget a call to defer s.free() ! 
func cstrNew(size int) (* C.char) {
  return (*C.char)(malloc(size))  
}

// free is a method on C char * strings to method to free the associated memory 
func (self * C.char) free() {
  free(unsafe.Pointer(self))
}

/*
// free is a method on C int * pointers to method to free the associated memory 
func (self * C.int) free() {
  C.free(unsafe.Pointer(self))
}
*/
// cstring converts a string to a C string. This allocates memory, 
// so don't forget to add a "defer s.free()"
func cstr(self string) (*C.char) {  
  /*
  buf := cstrNew(len(self) + 1)
  // Allocate buffer 
  if buf == nil { panic("Could not allocate memory for string") }
  // Some nice pointer math   
  for i:=0 ; i < len(self) ; i ++ {
    ch  := self[i]
    pto := (*byte)(ptr(uintptr(ptr(buf)) + uintptr(i)))
    *pto = ch
  }
  // Don't forget to zero-terminate
  ptoe := (*byte)(ptr(uintptr(ptr(buf)) + uintptr(len(self))))
  *ptoe = byte(0)
   
  return buf
  */
  // Strangely enough, C.String does NOT work for me. :p
  return C.CString(self)
}




func Open(name string) * Library {
	library 	:= &Library{}
	library.name 	 = name
	cname 		:= cstr(library.name)
 	defer cname.free()
 	library.ptr  	 = C.dlopen(cname, C.RTLD_LOCAL + C.RTLD_LAZY)
	return library;
}

func (library * Library) Close() (int) {
	if (library == nil) 	{ return -1 } 
	if (library.ptr == nil) { return -1 } 
    	res := C.dlclose(library.ptr)
	library.ptr = nil
    	return int(res)
}


func Error() (string) {
    return C.GoString(C.dlerror());
}

func (library * Library) Sym(name string) (*Function) {
	if (library == nil) { return nil } 
	if (library.ptr == nil) { return nil } 
	function 		:= &Function{}
	function.name 	 	 = name
	cname 			:= cstr(function.name) 
	defer cname.free()
	function.ptr  	 	 = C.dlsym(library.ptr, cname)
	return function
}


type VoidFunc func()();

func (fun * Function) CallVoid() {
  var vf * VoidFunc = nil;
  vf = (* VoidFunc)(fun.ptr)
  (*vf)()
}

