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
