package memory

type Mbc5 struct {
	ROM             []uint8
	RAM             [0x20000]uint8
	SelectedRomBank uint16
	SelectedRamBank uint16
}

func NewMbc5(ROM []uint8) *Mbc5 {
	m := new(Mbc5)
	m.ROM = ROM

	return m
}

func (m *Mbc5) Read(addr uint16) uint8 {
	if addr <= 0x3FFF {
		return m.ROM[addr]
	} else if addr >= 0xA000 && addr <= 0xBFFF {
		return m.RAM[m.SelectedRamBank*0x2000+(addr-0xA000)]
	} else {
		if addr > 0x7FFF {
			panic("INVALID ROM ADDRESS")
		}
		if m.SelectedRomBank == 0 {
			return m.ROM[addr]
		} else {
			return m.ROM[(m.SelectedRomBank*0x4000)+(addr-0x4000)]
		}
	}
}

func (m *Mbc5) Write(addr uint16, value uint8) {
	if addr <= 0x1FFF {
		// enable/disable RAM/RTC registers, NOP
	} else if addr >= 0x2000 && addr <= 0x2FFF {
		// ROM bank select
		m.SelectedRomBank = uint16(value)
	} else if addr >= 0x3000 && addr <= 0x3FFF {
		m.SelectedRomBank = ((uint16(value&0x1) << 8) | (m.SelectedRomBank & 0xFF))
	} else if addr >= 0x4000 && addr <= 0x5FFF {
		// RAM bank select
		m.SelectedRamBank = uint16(value & 0xF)
	} else if addr >= 0xA000 && addr <= 0xBFFF {
		m.RAM[addr-0xA000] = value
	}
}
