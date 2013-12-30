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

func TestLdaImmediate(t *testing.T) {
	Setup()

	cpu.registers.PC = 0x0100

	cpu.memory.store(0x0100, 0xa9)
	cpu.memory.store(0x0101, 0xff)

	cpu.Execute()

	if cpu.registers.A != 0xff {
		t.Error("LDA immediate: Register A is not 0xff")
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
		t.Error("LDA zero page: Register A is not 0xff")
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
		t.Error("LDA zero page,x: Register A is not 0xff")
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
		t.Error("LDA absolute: Register A is not 0xff")
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
		t.Error("LDA absolute,x: Register A is not 0xff")
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
		t.Error("LDA absolute,y: Register A is not 0xff")
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
		t.Error("LDA (indirect,x): Register A is not 0xff")
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
		t.Error("LDA (indirect),y: Register A is not 0xff")
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
		t.Error("LDA Z Flag Set: Z flag is not set")
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
		t.Error("LDA Z Flag Unset: Z flag is set")
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
		t.Error("LDA N Flag Set: N flag is not set")
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
		t.Error("LDA N Flag Unset: N flag is set")
	}

	Teardown()
}
