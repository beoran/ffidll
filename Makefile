# Copyright 2010 Beoran. 
# 
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

include $(GOROOT)/src/Make.$(GOARCH)

all: libs test-ffidll

libs:
	make -C ffidll install

test-ffidll: test-ffidll.go libs
	$(GC) test-ffidll.go
	$(LD) -o $@ test-ffidll.$(O)
	
clean:
	make -C ffidll clean
	rm -f -r *.8 *.6 *.o */*.8 */*.6 */*.o */_obj test-fungo test-gui test-midi
	
