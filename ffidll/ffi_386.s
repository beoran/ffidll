
/**
* ffi Foreign Function Interface for CGO.
* Borrowed partially from libffi, and from libcgo and syscall package,
* all under a BSD style license.
*/


/* Macro that prints an integer. Uses the go runtime. Trashes BX. */
#define printi(reg)   PUSHL   $0 ;                    \
                      PUSHL   reg ;                   \
                      CALL    runtime·printint(SB) ;  \
                      CALL    runtime·printnl(SB) ;   \
                      POPL    BX ;                    \
                      POPL    BX
                     
                     
/* Macro that saves all needed registers to the stack before a call 
   and copîes the original SP to BP. 
*/
#define saveregs()    PUSHL   BP       ;     \
                      MOVL    SP, BP   ;     \
                      PUSHL   BX       ;     \
                      PUSHL   SI       ;     \
                      PUSHL   DI       
  

/* Macro that restores all needed registers from the stack after a call.  
*/
#define loadregs()    POPL    DI       ;     \
                      POPL    SI       ;     \
                      POPL    BX       ;     \
                      POPL    DX  


/*
 * ·CallCDECLVoid(unsafe.Pointer function) ()
 *
 * Calling into the x86 CDECL ABI, where %ebp, %ebx, %esi,
 * and %edi are callee-save, so they must be saved explicitly.
 *
 */ 
TEXT  ·CallCDECLVoid(SB),7,$0
  // save all needed registers on the stack
  saveregs()
  
  // Call function, the pointer to which was on top of the stack 
  MOVL    (BP), AX       
  CALL    *AX
  
  // restore all needed registers from the stack  
  loadregs()
  RET


/* Does nothing for now. */
TEXT  ·PrepareCDECL(SB),7,$0
  RET
  
  
  
  