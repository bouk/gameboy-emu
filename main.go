package main

import (
	"flag"
	"fmt"
	"github.com/boukevanderbijl/gameboy-emu/lr35902"
	"github.com/boukevanderbijl/gameboy-emu/memory"
	"io/ioutil"
	"log"
	"math"
	"strings"
)

var (
	debugEnabled bool
	filename     string
	showHelp     bool
)

func init() {
	flag.BoolVar(&debugEnabled, "debug", false, "enable debug logging to stderr")
	flag.BoolVar(&showHelp, "help", false, "show help")
	flag.BoolVar(&showHelp, "h", false, "show help")
}

func main() {
	flag.Parse()

	if showHelp {
		flag.Usage()
		return
	}

	if len(flag.Args()) != 1 {
		fmt.Println(flag.Args())
		fmt.Println("Please supply a ROM to load")
		return
	}
	filename = flag.Arg(0)

	if !debugEnabled {
		log.SetOutput(ioutil.Discard)
	}

	rom, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Println(err)
		fmt.Printf("Unable to open \"%s\"\n", filename)
		return
	}

	if len(rom) < 0x0150 {
		log.Println("Rom is too small")
		fmt.Println("Invalid ROM")
		return
	}

	var checksum uint16
	for i, b := range rom {
		if i != 0x14E && i != 0x14F {
			checksum += uint16(b)
		}
	}
	log.Printf("Checksum: 0x%04X", checksum)
	if ((uint16(rom[0x14E]) << 8) | uint16(rom[0x14F])) != checksum {
		log.Println("The ROM's checksum does not match")
	}

	gameName := string(rom[0x0134:0x0144])
	if p := strings.Index(gameName, "\000"); p != -1 {
		gameName = gameName[0:p]
	}
	log.Println("Name of the game:", gameName)

	log.Printf("ROM type: 0x%02X", rom[0x0148])
	var romSize int
	switch rom[0x0148] {
	case 0x0:
		romSize = 32 * 1024
		log.Println("256Kbit = 32KByte = 2 banks")
	case 0x1:
		romSize = 64 * 1024
		log.Println("512Kbit = 64KByte = 4 banks")
	case 0x2:
		romSize = 128 * 1024
		log.Println("1Mbit = 128KByte = 8 banks")
	case 0x3:
		romSize = 256 * 1024
		log.Println("2Mbit = 256KByte = 16 banks")
	case 0x4:
		romSize = 512 * 1024
		log.Println("4Mbit = 512KByte = 32 banks")
	case 0x5:
		romSize = 1024 * 1024
		log.Println("8Mbit = 1MByte = 64 banks")
	case 0x6:
		romSize = 2 * 1024 * 1024
		log.Println("16Mbit = 2MByte = 128 banks")
	case 0x7:
		romSize = 4 * 1024 * 1024
		log.Println("32Mbit = 4MByte = 256 banks")
	case 0x8:
		romSize = 8 * 1024 * 1024
		log.Println("64Mbit = 8MByte = 512 banks")
	case 0x52:
		romSize = 1179648
		log.Println("9Mbit = 1.1MByte = 72 banks")
	case 0x53:
		romSize = 1310720
		log.Println("10Mbit = 1.2MByte = 80 banks")
	case 0x54:
		romSize = 1572864
		log.Println("12Mbit = 1.5MByte = 96 banks")
	default:
		log.Println("Unknown ROM size")
	}
	log.Printf("ROM is %d bytes", romSize)
	log.Printf("RAM type: 0x%02X", rom[0x0149])

	var ramSize int
	switch rom[0x0149] {
	case 0x0:
		ramSize = 0
		log.Println("None")
	case 0x1:
		ramSize = 2 * 1024
		log.Println("16kBit = 2kB = 1 bank")
	case 0x2:
		ramSize = 8 * 1024
		log.Println("64kBit = 8kB = 1 bank")
	case 0x3:
		ramSize = 32 * 1024
		log.Println("256kBit = 32kB = 4 banks")
	case 0x4:
		ramSize = 128 * 1024
		log.Println("1MBit = 128kB = 16 banks")
	default:
		var ramBanks int = int(math.Pow(4, float64(int(rom[0x0149])-0x2)))
		log.Println("Unknown RAM size, using formula")
		ramSize = ramBanks * 8 * 1024
	}

	log.Printf("RAM is %d bytes", ramSize)
	ram := make([]uint8, ramSize)

	log.Printf("Cartridge type: 0x%02X", rom[0x0147])

	var mbc memory.Memory

	switch rom[0x0147] {
	case 0x0:
		log.Println("ROM ONLY")
		mbc = memory.NewSimple(rom, ram)
	case 0x1:
		log.Println("ROM+MBC1")
		mbc = memory.NewMBC1(rom, ram)
	case 0x2:
		log.Println("ROM+MBC1+RAM")
		mbc = memory.NewMBC1(rom, ram)
	case 0x3:
		log.Println("ROM+MBC1+RAM+BATTERY")
		mbc = memory.NewMBC1(rom, ram)
	case 0x5:
		log.Println("ROM+MBC2")
		mbc = memory.NewMBC2(rom, ram)
	case 0x6:
		log.Println("ROM+MBC2+BATTERY")
		mbc = memory.NewMBC2(rom, ram)
	case 0x8:
		log.Println("ROM+RAM")
		mbc = memory.NewSimple(rom, ram)
	case 0x9:
		log.Println("ROM+RAM+BATTERY")
		mbc = memory.NewSimple(rom, ram)
	case 0xB:
		log.Println("ROM+MMM01")
		fmt.Println("Unimplemented memory controller")
		return
	case 0xC:
		log.Println("ROM+MMM01+SRAM")
		fmt.Println("Unimplemented memory controller")
		return
	case 0xD:
		log.Println("ROM+MMM01+SRAM+BATTERY")
		fmt.Println("Unimplemented memory controller")
		return
	case 0xF:
		log.Println("ROM+MBC3+TIMER+BATTERY")
		mbc = memory.NewMBC3(rom, ram)
	case 0x10:
		log.Println("ROM+MBC3+TIMER+RAM+BATTERY")
		mbc = memory.NewMBC3(rom, ram)
	case 0x11:
		log.Println("ROM+MBC3")
		mbc = memory.NewMBC3(rom, ram)
	case 0x12:
		log.Println("ROM+MBC3+RAM")
		mbc = memory.NewMBC3(rom, ram)
	case 0x13:
		log.Println("ROM+MBC3+RAM+BATTERY")
		mbc = memory.NewMBC3(rom, ram)
	case 0x19:
		log.Println("ROM+MBC5")
		mbc = memory.NewMBC5(rom, ram)
	case 0x1A:
		log.Println("ROM+MBC5+RAM")
		mbc = memory.NewMBC5(rom, ram)
	case 0x1B:
		log.Println("ROM+MBC5+RAM+BATTERY")
		mbc = memory.NewMBC5(rom, ram)
	case 0x1C:
		log.Println("ROM+MBC5+RUMBLE")
		mbc = memory.NewMBC5(rom, ram)
	case 0x1D:
		log.Println("ROM+MBC5+RUMBLE+SRAM")
		mbc = memory.NewMBC5(rom, ram)
	case 0x1E:
		log.Println("ROM+MBC5+RUMBLE+SRAM+BATTERY")
		mbc = memory.NewMBC5(rom, ram)
	case 0x1F:
		log.Println("Pocket Camera")
		fmt.Println("Unimplemented memory controller")
		return
	case 0xFD:
		log.Println("Bandai TM5")
		fmt.Println("Unimplemented memory controller")
		return
	case 0xFE:
		log.Println("Hudson HuC-3")
		fmt.Println("Unimplemented memory controller")
		return
	case 0xFF:
		log.Println("Hudson HuC-1")
		fmt.Println("Unimplemented memory controller")
		return
	default:
		log.Println("Unknown memory controller")
		fmt.Println("Unimplemented memory controller")
		return
	}

	m := memory.NewMMC(mbc)
	cpu := lr35902.NewCPU(m)

	if m.Bios, err = ioutil.ReadFile("DMG_ROM.bin"); err == nil {
		log.Println("BIOS file found, loading...")
		m.BiosEnabled = true
	} else {
		log.Println("BIOS file not found, emulating state after BIOS")

		cpu.BC = 0x0013
		cpu.DE = 0x00D8
		cpu.HL = 0x014D
		cpu.PC = 0x0100
		cpu.A = 0x01
		cpu.Flags.Z = true
		cpu.Flags.H = true
		cpu.Flags.C = true
	}

	for !cpu.Stopped && cpu.PC <= 0xFF {
		cpu.Step()
	}
	cpu.DumpState()
}
