package gb

// We will define the following variables extracted from the opcode byte to be
// used in determining both the instruction to execute and the registers to
// operate on:
// 	bits     :  7  6  5  4  3  2  1  0
// 	variables:  (xx)  ( yyy )  ( zzz )
// 	                  (pp) (q)
//

// bit masks
const (
	MASK_XX  byte = 0b11 << 6
	MASK_YYY      = 0b111 << 3
	MASK_ZZZ      = 0b111
	MASK_PP       = 0b11 << 4
	MASK_Q        = 0b1 << 3
)

// xxBits returns bits 6-7 of the given byte
func xxBits(op byte) byte {
	return (op & MASK_XX) >> 6
}

// yyyBits returns bits 3-5 of the given byte
func yyyBits(op byte) byte {
	return (op & MASK_YYY) >> 3
}

// zzzBits returns bits 0-1 of the given byte
func zzzBits(op byte) byte {
	return op & MASK_ZZZ
}

// ppBits returns bits 4-5 of the given byte
func ppBits(op byte) byte {
	return (op & MASK_PP) >> 4
}

// qBits returns bit 3 of the given byte
func qBits(op byte) byte {
	return (op & MASK_Q) >> 3
}

// setupRegLookupTables sets up the lookup tables used while decoding opcodes.
// This should be called on a newly instantiated CPU instance.
func (cpu *CPU) setupRegLookupTables() {
	// TODO!!!
}
