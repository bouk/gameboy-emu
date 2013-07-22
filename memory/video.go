package memory

import (
	"log"
)

type Video struct {
	RAM, OAM                                         []uint8
	LCDC, SCX, SCY, LY, LYC, BGP, OBP0, OBP1, WX, WY uint8
	CycleStep                                        int
	Pixels                                           []uint8
}

type Sprite []uint8

func (s Sprite) GetPixel(x, y uint8) uint8 {
	if len(s) == 16 { // 8x8
		x = 7 - x
		row := y * 2
		return ((s[row] & (1 << x)) >> x) | (((s[row+1] & (1 << x)) >> x) << 1)
	} else if len(s) == 32 { // 16x8
		// FIXME
		return 0
	}
	panic("Invalid sprite size")
}

const (
	WIDTH  = 160
	HEIGHT = 144

	VRAM_START = 0x8000

	TILEMAP_1_START = 0x9800
	TILEMAP_2_START = 0x9C00

	TILE_DATA_1_START = 0x8000
	TILE_DATA_2_START = 0x8800

	BG_TILEMAP_SELECT      = 1 << 3
	TILE_DATA_TABLE_SELECT = 1 << 4
)

var (
	bgColors = [][]byte{{0xFF, 0xFF, 0xFF, 0xFF},
		{0xDD, 0xDD, 0xDD, 0xFF},
		{0xAA, 0xAA, 0xAA, 0xFF},
		{0, 0, 0, 0xFF}}
	spriteColors = [][]byte{{0, 0, 0, 0},
		{0, 0, 0, 0xFF},
		{0, 0, 0, 0xFF},
		{0, 0, 0, 0xFF}}
)

func NewVideo() *Video {
	v := new(Video)
	v.RAM = make([]uint8, 8*1024)
	v.OAM = make([]uint8, 4*40)
	v.Pixels = make([]uint8, WIDTH*HEIGHT*4)
	return v
}

func (v *Video) GetBgTile(id uint8) Sprite {
	var tileStart uint16
	if v.LCDC&TILE_DATA_TABLE_SELECT != 0 {
		tileStart = 0x8000 + uint16(id)*16
	} else {
		id += 128
		tileStart = 0x8800 + uint16(id)*16
	}
	return Sprite(v.RAM[(tileStart - VRAM_START) : (tileStart-VRAM_START)+16])
}

func (v *Video) GetBgTileId(x, y uint8) uint8 {
	var tileMapAddr uint16
	if v.LCDC&BG_TILEMAP_SELECT != 0 {
		tileMapAddr = TILEMAP_2_START
	} else {
		tileMapAddr = TILEMAP_1_START
	}
	tileMapAddr += uint16(x)
	tileMapAddr += uint16(y) * 32

	return v.RAM[tileMapAddr-VRAM_START]
}

// Render a certain line to the pixel buffer
func (v *Video) renderLine(line uint8) {
	for pos, x := WIDTH*int(line)*4, uint8(0); x < WIDTH; x++ {
		// background
		var color uint8 = 0
		if v.LCDC&1 != 0 {
			bgX := x + v.SCX
			bgY := line + v.SCY
			tile := v.GetBgTileId(bgX/8, bgY/8)
			s := v.GetBgTile(tile)
			color = s.GetPixel(bgX%8, bgY%8)
		}
		// window
		copy(v.Pixels[pos:pos+4], bgColors[color])
		pos += 4
	}
	// sprites
}

func (v *Video) Step() {
	// If there has been enough time
	v.CycleStep++
	if v.CycleStep == 486 {
		v.CycleStep = 0

		if v.LY < 144 {
			v.renderLine(v.LY)
		}
		v.LY++
		if v.LY >= 154 {
			v.LY = 0
		}
	}
}

func (v *Video) Write(addr uint16, value uint8) {
	switch {
	case addr < VRAM_START:
		log.Println("Invalid write to Video 0x%04X 0x%02X", addr, value)
	case addr < 0xA000:
		v.RAM[addr-VRAM_START] = value
	case addr < 0xFE00:
		log.Println("Invalid write to Video 0x%04X 0x%02X", addr, value)
	case addr < 0xFEA0:
		v.OAM[addr-0xFE00] = value
	default:
		log.Println("Invalid write to Video 0x%04X 0x%02X", addr, value)
	}
}

func (v *Video) Read(addr uint16) uint8 {
	switch {
	case addr < VRAM_START:
		log.Println("Invalid read from Video 0x%04X", addr)
		return 0
	case addr < 0xA000:
		return v.RAM[addr-VRAM_START]
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
