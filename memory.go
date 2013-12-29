package _65go2

type Memory interface {
	reset()
	fetch(address uint16) (value uint8)
	store(address uint16, value uint8) (oldValue uint8)
}

type BasicMemory [65536]uint8

func NewBasicMemory() *BasicMemory {
	return &BasicMemory{}
}

func (mem *BasicMemory) reset() {
	for i := range mem {
		mem[i] = 0
	}
}

func (mem *BasicMemory) fetch(address uint16) (value uint8) {
	value = mem[address]
	return
}

func (mem *BasicMemory) store(address uint16, value uint8) (oldValue uint8) {
	oldValue = mem[address]
	mem[address] = value
	return
}

func SamePage(addr1 uint16, addr2 uint16) bool {
	return (0xff00 & addr1) == (0xff00 & addr2)
}
