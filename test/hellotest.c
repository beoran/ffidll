#include <stdio.h>
#include <dlfcn.h>
#include <ffi.h>


int main(void) {
  void *lib;
  void *func;
  lib 	= dlopen("hellolib.so", RTLD_LAZY);
  if(!lib) { 
    puts("Could not load library");
    puts(dlerror());
    return 1;
  }  
  func 	= dlsym(lib, "hello");
  if(!func) { 
    puts("Could not load function");
    puts(dlerror());
    return 1;
  }  
  
  dlclose(lib);
  puts(dlerror());
  return 0;
}