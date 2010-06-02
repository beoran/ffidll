#include <stdio.h>
#include <stdlib.h>
#include <dlfcn.h>
#include <ffi.h>

void pdlerror() {
  char *err;
  err = dlerror();
  if(err) puts(err);
}


typedef void (*voidfunc)(void);


int main(void) {
  void *lib;
  void *func;
  voidfunc cfun;
  char *err;
   
  
  lib 	= dlopen("./hellolib.so", RTLD_LAZY);
  if(!lib) { 
    puts("Could not load library");
    pdlerror();    
    return 1;
  }  
  func 	= dlsym(lib, "say_nothing_at_all");
  if(!func) { 
    puts("Could not load function");
    pdlerror();
    return 1;
  } else {
    int status;
    int reps;
    ffi_cif cif;
    printf("Size of cif: %d\n", sizeof(cif)); 
    // call function manually.
    cfun = (voidfunc)(func),
    (*cfun)();
    // And call though FFI
    status = ffi_prep_cif(&cif, FFI_DEFAULT_ABI, 0, &ffi_type_void, NULL);
    // the return type may NOT be null.
    printf("FFI prep status: %d. Calling through FFI:\n", status);
    ffi_call(&cif, func, NULL, NULL);    
  }
  
  dlclose(lib);
  pdlerror();
  return 0;
}
