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
		t.Error("Memory is not 0xff")
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
		t.Errorf("Register A is 0x%x not 0x01", cpu.registers.A)
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

	cpu.registers.A = 0xd0 // -60
	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xe9)
	cpu.memory.store(0x0101, 0x40) // +60

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
