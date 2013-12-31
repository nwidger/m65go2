package _65go2

import (
	"testing"
	"time"
)

const rate time.Duration = 6708 * time.Nanosecond
const divisor = 12

var cpu *Cpu

func Setup() {
	cpu = NewCpu(NewBasicMemory(), NewClock(rate, divisor))
	cpu.Reset()
	go cpu.clock.start()
}

func Teardown() {
	cpu.clock.stop()
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

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
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

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
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

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("Register A is not 0xff")
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

func TestPlp(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100
	cpu.push(0xff)

	cpu.memory.store(0x0100, 0x28)

	cpu.Execute()

	if cpu.registers.P != 0xff {
		t.Error("Memory is not 0xff")
	}

	Teardown()
}
