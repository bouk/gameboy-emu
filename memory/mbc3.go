package memory

// TODO: implement Real Time Clock

type Mbc3 struct {
	ROM             []uint8
	RAM             [0x8000]uint8
	SelectedRomBank uint16
	SelectedRamBank uint16
}

func NewMbc3(ROM []uint8) *Mbc3 {
	m := new(Mbc3)
	m.ROM = ROM
	m.SelectedRomBank = 1

	return m
}

func (m *Mbc3) Read(addr uint16) uint8 {
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

func (m *Mbc3) Write(addr uint16, value uint8) {
	if addr <= 0x1FFF {
		// enable/disable RAM/RTC registers, NOP
	} else if addr >= 0x2000 && addr <= 0x3FFF {
		// ROM bank select
		m.SelectedRomBank = uint16(value & 0x7F)
		if m.SelectedRomBank == 0 {
			m.SelectedRomBank++
		}
	} else if addr >= 0x4000 && addr <= 0x5FFF {
		// RAM bank select / RTC register select, RTC register not implemented
		m.SelectedRamBank = uint16(value & 0x3)
	} else if addr >= 0xA000 && addr <= 0xBFFF {
		m.RAM[addr-0xA000] = value
	}
}
