
typedef struct { char *p; int n; } _GoString_;
_GoString_ GoString(char *p);
char *CString(_GoString_);
#include <stdlib.h>
#include <dlfcn.h>



// Usual nonsense: if x and y are not equal, the type will be invalid
// (have a negative array count) and an inscrutable error will come
// out of the compiler and hopefully mention "name".
#define __cgo_compile_assert_eq(x, y, name) typedef char name[(x-y)*(x-y)*-2+1];

// Check at compile time that the sizes we use match our expectations.
#define __cgo_size_assert(t, n) __cgo_compile_assert_eq(sizeof(t), n, _cgo_sizeof_##t##_is_not_##n)

__cgo_size_assert(char, 1)
__cgo_size_assert(short, 2)
__cgo_size_assert(int, 4)
typedef long long __cgo_long_long;
__cgo_size_assert(__cgo_long_long, 8)
__cgo_size_assert(float, 4)
__cgo_size_assert(double, 8)

void
_cgo_calloc(void *v)
{
	struct {
		size_t p0;
		size_t p1;
		void* r;
	} *a = v;
	a->r = calloc(a->p0, a->p1);
}

void
_cgo_dlerror(void *v)
{
	struct {
		char* r;
	} *a = v;
	a->r = dlerror();
}

void
_cgo_dlsym(void *v)
{
	struct {
		void* p0;
		char* p1;
		void* r;
	} *a = v;
	a->r = dlsym(a->p0, a->p1);
}

void
_cgo_free(void *v)
{
	struct {
		void* p0;
	} *a = v;
	free(a->p0);
}

void
_cgo_dlclose(void *v)
{
	struct {
		void* p0;
		int r;
	} *a = v;
	a->r = dlclose(a->p0);
}

void
_cgo_dlopen(void *v)
{
	struct {
		char* p0;
		int p1;
		void* r;
	} *a = v;
	a->r = dlopen(a->p0, a->p1);
}

