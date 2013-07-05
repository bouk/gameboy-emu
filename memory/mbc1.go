package memory

import "log"

type Mbc1 struct {
	ROM             []uint8
	RAM             []uint8
	SelectedRomBank uint16
	SelectedRamBank uint16
	Mode            uint8
}

func NewMBC1(ROM, RAM []uint8) *Mbc1 {
	m := new(Mbc1)
	m.ROM = ROM
	m.RAM = RAM

	m.SelectedRomBank = 1
	return m
}

func (m *Mbc1) Read(addr uint16) uint8 {
	if addr <= 0x3FFF {
		return m.ROM[addr]
	} else if addr >= 0xA000 && addr <= 0xBFFF {
		return m.RAM[m.SelectedRamBank*0x2000+(addr-0xA000)]
	} else if addr < 0x8000 {
		return m.ROM[(m.SelectedRomBank*0x4000)+(addr-0x4000)]
	} else {
		log.Printf("Invalid read for Mbc1 0x%04X", addr)
		return 0
	}
}

func (m *Mbc1) Write(addr uint16, value uint8) {
	if addr <= 0x1FFF {
		// enable/disable RAM, NOP
	} else if addr >= 0x2000 && addr <= 0x3FFF {
		// ROM bank select
		m.SelectedRomBank = uint16(value & 0x1F)
		if m.SelectedRomBank == 0x00 || m.SelectedRomBank == 0x20 || m.SelectedRomBank == 0x40 || m.SelectedRomBank == 0x60 {
			m.SelectedRomBank++
		}
	} else if addr >= 0x4000 && addr <= 0x5FFF {
		// RAM bank select / 2 most significant bits of ROM bank selection
		if m.Mode == 0 {
			m.SelectedRamBank = uint16(value & 0x3)
		} else if m.Mode == 1 {
			m.SelectedRomBank = (uint16(value&0x3) << 5) | (m.SelectedRomBank & 0x1F)
		}
	} else if addr >= 0x6000 && addr <= 0x7FFF {
		// Mode selection
		if !(value&0x1 == 0x0 || value&0x1 == 0x1) {
			panic("Invalid mode for MBC1")
		}
		m.Mode = value & 0x1
	} else if addr >= 0xA000 && addr <= 0xBFFF {
		m.RAM[addr-0xA000] = value
	} else {
		log.Println("Invalid write for Mbc1 0x%04X 0x%02X", addr, value)
	}
}
