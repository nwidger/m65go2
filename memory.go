package _65go2

import (
	"io"
	"os"
)

type Memory interface {
	reset()
	load(path string)
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

func (mem *BasicMemory) load(path string) {
	fi, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := fi.Close(); err != nil {
			panic(err)
		}
	}()

	buf := make([]byte, 65536)

	for {
		n, err := fi.Read(buf)

		if err != nil && err != io.EOF {
			panic(err)
		}

		if n == 0 {
			break
		}
	}

	for i, b := range buf {
		mem[i] = b
	}

	return
}

func SamePage(addr1 uint16, addr2 uint16) bool {
	return (0xff00 & addr1) == (0xff00 & addr2)
}
