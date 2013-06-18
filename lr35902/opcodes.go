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
		c.Flags.C = (c.A & (1 << 7)) == (1 << 7)
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
		c.Flags.C = (c.A & 0x1) == 0x1
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
		c.A = c.rl(c.A)
		c.Flags.Z = false
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
	c.opcodes[0x80] = func() {
		c.add(getUpper(c.BC))
	}
	c.opcodes[0x81] = func() {
		c.add(getLower(c.BC))
	}
	c.opcodes[0x82] = func() {
		c.add(getUpper(c.DE))
	}
	c.opcodes[0x83] = func() {
		c.add(getLower(c.DE))
	}
	c.opcodes[0x84] = func() {
		c.add(getUpper(c.HL))
	}
	c.opcodes[0x85] = func() {
		c.add(getLower(c.HL))
	}
	c.opcodes[0x86] = func() {
		c.add(c.Memory.Read(c.HL))
	}
	c.opcodes[0x87] = func() {
		c.add(c.A)
	}
	c.opcodes[0x88] = func() {
		c.adc(getUpper(c.BC))
	}
	c.opcodes[0x89] = func() {
		c.adc(getLower(c.BC))
	}
	c.opcodes[0x8A] = func() {
		c.adc(getUpper(c.DE))
	}
	c.opcodes[0x8B] = func() {
		c.adc(getLower(c.DE))
	}
	c.opcodes[0x8C] = func() {
		c.adc(getUpper(c.HL))
	}
	c.opcodes[0x8D] = func() {
		c.adc(getLower(c.HL))
	}
	c.opcodes[0x8E] = func() {
		c.adc(c.Memory.Read(c.HL))
	}
	c.opcodes[0x8F] = func() {
		c.adc(c.A)
	}
	c.opcodes[0x90] = func() {
		c.sub(getUpper(c.BC))
	}
	c.opcodes[0x91] = func() {
		c.sub(getLower(c.BC))
	}
	c.opcodes[0x92] = func() {
		c.sub(getUpper(c.DE))
	}
	c.opcodes[0x93] = func() {
		c.sub(getLower(c.DE))
	}
	c.opcodes[0x94] = func() {
		c.sub(getUpper(c.HL))
	}
	c.opcodes[0x95] = func() {
		c.sub(getLower(c.HL))
	}
	c.opcodes[0x96] = func() {
		c.sub(c.Memory.Read(c.HL))
	}
	c.opcodes[0x97] = func() {
		c.sub(c.A)
	}
	c.opcodes[0x98] = func() {
		c.sbc(getUpper(c.BC))
	}
	c.opcodes[0x99] = func() {
		c.sbc(getLower(c.BC))
	}
	c.opcodes[0x9A] = func() {
		c.sbc(getUpper(c.DE))
	}
	c.opcodes[0x9B] = func() {
		c.sbc(getLower(c.DE))
	}
	c.opcodes[0x9C] = func() {
		c.sbc(getUpper(c.HL))
	}
	c.opcodes[0x9D] = func() {
		c.sbc(getLower(c.HL))
	}
	c.opcodes[0x9E] = func() {
		c.sbc(c.Memory.Read(c.HL))
	}
	c.opcodes[0x9F] = func() {
		c.sbc(c.A)
	}
	c.opcodes[0xA0] = func() {
		c.and(getUpper(c.BC))
	}
	c.opcodes[0xA1] = func() {
		c.and(getLower(c.BC))
	}
	c.opcodes[0xA2] = func() {
		c.and(getUpper(c.DE))
	}
	c.opcodes[0xA3] = func() {
		c.and(getLower(c.DE))
	}
	c.opcodes[0xA4] = func() {
		c.and(getUpper(c.HL))
	}
	c.opcodes[0xA5] = func() {
		c.and(getLower(c.HL))
	}
	c.opcodes[0xA6] = func() {
		c.and(c.Memory.Read(c.HL))
	}
	c.opcodes[0xA7] = func() {
		c.and(c.A)
	}
	c.opcodes[0xA8] = func() {
		c.xor(getUpper(c.BC))
	}
	c.opcodes[0xA9] = func() {
		c.xor(getLower(c.BC))
	}
	c.opcodes[0xAA] = func() {
		c.xor(getUpper(c.DE))
	}
	c.opcodes[0xAB] = func() {
		c.xor(getLower(c.DE))
	}
	c.opcodes[0xAC] = func() {
		c.xor(getUpper(c.HL))
	}
	c.opcodes[0xAD] = func() {
		c.xor(getLower(c.HL))
	}
	c.opcodes[0xAE] = func() {
		c.xor(c.Memory.Read(c.HL))
	}
	c.opcodes[0xAF] = func() {
		c.xor(c.A)
	}
	c.opcodes[0xB0] = func() {
		c.or(getUpper(c.BC))
	}
	c.opcodes[0xB1] = func() {
		c.or(getLower(c.BC))
	}
	c.opcodes[0xB2] = func() {
		c.or(getUpper(c.DE))
	}
	c.opcodes[0xB3] = func() {
		c.or(getLower(c.DE))
	}
	c.opcodes[0xB4] = func() {
		c.or(getUpper(c.HL))
	}
	c.opcodes[0xB5] = func() {
		c.or(getLower(c.HL))
	}
	c.opcodes[0xB6] = func() {
		c.or(c.Memory.Read(c.HL))
	}
	c.opcodes[0xB7] = func() {
		c.or(c.A)
	}
	c.opcodes[0xB8] = func() {
		c.cp(getUpper(c.BC))
	}
	c.opcodes[0xB9] = func() {
		c.cp(getLower(c.BC))
	}
	c.opcodes[0xBA] = func() {
		c.cp(getUpper(c.DE))
	}
	c.opcodes[0xBB] = func() {
		c.cp(getLower(c.DE))
	}
	c.opcodes[0xBC] = func() {
		c.cp(getUpper(c.HL))
	}
	c.opcodes[0xBD] = func() {
		c.cp(getLower(c.HL))
	}
	c.opcodes[0xBE] = func() {
		c.cp(c.Memory.Read(c.HL))
	}
	c.opcodes[0xBF] = func() {
		c.cp(c.A)
	}
	c.opcodes[0xC0] = func() {
		if !c.Flags.Z {
			c.PC = c.PopWord()
		}
	}
	c.opcodes[0xC1] = func() {
		c.BC = c.PopWord()
	}
	c.opcodes[0xC2] = func() {
		addr := c.NextWord()
		if !c.Flags.Z {
			c.PC = addr
		}
	}
	c.opcodes[0xC3] = func() {
		c.PC = c.NextWord()
	}
	c.opcodes[0xC4] = func() {
		addr := c.NextWord()
		if !c.Flags.Z {
			c.PushWord(c.PC)
			c.PC = addr
		}
	}
	c.opcodes[0xC5] = func() {
		c.PushWord(c.BC)
	}
	c.opcodes[0xC6] = func() {
		c.add(c.NextByte())
	}
	c.opcodes[0xC7] = func() {
		c.rst(0x00)
	}
	c.opcodes[0xC8] = func() {
		if c.Flags.Z {
			c.PC = c.PopWord()
		}
	}
	c.opcodes[0xC9] = func() {
		c.PC = c.PopWord()
	}
	c.opcodes[0xCA] = func() {
		addr := c.NextWord()
		if c.Flags.Z {
			c.PC = addr
		}
	}
	c.opcodes[0xCB] = func() {
		b := c.NextByte()
		c.cbOpcodes[b]()
	}
	c.opcodes[0xCC] = func() {
		addr := c.NextWord()
		if c.Flags.Z {
			c.PushWord(c.PC)
			c.PC = addr
		}
	}
	c.opcodes[0xCD] = func() {
		addr := c.NextWord()
		c.PushWord(c.PC)
		c.PC = addr
	}
	c.opcodes[0xCE] = func() {
		c.adc(c.NextByte())
	}
	c.opcodes[0xCF] = func() {
		c.rst(0x08)
	}
	c.opcodes[0xD0] = func() {
		if !c.Flags.C {
			c.PC = c.PopWord()
		}
	}
	c.opcodes[0xD1] = func() {
		c.DE = c.PopWord()
	}
	c.opcodes[0xD2] = func() {
		addr := c.NextWord()
		if !c.Flags.C {
			c.PC = addr
		}
	}
	c.opcodes[0xD4] = func() {
		addr := c.NextWord()
		if !c.Flags.C {
			c.PushWord(c.PC)
			c.PC = addr
		}
	}
	c.opcodes[0xD5] = func() {
		c.PushWord(c.DE)
	}
	c.opcodes[0xD6] = func() {
		c.sub(c.NextByte())
	}
	c.opcodes[0xD7] = func() {
		c.rst(0x10)
	}
	c.opcodes[0xD8] = func() {
		if c.Flags.C {
			c.PC = c.PopWord()
		}
	}
	c.opcodes[0xD9] = func() {
		c.PC = c.PopWord()
		c.InterruptsEnabled = true
	}
	c.opcodes[0xDA] = func() {
		addr := c.NextWord()
		if c.Flags.C {
			c.PC = addr
		}
	}
	c.opcodes[0xDC] = func() {
		addr := c.NextWord()
		if c.Flags.C {
			c.PushWord(c.PC)
			c.PC = addr
		}
	}
	c.opcodes[0xDE] = func() {
		c.sbc(c.NextByte())
	}
	c.opcodes[0xDF] = func() {
		c.rst(0x18)
	}
	c.opcodes[0xE0] = func() {
		c.Memory.Write((0xFF00 + uint16(c.NextByte())), c.A)
	}
	c.opcodes[0xE1] = func() {
		c.HL = c.PopWord()
	}
	c.opcodes[0xE2] = func() {
		c.Memory.Write((0xFF00 + uint16(getLower(c.BC))), c.A)
	}
	c.opcodes[0xE5] = func() {
		c.PushWord(c.HL)
	}
	c.opcodes[0xE6] = func() {
		c.and(c.NextByte())
	}
	c.opcodes[0xE7] = func() {
		c.rst(0x20)
	}
	c.opcodes[0xE8] = func() {
		val := c.NextByte()
		c.Flags.H = ((c.PC & 0x0FFF) + uint16(val)) > 0x0FFF
		c.Flags.C = (uint32(val) + uint32(c.PC)) > 0xFFFF
		c.Flags.Z = false
		c.Flags.N = false
		c.SP = signedAdd(c.SP, val)
	}
	// TODO does JP (HL) mean jump to HL or jump to an address pointed to by HL
	c.opcodes[0xE9] = func() {
		c.PC = c.HL
	}
	c.opcodes[0xEA] = func() {
		c.Memory.Write(c.NextWord(), c.A)
	}
	c.opcodes[0xEE] = func() {
		c.xor(c.NextByte())
	}
	c.opcodes[0xEF] = func() {
		c.rst(0x28)
	}
	c.opcodes[0xF0] = func() {
		c.A = c.Memory.Read(0xFF00 + uint16(c.NextByte()))
	}
	c.opcodes[0xF1] = func() {
		val := c.PopWord()
		c.A = getUpper(val)
		c.Flags.Z = val&(1<<7) != 0
		c.Flags.N = val&(1<<6) != 0
		c.Flags.H = val&(1<<5) != 0
		c.Flags.C = val&(1<<4) != 0
	}
	c.opcodes[0xF2] = func() {
		c.A = c.Memory.Read(0xFF00 + (c.BC & 0x00FF))
	}
	c.opcodes[0xF3] = func() {
		c.InterruptsEnabled = false
	}
	c.opcodes[0xF5] = func() {
		val := uint16(c.A) << 8
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
		c.PushWord(val)
	}
	c.opcodes[0xF6] = func() {
		c.or(c.NextByte())
	}
	c.opcodes[0xF7] = func() {
		c.rst(0x30)
	}
	c.opcodes[0xF8] = func() {
		val := c.NextByte()
		c.Flags.H = ((c.PC & 0x0FFF) + uint16(val)) > 0x0FFF
		c.Flags.C = (uint32(val) + uint32(c.PC)) > 0xFFFF
		c.Flags.Z = false
		c.Flags.N = false
		c.HL = signedAdd(c.SP, val)
	}
	c.opcodes[0xF9] = func() {
		c.SP = c.HL
	}
	c.opcodes[0xFA] = func() {
		c.A = c.Memory.Read(c.NextWord())
	}
	c.opcodes[0xFB] = func() {
		c.InterruptsEnabled = true
	}
	c.opcodes[0xFE] = func() {
		c.cp(c.NextByte())
	}
	c.opcodes[0xFF] = func() {
		c.rst(0x38)
	}

	c.cbOpcodes[0x7C] = func() {
		c.Flags.Z = getUpper(c.HL)&(1<<7) != 0
		c.Flags.N = false
		c.Flags.H = true
	}

	c.cbOpcodes[0x11] = func() {
		setLower(&c.BC, c.rl(getLower(c.BC)))
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

func (c *CPU) rst(addr uint8) {
	c.PushWord(c.PC)
	c.PC = uint16(addr)
}

func (c *CPU) add(b uint8) {
	n := uint16(c.A) + uint16(b)
	c.Flags.Z = (n & 0xFF) == 0
	c.Flags.H = ((c.A & 0x0F) + (b & 0x0F)) > 0x0F
	c.Flags.C = n > 0xFF
	c.Flags.N = false
	c.A = uint8(n & 0xFF)
}

func (c *CPU) adc(b uint8) {
	n := uint16(c.A) + uint16(b)
	if c.Flags.C {
		n++
	}
	c.Flags.Z = (n & 0xFF) == 0
	c.Flags.H = ((c.A & 0x0F) + (b & 0x0F)) > 0x0F
	c.Flags.C = n > 0xFF
	c.Flags.N = false
	c.A = uint8(n & 0xFF)
}

func (c *CPU) sub(b uint8) {
	c.cp(b)
	c.A -= b
}

func (c *CPU) sbc(b uint8) {
	c.Flags.H = (int8(c.A&0xF) - int8(b&0xF)) < 0
	c.Flags.C = (int16(c.A) - int16(b) - 1) < 0
	c.Flags.N = true
	c.A -= b
	c.A--
	c.Flags.Z = c.A == 0
}

func (c *CPU) and(b uint8) {
	c.A &= b
	c.Flags.Z = c.A == 0
	c.Flags.H = true
	c.Flags.N = false
	c.Flags.C = false
}

func (c *CPU) xor(b uint8) {
	c.A ^= b
	c.Flags.Z = c.A == 0
	c.Flags.H = false
	c.Flags.N = false
	c.Flags.C = false
}

func (c *CPU) or(b uint8) {
	c.A |= b
	c.Flags.Z = c.A == 0
	c.Flags.H = false
	c.Flags.N = false
	c.Flags.C = false
}

func (c *CPU) cp(b uint8) {
	c.Flags.Z = c.A == b
	c.Flags.H = (c.A & 0xF) < (b & 0xF)
	c.Flags.N = true
	c.Flags.C = c.A < b
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

func (c *CPU) rl(val uint8) uint8 {
	oldcarry := c.Flags.C
	c.Flags.C = (val & (1 << 7)) == (1 << 7)
	c.Flags.H = false
	c.Flags.N = false
	val <<= 1
	if oldcarry {
		val |= 0x1
	}
	c.Flags.Z = val == 0
	return val
}

func (c *CPU) addRegs(a *uint16, b *uint16) {
	temp := uint32(*a) + uint32(*b)
	c.Flags.N = false
	c.Flags.C = temp > 0xFFFF
	c.Flags.H = ((*a & 0x0FFF) + (*b & 0x0FFF)) > 0x0FFF
	*a = uint16(temp & 0xFFFF)
}

func (c *CPU) relativeJump(dist uint8) {
	c.PC = signedAdd(c.PC, dist)
}

func signedAdd(a uint16, b uint8) uint16 {
	bSigned := int8(b)
	if bSigned >= 0 {
		return a + uint16(bSigned)
	} else {
		return a - uint16(-bSigned)
	}
}
