package gb

// u16 creates and returns a unsigned word from the given LSB and MSB
func u16(lsb, msb byte) uint16 {
	hi := uint16(msb) << 8
	lo := uint16(lsb)
	return hi | lo
}

// halfCarryOccurs returns whether a half carry occurs when adding together the
// two given bytes
func halfCarryOccurs(b1, b2 byte) bool {
	n1 := b1 & 0xf
	n2 := b2 & 0xf

	return ((n1 + n2) & 0x10) == 0x10
}

// halfCarryOccurs16 returns whether a carry occurs in the high byte, from bit
// 11 to 12, when adding together the two given words. This function is used
// when determining whether to set the half-carry flags for opcodes 0x09, 0x19,
// 0x29, 0x39.
//
// ref: (https://newbedev.com/game-boy-half-carry-flag-and-16-bit-instructions-especially-opcode-0xe8)
//	"ADD HL, rr: H from bit 11, C from bit 15 (flags from high byte op)"
func halfCarryOccurs16(w1, w2 uint16) bool {
	b1 := byte(w1 >> 8)
	b2 := byte(w2 >> 8)

	return halfCarryOccurs(b1, b2)
}
