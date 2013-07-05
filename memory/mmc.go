package memory

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
	m.UpperInternalRam = make([]uint8, 128)
	return m
}

func (m *Mmc) Read(addr uint16) uint8 {
	switch {
	case addr <= 0xFF && m.BiosEnabled:
		return m.Bios[addr]
	case addr >= 0xC000 && addr < 0xE000:
		return m.InternalRam[addr-0xC000]
	case addr >= 0xE000 && addr < 0xF000:
		return m.InternalRam[addr-0xE000]
	case addr >= 0xFF80 && addr < 0xFFFF:
		return m.UpperInternalRam[addr-0xFF80]
	default:
		return m.MemoryBank.Read(addr)
	}

}

func (m *Mmc) Write(addr uint16, value uint8) {
	switch {
	case addr == 0xFF50:
		m.BiosEnabled = false
	case addr >= 0xC000 && addr < 0xE000:
		m.InternalRam[addr-0xC000] = value
	case addr >= 0xE000 && addr < 0xF000:
		m.InternalRam[addr-0xE000] = value
	case addr >= 0xFF80 && addr < 0xFFFF:
		m.UpperInternalRam[addr-0xFF80] = value
	default:
		m.MemoryBank.Write(addr, value)
	}
}
