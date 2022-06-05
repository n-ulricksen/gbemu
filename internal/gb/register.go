package gb

// register represents a 16-bit CPU register accessible by either its full 16
// bits, or as 2 separate 8 bit registers (hi/lo)
type register struct {
	value uint16
}

// val returns the 16-bit register value
func (r *register) val() uint16 {
	return r.value
}

// hi returns the register's upper byte
func (r *register) hi() byte {
	return byte(r.value >> 8)
}

// hi returns the register's lower byte
func (r *register) lo() byte {
	return byte(r.value & 0xFF)
}

// set sets the register's 16-bit value to the given val
func (r *register) set(val uint16) {
	r.value = val
}

// setHi sets the register's upper byte to the given val
func (r *register) setHi(val byte) {
	hi := uint16(val) << 8
	lo := r.value & 0xFF
	r.value = hi | lo
}

// setLo sets the register's lower byte to the given val
func (r *register) setLo(val byte) {
	hi := r.value & 0xFF00
	lo := uint16(val)
	r.value = hi | lo
}
