package gb

// CPU represents a Sharp SM83 CPU core
type CPU struct {
	AF register // Accumulator and flags
	BC register
	DE register
	HL register

	SP register // Stack pointer
	PC register // Program counter
}

// incPC increments the CPU's program counter register
func (cpu *CPU) incPC() {
	cpu.PC.value++
}
