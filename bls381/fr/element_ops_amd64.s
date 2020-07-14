#include "textflag.h"
#include "funcdata.h"

TEXT ·_mulADXElement(SB), NOSPLIT, $0-24

	// the algorithm is described here
	// https://hackmd.io/@zkteam/modular_multiplication
	// however, to benefit from the ADCX and ADOX carry chains
	// we split the inner loops in 2:
	// for i=0 to N-1
	// 		for j=0 to N-1
	// 		    (A,t[j])  := t[j] + x[j]*y[i] + A
	// 		m := t[0]*q'[0] mod W
	// 		C,_ := t[0] + m*q[0]
	// 		for j=1 to N-1
	// 		    (C,t[j-1]) := t[j] + m*q[j] + C
	// 		t[N-1] = C + A
	
    CMPB ·supportAdx(SB), $0x0000000000000001
    JNE no_adx
    MOVQ x+8(FP), DI
    MOVQ y+16(FP), R8
    XORQ DX, DX
    MOVQ 0(R8), DX
    MULXQ 0(DI), CX, BX
    MULXQ 8(DI), AX, BP
    ADOXQ AX, BX
    MULXQ 16(DI), AX, SI
    ADOXQ AX, BP
    MULXQ 24(DI), AX, R9
    ADOXQ AX, SI
    // add the last carries to R9
    MOVQ $0x0000000000000000, DX
    ADCXQ DX, R9
    ADOXQ DX, R9
    MOVQ CX, DX
    MULXQ ·qElementInv0(SB), DX, AX                        // m := t[0]*q'[0] mod W
    XORQ AX, AX
    // C,_ := t[0] + m*q[0]
    MULXQ ·qElement+0(SB), AX, R10
    ADCXQ CX, AX
    MOVQ R10, CX
    // for j=1 to N-1
    //     (C,t[j-1]) := t[j] + m*q[j] + C
    ADCXQ BX, CX
    MULXQ ·qElement+8(SB), AX, BX
    ADOXQ AX, CX
    ADCXQ BP, BX
    MULXQ ·qElement+16(SB), AX, BP
    ADOXQ AX, BX
    ADCXQ SI, BP
    MULXQ ·qElement+24(SB), AX, SI
    ADOXQ AX, BP
    MOVQ $0x0000000000000000, AX
    ADCXQ AX, SI
    ADOXQ R9, SI
    XORQ DX, DX
    MOVQ 8(R8), DX
    MULXQ 0(DI), AX, R9
    ADOXQ AX, CX
    ADCXQ R9, BX
    MULXQ 8(DI), AX, R9
    ADOXQ AX, BX
    ADCXQ R9, BP
    MULXQ 16(DI), AX, R9
    ADOXQ AX, BP
    ADCXQ R9, SI
    MULXQ 24(DI), AX, R9
    ADOXQ AX, SI
    // add the last carries to R9
    MOVQ $0x0000000000000000, DX
    ADCXQ DX, R9
    ADOXQ DX, R9
    MOVQ CX, DX
    MULXQ ·qElementInv0(SB), DX, AX                        // m := t[0]*q'[0] mod W
    XORQ AX, AX
    // C,_ := t[0] + m*q[0]
    MULXQ ·qElement+0(SB), AX, R10
    ADCXQ CX, AX
    MOVQ R10, CX
    // for j=1 to N-1
    //     (C,t[j-1]) := t[j] + m*q[j] + C
    ADCXQ BX, CX
    MULXQ ·qElement+8(SB), AX, BX
    ADOXQ AX, CX
    ADCXQ BP, BX
    MULXQ ·qElement+16(SB), AX, BP
    ADOXQ AX, BX
    ADCXQ SI, BP
    MULXQ ·qElement+24(SB), AX, SI
    ADOXQ AX, BP
    MOVQ $0x0000000000000000, AX
    ADCXQ AX, SI
    ADOXQ R9, SI
    XORQ DX, DX
    MOVQ 16(R8), DX
    MULXQ 0(DI), AX, R9
    ADOXQ AX, CX
    ADCXQ R9, BX
    MULXQ 8(DI), AX, R9
    ADOXQ AX, BX
    ADCXQ R9, BP
    MULXQ 16(DI), AX, R9
    ADOXQ AX, BP
    ADCXQ R9, SI
    MULXQ 24(DI), AX, R9
    ADOXQ AX, SI
    // add the last carries to R9
    MOVQ $0x0000000000000000, DX
    ADCXQ DX, R9
    ADOXQ DX, R9
    MOVQ CX, DX
    MULXQ ·qElementInv0(SB), DX, AX                        // m := t[0]*q'[0] mod W
    XORQ AX, AX
    // C,_ := t[0] + m*q[0]
    MULXQ ·qElement+0(SB), AX, R10
    ADCXQ CX, AX
    MOVQ R10, CX
    // for j=1 to N-1
    //     (C,t[j-1]) := t[j] + m*q[j] + C
    ADCXQ BX, CX
    MULXQ ·qElement+8(SB), AX, BX
    ADOXQ AX, CX
    ADCXQ BP, BX
    MULXQ ·qElement+16(SB), AX, BP
    ADOXQ AX, BX
    ADCXQ SI, BP
    MULXQ ·qElement+24(SB), AX, SI
    ADOXQ AX, BP
    MOVQ $0x0000000000000000, AX
    ADCXQ AX, SI
    ADOXQ R9, SI
    XORQ DX, DX
    MOVQ 24(R8), DX
    MULXQ 0(DI), AX, R9
    ADOXQ AX, CX
    ADCXQ R9, BX
    MULXQ 8(DI), AX, R9
    ADOXQ AX, BX
    ADCXQ R9, BP
    MULXQ 16(DI), AX, R9
    ADOXQ AX, BP
    ADCXQ R9, SI
    MULXQ 24(DI), AX, R9
    ADOXQ AX, SI
    // add the last carries to R9
    MOVQ $0x0000000000000000, DX
    ADCXQ DX, R9
    ADOXQ DX, R9
    MOVQ CX, DX
    MULXQ ·qElementInv0(SB), DX, AX                        // m := t[0]*q'[0] mod W
    XORQ AX, AX
    // C,_ := t[0] + m*q[0]
    MULXQ ·qElement+0(SB), AX, R10
    ADCXQ CX, AX
    MOVQ R10, CX
    // for j=1 to N-1
    //     (C,t[j-1]) := t[j] + m*q[j] + C
    ADCXQ BX, CX
    MULXQ ·qElement+8(SB), AX, BX
    ADOXQ AX, CX
    ADCXQ BP, BX
    MULXQ ·qElement+16(SB), AX, BP
    ADOXQ AX, BX
    ADCXQ SI, BP
    MULXQ ·qElement+24(SB), AX, SI
    ADOXQ AX, BP
    MOVQ $0x0000000000000000, AX
    ADCXQ AX, SI
    ADOXQ R9, SI
    MOVQ res+0(FP), DI
    MOVQ CX, R11
    SUBQ ·qElement+0(SB), R11
    MOVQ BX, R12
    SBBQ ·qElement+8(SB), R12
    MOVQ BP, R13
    SBBQ ·qElement+16(SB), R13
    MOVQ SI, R14
    SBBQ ·qElement+24(SB), R14
    CMOVQCC R11, CX
    CMOVQCC R12, BX
    CMOVQCC R13, BP
    CMOVQCC R14, SI
    MOVQ CX, 0(DI)
    MOVQ BX, 8(DI)
    MOVQ BP, 16(DI)
    MOVQ SI, 24(DI)
    RET
no_adx:
    MOVQ x+8(FP), CX
    MOVQ y+16(FP), BX
    MOVQ 0(CX), AX
    MOVQ 0(BX), R10
    MULQ R10
    MOVQ AX, BP
    MOVQ DX, R11
    MOVQ ·qElementInv0(SB), R12
    IMULQ BP, R12
    MOVQ $0xffffffff00000001, AX
    MULQ R12
    ADDQ BP, AX
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R9
    MOVQ 8(CX), AX
    MULQ R10
    MOVQ R11, SI
    ADDQ AX, SI
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R11
    MOVQ $0x53bda402fffe5bfe, AX
    MULQ R12
    ADDQ SI, R9
    ADCQ $0x0000000000000000, DX
    ADDQ AX, R9
    ADCQ $0x0000000000000000, DX
    MOVQ R9, BP
    MOVQ DX, R9
    MOVQ 16(CX), AX
    MULQ R10
    MOVQ R11, DI
    ADDQ AX, DI
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R11
    MOVQ $0x3339d80809a1d805, AX
    MULQ R12
    ADDQ DI, R9
    ADCQ $0x0000000000000000, DX
    ADDQ AX, R9
    ADCQ $0x0000000000000000, DX
    MOVQ R9, SI
    MOVQ DX, R9
    MOVQ 24(CX), AX
    MULQ R10
    MOVQ R11, R8
    ADDQ AX, R8
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R11
    MOVQ $0x73eda753299d7d48, AX
    MULQ R12
    ADDQ R8, R9
    ADCQ $0x0000000000000000, DX
    ADDQ AX, R9
    ADCQ $0x0000000000000000, DX
    MOVQ R9, DI
    MOVQ DX, R9
    ADDQ R9, R11
    MOVQ R11, R8
    MOVQ 0(CX), AX
    MOVQ 8(BX), R10
    MULQ R10
    ADDQ AX, BP
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R11
    MOVQ ·qElementInv0(SB), R12
    IMULQ BP, R12
    MOVQ $0xffffffff00000001, AX
    MULQ R12
    ADDQ BP, AX
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R9
    MOVQ 8(CX), AX
    MULQ R10
    ADDQ R11, SI
    ADCQ $0x0000000000000000, DX
    ADDQ AX, SI
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R11
    MOVQ $0x53bda402fffe5bfe, AX
    MULQ R12
    ADDQ SI, R9
    ADCQ $0x0000000000000000, DX
    ADDQ AX, R9
    ADCQ $0x0000000000000000, DX
    MOVQ R9, BP
    MOVQ DX, R9
    MOVQ 16(CX), AX
    MULQ R10
    ADDQ R11, DI
    ADCQ $0x0000000000000000, DX
    ADDQ AX, DI
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R11
    MOVQ $0x3339d80809a1d805, AX
    MULQ R12
    ADDQ DI, R9
    ADCQ $0x0000000000000000, DX
    ADDQ AX, R9
    ADCQ $0x0000000000000000, DX
    MOVQ R9, SI
    MOVQ DX, R9
    MOVQ 24(CX), AX
    MULQ R10
    ADDQ R11, R8
    ADCQ $0x0000000000000000, DX
    ADDQ AX, R8
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R11
    MOVQ $0x73eda753299d7d48, AX
    MULQ R12
    ADDQ R8, R9
    ADCQ $0x0000000000000000, DX
    ADDQ AX, R9
    ADCQ $0x0000000000000000, DX
    MOVQ R9, DI
    MOVQ DX, R9
    ADDQ R9, R11
    MOVQ R11, R8
    MOVQ 0(CX), AX
    MOVQ 16(BX), R10
    MULQ R10
    ADDQ AX, BP
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R11
    MOVQ ·qElementInv0(SB), R12
    IMULQ BP, R12
    MOVQ $0xffffffff00000001, AX
    MULQ R12
    ADDQ BP, AX
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R9
    MOVQ 8(CX), AX
    MULQ R10
    ADDQ R11, SI
    ADCQ $0x0000000000000000, DX
    ADDQ AX, SI
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R11
    MOVQ $0x53bda402fffe5bfe, AX
    MULQ R12
    ADDQ SI, R9
    ADCQ $0x0000000000000000, DX
    ADDQ AX, R9
    ADCQ $0x0000000000000000, DX
    MOVQ R9, BP
    MOVQ DX, R9
    MOVQ 16(CX), AX
    MULQ R10
    ADDQ R11, DI
    ADCQ $0x0000000000000000, DX
    ADDQ AX, DI
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R11
    MOVQ $0x3339d80809a1d805, AX
    MULQ R12
    ADDQ DI, R9
    ADCQ $0x0000000000000000, DX
    ADDQ AX, R9
    ADCQ $0x0000000000000000, DX
    MOVQ R9, SI
    MOVQ DX, R9
    MOVQ 24(CX), AX
    MULQ R10
    ADDQ R11, R8
    ADCQ $0x0000000000000000, DX
    ADDQ AX, R8
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R11
    MOVQ $0x73eda753299d7d48, AX
    MULQ R12
    ADDQ R8, R9
    ADCQ $0x0000000000000000, DX
    ADDQ AX, R9
    ADCQ $0x0000000000000000, DX
    MOVQ R9, DI
    MOVQ DX, R9
    ADDQ R9, R11
    MOVQ R11, R8
    MOVQ 0(CX), AX
    MOVQ 24(BX), R10
    MULQ R10
    ADDQ AX, BP
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R11
    MOVQ ·qElementInv0(SB), R12
    IMULQ BP, R12
    MOVQ $0xffffffff00000001, AX
    MULQ R12
    ADDQ BP, AX
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R9
    MOVQ 8(CX), AX
    MULQ R10
    ADDQ R11, SI
    ADCQ $0x0000000000000000, DX
    ADDQ AX, SI
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R11
    MOVQ $0x53bda402fffe5bfe, AX
    MULQ R12
    ADDQ SI, R9
    ADCQ $0x0000000000000000, DX
    ADDQ AX, R9
    ADCQ $0x0000000000000000, DX
    MOVQ R9, BP
    MOVQ DX, R9
    MOVQ 16(CX), AX
    MULQ R10
    ADDQ R11, DI
    ADCQ $0x0000000000000000, DX
    ADDQ AX, DI
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R11
    MOVQ $0x3339d80809a1d805, AX
    MULQ R12
    ADDQ DI, R9
    ADCQ $0x0000000000000000, DX
    ADDQ AX, R9
    ADCQ $0x0000000000000000, DX
    MOVQ R9, SI
    MOVQ DX, R9
    MOVQ 24(CX), AX
    MULQ R10
    ADDQ R11, R8
    ADCQ $0x0000000000000000, DX
    ADDQ AX, R8
    ADCQ $0x0000000000000000, DX
    MOVQ DX, R11
    MOVQ $0x73eda753299d7d48, AX
    MULQ R12
    ADDQ R8, R9
    ADCQ $0x0000000000000000, DX
    ADDQ AX, R9
    ADCQ $0x0000000000000000, DX
    MOVQ R9, DI
    MOVQ DX, R9
    ADDQ R9, R11
    MOVQ R11, R8
    MOVQ res+0(FP), CX
    MOVQ BP, R13
    SUBQ ·qElement+0(SB), R13
    MOVQ SI, R14
    SBBQ ·qElement+8(SB), R14
    MOVQ DI, R15
    SBBQ ·qElement+16(SB), R15
    MOVQ R8, R9
    SBBQ ·qElement+24(SB), R9
    CMOVQCC R13, BP
    CMOVQCC R14, SI
    CMOVQCC R15, DI
    CMOVQCC R9, R8
    MOVQ BP, 0(CX)
    MOVQ SI, 8(CX)
    MOVQ DI, 16(CX)
    MOVQ R8, 24(CX)
    RET

TEXT ·subElement(SB), NOSPLIT, $0-24
    XORQ DX, DX
    MOVQ x+8(FP), DI
    MOVQ 0(DI), CX
    MOVQ 8(DI), BX
    MOVQ 16(DI), BP
    MOVQ 24(DI), SI
    MOVQ y+16(FP), DI
    SUBQ 0(DI), CX
    SBBQ 8(DI), BX
    SBBQ 16(DI), BP
    SBBQ 24(DI), SI
    MOVQ $0xffffffff00000001, R8
    CMOVQCC DX, R8
    MOVQ $0x53bda402fffe5bfe, R9
    CMOVQCC DX, R9
    MOVQ $0x3339d80809a1d805, R10
    CMOVQCC DX, R10
    MOVQ $0x73eda753299d7d48, R11
    CMOVQCC DX, R11
    MOVQ res+0(FP), DI
    ADDQ R8, CX
    MOVQ CX, 0(DI)
    ADCQ R9, BX
    MOVQ BX, 8(DI)
    ADCQ R10, BP
    MOVQ BP, 16(DI)
    ADCQ R11, SI
    MOVQ SI, 24(DI)
    RET

TEXT ·_fromMontADXElement(SB), $8-8
NO_LOCAL_POINTERS

	// the algorithm is described here
	// https://hackmd.io/@zkteam/modular_multiplication
	// when y = 1 we have: 
	// for i=0 to N-1
	// 		t[i] = x[i]
	// for i=0 to N-1
	// 		m := t[0]*q'[0] mod W
	// 		C,_ := t[0] + m*q[0]
	// 		for j=1 to N-1
	// 		    (C,t[j-1]) := t[j] + m*q[j] + C
	// 		t[N-1] = C
    CMPB ·supportAdx(SB), $0x0000000000000001
    JNE no_adx
    MOVQ res+0(FP), DI
    MOVQ 0(DI), CX
    MOVQ 8(DI), BX
    MOVQ 16(DI), BP
    MOVQ 24(DI), SI
    XORQ DX, DX
    MOVQ CX, DX
    MULXQ ·qElementInv0(SB), DX, AX                        // m := t[0]*q'[0] mod W
    XORQ AX, AX
    // C,_ := t[0] + m*q[0]
    MULXQ ·qElement+0(SB), AX, R8
    ADCXQ CX, AX
    MOVQ R8, CX
    // for j=1 to N-1
    //     (C,t[j-1]) := t[j] + m*q[j] + C
    ADCXQ BX, CX
    MULXQ ·qElement+8(SB), AX, BX
    ADOXQ AX, CX
    ADCXQ BP, BX
    MULXQ ·qElement+16(SB), AX, BP
    ADOXQ AX, BX
    ADCXQ SI, BP
    MULXQ ·qElement+24(SB), AX, SI
    ADOXQ AX, BP
    MOVQ $0x0000000000000000, AX
    ADCXQ AX, SI
    ADOXQ AX, SI
    XORQ DX, DX
    MOVQ CX, DX
    MULXQ ·qElementInv0(SB), DX, AX                        // m := t[0]*q'[0] mod W
    XORQ AX, AX
    // C,_ := t[0] + m*q[0]
    MULXQ ·qElement+0(SB), AX, R8
    ADCXQ CX, AX
    MOVQ R8, CX
    // for j=1 to N-1
    //     (C,t[j-1]) := t[j] + m*q[j] + C
    ADCXQ BX, CX
    MULXQ ·qElement+8(SB), AX, BX
    ADOXQ AX, CX
    ADCXQ BP, BX
    MULXQ ·qElement+16(SB), AX, BP
    ADOXQ AX, BX
    ADCXQ SI, BP
    MULXQ ·qElement+24(SB), AX, SI
    ADOXQ AX, BP
    MOVQ $0x0000000000000000, AX
    ADCXQ AX, SI
    ADOXQ AX, SI
    XORQ DX, DX
    MOVQ CX, DX
    MULXQ ·qElementInv0(SB), DX, AX                        // m := t[0]*q'[0] mod W
    XORQ AX, AX
    // C,_ := t[0] + m*q[0]
    MULXQ ·qElement+0(SB), AX, R8
    ADCXQ CX, AX
    MOVQ R8, CX
    // for j=1 to N-1
    //     (C,t[j-1]) := t[j] + m*q[j] + C
    ADCXQ BX, CX
    MULXQ ·qElement+8(SB), AX, BX
    ADOXQ AX, CX
    ADCXQ BP, BX
    MULXQ ·qElement+16(SB), AX, BP
    ADOXQ AX, BX
    ADCXQ SI, BP
    MULXQ ·qElement+24(SB), AX, SI
    ADOXQ AX, BP
    MOVQ $0x0000000000000000, AX
    ADCXQ AX, SI
    ADOXQ AX, SI
    XORQ DX, DX
    MOVQ CX, DX
    MULXQ ·qElementInv0(SB), DX, AX                        // m := t[0]*q'[0] mod W
    XORQ AX, AX
    // C,_ := t[0] + m*q[0]
    MULXQ ·qElement+0(SB), AX, R8
    ADCXQ CX, AX
    MOVQ R8, CX
    // for j=1 to N-1
    //     (C,t[j-1]) := t[j] + m*q[j] + C
    ADCXQ BX, CX
    MULXQ ·qElement+8(SB), AX, BX
    ADOXQ AX, CX
    ADCXQ BP, BX
    MULXQ ·qElement+16(SB), AX, BP
    ADOXQ AX, BX
    ADCXQ SI, BP
    MULXQ ·qElement+24(SB), AX, SI
    ADOXQ AX, BP
    MOVQ $0x0000000000000000, AX
    ADCXQ AX, SI
    ADOXQ AX, SI
    MOVQ CX, R9
    SUBQ ·qElement+0(SB), R9
    MOVQ BX, R10
    SBBQ ·qElement+8(SB), R10
    MOVQ BP, R11
    SBBQ ·qElement+16(SB), R11
    MOVQ SI, R12
    SBBQ ·qElement+24(SB), R12
    CMOVQCC R9, CX
    CMOVQCC R10, BX
    CMOVQCC R11, BP
    CMOVQCC R12, SI
    MOVQ CX, 0(DI)
    MOVQ BX, 8(DI)
    MOVQ BP, 16(DI)
    MOVQ SI, 24(DI)
    RET
no_adx:
    MOVQ res+0(FP), AX
    MOVQ AX, (SP)
CALL ·_fromMontGenericElement(SB)
    RET

TEXT ·reduceElement(SB), NOSPLIT, $0-8
    MOVQ res+0(FP), CX
    MOVQ 0(CX), BX
    MOVQ 8(CX), BP
    MOVQ 16(CX), SI
    MOVQ 24(CX), DI
    MOVQ BX, R8
    SUBQ ·qElement+0(SB), R8
    MOVQ BP, R9
    SBBQ ·qElement+8(SB), R9
    MOVQ SI, R10
    SBBQ ·qElement+16(SB), R10
    MOVQ DI, R11
    SBBQ ·qElement+24(SB), R11
    CMOVQCC R8, BX
    CMOVQCC R9, BP
    CMOVQCC R10, SI
    CMOVQCC R11, DI
    MOVQ BX, 0(CX)
    MOVQ BP, 8(CX)
    MOVQ SI, 16(CX)
    MOVQ DI, 24(CX)
    RET

TEXT ·addElement(SB), NOSPLIT, $0-24
    MOVQ x+8(FP), AX
    MOVQ 0(AX), CX
    MOVQ 8(AX), BX
    MOVQ 16(AX), BP
    MOVQ 24(AX), SI
    MOVQ y+16(FP), AX
    ADDQ 0(AX), CX
    ADCQ 8(AX), BX
    ADCQ 16(AX), BP
    ADCQ 24(AX), SI
    // note that we don't check for the carry here, as this code was generated assuming F.NoCarry condition is set
    // (see goff for more details)
    MOVQ res+0(FP), AX
    MOVQ CX, DI
    SUBQ ·qElement+0(SB), DI
    MOVQ BX, R8
    SBBQ ·qElement+8(SB), R8
    MOVQ BP, R9
    SBBQ ·qElement+16(SB), R9
    MOVQ SI, R10
    SBBQ ·qElement+24(SB), R10
    CMOVQCC DI, CX
    CMOVQCC R8, BX
    CMOVQCC R9, BP
    CMOVQCC R10, SI
    MOVQ CX, 0(AX)
    MOVQ BX, 8(AX)
    MOVQ BP, 16(AX)
    MOVQ SI, 24(AX)
    RET

TEXT ·doubleElement(SB), NOSPLIT, $0-16
    MOVQ x+8(FP), AX
    MOVQ 0(AX), CX
    MOVQ 8(AX), BX
    MOVQ 16(AX), BP
    MOVQ 24(AX), SI
    ADDQ CX, CX
    ADCQ BX, BX
    ADCQ BP, BP
    ADCQ SI, SI
    // note that we don't check for the carry here, as this code was generated assuming F.NoCarry condition is set
    // (see goff for more details)
    MOVQ res+0(FP), AX
    MOVQ CX, DI
    SUBQ ·qElement+0(SB), DI
    MOVQ BX, R8
    SBBQ ·qElement+8(SB), R8
    MOVQ BP, R9
    SBBQ ·qElement+16(SB), R9
    MOVQ SI, R10
    SBBQ ·qElement+24(SB), R10
    CMOVQCC DI, CX
    CMOVQCC R8, BX
    CMOVQCC R9, BP
    CMOVQCC R10, SI
    MOVQ CX, 0(AX)
    MOVQ BX, 8(AX)
    MOVQ BP, 16(AX)
    MOVQ SI, 24(AX)
    RET
