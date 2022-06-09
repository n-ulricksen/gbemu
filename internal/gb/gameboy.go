package gb

import (
	"log"
	"os"
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
	logger    *log.Logger

	debugMode bool

	disassembly []string
}

// New creates and returns a GameBoy instance, ready to start by running its
// 'Start' method
func New(romPath string, debug bool) *GameBoy {
	logger := log.Default()
	logger.SetFlags(0)

	gb := &GameBoy{
		logger:    log.Default(),
		debugMode: debug,
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
		gb.printDebug()
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

func (gb *GameBoy) initPowerUpSequence() {
	gb.Cpu.PC = ENTRY_POINT
	gb.isRunning = true
}
func (gb *GameBoy) printDebug() {
}