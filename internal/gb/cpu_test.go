package gb

import "testing"

// TestCpuInstrs runs the third-party 'cpu_instrs.gb' test package in an
// emulated SM83 CPU
func TestCpuInstrs(t *testing.T) {
	filepath := "../../test/cpu_instrs/cpu_instrs.gb"
	gb := New(filepath, true)
	gb.Start()
}
