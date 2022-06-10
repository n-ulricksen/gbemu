package gb

// u16 creates and returns a unsigned word from the given LSB and MSB
func u16(lsb, msb byte) uint16 {
	hi := uint16(msb) << 8
	lo := uint16(lsb)
	return hi | lo
}
