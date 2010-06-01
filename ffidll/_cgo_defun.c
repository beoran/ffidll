
#include "runtime.h"
#include "cgocall.h"

#pragma dynimport initcgo initcgo "/home/bjorn/arch/dl/go/gc/pkg/linux_386/libcgo.so"
#pragma dynimport libcgo_thread_start libcgo_thread_start "/home/bjorn/arch/dl/go/gc/pkg/linux_386/libcgo.so"
#pragma dynimport libcgo_set_scheduler libcgo_set_scheduler "/home/bjorn/arch/dl/go/gc/pkg/linux_386/libcgo.so"
#pragma dynimport _cgo_malloc _cgo_malloc "/home/bjorn/arch/dl/go/gc/pkg/linux_386/libcgo.so"
#pragma dynimport _cgo_free _cgo_free "/home/bjorn/arch/dl/go/gc/pkg/linux_386/libcgo.so"

void
·_C_GoString(int8 *p, String s)
{
	s = gostring((byte*)p);
	FLUSH(&s);
}

void
·_C_CString(String s, int8 *p)
{
	p = cmalloc(s.len+1);
	mcpy((byte*)p, s.str, s.len);
	p[s.len] = 0;
	FLUSH(&p);
}

#pragma dynimport _cgo_calloc _cgo_calloc "/home/bjorn/arch/dl/go/gc/pkg/linux_386/ffidll/dll.so"
void (*_cgo_calloc)(void*);

void
·_C_calloc(struct{uint8 x[12];}p)
{
	cgocall(_cgo_calloc, &p);
}

#pragma dynimport _cgo_dlerror _cgo_dlerror "/home/bjorn/arch/dl/go/gc/pkg/linux_386/ffidll/dll.so"
void (*_cgo_dlerror)(void*);

void
·_C_dlerror(struct{uint8 x[4];}p)
{
	cgocall(_cgo_dlerror, &p);
}

#pragma dynimport _cgo_dlsym _cgo_dlsym "/home/bjorn/arch/dl/go/gc/pkg/linux_386/ffidll/dll.so"
void (*_cgo_dlsym)(void*);

void
·_C_dlsym(struct{uint8 x[12];}p)
{
	cgocall(_cgo_dlsym, &p);
}

#pragma dynimport _cgo_free _cgo_free "/home/bjorn/arch/dl/go/gc/pkg/linux_386/ffidll/dll.so"
void (*_cgo_free)(void*);

void
·_C_free(struct{uint8 x[4];}p)
{
	cgocall(_cgo_free, &p);
}

#pragma dynimport _cgo_dlclose _cgo_dlclose "/home/bjorn/arch/dl/go/gc/pkg/linux_386/ffidll/dll.so"
void (*_cgo_dlclose)(void*);

void
·_C_dlclose(struct{uint8 x[8];}p)
{
	cgocall(_cgo_dlclose, &p);
}

#pragma dynimport _cgo_dlopen _cgo_dlopen "/home/bjorn/arch/dl/go/gc/pkg/linux_386/ffidll/dll.so"
void (*_cgo_dlopen)(void*);

void
·_C_dlopen(struct{uint8 x[12];}p)
{
	cgocall(_cgo_dlopen, &p);
}

