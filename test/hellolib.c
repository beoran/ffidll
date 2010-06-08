/* Dynamic library for tesing */
#include <stdio.h>

void init() {
  puts("Init called from hellolib!");
}

void hello(void) {
  puts("Hello from hellolib!");
}


int sum(int a, int b) {
  puts("Hello from hellolib: sum!");
  return a + b;  
} 

int sumall(int a, int b, int c, int d, int e, int f,
int g, int h, int i, int j, int k, int l)  {
  return a + b + c + d + e + f + g + h + i + j + k + l;  
} 

void say_nothing_at_all(void) {
  
}

