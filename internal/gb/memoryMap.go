package gb

const (
	CARTRIDGE_ROM_00_START uint16 = 0x0000 // 16KB from cartridge
	CARTRIDGE_ROM_00_END          = 0x3FFF
	CARTRIDGE_ROM_01_START        = 0x4000 // 16KB from cartridge via mapper
	CARTRIDGE_ROM_01_END          = 0x7FFF
	CARTRIDGE_HEADER_START        = 0x0100 // 80B
	CARTRIDGE_HEADER_END          = 0x014F

	VRAM_START          = 0x8000 // 8KB
	VRAM_END            = 0x9FFF
	CHARACTER_RAM_START = 0x8000 // 6KB
	CHARACTER_RAM_END   = 0x97FF
	BG_MAP_1_START      = 0x9800 // 1KB
	BG_MAP_1_END        = 0x9BFF
	BG_MAP_2_START      = 0x9C00 // 1KB
	BG_MAP_2_END        = 0x9FFF

	CARTRIDGE_RAM_START = 0xA000 // 8KB
	CARTRIDGE_RAM_END   = 0xBFFF

	INTERNAL_RAM_START = 0xC000 // 8KB
	INTERNAL_RAM_END   = 0xDFFF

	OAM_START = 0xFE00 // 160B
	OAM_END   = 0xFE9F

	IO_REGISTERS_START = 0xFF00 // 128B
	IO_REGISTERS_END   = 0xFF7F

	HRAM_BEGIN = 0xFF80 // 127B
	HRAM_END   = 0xFFFE

	INTERRUPT_ENABLE = 0xFFFF // 1B
)