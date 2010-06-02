gcc -shared -fPIC hellolib.c -o hellolib.so
gcc -I/usr/local/include -c hellotest.c -o hellotest.o 
gcc -L/usr/local/lib -ldl -lffi  hellotest.o  -o hellotest 
./hellotest