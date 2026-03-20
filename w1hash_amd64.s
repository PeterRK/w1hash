#include "textflag.h"

// Derived from clang -O3 codegen for stub.c on amd64 and translated to Go's Plan 9 assembler.

TEXT ·Hash(SB), NOSPLIT, $32-32
	MOVQ key_base+0(FP), AX
	MOVQ AX, 0(SP)
	MOVQ key_len+8(FP), AX
	MOVQ AX, 8(SP)
	MOVQ $0, 16(SP)
	CALL ·HashWithSeed(SB)
	MOVQ 24(SP), AX
	MOVQ AX, ret+24(FP)
	RET

TEXT ·Hash64(SB), NOSPLIT, $0-16
	MOVQ x+0(FP), AX
	MOVQ $0x8bb84b93962eacc9, CX
	XORQ CX, AX
	MOVQ $0x702daa6e740fb546, DX
	MULQ DX
	MOVQ $0x2d358dccaa6c78ad, SI
	XORQ AX, SI
	XORQ DX, CX
	MOVQ CX, AX
	MULQ SI
	XORQ DX, AX
	MOVQ AX, ret+8(FP)
	RET

TEXT ·HashWithSeed(SB), NOSPLIT, $0-40
	MOVQ key_base+0(FP), DI
	MOVQ key_len+8(FP), DX
	MOVQ DX, R13
	MOVQ DX, R14
	MOVQ seed+24(FP), R8

	MOVQ $0x8bb84b93962eacc9, CX
	MOVQ $0x2d358dccaa6c78a5, BX
	MOVQ R8, AX
	XORQ BX, AX
	XORQ CX, DX
	MULQ DX
	MOVQ DX, R8
	XORQ seed+24(FP), R8
	XORQ AX, R8

	MOVQ $0x4b33a62ed433d4a3, R15
	MOVQ $0x4d5a2da51de1aa47, R12

loop_top:
	CMPQ R14, $16
	JLS tail_dispatch

	CMPQ R14, $64
	JLS maybe_32

	MOVQ R8, R9
	MOVQ R8, R10
	MOVQ R8, R11

loop64:
	MOVQ (DI), DX
	XORQ 8(DI), R11
	XORQ BX, DX
	MOVQ R11, AX
	MULQ DX
	MOVQ DX, R11
	XORQ AX, R11

	MOVQ 16(DI), DX
	XORQ CX, DX
	XORQ 24(DI), R8
	MOVQ R8, AX
	MULQ DX
	MOVQ DX, R8
	XORQ AX, R8

	MOVQ 32(DI), DX
	XORQ 40(DI), R9
	XORQ R15, DX
	MOVQ R9, AX
	MULQ DX
	MOVQ DX, R9
	XORQ AX, R9

	MOVQ 48(DI), DX
	XORQ R12, DX
	XORQ 56(DI), R10
	MOVQ R10, AX
	MULQ DX
	MOVQ DX, R10
	XORQ AX, R10

	ADDQ $64, DI
	SUBQ $64, R14
	CMPQ R14, $64
	JHI loop64

	XORQ R11, R8
	XORQ R10, R9
	XORQ R9, R8

maybe_32:
	CMPQ R14, $32
	JLS maybe_16

	MOVQ (DI), DX
	XORQ BX, DX
	MOVQ 8(DI), AX
	XORQ R8, AX
	MULQ DX
	MOVQ AX, R9
	MOVQ DX, R10

	MOVQ 16(DI), DX
	XORQ CX, DX
	XORQ 24(DI), R8
	MOVQ R8, AX
	MULQ DX
	MOVQ DX, R8
	XORQ R10, R8
	XORQ R9, R8
	XORQ AX, R8

	ADDQ $32, DI
	SUBQ $32, R14

maybe_16:
	CMPQ R14, $16
	JLS tail_dispatch

	MOVQ (DI), DX
	XORQ BX, DX
	XORQ 8(DI), R8
	MOVQ R8, AX
	MULQ DX
	MOVQ DX, R8
	XORQ AX, R8

	ADDQ $16, DI
	SUBQ $16, R14
	JMP loop_top

tail_dispatch:
	CMPQ R14, $0
	JEQ tail0
	CMPQ R14, $1
	JEQ tail1
	CMPQ R14, $2
	JEQ tail2
	CMPQ R14, $3
	JEQ tail3
	CMPQ R14, $4
	JEQ tail4
	CMPQ R14, $5
	JEQ tail5
	CMPQ R14, $6
	JEQ tail6
	CMPQ R14, $7
	JEQ tail7
	CMPQ R14, $8
	JEQ tail8
	CMPQ R14, $9
	JEQ tail9
	CMPQ R14, $10
	JEQ tail10
	CMPQ R14, $11
	JEQ tail11
	CMPQ R14, $12
	JEQ tail12
	CMPQ R14, $13
	JEQ tail13
	CMPQ R14, $14
	JEQ tail14
	CMPQ R14, $15
	JEQ tail15
	JMP tail16

tail0:
	XORQ R14, R14
	XORQ AX, AX
	JMP final_mix

tail1:
	MOVBLZX (DI), R14
	XORQ AX, AX
	JMP final_mix

tail2:
	MOVWLZX (DI), R14
	XORQ AX, AX
	JMP final_mix

tail3:
	MOVQ DI, SI
	ANDQ $4095, SI
	CMPQ SI, $4092
	JHI tail3slow
	MOVL (DI), R14
	ANDQ $0xffffff, R14
	XORQ AX, AX
	JMP final_mix

tail3slow:
	MOVWLZX (DI), AX
	MOVBLZX 2(DI), R14
	SHLQ $16, R14
	ORQ AX, R14
	XORQ AX, AX
	JMP final_mix

tail4:
	MOVL (DI), R14
	XORQ AX, AX
	JMP final_mix

tail5:
	MOVQ DI, SI
	ANDQ $4095, SI
	CMPQ SI, $4088
	JHI tail5slow
	MOVQ $0xffffffffff, R14
	ANDQ (DI), R14
	XORQ AX, AX
	JMP final_mix

tail5slow:
	MOVL (DI), AX
	MOVBLZX 4(DI), R14
	SHLQ $32, R14
	ORQ AX, R14
	XORQ AX, AX
	JMP final_mix

tail6:
	MOVQ DI, SI
	ANDQ $4095, SI
	CMPQ SI, $4088
	JHI tail6slow
	MOVQ $0xffffffffffff, R14
	ANDQ (DI), R14
	XORQ AX, AX
	JMP final_mix

tail6slow:
	MOVL (DI), AX
	MOVWLZX 4(DI), R14
	SHLQ $32, R14
	ORQ AX, R14
	XORQ AX, AX
	JMP final_mix

tail7:
	MOVQ DI, SI
	ANDQ $4095, SI
	CMPQ SI, $4088
	JHI tail7slow
	MOVQ $0xffffffffffffff, R14
	ANDQ (DI), R14
	XORQ AX, AX
	JMP final_mix

tail7slow:
	MOVL (DI), AX
	MOVWLZX 4(DI), DX
	SHLQ $32, DX
	ORQ DX, AX
	MOVBLZX 6(DI), R14
	SHLQ $48, R14
	ORQ AX, R14
	XORQ AX, AX
	JMP final_mix

tail8:
	MOVQ (DI), R14
	XORQ AX, AX
	JMP final_mix

tail9:
	MOVQ (DI), R14
	MOVBLZX 8(DI), AX
	JMP final_mix

tail10:
	MOVQ (DI), R14
	MOVWLZX 8(DI), AX
	JMP final_mix

tail11:
	MOVQ (DI), R14
	LEAQ 8(DI), SI
	MOVQ SI, DX
	ANDQ $4095, DX
	CMPQ DX, $4092
	JHI tail11slow
	MOVL (SI), AX
	ANDQ $0xffffff, AX
	JMP final_mix

tail11slow:
	MOVWLZX 8(DI), DX
	MOVBLZX 10(DI), AX
	SHLQ $16, AX
	ORQ DX, AX
	JMP final_mix

tail12:
	MOVQ (DI), R14
	MOVL 8(DI), AX
	JMP final_mix

tail13:
	MOVQ (DI), R14
	LEAQ 8(DI), SI
	MOVQ SI, DX
	ANDQ $4095, DX
	CMPQ DX, $4088
	JHI tail13slow
	MOVQ $0xffffffffff, AX
	ANDQ (SI), AX
	JMP final_mix

tail13slow:
	MOVL 8(DI), DX
	MOVBLZX 12(DI), AX
	SHLQ $32, AX
	ORQ DX, AX
	JMP final_mix

tail14:
	MOVQ (DI), R14
	LEAQ 8(DI), SI
	MOVQ SI, DX
	ANDQ $4095, DX
	CMPQ DX, $4088
	JHI tail14slow
	MOVQ $0xffffffffffff, AX
	ANDQ (SI), AX
	JMP final_mix

tail14slow:
	MOVL 8(DI), DX
	MOVWLZX 12(DI), AX
	SHLQ $32, AX
	ORQ DX, AX
	JMP final_mix

tail15:
	MOVQ (DI), R14
	LEAQ 8(DI), SI
	MOVQ SI, DX
	ANDQ $4095, DX
	CMPQ DX, $4088
	JHI tail15slow
	MOVQ $0xffffffffffffff, AX
	ANDQ (SI), AX
	JMP final_mix

tail15slow:
	MOVL 8(DI), AX
	MOVWLZX 12(DI), DX
	SHLQ $32, DX
	ORQ DX, AX
	MOVBLZX 14(DI), DX
	SHLQ $48, DX
	ORQ DX, AX
	JMP final_mix

tail16:
	MOVQ (DI), R14
	MOVQ 8(DI), AX

final_mix:
	XORQ CX, R14
	XORQ R8, AX
	MULQ R14
	MOVQ R13, SI
	XORQ BX, SI
	XORQ AX, SI
	MOVQ CX, AX
	XORQ DX, AX
	MULQ SI
	XORQ DX, AX
	MOVQ AX, ret+32(FP)
	RET
