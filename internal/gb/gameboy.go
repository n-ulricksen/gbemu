package gb

import (
	"fmt"
	"log"
	"os"
	"text/tabwriter"
)

// GameBoy represents a software emulation of Nintendo's Game Boy handheld game
// console. This struct will also be used as the system bus, responsible for
// all communication between attached devices.
type GameBoy struct {
	// Sharp SM83 CPU
	Cpu *CPU

	// Cartridge ROM
	CartRom []byte

	// Internal variables
	isRunning bool

	// Logging
	logger *log.Logger
	tw     *tabwriter.Writer

	debugMode  bool
	debugState *debugState

	disassembly []string
}

// debugState stores data relevant to CPU debugging
type debugState struct {
	iosb byte // IO serial bus
	iosc byte // IO serial control
}

// New creates and returns a GameBoy instance, ready to start by running its
// 'Start' method
func New(romPath string, debug bool) *GameBoy {
	logger := log.Default()
	logger.SetFlags(0)
	tw := tabwriter.NewWriter(logger.Writer(), 12, 4, 2, ' ', 0)

	gb := &GameBoy{
		logger:     logger,
		tw:         tw,
		debugMode:  debug,
		debugState: new(debugState),
	}
	cpu := newCPU()
	gb.attachCPU(cpu)

	// TODO: load ROM, insert cartridge
	gb.insertCartridge(romPath)
	gb.disassemble(0x0000, 0xFFFF)

	return gb
}

// Start "powers on" the GameBoy console. This will run the power-up sequence,
// and begin CPU execution.
func (gb *GameBoy) Start() {
	gb.initPowerUpSequence()

	for gb.isRunning {
		gb.Cpu.execNextInst()

		// DEBUG
		// if gb.debugMode {
		// 	gb.logIOSerialTransfers()
		// }
	}
}

// log logs a message to the GameBoy's logger
func (gb *GameBoy) log(msg string) {
	fmt.Fprintln(gb.tw, msg)
	gb.tw.Flush()
}

// logIOSerialTransfers logs write to memory at `IO_SB` or `IO_SC`
func (gb *GameBoy) logIOSerialTransfers() {
	iosb := gb.CartRom[IO_SB]
	iosc := gb.CartRom[IO_SC]

	if gb.debugState.iosb != iosb {
		gb.debugState.iosb = iosb
		gb.logger.Printf("[DEBUG_IO] IO_SB (%#04x): %#02x\n", IO_SB, iosb)
	}
	if gb.debugState.iosc != iosc {
		gb.debugState.iosc = iosc
		gb.logger.Printf("[DEBUG_IO] IO_SC (%#04x): %#02x\n", IO_SC, iosc)
	}
}

// TODO: accept a cartridge type, extract loading the file from this method
func (gb *GameBoy) insertCartridge(romPath string) error {
	cartRom, err := os.ReadFile(romPath)
	if err != nil {
		return errGameFilepathNotFound
	}

	// XXX: for now, copy the entire file into memory
	gb.CartRom = make([]byte, len(cartRom))
	for i, data := range cartRom {
		gb.CartRom[i] = data
	}
	return nil
}

// initPowerUpSequence performs the DMG boot sequence, leaving the CPU ready to
// begin executing the loaded game ROM
// 	reference: https://gbdev.io/pandocs/Power_Up_Sequence.html
func (gb *GameBoy) initPowerUpSequence() {
	// CPU
	gb.Cpu.AF.setHi(0x01)
	gb.Cpu.AF.setLo(0b10000000) // TODO: calculate bits 4 and 5 (see ref.)
	gb.Cpu.BC.set(0x0013)
	gb.Cpu.DE.set(0x00D8)
	gb.Cpu.HL.set(0x014D)
	gb.Cpu.PC = ENTRY_POINT
	gb.Cpu.SP = 0xFFFE

	// Hardware registers
	for addr, val := range hardwareRegisterInit {
		gb.CartRom[addr] = val
	}

	gb.isRunning = true
}
