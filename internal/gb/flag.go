package gb

// flag is used to represent a CPU flag (bit in lower AF register)
type flag byte

// Flag bits found in the lower 8 bits of the AF register
const (
	FLAG_C flag = 1 << (iota + 4) // Carry
	FLAG_H                        // Half Carry
	FLAG_N                        // Subraction
	FLAG_Z                        // Zero
)

// setFlag sets a bit in the flag register to the given val
func (cpu *CPU) setFlag(f flag, val bool) {
	if val {
		cpu.AF.setLo(cpu.AF.getLo() | byte(f))
	} else {
		cpu.AF.setLo(cpu.AF.getLo() &^ byte(f))
	}
}
