package lr35902

import (
	"fmt"
)

type Memory interface {
	Read(address uint16) uint8
	Write(address uint16, value uint8)
}

type CPU struct {
	PC uint16
	SP uint16

	A          uint8
	BC, DE, HL uint16

	Stopped, Halted bool

	IME bool

	Flags struct {
		Z, N, H, C bool
	}

	Memory

	// in HZ
	ClockSpeed uint

	// Variable that's filled in by the opcode functions to make sure the processor runs on the correct speed.
	// 1 machine cycle = 4 clock cycles
	// is set to 4 and overwritten on a per-function basis
	clockCycles uint

	opcodes   [256]func()
	cbOpcodes [256]func()
}

func NewCPU(m Memory) *CPU {
	c := new(CPU)
	c.Memory = m
	c.ClockSpeed = 4194304
	c.setupOpcodes()
	return c
}

func (c *CPU) Step() {
	c.clockCycles = 4
	opcode := c.NextByte()
	c.opcodes[opcode]()
}

func (c *CPU) NextByte() uint8 {
	b := c.Memory.Read(c.PC)
	c.PC++
	return b
}

func (c *CPU) NextWord() uint16 {
	return uint16(c.NextByte()) | (uint16(c.NextByte()) << 8)
}

func (c *CPU) PushWord(val uint16) {
	c.SP -= 2
	c.WriteWord(c.SP, val)
}

func (c *CPU) PopWord() uint16 {
	val := uint16(c.Memory.Read(c.SP)) | (uint16(c.Memory.Read(c.SP+1)) << 8)
	c.SP += 2
	return val
}

func (c *CPU) WriteWord(pos uint16, val uint16) {
	c.Memory.Write(pos, uint8(val&0xFF))
	c.Memory.Write(pos+1, uint8((val>>8)&0xFF))
}

func (c *CPU) GetAF() uint16 {
	val := (uint16(c.A) << 8)
	if c.Flags.Z {
		val |= (1 << 7)
	}
	if c.Flags.N {
		val |= (1 << 6)
	}
	if c.Flags.H {
		val |= (1 << 5)
	}
	if c.Flags.C {
		val |= (1 << 4)
	}
	return val
}

func (c *CPU) DumpState() {
	fmt.Println("AF   BC   DE   HL   SP   PC")
	fmt.Printf("%04X %04X %04X %04X %04X %04X\n", c.GetAF(), c.BC, c.DE, c.HL, c.SP, c.PC)
}
