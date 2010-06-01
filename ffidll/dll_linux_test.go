package dll

import "testing"
import "ffidll/dll"
import . "fmt"

func TestLoad(t *testing.T) {
  lib := dll.Open("/usr/lib/libSDL-1.2.so.0")
  Println(dll.Error())
  fun := lib.Sym("SDL_GetError")
  _ = fun
  Println(dll.Error()) 
  // fun.CallVoid(); // Doesn't work yet. We need libffi
  lib.Close()
  Println(dll.Error())
  
}
