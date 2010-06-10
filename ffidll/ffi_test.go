package ffidll

import "testing"
import dll "ffidll"
import . "fmt"
import _ "time"

func say_nothing_at_all() () {

}

/* Some speed measurements: 
Call Speed through ffi:
2358867000 10000000 235
Call Speed through cgo:
1880411000 10000000 188
Call Speed of similar go function:
52813000 10000000 5

Call Speed through ffi:
2358867000 10000000 235
Call Speed through cgo:
1880411000 10000000 188
Call Speed of similar go function:
52813000 10000000 5

Call Speed through ffi:
2310437000 10000000 231
Call Speed through cgo:
2099997000 10000000 209
Call Speed of similar go function:
51935000 10000000 5

Call Speed through ffi:
2409277000 10000000 240
Call Speed through cgo:
2143419000 10000000 214
Call Speed of similar go function:
51530000 10000000 5

Call Speed through ffi:
2439886000 10000000 243
Call Speed through cgo:
2134587000 10000000 213
Call Speed of similar go function:
52339000 10000000 5

all Speed through ffi:
2383757000 10000000 238
Call Speed through cgo:
2148135000 10000000 214
Call Speed of similar go function:
51105000 10000000 5

Call Speed through ffi:
2395475000 10000000 239
Call Speed through cgo:
2134966000 10000000 213
Call Speed of similar go function:
51441000 10000000 5


*/



func TestLoad(t *testing.T) {
  lib := dll.Open("../test/hellolib.so")
  Println(dll.Error())
  /*
  i     := 0
  reps  := 1000000
  */
  
  fun := lib.Sym("say_nothing_at_all")
  _ = fun
  Println(dll.Error())
  Println("CallVoid say_nothing_at_all")
  dll.CallCDECLVoid(fun.Pointer())
/*  
  // fun.CallVoid() // Doesn't work. We need libffi
  Println("Prepare")
  call, status := dll.Prepare(fun, 0);
  Println("Call", status)
  call.Call()
*/
  Println("CallVoid sum")
  sum := lib.Sym("sum")
  dll.CallCDECLVoid(sum.Pointer())
  
  
/*  
  Println("Call Speed through ffi:")
  start := time.Nanoseconds() 
  for i = 0 ; i < reps ; i ++ {
    call.Call()
  }
  stop := time.Nanoseconds()
  delta := stop - start
  Println(delta, reps, delta / int64(reps))
  

  Println("Call Speed through assembly based go FFI:")
  start = time.Nanoseconds() 
  for i = 0 ; i < reps ; i ++ {
    dll.CallCDECLVoid(fun.Pointer())
  }
  stop = time.Nanoseconds()
  delta = stop - start
  Println(delta, reps, delta / int64(reps))  

  
  Println("Call Speed of similar go function:")
  start = time.Nanoseconds() 
  for i = 0 ; i < reps ; i ++ {
    say_nothing_at_all()
  }
  stop = time.Nanoseconds()
  delta = stop - start
  Println(delta, reps, delta / int64(reps))  
*/  
  lib.Close()
  Println(dll.Error())
}
