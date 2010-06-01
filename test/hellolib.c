/* Dynamic library for tesing */
#include <stdio.h>

void init() {
  puts("Init called from hellolib!");
}

void hello(void) {
  puts("Hello from hellolib!");
}