package memory

import "log"

type Mbc2 struct {
	ROM             []uint8
	RAM             []uint8
	SelectedRomBank uint32
}

func NewMBC2(ROM, RAM []uint8) *Mbc2 {
	m := new(Mbc2)
	m.ROM = ROM
	m.RAM = RAM

	m.SelectedRomBank = 1
	return m
}

func (m *Mbc2) Read(addr uint16) uint8 {
	if addr <= 0x3FFF {
		return m.ROM[addr]
	} else if addr >= 0xA000 && addr <= 0xA1FF {
		return m.RAM[addr-0xA000]
	} else if addr < 0x8000 {
		return m.ROM[(m.SelectedRomBank*0x4000)+uint32(addr-0x4000)]
	} else {
		log.Printf("Invalid read for Mbc2 0x%04X", addr)
		return 0
	}
}

func (m *Mbc2) Write(addr uint16, value uint8) {
	if addr <= 0x1FFF {
		// enable/disable RAM, NOP
	} else if addr >= 0x2000 && addr <= 0x3FFF {
		// ROM bank select

		// The least significant bit of the upper address byte must be one to select a ROM bank
		if (addr>>8)&0x1 == 0x1 {
			m.SelectedRomBank = uint32(value & 0xF)
			if m.SelectedRomBank == 0x00 || m.SelectedRomBank == 0x20 || m.SelectedRomBank == 0x40 || m.SelectedRomBank == 0x60 {
				m.SelectedRomBank++
			}
		}
	} else if addr >= 0xA000 && addr <= 0xA1FF {
		m.RAM[addr-0xA000] = (value & 0xF)
	} else {
		log.Println("Invalid write for Mbc2 0x%04X 0x%02X", addr, value)
	}
}
