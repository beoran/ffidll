# Copyright 2009 The Go Authors.  All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

TARG=ffidll

GOFILES:=helper.go ffi_raw.go
OFILES=ffi_$(GOARCH).$O
CGOFILES:=dll_$(GOOS).go ffi.go

# CGO_CFLAGS:=`pkg-config --cflags sdl`
#CGO_CFLAGS:=-I../test 
CGO_LDFLAGS:=-ldl -lffi 
# CGO_LDFLAGS:=`sdl-config --libs` -lSDL_image -lSDL_ttf -lSDL_mixer

CLEANFILES+=ffidll

include $(GOROOT)/src/Make.pkg
