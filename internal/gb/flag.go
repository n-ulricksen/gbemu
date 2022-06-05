package gb

// flag is used to represent a CPU flag (bit in lower AF register)
type flag byte

// Flag bits found in the lower 8 bits of the AF register
const (
	c flag = 1 << (iota + 4) // Carry
	h                        // Half Carry
	n                        // Subraction
	z                        // Zero
)

// setFlag sets a bit in the flag register to the given val
func (r *register) setFlag(f flag, val bool) {
	if val {
		r.value |= uint16(f)
	} else {
		r.value &^= uint16(f)
	}
}
