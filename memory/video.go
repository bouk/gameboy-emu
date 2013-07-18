package memory

import (
	"log"
	"sync"
)

type Video struct {
	sync.Mutex
	RAM, OAM []uint8
}

func NewVideo() *Video {
	v := new(Video)
	v.RAM = make([]uint8, 8*1024)
	v.OAM = make([]uint8, 4*40)
	return v
}

func (v *Video) Write(addr uint16, value uint8) {
	v.Lock()
	defer v.Unlock()
	switch {
	case addr < 0x8000:
		log.Println("Invalid write to Video 0x%04X 0x%02X", addr, value)
	case addr < 0xA000:
		v.RAM[addr-0x8000] = value
	case addr < 0xFE00:
		log.Println("Invalid write to Video 0x%04X 0x%02X", addr, value)
	case addr < 0xFEA0:
		v.OAM[addr-0xFE00] = value
	default:
		log.Println("Invalid write to Video 0x%04X 0x%02X", addr, value)
	}
}

func (v *Video) Read(addr uint16) uint8 {
	v.Lock()
	defer v.Unlock()
	switch {
	case addr < 0x8000:
		log.Println("Invalid read from Video 0x%04X", addr)
		return 0
	case addr < 0xA000:
		return v.RAM[addr-0x8000]
	case addr < 0xFE00:
		log.Println("Invalid read from Video 0x%04X", addr)
		return 0
	case addr < 0xFEA0:
		return v.OAM[addr-0xFE00]
	default:
		log.Println("Invalid read from Video 0x%04X", addr)
		return 0
	}
}
