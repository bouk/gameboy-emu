package memory

import (
	"fmt"
	"log"
)

type Memory interface {
	Read(address uint16) uint8
	Write(address uint16, value uint8)
}

type RAM []uint8

func (r *RAM) Read(addr uint16) uint8 {
	return (*r)[addr]
}

func (r *RAM) Write(addr uint16, value uint8) {
	(*r)[addr] = value
}

func MbcType(id uint8, rom, ram []uint8) (Memory, string, error) {
	switch id {
	case 0x0:
		log.Println()
		return NewSimple(rom, ram), "ROM ONLY", nil
	case 0x1:
		log.Println()
		return NewMBC1(rom, ram), "ROM+MBC1", nil
	case 0x2:
		log.Println()
		return NewMBC1(rom, ram), "ROM+MBC1+RAM", nil
	case 0x3:
		log.Println()
		return NewMBC1(rom, ram), "ROM+MBC1+RAM+BATTERY", nil
	case 0x5:
		log.Println()
		return NewMBC2(rom, ram), "ROM+MBC2", nil
	case 0x6:
		log.Println()
		return NewMBC2(rom, ram), "ROM+MBC2+BATTERY", nil
	case 0x8:
		log.Println()
		return NewSimple(rom, ram), "ROM+RAM", nil
	case 0x9:
		log.Println()
		return NewSimple(rom, ram), "ROM+RAM+BATTERY", nil
	case 0xB:
		log.Println()
		return nil, "ROM+MMM01", fmt.Errorf("Unimplemented memory controller")
	case 0xC:
		log.Println()
		return nil, "ROM+MMM01+SRAM", fmt.Errorf("Unimplemented memory controller")
	case 0xD:
		log.Println()
		return nil, "ROM+MMM01+SRAM+BATTERY", fmt.Errorf("Unimplemented memory controller")
	case 0xF:
		log.Println()
		return NewMBC3(rom, ram), "ROM+MBC3+TIMER+BATTERY", nil
	case 0x10:
		log.Println()
		return NewMBC3(rom, ram), "ROM+MBC3+TIMER+RAM+BATTERY", nil
	case 0x11:
		log.Println()
		return NewMBC3(rom, ram), "ROM+MBC3", nil
	case 0x12:
		log.Println()
		return NewMBC3(rom, ram), "ROM+MBC3+RAM", nil
	case 0x13:
		log.Println()
		return NewMBC3(rom, ram), "ROM+MBC3+RAM+BATTERY", nil
	case 0x19:
		log.Println()
		return NewMBC5(rom, ram), "ROM+MBC5", nil
	case 0x1A:
		log.Println()
		return NewMBC5(rom, ram), "ROM+MBC5+RAM", nil
	case 0x1B:
		log.Println()
		return NewMBC5(rom, ram), "ROM+MBC5+RAM+BATTERY", nil
	case 0x1C:
		log.Println()
		return NewMBC5(rom, ram), "ROM+MBC5+RUMBLE", nil
	case 0x1D:
		log.Println()
		return NewMBC5(rom, ram), "ROM+MBC5+RUMBLE+SRAM", nil
	case 0x1E:
		log.Println()
		return NewMBC5(rom, ram), "ROM+MBC5+RUMBLE+SRAM+BATTERY", nil
	case 0x1F:
		log.Println()
		return nil, "Pocket Camera", fmt.Errorf("Unimplemented memory controller")
	case 0xFD:
		log.Println()
		return nil, "Bandai TM5", fmt.Errorf("Unimplemented memory controller")
	case 0xFE:
		log.Println()
		return nil, "Hudson HuC-3", fmt.Errorf("Unimplemented memory controller")
	case 0xFF:
		log.Println()
		return nil, "Hudson HuC-1", fmt.Errorf("Unimplemented memory controller")
	default:
		return nil, "", fmt.Errorf("Unknown memory controller")
	}
}
