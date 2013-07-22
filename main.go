package main

import (
	sf "bitbucket.org/krepa098/gosfml2"
	"flag"
	"fmt"
	"github.com/boukevanderbijl/gameboy-emu/lr35902"
	"github.com/boukevanderbijl/gameboy-emu/memory"
	"io/ioutil"
	"log"
	"math"
	"runtime"
	"strings"
	"time"
)

var (
	debugEnabled bool
	filename     string
	showHelp     bool
	cpuDump      bool
	turbo        bool
)

func init() {
	flag.BoolVar(&debugEnabled, "debug", false, "enable debug logging to stderr")
	flag.BoolVar(&showHelp, "help", false, "show help")
	flag.BoolVar(&showHelp, "h", false, "show help")
	flag.BoolVar(&cpuDump, "dumpcpu", false, "dump the state of the CPU at every step")
	flag.BoolVar(&turbo, "turbo", false, "enables 4x TURBO MODE")
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	runtime.LockOSThread()

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

	mbc, mbcName, err := memory.MbcType(rom[0x0147], rom, ram)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println(mbcName)

	m := memory.NewMMC(mbc)
	video := m.Video
	cpu := lr35902.NewCPU(m)

	if turbo {
		cpu.RealisticSteps = false
	}

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

	window := sf.NewRenderWindow(sf.VideoMode{160, 144, 32}, "Goboy", sf.StyleDefault, nil)
	texture, _ := sf.NewTexture(memory.WIDTH, memory.HEIGHT)
	renderingSprite := sf.NewSprite(texture)
	window.Clear(sf.ColorWhite())
	window.Display()

	screenUpdates := make(chan []uint8)
	go func() {

		for !cpu.Stopped {
			start := time.Now()
			cycles := cpu.Step()
			if cpuDump {
				cpu.DumpState()
			}
			for i := uint(0); i < cycles; i++ {
				video.Step()
				if video.LY == 144 && video.CycleStep == 0 {
					screenUpdates <- video.Pixels
				}
			}
			timeItShouldHaveTaken := int64((float64(cycles) / float64(cpu.ClockSpeed)) * 1e9)
			timeItTook := time.Now().Sub(start).Nanoseconds()
			if cpu.RealisticSteps {
				if timediff := timeItShouldHaveTaken - timeItTook; timediff > 0 {
					log.Printf("Sleeping for %d", timediff)
					time.Sleep(time.Duration(timediff) * time.Nanosecond)
				}
			}
		}
	}()
main:
	for {
		t := <-screenUpdates
		texture.UpdateFromPixels(t, memory.WIDTH, memory.HEIGHT, 0, 0)
		// VBLANK interrupt & render
		window.Clear(sf.ColorWhite())
		window.Draw(renderingSprite, nil)
		window.Display()
		for event := window.PollEvent(); event != nil; event = window.PollEvent() {
			switch ev := event.(type) {
			case sf.EventKeyPressed:
				//exit on ESC
				if ev.Code == sf.KeyEscape {
					break main
				}
			case sf.EventClosed:
				break main
			}
		}
	}
	cpu.DumpState()
}
