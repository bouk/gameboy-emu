package main

import (
	"fmt"
	"github.com/boukevanderbijl/gameboy-emu/lr35902"
	"os"
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
	file, _ := os.Open("DMG_ROM.bin")
	file.Read(m[:255])
	fmt.Println(m[:255])
	cpu := lr35902.NewCPU(m)
	for !cpu.Stopped && cpu.PC <= 0xFF {
		cpu.Step()
		cpu.DumpState()
	}
}
