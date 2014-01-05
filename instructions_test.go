package go6502

import (
	"testing"
	"time"
)

const rate time.Duration = 46 * time.Nanosecond // 21.477272Mhz
const divisor = 12

var cpu *CPU

func Setup() {
	clock := NewClock(rate)
	cpu = NewCPU(NewBasicMemory(), divisor, clock)
	cpu.Reset()
	// cpu.decode = true
	go clock.start()
}

func Teardown() {
	cpu.clock.stop()
}

// BadOpCodeError

func TestBadOpCodeError(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x02)

	_, error := cpu.Execute()

	if error == nil {
		t.Error("No error returned")
	}

	if _, ok := error.(BadOpCodeError); !ok {
		t.Error("Did not receive expected error type BadOpCodeError")
	}

	Teardown()
}

// LDA

func TestLdaImmediate(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa9)
	cpu.memory.store(0x0101, 0xff)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
	}

	Teardown()
}

func TestLdaZeroPage(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa5)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0xff)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
	}

	Teardown()
}

func TestLdaZeroPageX(t *testing.T) {
	Setup()

	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xb5)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0xff)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
	}

	Teardown()
}

func TestLdaAbsolute(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xad)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0xff)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
	}

	Teardown()
}

func TestLdaAbsoluteX(t *testing.T) {
	Setup()

	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xbd)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0xff)

	cycles, _ := cpu.Execute()

	if cycles != 4 {
		t.Error("Cycles is not 4")
	}

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
	}

	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xbd)
	cpu.memory.store(0x0101, 0xff)
	cpu.memory.store(0x0102, 0x02)
	cpu.memory.store(0x0300, 0xff)

	cycles, _ = cpu.Execute()

	if cycles != 5 {
		t.Error("Cycles is not 5")
	}

	Teardown()
}

func TestLdaAbsoluteY(t *testing.T) {
	Setup()

	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xb9)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0xff)

	cycles, _ := cpu.Execute()

	if cycles != 4 {
		t.Error("Cycles is not 4")
	}

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
	}

	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xb9)
	cpu.memory.store(0x0101, 0xff)
	cpu.memory.store(0x0102, 0x02)
	cpu.memory.store(0x0300, 0xff)

	cycles, _ = cpu.Execute()

	if cycles != 5 {
		t.Error("Cycles is not 5")
	}

	Teardown()
}

func TestLdaIndirectX(t *testing.T) {
	Setup()

	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa1)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x87)
	cpu.memory.store(0x0086, 0x00)
	cpu.memory.store(0x0087, 0xff)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
	}

	Teardown()
}

func TestLdaIndirectY(t *testing.T) {
	Setup()

	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xb1)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x86)
	cpu.memory.store(0x0085, 0x00)
	cpu.memory.store(0x0087, 0xff)

	cycles, _ := cpu.Execute()

	if cycles != 5 {
		t.Error("Cycles is not 5")
	}

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
	}

	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xb1)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0xff)
	cpu.memory.store(0x0085, 0x02)
	cpu.memory.store(0x0300, 0xff)

	cycles, _ = cpu.Execute()

	if cycles != 6 {
		t.Error("Cycles is not 6")
	}

	Teardown()
}

func TestLdaZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa9)
	cpu.memory.store(0x0101, 0x00)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestLdaZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa9)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestLdaNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa9)
	cpu.memory.store(0x0101, 0x81)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestLdaNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa9)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// LDX

func TestLdxImmediate(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa2)
	cpu.memory.store(0x0101, 0xff)

	cpu.Execute()

	if cpu.registers.X != 0xff {
		t.Error("Register X is not 0xff")
	}

	Teardown()
}

func TestLdxZeroPage(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa6)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0xff)

	cpu.Execute()

	if cpu.registers.X != 0xff {
		t.Error("Register X is not 0xff")
	}

	Teardown()
}

func TestLdxZeroPageY(t *testing.T) {
	Setup()

	cpu.registers.Y = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xb6)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0xff)

	cpu.Execute()

	if cpu.registers.X != 0xff {
		t.Error("Register X is not 0xff")
	}

	Teardown()
}

func TestLdxAbsolute(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xae)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0xff)

	cpu.Execute()

	if cpu.registers.X != 0xff {
		t.Error("Register X is not 0xff")
	}

	Teardown()
}

func TestLdxAbsoluteY(t *testing.T) {
	Setup()

	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xbe)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0xff)

	cpu.Execute()

	if cpu.registers.X != 0xff {
		t.Error("Register X is not 0xff")
	}

	Teardown()
}

func TestLdxZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa2)
	cpu.memory.store(0x0101, 0x00)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestLdxZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa2)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestLdxNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa2)
	cpu.memory.store(0x0101, 0x81)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestLdxNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa2)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// LDY

func TestLdyImmediate(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa0)
	cpu.memory.store(0x0101, 0xff)

	cpu.Execute()

	if cpu.registers.Y != 0xff {
		t.Error("Register Y is not 0xff")
	}

	Teardown()
}

func TestLdyZeroPage(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa4)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0xff)

	cpu.Execute()

	if cpu.registers.Y != 0xff {
		t.Error("Register Y is not 0xff")
	}

	Teardown()
}

func TestLdyZeroPageX(t *testing.T) {
	Setup()

	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xb4)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0xff)

	cpu.Execute()

	if cpu.registers.Y != 0xff {
		t.Error("Register Y is not 0xff")
	}

	Teardown()
}

func TestLdyAbsolute(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xac)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0xff)

	cpu.Execute()

	if cpu.registers.Y != 0xff {
		t.Error("Register Y is not 0xff")
	}

	Teardown()
}

func TestLdyAbsoluteX(t *testing.T) {
	Setup()

	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xbc)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0xff)

	cpu.Execute()

	if cpu.registers.Y != 0xff {
		t.Error("Register Y is not 0xff")
	}

	Teardown()
}

func TestLdyZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa0)
	cpu.memory.store(0x0101, 0x00)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestLdyZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa0)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestLdyNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa0)
	cpu.memory.store(0x0101, 0x81)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestLdyNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa0)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// STA

func TestStaZeroPage(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x85)
	cpu.memory.store(0x0101, 0x84)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

func TestStaZeroPageX(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x95)
	cpu.memory.store(0x0101, 0x84)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

func TestStaAbsolute(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x8d)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

func TestStaAbsoluteX(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x9d)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

func TestStaAbsoluteY(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x99)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

func TestStaIndirectX(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x81)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x87)
	cpu.memory.store(0x0086, 0x00)

	cpu.Execute()

	if cpu.memory.fetch(0x0087) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

func TestStaIndirectY(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x91)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x86)
	cpu.memory.store(0x0085, 0x00)

	cpu.Execute()

	if cpu.memory.fetch(0x0087) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

// STX

func TestStxZeroPage(t *testing.T) {
	Setup()

	cpu.registers.X = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x86)
	cpu.memory.store(0x0101, 0x84)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

func TestStxZeroPageY(t *testing.T) {
	Setup()

	cpu.registers.X = 0xff
	cpu.registers.Y = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x96)
	cpu.memory.store(0x0101, 0x84)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

func TestStxAbsolute(t *testing.T) {
	Setup()

	cpu.registers.X = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x8e)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

// STY

func TestStyZeroPage(t *testing.T) {
	Setup()

	cpu.registers.Y = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x84)
	cpu.memory.store(0x0101, 0x84)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

func TestStyZeroPageY(t *testing.T) {
	Setup()

	cpu.registers.Y = 0xff
	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x94)
	cpu.memory.store(0x0101, 0x84)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

func TestStyAbsolute(t *testing.T) {
	Setup()

	cpu.registers.Y = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x8c)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

// TAX

func TestTax(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xaa)

	cpu.Execute()

	if cpu.registers.X != 0xff {
		t.Error("Register is not 0xff")
	}

	Teardown()
}

func TestTaxZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x00
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xaa)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestTaxZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xaa)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestTaxNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x81
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xaa)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestTaxNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xaa)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// TAY

func TestTay(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa8)

	cpu.Execute()

	if cpu.registers.Y != 0xff {
		t.Error("Register is not 0xff")
	}

	Teardown()
}

// TXA

func TestTxa(t *testing.T) {
	Setup()

	cpu.registers.X = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x8a)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register is not 0xff")
	}

	Teardown()
}

// TYA

func TestTya(t *testing.T) {
	Setup()

	cpu.registers.Y = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x98)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register is not 0xff")
	}

	Teardown()
}

// TSX

func TestTsx(t *testing.T) {
	Setup()

	cpu.registers.SP = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xba)

	cpu.Execute()

	if cpu.registers.X != 0xff {
		t.Error("Register is not 0xff")
	}

	Teardown()
}

// TXS

func TestTxs(t *testing.T) {
	Setup()

	cpu.registers.X = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x9a)

	cpu.Execute()

	if cpu.registers.SP != 0xff {
		t.Error("Register is not 0xff")
	}

	Teardown()
}

// PHA

func TestPha(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x48)

	cpu.Execute()

	if cpu.pull() != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

// PHP

func TestPhp(t *testing.T) {
	Setup()

	cpu.registers.P = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x08)

	cpu.Execute()

	if cpu.pull() != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

// PLA

func TestPla(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100
	cpu.push(0xff)

	cpu.memory.store(0x0100, 0x68)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

func TestPlaZFlagSet(t *testing.T) {
	Setup()

	cpu.push(0x00)
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x68)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestPlaZFlagUnset(t *testing.T) {
	Setup()

	cpu.push(0x01)
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x68)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestPlaNFlagSet(t *testing.T) {
	Setup()

	cpu.push(0x81)
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x68)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestPlaNFlagUnset(t *testing.T) {
	Setup()

	cpu.push(0x01)
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x68)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// PLP

func TestPlp(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100
	cpu.push(0xff)

	cpu.memory.store(0x0100, 0x28)

	cpu.Execute()

	if cpu.registers.P != 0xff {
		t.Error("Status is not 0xff")
	}

	Teardown()
}

// AND

func TestAndImmediate(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x29)
	cpu.memory.store(0x0101, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0x0f {
		t.Error("Register A is not 0x0f")
	}

	Teardown()
}

func TestAndZeroPage(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x25)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0x0f {
		t.Error("Register A is not 0x0f")
	}

	Teardown()
}

func TestAndZeroPageX(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x35)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0x0f {
		t.Error("Register A is not 0x0f")
	}

	Teardown()
}

func TestAndAbsolute(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x2d)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0x0f {
		t.Error("Register A is not 0x0f")
	}

	Teardown()
}

func TestAndAbsoluteX(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x3d)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0x0f {
		t.Error("Register A is not 0x0f")
	}

	Teardown()
}

func TestAndAbsoluteY(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x39)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0x0f {
		t.Error("Register A is not 0x0f")
	}

	Teardown()
}

func TestAndIndirectX(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x21)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x87)
	cpu.memory.store(0x0086, 0x00)
	cpu.memory.store(0x0087, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0x0f {
		t.Error("Register A is not 0x0f")
	}

	Teardown()
}

func TestAndIndirectY(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x31)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x86)
	cpu.memory.store(0x0085, 0x00)
	cpu.memory.store(0x0087, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0x0f {
		t.Error("Register A is not 0x0f")
	}

	Teardown()
}

func TestAndZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x29)
	cpu.memory.store(0x0101, 0x00)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestAndZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x29)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestAndNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x81
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x29)
	cpu.memory.store(0x0101, 0x81)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestAndNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x29)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// EOR

func TestEorImmediate(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x49)
	cpu.memory.store(0x0101, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0xf0 {
		t.Error("Register A is not 0xf0")
	}

	Teardown()
}

func TestEorZeroPage(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x45)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0xf0 {
		t.Error("Register A is not 0xf0")
	}

	Teardown()
}

func TestEorZeroPageX(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x55)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0xf0 {
		t.Error("Register A is not 0xf0")
	}

	Teardown()
}

func TestEorAbsolute(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x4d)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0xf0 {
		t.Error("Register A is not 0xf0")
	}

	Teardown()
}

func TestEorAbsoluteX(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x5d)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0xf0 {
		t.Error("Register A is not 0xf0")
	}

	Teardown()
}

func TestEorAbsoluteY(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x59)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0xf0 {
		t.Error("Register A is not 0xf0")
	}

	Teardown()
}

func TestEorIndirectX(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x41)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x87)
	cpu.memory.store(0x0086, 0x00)
	cpu.memory.store(0x0087, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0xf0 {
		t.Error("Register A is not 0xf0")
	}

	Teardown()
}

func TestEorIndirectY(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x51)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x86)
	cpu.memory.store(0x0085, 0x00)
	cpu.memory.store(0x0087, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0xf0 {
		t.Error("Register A is not 0xf0")
	}

	Teardown()
}

func TestEorZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x49)
	cpu.memory.store(0x0101, 0x00)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestEorZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x00
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x49)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestEorNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x00
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x49)
	cpu.memory.store(0x0101, 0x81)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestEorNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x49)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// ORA

func TestOraImmediate(t *testing.T) {
	Setup()

	cpu.registers.A = 0xf0
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x09)
	cpu.memory.store(0x0101, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
	}

	Teardown()
}

func TestOraZeroPage(t *testing.T) {
	Setup()

	cpu.registers.A = 0xf0
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x05)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
	}

	Teardown()
}

func TestOraZeroPageX(t *testing.T) {
	Setup()

	cpu.registers.A = 0xf0
	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x15)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
	}

	Teardown()
}

func TestOraAbsolute(t *testing.T) {
	Setup()

	cpu.registers.A = 0xf0
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x0d)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
	}

	Teardown()
}

func TestOraAbsoluteX(t *testing.T) {
	Setup()

	cpu.registers.A = 0xf0
	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x1d)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
	}

	Teardown()
}

func TestOraAbsoluteY(t *testing.T) {
	Setup()

	cpu.registers.A = 0xf0
	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x19)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
	}

	Teardown()
}

func TestOraIndirectX(t *testing.T) {
	Setup()

	cpu.registers.A = 0xf0
	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x01)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x87)
	cpu.memory.store(0x0086, 0x00)
	cpu.memory.store(0x0087, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
	}

	Teardown()
}

func TestOraIndirectY(t *testing.T) {
	Setup()

	cpu.registers.A = 0xf0
	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x11)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x86)
	cpu.memory.store(0x0085, 0x00)
	cpu.memory.store(0x0087, 0x0f)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
	}

	Teardown()
}

func TestOraZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x09)
	cpu.memory.store(0x0101, 0x00)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestOraZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x09)
	cpu.memory.store(0x0101, 0x00)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestOraNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x81
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x09)
	cpu.memory.store(0x0101, 0x00)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestOraNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x09)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

//  BIT

func TestBitZeroPage(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x24)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x7f)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

func TestBitAbsolute(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x2c)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0x7f)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

func TestBitNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x24)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0xff)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestBitNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x24)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x7f)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

func TestBitVFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x24)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0xff)

	cpu.Execute()

	if cpu.registers.P&V == 0 {
		t.Error("V flag is not set")
	}

	Teardown()
}

func TestBitVFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x24)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x3f)

	cpu.Execute()

	if cpu.registers.P&V != 0 {
		t.Error("V flag is set")
	}

	Teardown()
}

func TestBitZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x00
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x24)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0xff)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestBitZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x24)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x3f)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

// ADC

func TestAdcImmediate(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x69)
	cpu.memory.store(0x0101, 0x02)

	cpu.Execute()

	if cpu.registers.A != 0x03 {
		t.Error("Register A is not 0x03")
	}

	Teardown()
}

func TestAdcZeroPage(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x65)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x02)

	cpu.Execute()

	if cpu.registers.A != 0x03 {
		t.Error("Register A is not 0x03")
	}

	Teardown()
}

func TestAdcZeroPageX(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x75)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x02)

	cpu.Execute()

	if cpu.registers.A != 0x03 {
		t.Error("Register A is not 0x03")
	}

	Teardown()
}

func TestAdcAbsolute(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x6d)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0x02)

	cpu.Execute()

	if cpu.registers.A != 0x03 {
		t.Error("Register A is not 0x03")
	}

	Teardown()
}

func TestAdcAbsoluteX(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x7d)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0x02)

	cpu.Execute()

	if cpu.registers.A != 0x03 {
		t.Error("Register A is not 0x03")
	}

	Teardown()
}

func TestAdcAbsoluteY(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x79)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0x02)

	cpu.Execute()

	if cpu.registers.A != 0x03 {
		t.Error("Register A is not 0x03")
	}

	Teardown()
}

func TestAdcIndirectX(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x61)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x87)
	cpu.memory.store(0x0086, 0x00)
	cpu.memory.store(0x0087, 0x02)

	cpu.Execute()

	if cpu.registers.A != 0x03 {
		t.Error("Register A is not 0x03")
	}

	Teardown()
}

func TestAdcIndirectY(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x71)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x86)
	cpu.memory.store(0x0085, 0x00)
	cpu.memory.store(0x0087, 0x02)

	cpu.Execute()

	if cpu.registers.A != 0x03 {
		t.Error("Register A is not 0x03")
	}

	Teardown()
}

func TestAdcCFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff // -1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x69)
	cpu.memory.store(0x0101, 0x01) // +1

	cpu.Execute()

	if cpu.registers.P&C == 0 {
		t.Error("C flag is not set")
	}

	Teardown()
}

func TestAdcCFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x00 // +0
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x69)
	cpu.memory.store(0x0101, 0x01) // +1

	cpu.Execute()

	if cpu.registers.P&C != 0 {
		t.Error("C flag is set")
	}

	Teardown()
}

func TestAdcZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x00 // +0
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x69)
	cpu.memory.store(0x0101, 0x00) // +0

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestAdcZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x00 // +0
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x69)
	cpu.memory.store(0x0101, 0xff) // -1

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestAdcVFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x7f // +127
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x69)
	cpu.memory.store(0x0101, 0x01) // +1

	cpu.Execute()

	if cpu.registers.P&V == 0 {
		t.Error("V flag is not set")
	}

	Teardown()
}

func TestAdcVFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01 // +1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x69)
	cpu.memory.store(0x0101, 0x01) // +1

	cpu.Execute()

	if cpu.registers.P&V != 0 {
		t.Error("V flag is set")
	}

	Teardown()
}

func TestAdcNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01 // +1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x69)
	cpu.memory.store(0x0101, 0xfe) // -2

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestAdcNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01 // +1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x69)
	cpu.memory.store(0x0101, 0x01) // +1

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// SBC

func TestSbcImmediate(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe9)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.A != 0x01 {
		t.Errorf("Register A is not 0x01")
	}

	Teardown()
}

func TestSbcZeroPage(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe5)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x01)

	cpu.Execute()

	if cpu.registers.A != 0x01 {
		t.Error("Register A is not 0x01")
	}

	Teardown()
}

func TestSbcZeroPageX(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02
	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xf5)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x01)

	cpu.Execute()

	if cpu.registers.A != 0x01 {
		t.Error("Register A is not 0x01")
	}

	Teardown()
}

func TestSbcAbsolute(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xed)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0x01)

	cpu.Execute()

	if cpu.registers.A != 0x01 {
		t.Error("Register A is not 0x01")
	}

	Teardown()
}

func TestSbcAbsoluteX(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02
	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xfd)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0x01)

	cpu.Execute()

	if cpu.registers.A != 0x01 {
		t.Error("Register A is not 0x01")
	}

	Teardown()
}

func TestSbcAbsoluteY(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02
	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xf9)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0x01)

	cpu.Execute()

	if cpu.registers.A != 0x01 {
		t.Error("Register A is not 0x01")
	}

	Teardown()
}

func TestSbcIndirectX(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02
	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe1)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x87)
	cpu.memory.store(0x0086, 0x00)
	cpu.memory.store(0x0087, 0x01)

	cpu.Execute()

	if cpu.registers.A != 0x01 {
		t.Error("Register A is not 0x01")
	}

	Teardown()
}

func TestSbcIndirectY(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02
	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xf1)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x86)
	cpu.memory.store(0x0085, 0x00)
	cpu.memory.store(0x0087, 0x01)

	cpu.Execute()

	if cpu.registers.A != 0x01 {
		t.Error("Register A is not 0x01")
	}

	Teardown()
}

func TestSbcCFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0xc4 // -60
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe9)
	cpu.memory.store(0x0101, 0x3c) // +60

	cpu.Execute()

	if cpu.registers.P&C == 0 {
		t.Error("C flag is not set")
	}

	Teardown()
}

func TestSbcCFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02 // +2
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe9)
	cpu.memory.store(0x0101, 0x04) // +4

	cpu.Execute()

	if cpu.registers.P&C != 0 {
		t.Error("C flag is set")
	}

	Teardown()
}

func TestSbcZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02 // +2
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe9)
	cpu.memory.store(0x0101, 0x02) // +2

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestSbcZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02 // +2
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe9)
	cpu.memory.store(0x0101, 0x01) // +1

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestSbcVFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x80 // -128
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe9)
	cpu.memory.store(0x0101, 0x01) // +1

	cpu.Execute()

	if cpu.registers.P&V == 0 {
		t.Error("V flag is not set")
	}

	Teardown()
}

func TestSbcVFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01 // +1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe9)
	cpu.memory.store(0x0101, 0x01) // +1

	cpu.Execute()

	if cpu.registers.P&V != 0 {
		t.Error("V flag is set")
	}

	Teardown()
}

func TestSbcNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0xfd // -3
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe9)
	cpu.memory.store(0x0101, 0x01) // +1

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestSbcNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02 // +2
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe9)
	cpu.memory.store(0x0101, 0x01) // +1

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// CMP

func TestCmpImmediate(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc9)
	cpu.memory.store(0x0101, 0xff)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCmpZeroPage(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc5)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0xff)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCmpZeroPageX(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xd5)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0xff)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCmpAbsolute(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xcd)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0xff)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCmpAbsoluteX(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xdd)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0xff)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCmpAbsoluteY(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xd9)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0xff)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCmpIndirectX(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc1)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x87)
	cpu.memory.store(0x0086, 0x00)
	cpu.memory.store(0x0087, 0xff)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCmpIndirectY(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.Y = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xd1)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x86)
	cpu.memory.store(0x0085, 0x00)
	cpu.memory.store(0x0087, 0xff)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCmpNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc9)
	cpu.memory.store(0x0101, 0x02)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestCmpNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc9)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

func TestCmpZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc9)
	cpu.memory.store(0x0101, 0x02)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCmpZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc9)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestCmpCFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc9)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&C == 0 {
		t.Error("C flag is not set")
	}

	Teardown()
}

func TestCmpCFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc9)
	cpu.memory.store(0x0101, 0x02)

	cpu.Execute()

	if cpu.registers.P&C != 0 {
		t.Error("C flag is set")
	}

	Teardown()
}

// CPX

func TestCpxImmediate(t *testing.T) {
	Setup()

	cpu.registers.X = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe0)
	cpu.memory.store(0x0101, 0xff)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCpxZeroPage(t *testing.T) {
	Setup()

	cpu.registers.X = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe4)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0xff)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCpxAbsolute(t *testing.T) {
	Setup()

	cpu.registers.X = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xec)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0xff)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCpxNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe0)
	cpu.memory.store(0x0101, 0x02)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestCpxNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe0)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

func TestCpxZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.X = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe0)
	cpu.memory.store(0x0101, 0x02)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCpxZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.X = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe0)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestCpxCFlagSet(t *testing.T) {
	Setup()

	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe0)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&C == 0 {
		t.Error("C flag is not set")
	}

	Teardown()
}

func TestCpxCFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe0)
	cpu.memory.store(0x0101, 0x02)

	cpu.Execute()

	if cpu.registers.P&C != 0 {
		t.Error("C flag is set")
	}

	Teardown()
}

// CPY

func TestCpyImmediate(t *testing.T) {
	Setup()

	cpu.registers.Y = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc0)
	cpu.memory.store(0x0101, 0xff)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCpyZeroPage(t *testing.T) {
	Setup()

	cpu.registers.Y = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc4)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0xff)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCpyAbsolute(t *testing.T) {
	Setup()

	cpu.registers.Y = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xcc)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0xff)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCpyNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.Y = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc0)
	cpu.memory.store(0x0101, 0x02)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestCpyNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.Y = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc0)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

func TestCpyZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.Y = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc0)
	cpu.memory.store(0x0101, 0x02)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestCpyZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.Y = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc0)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestCpyCFlagSet(t *testing.T) {
	Setup()

	cpu.registers.Y = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc0)
	cpu.memory.store(0x0101, 0x01)

	cpu.Execute()

	if cpu.registers.P&C == 0 {
		t.Error("C flag is not set")
	}

	Teardown()
}

func TestCpyCFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.Y = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc0)
	cpu.memory.store(0x0101, 0x02)

	cpu.Execute()

	if cpu.registers.P&C != 0 {
		t.Error("C flag is set")
	}

	Teardown()
}

// INC

func TestIncZeroPage(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe6)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0xfe)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

func TestIncZeroPageX(t *testing.T) {
	Setup()

	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xf6)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0xfe)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

func TestIncAbsolute(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xee)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0xfe)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

func TestIncAbsoluteX(t *testing.T) {
	Setup()

	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xfe)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0xfe)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}

func TestIncZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe6)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0xff) // -1

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestIncZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe6)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x00)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestIncNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe6)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0xfe) // -2

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestIncNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe6)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x00)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// INX

func TestInx(t *testing.T) {
	Setup()

	cpu.registers.X = 0xfe
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe8)

	cpu.Execute()

	if cpu.registers.X != 0xff {
		t.Error("Register X is not 0xff")
	}

	Teardown()
}

func TestInxZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.X = 0xff // -1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe8)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestInxZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe8)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestInxNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.X = 0xfe // -2
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe8)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestInxNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe8)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// INY

func TestIny(t *testing.T) {
	Setup()

	cpu.registers.Y = 0xfe // -2
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc8)

	cpu.Execute()

	if cpu.registers.Y != 0xff {
		t.Error("Register X is not 0xff")
	}

	Teardown()
}

func TestInyZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.Y = 0xff // -1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc8)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestInyZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.Y = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc8)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestInyNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.Y = 0xfe // -2
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc8)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestInyNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.Y = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc8)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// DEC

func TestDecZeroPage(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc6)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x02)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0x01 {
		t.Error("Memory is not 0x01")
	}

	Teardown()
}

func TestDecZeroPageX(t *testing.T) {
	Setup()

	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xd6)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x02)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0x01 {
		t.Error("Memory is not 0x01")
	}

	Teardown()
}

func TestDecAbsolute(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xce)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0x02)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0x01 {
		t.Error("Memory is not 0x01")
	}

	Teardown()
}

func TestDecAbsoluteX(t *testing.T) {
	Setup()

	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xde)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0x02)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0x01 {
		t.Error("Memory is not 0x01")
	}

	Teardown()
}

func TestDecZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc6)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x01)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestDecZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc6)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x02)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestDecNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc6)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x00)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestDecNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xc6)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x01)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// DEX

func TestDex(t *testing.T) {
	Setup()

	cpu.registers.X = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xca)

	cpu.Execute()

	if cpu.registers.X != 0x01 {
		t.Error("Register X is not 0x01")
	}

	Teardown()
}

func TestDexZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xca)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestDexZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.X = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xca)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestDexNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.X = 0x00
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xca)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestDexNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xca)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// DEY

func TestDey(t *testing.T) {
	Setup()

	cpu.registers.Y = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x88)

	cpu.Execute()

	if cpu.registers.Y != 0x01 {
		t.Error("Register X is not 0x01")
	}

	Teardown()
}

func TestDeyZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.Y = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x88)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestDeyZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.Y = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x88)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestDeyNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.Y = 0x00
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x88)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestDeyNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.Y = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x88)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// ASL

func TestAslAccumulator(t *testing.T) {
	Setup()

	cpu.registers.A = 0x2
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x0a)

	cpu.Execute()

	if cpu.registers.A != 0x04 {
		t.Error("Register A is not 0x04")
	}

	Teardown()
}

func TestAslZeroPage(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x06)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x02)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0x04 {
		t.Error("Memory is not 0x04")
	}

	Teardown()
}

func TestAslZeroPageX(t *testing.T) {
	Setup()

	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x16)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x02)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0x04 {
		t.Error("Memory is not 0x04")
	}

	Teardown()
}

func TestAslAbsolute(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x0e)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0x02)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0x04 {
		t.Error("Memory is not 0x04")
	}

	Teardown()
}

func TestAslAbsoluteX(t *testing.T) {
	Setup()

	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x1e)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0x02)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0x04 {
		t.Error("Memory is not 0x04")
	}

	Teardown()
}

func TestAslCFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x0a)

	cpu.Execute()

	if cpu.registers.P&C == 0 {
		t.Error("C flag is not set")
	}

	Teardown()
}

func TestAslCFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x0a)

	cpu.Execute()

	if cpu.registers.P&C != 0 {
		t.Error("C flag is set")
	}

	Teardown()
}

func TestAslZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x00
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x0a)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestAslZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x0a)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestAslNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0xfe
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x0a)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestAslNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x0a)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// LSR

func TestLsrAccumulator(t *testing.T) {
	Setup()

	cpu.registers.A = 0x2
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x4a)

	cpu.Execute()

	if cpu.registers.A != 0x01 {
		t.Error("Register A is not 0x01")
	}

	Teardown()
}

func TestLsrZeroPage(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x46)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x02)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0x01 {
		t.Error("Memory is not 0x01")
	}

	Teardown()
}

func TestLsrZeroPageX(t *testing.T) {
	Setup()

	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x56)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x02)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0x01 {
		t.Error("Memory is not 0x01")
	}

	Teardown()
}

func TestLsrAbsolute(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x4e)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0x02)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0x01 {
		t.Error("Memory is not 0x01")
	}

	Teardown()
}

func TestLsrAbsoluteX(t *testing.T) {
	Setup()

	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x5e)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0x02)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0x01 {
		t.Error("Memory is not 0x01")
	}

	Teardown()
}

func TestLsrCFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0xff
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x4a)

	cpu.Execute()

	if cpu.registers.P&C == 0 {
		t.Error("C flag is not set")
	}

	Teardown()
}

func TestLsrCFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x10
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x4a)

	cpu.Execute()

	if cpu.registers.P&C != 0 {
		t.Error("C flag is set")
	}

	Teardown()
}

func TestLsrZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x4a)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestLsrZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x4a)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

// func TestLsrNFlagSet(t *testing.T) { }
// not tested, N bit always set to 0

func TestLsrNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x4a)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// ROL

func TestRolAccumulator(t *testing.T) {
	Setup()

	cpu.registers.P |= C
	cpu.registers.A = 0x2
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x2a)

	cpu.Execute()

	if cpu.registers.A != 0x05 {
		t.Error("Register A is not 0x05")
	}

	Teardown()
}

func TestRolZeroPage(t *testing.T) {
	Setup()

	cpu.registers.P |= C
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x26)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x02)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0x05 {
		t.Error("Memory is not 0x05")
	}

	Teardown()
}

func TestRolZeroPageX(t *testing.T) {
	Setup()

	cpu.registers.P |= C
	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x36)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x02)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0x05 {
		t.Error("Memory is not 0x05")
	}

	Teardown()
}

func TestRolAbsolute(t *testing.T) {
	Setup()

	cpu.registers.P |= C
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x2e)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0x02)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0x05 {
		t.Error("Memory is not 0x05")
	}

	Teardown()
}

func TestRolAbsoluteX(t *testing.T) {
	Setup()

	cpu.registers.P |= C
	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x3e)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0x02)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0x05 {
		t.Error("Memory is not 0x05")
	}

	Teardown()
}

func TestRolCFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x80
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x2a)

	cpu.Execute()

	if cpu.registers.P&C == 0 {
		t.Error("C flag is not set")
	}

	Teardown()
}

func TestRolCFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x2a)

	cpu.Execute()

	if cpu.registers.P&C != 0 {
		t.Error("C flag is set")
	}

	Teardown()
}

func TestRolZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x00
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x2a)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestRolZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x2a)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestRolNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0xfe
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x2a)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestRolNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x2a)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// ROR

func TestRorAccumulator(t *testing.T) {
	Setup()

	cpu.registers.P |= C
	cpu.registers.A = 0x08
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x6a)

	cpu.Execute()

	if cpu.registers.A != 0x84 {
		t.Error("Register A is not 0x84")
	}

	Teardown()
}

func TestRorZeroPage(t *testing.T) {
	Setup()

	cpu.registers.P |= C
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x66)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0084, 0x08)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0x84 {
		t.Error("Memory is not 0x84")
	}

	Teardown()
}

func TestRorZeroPageX(t *testing.T) {
	Setup()

	cpu.registers.P |= C
	cpu.registers.X = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x76)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0085, 0x08)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0x84 {
		t.Error("Memory is not 0x84")
	}

	Teardown()
}

func TestRorAbsolute(t *testing.T) {
	Setup()

	cpu.registers.P |= C
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x6e)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0084, 0x08)

	cpu.Execute()

	if cpu.memory.fetch(0x0084) != 0x84 {
		t.Error("Memory is not 0x84")
	}

	Teardown()
}

func TestRorAbsoluteX(t *testing.T) {
	Setup()

	cpu.registers.P |= C
	cpu.registers.X = 1
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x7e)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x00)
	cpu.memory.store(0x0085, 0x08)

	cpu.Execute()

	if cpu.memory.fetch(0x0085) != 0x84 {
		t.Error("Memory is not 0x84")
	}

	Teardown()
}

func TestRorCFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x6a)

	cpu.Execute()

	if cpu.registers.P&C == 0 {
		t.Error("C flag is not set")
	}

	Teardown()
}

func TestRorCFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x10
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x6a)

	cpu.Execute()

	if cpu.registers.P&C != 0 {
		t.Error("C flag is set")
	}

	Teardown()
}

func TestRorZFlagSet(t *testing.T) {
	Setup()

	cpu.registers.A = 0x00
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x6a)

	cpu.Execute()

	if cpu.registers.P&Z == 0 {
		t.Error("Z flag is not set")
	}

	Teardown()
}

func TestRorZFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.A = 0x02
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x6a)

	cpu.Execute()

	if cpu.registers.P&Z != 0 {
		t.Error("Z flag is set")
	}

	Teardown()
}

func TestRorNFlagSet(t *testing.T) {
	Setup()

	cpu.registers.P |= C
	cpu.registers.A = 0xfe
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x6a)

	cpu.Execute()

	if cpu.registers.P&N == 0 {
		t.Error("N flag is not set")
	}

	Teardown()
}

func TestRorNFlagUnset(t *testing.T) {
	Setup()

	cpu.registers.P &^= C
	cpu.registers.A = 0x01
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x6a)

	cpu.Execute()

	if cpu.registers.P&N != 0 {
		t.Error("N flag is set")
	}

	Teardown()
}

// JMP

func TestJmpAbsolute(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x4c)
	cpu.memory.store(0x0101, 0xff)
	cpu.memory.store(0x0102, 0x01)

	cpu.Execute()

	if cpu.registers.PC != 0x01ff {
		t.Error("Register PC is not 0x01ff")
	}

	Teardown()
}

func TestJmpIndirect(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x6c)
	cpu.memory.store(0x0101, 0x84)
	cpu.memory.store(0x0102, 0x01)
	cpu.memory.store(0x0184, 0xff)
	cpu.memory.store(0x0185, 0xff)

	cpu.Execute()

	if cpu.registers.PC != 0xffff {
		t.Error("Register PC is not 0xffff")
	}

	Teardown()
}

// JSR

func TestJsr(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x20)
	cpu.memory.store(0x0101, 0xff)
	cpu.memory.store(0x0102, 0x01)

	cpu.Execute()

	if cpu.registers.PC != 0x01ff {
		t.Error("Register PC is not 0x01ff")
	}

	if cpu.memory.fetch(0x01fd) != 0x01 {
		t.Error("Memory is not 0x01")
	}

	if cpu.memory.fetch(0x01fc) != 0x02 {
		t.Error("Memory is not 0x02")
	}

	Teardown()
}

// RTS

func TestRts(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100
	cpu.push(0x01)
	cpu.push(0x02)

	cpu.memory.store(0x0100, 0x60)

	cpu.Execute()

	if cpu.registers.PC != 0x0103 {
		t.Error("Register PC is not 0x0103")
	}

	Teardown()
}

// BCC

func TestBcc(t *testing.T) {
	Setup()

	cpu.registers.P |= C
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x90)

	cycles, _ := cpu.Execute()

	if cycles != 2 {
		t.Error("Cycles is not 2")
	}

	if cpu.registers.PC != 0x0102 {
		t.Error("Register PC is not 0x0102")
	}

	cpu.registers.P &^= C
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x90)
	cpu.memory.store(0x0101, 0x02) // +2

	cycles, _ = cpu.Execute()

	if cycles != 3 {
		t.Error("Cycles is not 3")
	}

	if cpu.registers.PC != 0x0104 {
		t.Error("Register PC is not 0x0104")
	}

	cpu.registers.P &^= C
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x90)
	cpu.memory.store(0x0101, 0xfd) // -3

	cycles, _ = cpu.Execute()

	if cycles != 4 {
		t.Error("Cycles is not 4")
	}

	if cpu.registers.PC != 0x00ff {
		t.Error("Register PC is not 0x00ff")
	}

	Teardown()
}

// BCS

func TestBcs(t *testing.T) {
	Setup()

	cpu.registers.P |= C
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xb0)
	cpu.memory.store(0x0101, 0x02) // +2

	cpu.Execute()

	if cpu.registers.PC != 0x0104 {
		t.Error("Register PC is not 0x0104")
	}

	cpu.registers.P |= C
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xb0)
	cpu.memory.store(0x0101, 0xfe) // -2

	cpu.Execute()

	if cpu.registers.PC != 0x0100 {
		t.Error("Register PC is not 0x0100")
	}

	Teardown()
}

// BEQ

func TestBeq(t *testing.T) {
	Setup()

	cpu.registers.P |= Z
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xf0)
	cpu.memory.store(0x0101, 0x02) // +2

	cpu.Execute()

	if cpu.registers.PC != 0x0104 {
		t.Error("Register PC is not 0x0104")
	}

	cpu.registers.P |= Z
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xf0)
	cpu.memory.store(0x0101, 0xfe) // -2

	cpu.Execute()

	if cpu.registers.PC != 0x0100 {
		t.Error("Register PC is not 0x0100")
	}

	Teardown()
}

// BMI

func TestBmi(t *testing.T) {
	Setup()

	cpu.registers.P |= N
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x30)
	cpu.memory.store(0x0101, 0x02) // +2

	cpu.Execute()

	if cpu.registers.PC != 0x0104 {
		t.Error("Register PC is not 0x0104")
	}

	cpu.registers.P |= N
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x30)
	cpu.memory.store(0x0101, 0xfe) // -2

	cpu.Execute()

	if cpu.registers.PC != 0x0100 {
		t.Error("Register PC is not 0x0100")
	}

	Teardown()
}

// BNE

func TestBne(t *testing.T) {
	Setup()

	cpu.registers.P &^= Z
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xd0)
	cpu.memory.store(0x0101, 0x02) // +2

	cpu.Execute()

	if cpu.registers.PC != 0x0104 {
		t.Error("Register PC is not 0x0104")
	}

	cpu.registers.P &^= Z
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xd0)
	cpu.memory.store(0x0101, 0xfe) // -2

	cpu.Execute()

	if cpu.registers.PC != 0x0100 {
		t.Error("Register PC is not 0x0100")
	}

	Teardown()
}

// BPL

func TestBpl(t *testing.T) {
	Setup()

	cpu.registers.P &^= N
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x10)
	cpu.memory.store(0x0101, 0x02) // +2

	cpu.Execute()

	if cpu.registers.PC != 0x0104 {
		t.Error("Register PC is not 0x0104")
	}

	cpu.registers.P &^= N
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x10)
	cpu.memory.store(0x0101, 0xfe) // -2

	cpu.Execute()

	if cpu.registers.PC != 0x0100 {
		t.Error("Register PC is not 0x0100")
	}

	Teardown()
}

// BVC

func TestBvc(t *testing.T) {
	Setup()

	cpu.registers.P &^= V
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x50)
	cpu.memory.store(0x0101, 0x02) // +2

	cpu.Execute()

	if cpu.registers.PC != 0x0104 {
		t.Error("Register PC is not 0x0104")
	}

	cpu.registers.P &^= V
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x50)
	cpu.memory.store(0x0101, 0xfe) // -2

	cpu.Execute()

	if cpu.registers.PC != 0x0100 {
		t.Error("Register PC is not 0x0100")
	}

	Teardown()
}

// BVS

func TestBvs(t *testing.T) {
	Setup()

	cpu.registers.P |= V
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x70)
	cpu.memory.store(0x0101, 0x02) // +2

	cpu.Execute()

	if cpu.registers.PC != 0x0104 {
		t.Error("Register PC is not 0x0104")
	}

	cpu.registers.P |= V
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x70)
	cpu.memory.store(0x0101, 0xfe) // -2

	cpu.Execute()

	if cpu.registers.PC != 0x0100 {
		t.Error("Register PC is not 0x0100")
	}

	Teardown()
}

// CLC

func TestClc(t *testing.T) {
	Setup()

	cpu.registers.P &^= C
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x18)

	cpu.Execute()

	if cpu.registers.P&C != 0 {
		t.Error("C flag is set")
	}

	cpu.registers.P |= C
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x18)

	cpu.Execute()

	if cpu.registers.P&C != 0 {
		t.Error("C flag is set")
	}

	Teardown()
}

// CLD

func TestCld(t *testing.T) {
	Setup()

	cpu.registers.P &^= D
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xd8)

	cpu.Execute()

	if cpu.registers.P&D != 0 {
		t.Error("D flag is set")
	}

	cpu.registers.P |= D
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xd8)

	cpu.Execute()

	if cpu.registers.P&D != 0 {
		t.Error("D flag is set")
	}

	Teardown()
}

// CLI

func TestCli(t *testing.T) {
	Setup()

	cpu.registers.P &^= I
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x58)

	cpu.Execute()

	if cpu.registers.P&I != 0 {
		t.Error("I flag is set")
	}

	cpu.registers.P |= I
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x58)

	cpu.Execute()

	if cpu.registers.P&I != 0 {
		t.Error("I flag is set")
	}

	Teardown()
}

// CLV

func TestClv(t *testing.T) {
	Setup()

	cpu.registers.P &^= V
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xb8)

	cpu.Execute()

	if cpu.registers.P&V != 0 {
		t.Error("V flag is set")
	}

	cpu.registers.P |= V
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xb8)

	cpu.Execute()

	if cpu.registers.P&V != 0 {
		t.Error("V flag is set")
	}

	Teardown()
}

// SEC

func TestSec(t *testing.T) {
	Setup()

	cpu.registers.P &^= C
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x38)

	cpu.Execute()

	if cpu.registers.P&C == 0 {
		t.Error("C flag is not set")
	}

	cpu.registers.P |= C
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x38)

	cpu.Execute()

	if cpu.registers.P&C == 0 {
		t.Error("C flag is not set")
	}

	Teardown()
}

// SED

func TestSed(t *testing.T) {
	Setup()

	cpu.registers.P &^= D
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xf8)

	cpu.Execute()

	if cpu.registers.P&D == 0 {
		t.Error("D flag is not set")
	}

	cpu.registers.P |= D
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xf8)

	cpu.Execute()

	if cpu.registers.P&D == 0 {
		t.Error("D flag is not set")
	}

	Teardown()
}

// SEI

func TestSei(t *testing.T) {
	Setup()

	cpu.registers.P &^= I
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x78)

	cpu.Execute()

	if cpu.registers.P&I == 0 {
		t.Error("I flag is not set")
	}

	cpu.registers.P |= I
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x78)

	cpu.Execute()

	if cpu.registers.P&I == 0 {
		t.Error("I flag is not set")
	}

	Teardown()
}

// BRK

func TestBrk(t *testing.T) {
	Setup()

	cpu.registers.P = 0xff & (^B)
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0x00)
	cpu.memory.store(0xfffe, 0xff)
	cpu.memory.store(0xffff, 0x01)

	cpu.Execute()

	if cpu.pull() != 0xff {
		t.Error("Memory is not 0xff")
	}

	if cpu.pull() != 0x02 {
		t.Error("Memory is not 0x02")
	}

	if cpu.pull() != 0x01 {
		t.Error("Memory is not 0x01")
	}

	if cpu.registers.PC != 0x01ff {
		t.Error("Register PC is not 0x01ff")
	}

	Teardown()
}

// RTI

func TestRti(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100
	cpu.push(0x01)
	cpu.push(0x02)
	cpu.push(0x03)

	cpu.memory.store(0x0100, 0x40)

	cpu.Execute()

	if cpu.registers.P != 0x03 {
		t.Error("Register P is not 0x03")
	}

	if cpu.registers.PC != 0x0102 {
		t.Error("Register PC is not 0x0102")
	}

	Teardown()
}
