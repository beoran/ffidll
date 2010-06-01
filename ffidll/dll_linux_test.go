package ffidll

import "testing"
import dll "ffidll"
import . "fmt"

func TestLoad(t *testing.T) {
  lib := dll.Open("../test/hellolib.so")
  Println(dll.Error())
  fun := lib.Sym("hello")
  _ = fun
  Println(dll.Error())
  // fun.CallVoid() // Doesn't work. We need libffi
  call := dll.Prepare(fun, 0);
  call.Call();
  lib.Close()
  Println(dll.Error())
  
}
