package _65go2

import (
	"fmt"
	"os"
)

type Status uint8

const (
	C Status = 1 << iota // carry flag
	Z                    // zero flag
	I                    // interrupt disable
	D                    // decimal mode
	B                    // break command
	_                    // -UNUSED-
	V                    // overflow flag
	N                    // negative flag
)

type Registers struct {
	A  uint8  // accumulator
	X  uint8  // index register X
	Y  uint8  // index register Y
	P  Status // processor status
	SP uint8  // stack pointer
	PC uint16 // program counter
}

func NewRegisters() Registers {
	return Registers{}
}

func (reg *Registers) reset() {
	reg.A = 0
	reg.X = 0
	reg.Y = 0
	reg.P = 0
	reg.SP = 0
	reg.PC = 0
}

type Cpu struct {
	clock        Clock
	registers    Registers
	memory       Memory
	instructions InstructionTable
}

func NewCpu(mem Memory, clock Clock) *Cpu {
	return &Cpu{clock: clock, registers: NewRegisters(), memory: mem, instructions: NewInstructionTable()}
}

func (cpu *Cpu) Reset() {
	cpu.registers.reset()
	cpu.memory.reset()
}

func (cpu *Cpu) Execute() {
	ticks := cpu.clock.ticks

	// fetch
	opcode := OpCode(cpu.memory.fetch(cpu.registers.PC))
	inst, ok := cpu.instructions[opcode]

	if !ok {
		fmt.Printf("No such opcode 0x%x\n", opcode)
		os.Exit(1)
	}

	// execute
	cpu.registers.PC++
	ticks += uint64(inst.exec(cpu))

	// count cycles
	cpu.clock.await(ticks)
}

func (cpu *Cpu) Ld(address uint16, register *uint8) {
	value := cpu.memory.fetch(address)
	*register = value

	cpu.registers.P &= ^Z
	cpu.registers.P &= ^N

	switch {
	case value == 0:
		cpu.registers.P |= Z
	case value&(uint8(1)<<7) != 0:
		cpu.registers.P |= N
	}
}

func (cpu *Cpu) Lda(address uint16) {
	cpu.Ld(address, &cpu.registers.A)
}

func (cpu *Cpu) Ldx(address uint16) {
	cpu.Ld(address, &cpu.registers.X)
}

func (cpu *Cpu) Ldy(address uint16) {
	cpu.Ld(address, &cpu.registers.Y)
}

func (cpu *Cpu) St(address uint16, value uint8) {
	cpu.memory.store(address, value)
}

func (cpu *Cpu) Sta(address uint16) {
	cpu.St(address, cpu.registers.A)
}

func (cpu *Cpu) Stx(address uint16) {
	cpu.St(address, cpu.registers.X)
}

func (cpu *Cpu) Sty(address uint16) {
	cpu.St(address, cpu.registers.Y)
}

func (cpu *Cpu) immediateAddress() (result uint16) {
	result = cpu.registers.PC
	cpu.registers.PC++
	return
}

func (cpu *Cpu) zeroPageAddress() (result uint16) {
	result = uint16(cpu.memory.fetch(cpu.registers.PC))
	cpu.registers.PC++
	return
}

func (cpu *Cpu) zeroPageIndexedAddress(index uint8) (result uint16) {
	result = uint16(cpu.memory.fetch(cpu.registers.PC) + index)
	cpu.registers.PC++
	return
}

func (cpu *Cpu) absoluteAddress() (result uint16) {
	low := cpu.memory.fetch(cpu.registers.PC)
	high := cpu.memory.fetch(cpu.registers.PC + 1)
	cpu.registers.PC += 2

	result = (uint16(high) << 8) | uint16(low)
	return
}

func (cpu *Cpu) absoluteIndexedAddress(index uint8, cycles *uint16) (result uint16) {
	low := cpu.memory.fetch(cpu.registers.PC)
	high := cpu.memory.fetch(cpu.registers.PC + 1)
	cpu.registers.PC += 2

	address := (uint16(high) << 8) | uint16(low)
	result = address + uint16(index)

	if cycles != nil && !SamePage(address, result) {
		*cycles++
	}

	return
}

func (cpu *Cpu) indexedIndirectAddress() (result uint16) {
	address := uint16(cpu.memory.fetch(cpu.registers.PC) + cpu.registers.X)
	cpu.registers.PC++

	low := cpu.memory.fetch(address)
	high := cpu.memory.fetch(address + 1)

	result = (uint16(high) << 8) | uint16(low)
	return
}

func (cpu *Cpu) indirectIndexedAddress(cycles *uint16) (result uint16) {
	address := uint16(cpu.memory.fetch(cpu.registers.PC))
	cpu.registers.PC++

	low := cpu.memory.fetch(address)
	high := cpu.memory.fetch(address + 1)

	address = (uint16(high) << 8) | uint16(low)

	result = address + uint16(cpu.registers.Y)

	if cycles != nil && !SamePage(address, result) {
		*cycles++
	}

	return
}
