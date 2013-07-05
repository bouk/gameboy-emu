package memory

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
