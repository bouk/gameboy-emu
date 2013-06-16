package lr35902

func (c *CPU) setupOpcodes() {
	c.opcodes[0x00] = func() {}
	c.opcodes[0x01] = func() {
		c.BC = c.NextWord()
	}
	c.opcodes[0x02] = func() {
		c.Memory.Write(c.BC, c.A)
	}
	c.opcodes[0x03] = func() {
		c.BC++
	}
	c.opcodes[0x04] = func() {
		setUpper(&c.BC, c.inc(getUpper(c.BC)))
	}
	c.opcodes[0x05] = func() {
		setUpper(&c.BC, c.dec(getUpper(c.BC)))
	}
	c.opcodes[0x06] = func() {
		setUpper(&c.BC, c.NextByte())
	}
	c.opcodes[0x07] = func() {
		c.Flags.C = c.A&(1<<7) == (1 << 7)
		c.Flags.H = false
		c.Flags.N = false
		c.Flags.Z = false
		c.A <<= 1
		if c.Flags.C {
			c.A |= 0x1
		}
	}
	c.opcodes[0x08] = func() {
		c.WriteWord(c.NextWord(), c.SP)
	}
	c.opcodes[0x09] = func() {
		c.addRegs(&c.HL, &c.BC)
	}
	c.opcodes[0x0A] = func() {
		c.A = c.Memory.Read(c.BC)
	}
	c.opcodes[0x0B] = func() {
		c.BC--
	}
	c.opcodes[0x0C] = func() {
		setLower(&c.BC, c.inc(getLower(c.BC)))
	}
	c.opcodes[0x0D] = func() {
		setLower(&c.BC, c.dec(getLower(c.BC)))
	}
	c.opcodes[0x0E] = func() {
		setLower(&c.BC, c.NextByte())
	}
	c.opcodes[0x0F] = func() {
		c.Flags.C = c.A&0x1 == 0x1
		c.Flags.H = false
		c.Flags.N = false
		c.Flags.Z = false
		c.A >>= 1
		if c.Flags.C {
			c.A |= (1 << 7)
		}
	}
	c.opcodes[0x10] = func() {
		c.Stopped = true
	}
	c.opcodes[0x11] = func() {
		c.DE = c.NextWord()
	}
	c.opcodes[0x12] = func() {
		c.Memory.Write(c.DE, c.A)
	}
	c.opcodes[0x13] = func() {
		c.DE++
	}
	c.opcodes[0x14] = func() {
		setUpper(&c.DE, c.inc(getUpper(c.DE)))
	}
	c.opcodes[0x15] = func() {
		setUpper(&c.DE, c.dec(getUpper(c.DE)))
	}
	c.opcodes[0x16] = func() {
		setUpper(&c.DE, c.NextByte())
	}
	c.opcodes[0x17] = func() {
		oldcarry := c.Flags.C
		c.Flags.C = c.A&(1<<7) == (1 << 7)
		c.Flags.H = false
		c.Flags.N = false
		c.Flags.Z = false
		c.A <<= 1
		if oldcarry {
			c.A |= 0x1
		}
	}
	c.opcodes[0x18] = func() {
		c.relativeJump(c.NextByte())
	}
	c.opcodes[0x19] = func() {
		c.addRegs(&c.HL, &c.DE)
	}
	c.opcodes[0x1A] = func() {
		c.A = c.Memory.Read(c.DE)
	}
	c.opcodes[0x1B] = func() {
		c.DE--
	}
	c.opcodes[0x1C] = func() {
		setLower(&c.DE, c.inc(getLower(c.DE)))
	}
	c.opcodes[0x1D] = func() {
		setLower(&c.DE, c.dec(getLower(c.DE)))
	}
	c.opcodes[0x1E] = func() {
		setLower(&c.DE, c.NextByte())
	}
	c.opcodes[0x1F] = func() {
		oldcarry := c.Flags.C
		c.Flags.C = c.A&0x1 == 0x1
		c.Flags.H = false
		c.Flags.N = false
		c.Flags.Z = false
		c.A >>= 1
		if oldcarry {
			c.A |= (1 << 7)
		}
	}
	c.opcodes[0x20] = func() {
		dist := c.NextByte()
		if !c.Flags.Z {
			c.relativeJump(dist)
		}
	}
	c.opcodes[0x21] = func() {
		c.HL = c.NextWord()
	}
	c.opcodes[0x22] = func() {
		c.Memory.Write(c.HL, c.A)
		c.HL++
	}
	c.opcodes[0x23] = func() {
		c.HL++
	}
	c.opcodes[0x24] = func() {
		setUpper(&c.HL, c.inc(getUpper(c.HL)))
	}
	c.opcodes[0x25] = func() {
		setUpper(&c.HL, c.dec(getUpper(c.HL)))
	}
	c.opcodes[0x26] = func() {
		setUpper(&c.HL, c.NextByte())
	}
	c.opcodes[0x27] = func() {
		c.Flags.C = false
		if c.A&0x0F > 9 {
			c.A += 0x06
		}
		if ((c.A & 0xF0) >> 4) > 9 {
			c.Flags.C = true
			c.A += 0x60
		}
		c.Flags.H = false
		c.Flags.Z = c.A == 0x00
	}
	c.opcodes[0x28] = func() {
		dist := c.NextByte()
		if c.Flags.Z {
			c.relativeJump(dist)
		}
	}
	c.opcodes[0x29] = func() {
		c.addRegs(&c.HL, &c.HL)
	}
	c.opcodes[0x2A] = func() {
		c.A = c.Memory.Read(c.HL)
		c.HL++
	}
	c.opcodes[0x2B] = func() {
		c.HL--
	}
	c.opcodes[0x2C] = func() {
		setLower(&c.HL, c.inc(getLower(c.HL)))
	}
	c.opcodes[0x2D] = func() {
		setLower(&c.HL, c.dec(getLower(c.HL)))
	}
	c.opcodes[0x2E] = func() {
		setLower(&c.HL, c.NextByte())
	}
	c.opcodes[0x2F] = func() {
		c.A = ^c.A
	}
	c.opcodes[0x30] = func() {
		dist := c.NextByte()
		if !c.Flags.C {
			c.relativeJump(dist)
		}
	}
	c.opcodes[0x31] = func() {
		c.SP = c.NextWord()
	}
	c.opcodes[0x32] = func() {
		c.Memory.Write(c.HL, c.A)
		c.HL--
	}
	c.opcodes[0x33] = func() {
		c.SP++
	}
	c.opcodes[0x34] = func() {
		c.Memory.Write(c.HL, c.inc(c.Memory.Read(c.HL)))
	}
	c.opcodes[0x35] = func() {
		c.Memory.Write(c.HL, c.dec(c.Memory.Read(c.HL)))
	}
	c.opcodes[0x36] = func() {
		c.Memory.Write(c.HL, c.NextByte())
	}
	c.opcodes[0x37] = func() {
		c.Flags.C = true
		c.Flags.H = false
		c.Flags.N = false
	}
	c.opcodes[0x38] = func() {
		dist := c.NextByte()
		if c.Flags.C {
			c.relativeJump(dist)
		}
	}
	c.opcodes[0x39] = func() {
		c.addRegs(&c.HL, &c.SP)
	}
	c.opcodes[0x3A] = func() {
		c.A = c.Memory.Read(c.HL)
		c.HL--
	}
	c.opcodes[0x3B] = func() {
		c.SP--
	}
	c.opcodes[0x3C] = func() {
		c.A = c.inc(c.A)
	}
	c.opcodes[0x3D] = func() {
		c.A = c.dec(c.A)
	}
	c.opcodes[0x3E] = func() {
		c.A = c.NextByte()
	}
	c.opcodes[0x3F] = func() {
		c.Flags.C = !c.Flags.C
		c.Flags.H = false
		c.Flags.N = false
	}
	c.opcodes[0x40] = func() {
		setUpper(&c.BC, getUpper(c.BC))
	}
	c.opcodes[0x41] = func() {
		setUpper(&c.BC, getLower(c.BC))
	}
	c.opcodes[0x42] = func() {
		setUpper(&c.BC, getUpper(c.DE))
	}
	c.opcodes[0x43] = func() {
		setUpper(&c.BC, getLower(c.DE))
	}
	c.opcodes[0x44] = func() {
		setUpper(&c.BC, getUpper(c.HL))
	}
	c.opcodes[0x45] = func() {
		setUpper(&c.BC, getLower(c.HL))
	}
	c.opcodes[0x46] = func() {
		setUpper(&c.BC, c.Memory.Read(c.HL))
	}
	c.opcodes[0x47] = func() {
		setUpper(&c.BC, c.A)
	}
	c.opcodes[0x48] = func() {
		setLower(&c.BC, getUpper(c.BC))
	}
	c.opcodes[0x49] = func() {
		setLower(&c.BC, getLower(c.BC))
	}
	c.opcodes[0x4A] = func() {
		setLower(&c.BC, getUpper(c.DE))
	}
	c.opcodes[0x4B] = func() {
		setLower(&c.BC, getLower(c.DE))
	}
	c.opcodes[0x4C] = func() {
		setLower(&c.BC, getUpper(c.HL))
	}
	c.opcodes[0x4D] = func() {
		setLower(&c.BC, getLower(c.HL))
	}
	c.opcodes[0x4E] = func() {
		setLower(&c.BC, c.Memory.Read(c.HL))
	}
	c.opcodes[0x4F] = func() {
		setLower(&c.BC, c.A)
	}
	c.opcodes[0x50] = func() {
		setUpper(&c.DE, getUpper(c.BC))
	}
	c.opcodes[0x51] = func() {
		setUpper(&c.DE, getLower(c.BC))
	}
	c.opcodes[0x52] = func() {
		setUpper(&c.DE, getUpper(c.DE))
	}
	c.opcodes[0x53] = func() {
		setUpper(&c.DE, getLower(c.DE))
	}
	c.opcodes[0x54] = func() {
		setUpper(&c.DE, getUpper(c.HL))
	}
	c.opcodes[0x55] = func() {
		setUpper(&c.DE, getLower(c.HL))
	}
	c.opcodes[0x56] = func() {
		setUpper(&c.DE, c.Memory.Read(c.HL))
	}
	c.opcodes[0x57] = func() {
		setUpper(&c.DE, c.A)
	}
	c.opcodes[0x58] = func() {
		setLower(&c.DE, getUpper(c.BC))
	}
	c.opcodes[0x59] = func() {
		setLower(&c.DE, getLower(c.BC))
	}
	c.opcodes[0x5A] = func() {
		setLower(&c.DE, getUpper(c.DE))
	}
	c.opcodes[0x5B] = func() {
		setLower(&c.DE, getLower(c.DE))
	}
	c.opcodes[0x5C] = func() {
		setLower(&c.DE, getUpper(c.HL))
	}
	c.opcodes[0x5D] = func() {
		setLower(&c.DE, getLower(c.HL))
	}
	c.opcodes[0x5E] = func() {
		setLower(&c.DE, c.Memory.Read(c.HL))
	}
	c.opcodes[0x5F] = func() {
		setLower(&c.DE, c.A)
	}
	c.opcodes[0x60] = func() {
		setUpper(&c.HL, getUpper(c.BC))
	}
	c.opcodes[0x61] = func() {
		setUpper(&c.HL, getLower(c.BC))
	}
	c.opcodes[0x62] = func() {
		setUpper(&c.HL, getUpper(c.DE))
	}
	c.opcodes[0x63] = func() {
		setUpper(&c.HL, getLower(c.DE))
	}
	c.opcodes[0x64] = func() {
		setUpper(&c.HL, getUpper(c.HL))
	}
	c.opcodes[0x65] = func() {
		setUpper(&c.HL, getLower(c.HL))
	}
	c.opcodes[0x66] = func() {
		setUpper(&c.HL, c.Memory.Read(c.HL))
	}
	c.opcodes[0x67] = func() {
		setUpper(&c.HL, c.A)
	}
	c.opcodes[0x68] = func() {
		setLower(&c.HL, getUpper(c.BC))
	}
	c.opcodes[0x69] = func() {
		setLower(&c.HL, getLower(c.BC))
	}
	c.opcodes[0x6A] = func() {
		setLower(&c.HL, getUpper(c.DE))
	}
	c.opcodes[0x6B] = func() {
		setLower(&c.HL, getLower(c.DE))
	}
	c.opcodes[0x6C] = func() {
		setLower(&c.HL, getUpper(c.HL))
	}
	c.opcodes[0x6D] = func() {
		setLower(&c.HL, getLower(c.HL))
	}
	c.opcodes[0x6E] = func() {
		setLower(&c.HL, c.Memory.Read(c.HL))
	}
	c.opcodes[0x6F] = func() {
		setLower(&c.HL, c.A)
	}
	c.opcodes[0x70] = func() {
		c.Memory.Write(c.HL, getUpper(c.BC))
	}
	c.opcodes[0x71] = func() {
		c.Memory.Write(c.HL, getLower(c.BC))
	}
	c.opcodes[0x72] = func() {
		c.Memory.Write(c.HL, getUpper(c.DE))
	}
	c.opcodes[0x73] = func() {
		c.Memory.Write(c.HL, getLower(c.DE))
	}
	c.opcodes[0x74] = func() {
		c.Memory.Write(c.HL, getUpper(c.HL))
	}
	c.opcodes[0x75] = func() {
		c.Memory.Write(c.HL, getLower(c.HL))
	}
	c.opcodes[0x76] = func() {
		c.Halted = true
	}
	c.opcodes[0x77] = func() {
		c.Memory.Write(c.HL, c.A)
	}
	c.opcodes[0x78] = func() {
		c.A = getUpper(c.BC)
	}
	c.opcodes[0x79] = func() {
		c.A = getLower(c.BC)
	}
	c.opcodes[0x7A] = func() {
		c.A = getUpper(c.DE)
	}
	c.opcodes[0x7B] = func() {
		c.A = getLower(c.DE)
	}
	c.opcodes[0x7C] = func() {
		c.A = getUpper(c.HL)
	}
	c.opcodes[0x7D] = func() {
		c.A = getLower(c.HL)
	}
	c.opcodes[0x7E] = func() {
		c.A = c.Memory.Read(c.HL)
	}
	c.opcodes[0x7F] = func() {
		c.A = c.A
	}
}

func setLower(mem *uint16, val uint8) {
	*mem = (*mem & 0xFF00) | uint16(val)
}

func setUpper(mem *uint16, val uint8) {
	*mem = (*mem & 0x00FF) | (uint16(val) << 8)
}

func getLower(val uint16) uint8 {
	return uint8(val & 0xFF)
}

func getUpper(val uint16) uint8 {
	return uint8((val & 0xFF00) >> 8)
}

func (c *CPU) inc(mem uint8) uint8 {
	c.Flags.H = (mem & 0xF) == 0xF
	mem++
	c.Flags.Z = mem == 0
	c.Flags.N = false
	return mem
}

func (c *CPU) dec(mem uint8) uint8 {
	mem--
	c.Flags.Z = mem == 0
	c.Flags.N = true
	c.Flags.H = (mem & 0xF) == 0xF
	return mem
}

func (c *CPU) addRegs(a *uint16, b *uint16) {
	temp := uint32(*a) + uint32(*b)
	c.Flags.N = false
	c.Flags.C = temp > 0xFFFF
	// TODO
	c.Flags.H = (((*a) & 0x0F00) == 0x0F00) && ((temp & 0x0F00) == 0x0000)
	*a = uint16(temp & 0xFFFF)
}

func (c *CPU) relativeJump(dist uint8) {
	jump := int8(dist)
	if jump >= 0 {
		c.PC += uint16(jump)
	} else {
		c.PC -= uint16(-jump)
	}
}

func halfOp(mem *uint16, high bool, op func(*uint8)) {
	btoi := func(b bool) uint8 {
		if b {
			return 1
		}
		return 0
	}
	var half uint8 = uint8((*mem >> (btoi(high) * 8)) & 0xFF)
	op(&half)
	*mem = uint16((*mem)&(0xFF<<(btoi(!high)*8))) | (uint16(half) << (btoi(high) * 8))
}
