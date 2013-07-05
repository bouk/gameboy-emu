package memory

import "log"

const (
	P1 = 0xFF00
	SB = 0xFF01
	SC = 0xFF02

	DIV  = 0xFF04
	TIMA = 0xFF05
	TMA  = 0xFF06
	TAC  = 0xFF07

	IF   = 0xFF0F
	LCDC = 0xFF40
	STAT = 0xFF41

	SCY = 0xFF42
	SCX = 0xFF43

	LY  = 0xFF44
	LYC = 0xFF45
	DMA = 0xFF46
	BGP = 0xFF47

	OBP0 = 0xFF48
	OBP1 = 0xFF49

	WY = 0xFF4A
	WX = 0xFF4B

	KEY1 = 0xFF4D
	VBK  = 0xFF4F

	DISABLE_BIOS = 0xFF50

	HDMA1 = 0xFF51
	HDMA2 = 0xFF52
	HDMA3 = 0xFF53
	HDMA4 = 0xFF54
	HDMA5 = 0xFF55

	RP = 0xFF56

	BCPS = 0xFF68
	BCPD = 0xFF69

	OCPS = 0xFF6A
	OCPD = 0xFF6B

	SVBK = 0xFF70

	NR10 = 0xFF10
	NR11 = 0xFF11
	NR12 = 0xFF12
	NR13 = 0xFF13
	NR14 = 0xFF14
	NR21 = 0xFF16
	NR22 = 0xFF17
	NR23 = 0xFF18
	NR24 = 0xFF19
	NR30 = 0xFF1A
	NR31 = 0xFF1B
	NR32 = 0xFF1C
	NR33 = 0xFF1D
	NR34 = 0xFF1E
	NR41 = 0xFF20
	NR42 = 0xFF21
	NR43 = 0xFF22
	NR44 = 0xFF23
	NR50 = 0xFF24
	NR51 = 0xFF25
	NR52 = 0xFF26
)

type Mmc struct {
	Bios             []uint8
	BiosEnabled      bool
	MemoryBank       Memory
	InternalRam      []uint8
	UpperInternalRam []uint8
}

func NewMMC(memory Memory) *Mmc {
	m := new(Mmc)
	m.Bios = make([]uint8, 256)
	m.BiosEnabled = false

	m.MemoryBank = memory
	m.InternalRam = make([]uint8, 8*1024)
	m.UpperInternalRam = make([]uint8, 256)
	return m
}

func (m *Mmc) Read(addr uint16) uint8 {
	switch {
	case addr < 0x100 && m.BiosEnabled:
		return m.Bios[addr]
	case addr < 0x8000:
		return m.MemoryBank.Read(addr)
	case addr < 0xA000:
		// read VRAM
		return 0
	case addr < 0xC000:
		return m.MemoryBank.Read(addr)
	case addr < 0xE000:
		return m.InternalRam[addr-0xC000]
	case addr < 0xF000:
		return m.InternalRam[addr-0xE000]
	case addr < 0xFE00:
		log.Printf("Read from unused memory 0x%04X", addr)
		return 0
	case addr < 0xFF00:
		// Sprite attribute memory (OAM)
		return 0
	default:
		return m.UpperInternalRam[addr-0xFF00]
	}
}

func (m *Mmc) Write(addr uint16, value uint8) {
	switch {
	case addr < 0x8000:
		m.MemoryBank.Write(addr, value)
	case addr < 0xA000:
		// write video RAM
	case addr < 0xC000:
		m.MemoryBank.Write(addr, value)
	case addr < 0xE000:
		m.InternalRam[addr-0xC000] = value
	case addr < 0xF000:
		m.InternalRam[addr-0xE000] = value
	case addr < 0xFE00:
		log.Printf("Write to unused memory 0x%04X 0x%02X", addr, value)
	case addr < 0xFF00:
		// Sprite attribute memory (OAM)
	default:
		m.UpperInternalRam[addr-0xFF00] = value
		switch addr {
		case DISABLE_BIOS:
			m.BiosEnabled = false
			// something with io ports and stuff
		}
	}
}
