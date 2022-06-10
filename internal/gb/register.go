package gb

// register represents a 16-bit CPU register accessible by either its full 16
// bits, or as 2 separate 8 bit registers (hi/lo)
type register struct {
	loReg register8Bit
	hiReg register8Bit
	name  string
}

// register8Bit is an 8-bit CPU register
type register8Bit struct {
	value byte
	name  string
}

// newReg8Bit returns a new 8 bit register with the given name
func newReg8Bit(name string) register8Bit {
	return register8Bit{name: name}
}

// get returns the 16-bit register value
func (r *register) get() uint16 {
	return (uint16(r.hiReg.value) << 8) | uint16(r.loReg.value)
}

// hi returns the register's upper byte
func (r *register) getHi() byte {
	return r.hiReg.value
}

// hi returns the register's lower byte
func (r *register) getLo() byte {
	return r.loReg.value
}

// set sets the register's 16-bit value to the given val
func (r *register) set(val uint16) {
	r.loReg.value = byte(val)
	r.hiReg.value = byte(val >> 8)
}

// setHi sets the register's upper byte to the given val
func (r *register) setHi(val byte) {
	r.hiReg.value = val
}

// setLo sets the register's lower byte to the given val
func (r *register) setLo(val byte) {
	r.loReg.value = val
}
