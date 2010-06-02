package ffidll

import "testing"
import dll "ffidll"
import . "fmt"
import "time"

func TestLoad(t *testing.T) {
  lib := dll.Open("../test/hellolib.so")
  Println(dll.Error())
  fun := lib.Sym("say_nothing_at_all")
  _ = fun
  Println(dll.Error())
  Println("CallVoid")
  // fun.CallVoid() // Doesn't work. We need libffi
  Println("Prepare")
  call, status := dll.Prepare(fun, 0);
  Println("Call", status)
  call.Call()
  i := 0
  reps  := 1000000
  Println("Call Speed:", status)
  start := time.Nanoseconds() 
  for i = 0 ; i < reps ; i ++ {
    call.Call()
  }
  stop := time.Nanoseconds()
  delta := stop - start
  Println(delta, reps, delta / int64(reps))  
  lib.Close()
  Println(dll.Error())
}
