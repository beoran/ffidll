package ffidll
import "unsafe"


type TypeInC struct {
  size        uint32;
  alignment   uint32;
  kind        uint32;
  elements  []TypeInC;   
} 

type ABI int;

const ( 
  ABI_NONE    = ABI(iota)
  ABI_CDECL   = ABI(iota)
  ABI_STDCALL = ABI(iota)
  ABI_DEFAULT = ABI_CDECL
)

func (fun Function) PrecallABI(kinds [] TypeInC, abi ABI) {
  
}

type FuncInfo {
  function        Function
  resultkind      TypeInC
  argumentkind  []TypeInC
}

type FuncCallInfo {
  FuncInfo
  arguments      []byte;
}

 
/* Implemented in assembly. */ 
func PrepareCDECL(function unsafe.Pointer, kinds **TypeInC, nkinds int32);


/** Implemented in assembly. */
func CallCDECLVoid(function unsafe.Pointer) ();















