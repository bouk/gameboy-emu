package memory

import (
	"github.com/boukevanderbijl/Go-SDL/sdl"
	"log"
	"sync"
)

type Video struct {
	sync.Mutex
	RAM, OAM   []uint8
	Tiles      [256]*sdl.Surface
	Background *sdl.Surface
}

func NewVideo() *Video {
	v := new(Video)
	v.RAM = make([]uint8, 8*1024)
	v.OAM = make([]uint8, 4*40)
	for i := 0; i < 256; i++ {
		v.Tiles[i] = sdl.CreateRGBSurface(0, 8, 8, 32, 0, 0, 0, 0)
		v.Tiles[i].FillRect(&sdl.Rect{0, 0, 8, 8}, sdl.MapRGBA(v.Tiles[i].Format, 0xFF, 0xFF, 0xFF, 0xFF))
	}
	v.Background = sdl.CreateRGBSurface(0, 256, 256, 32, 0, 0, 0, 0)
	v.Background.FillRect(&sdl.Rect{0, 0, 256, 256}, sdl.MapRGBA(v.Background.Format, 0xFF, 0xFF, 0xFF, 0xFF))
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
		tilenumber := (addr - 0x8000) / 16
		if tilenumber < 256 {
			for line := uint16(0); line < 8; line++ {
				for pixel := uint16(0); pixel < 8; pixel++ {
					b := tilenumber*16 + line*2
					rect := sdl.Rect{int16(7 - pixel), int16(line), 1, 1}
					if (((v.RAM[b] >> pixel) & 0x1) | (((v.RAM[b+1] >> pixel) & 0x1) << 1)) > 0 {
						v.Tiles[tilenumber].FillRect(&rect, sdl.MapRGBA(v.Tiles[tilenumber].Format, 0, 0, 0, 0xFF))
					}
				}
			}
			for row := int16(0); row < 32; row++ {
				for col := int16(0); col < 32; col++ {
					dstRect := &sdl.Rect{(col * 8), (row * 8), 8, 8}
					v.Background.Blit(dstRect, v.Tiles[v.RAM[0x1800+row*32+col]], &sdl.Rect{0, 0, 8, 8})
				}
			}
		} else if addr >= 0x9800 && addr <= 0x9BFF {
			col := int16(addr-0x9800) % 32
			row := int16(addr-0x9800) / 32
			dstRect := &sdl.Rect{(col * 8), (row * 8), 8, 8}
			tile := v.RAM[0x1800+row*32+col]
			v.Background.Blit(dstRect, v.Tiles[tile], &sdl.Rect{0, 0, 8, 8})
		}
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
