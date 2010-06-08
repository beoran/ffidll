
/**
* Borrowed partially from libcgo and syscall package.
*/

/*
 * ·CallCDECLVoid(unsafe.Pointer function) ()
 *
 * Calling into the 8c tool chain, where all registers are caller save.
 * Called from standard x86 ABI, where %ebp, %ebx, %esi,
 * and %edi are callee-save, so they must be saved explicitly.
 */ 
TEXT  ·CallCDECLVoid(SB),7,$0
  PUSHL BP          // save all needed registers on the stack
  MOVL  SP, BP
  PUSHL BX
  PUSHL SI 
  PUSHL DI

  MOVL  (BP), AX   
  /* Call function, the pointer to which is on top of the stack  */
  CALL  *AX

  POPL  DI          // restore all needed registers from the stack
  POPL  SI
  POPL  BX
  POPL  DX
  RET
