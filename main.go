package main

import (
	"github.com/boukevanderbijl/gameboy-emu/lr35902"
	"os"
)

func main() {
	m := lr35902.RAM(make([]uint8, 0xFFFF))

	file, _ := os.Open("DMG_ROM.bin")
	file.Read(m[:255])

	cpu := lr35902.NewCPU(&m)

	for !cpu.Stopped && cpu.PC <= 0xFF {
		cpu.Step()
	}
	cpu.DumpState()
}
