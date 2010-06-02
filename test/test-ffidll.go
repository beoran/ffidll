package main

// import "testing"
import dll "ffidll"
import . "fmt"
import "time"

func TestLoad() {
  lib := dll.Open("../test/hellolib.so")
  Println(dll.Error())
  fun := lib.Sym("hello")
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
    Printf("%d ", i)
    call.Call() 
  }
  stop := time.Nanoseconds()
  Println(start - stop, reps, (start - stop) / int64(reps))  
  lib.Close()
  Println(dll.Error())
}

func main() {
  TestLoad();
}

