package memory

import "log"

type Simple struct {
	ROM []uint8
	RAM []uint8
}

func NewSimple(ROM, RAM []uint8) *Simple {
	m := new(Simple)
	m.ROM = ROM
	m.RAM = RAM

	return m
}

func (m *Simple) Read(addr uint16) uint8 {
	if addr < 0x8000 {
		return m.ROM[addr]
	} else if addr >= 0xA000 && addr <= 0xBFFF {
		return m.RAM[addr-0xA000]
	} else {
		log.Printf("Invalid read for Simple 0x%04X", addr)
		return 0
	}
}

func (m *Simple) Write(addr uint16, value uint8) {
	if addr >= 0xA000 && addr <= 0xBFFF {
		m.RAM[addr-0xA000] = value
	} else {
		log.Println("Invalid write for Simple memory 0x%04X 0x%02X", addr, value)
	}
}
