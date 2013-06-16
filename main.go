package main

import (
	"github.com/boukevanderbijl/gameboy-emu/lr35902"
)

type mem [0xFFFF]uint8

func (m *mem) Read(pos uint16) uint8 {
	return (*m)[pos]
}

func (m *mem) Write(pos uint16, val uint8) {
	(*m)[pos] = val
}

func main() {
	m := new(mem)
	m[0] = 0x3E
	m[1] = 0x00
	m[2] = 0x2F
	m[4] = 0x10
	cpu := lr35902.NewCPU(m)
	for !cpu.Stopped {
		cpu.Step()
	}
	cpu.DumpState()
}
