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
	reg.SP = 0xff
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

func (cpu *Cpu) setZNFlags(value uint8) uint8 {
	cpu.registers.P &= ^Z
	cpu.registers.P &= ^N

	switch {
	case value == 0:
		cpu.registers.P |= Z
	case value&(uint8(1)<<7) != 0:
		cpu.registers.P |= N
	}

	return value
}

func (cpu *Cpu) load(address uint16, register *uint8) {
	*register = cpu.setZNFlags(cpu.memory.fetch(address))
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

func (cpu *Cpu) Lda(address uint16) {
	cpu.load(address, &cpu.registers.A)
}

func (cpu *Cpu) Ldx(address uint16) {
	cpu.load(address, &cpu.registers.X)
}

func (cpu *Cpu) Ldy(address uint16) {
	cpu.load(address, &cpu.registers.Y)
}

func (cpu *Cpu) store(address uint16, value uint8) {
	cpu.memory.store(address, value)
}

func (cpu *Cpu) Sta(address uint16) {
	cpu.store(address, cpu.registers.A)
}

func (cpu *Cpu) Stx(address uint16) {
	cpu.store(address, cpu.registers.X)
}

func (cpu *Cpu) Sty(address uint16) {
	cpu.store(address, cpu.registers.Y)
}

func (cpu *Cpu) transfer(from uint8, to *uint8) {
	*to = cpu.setZNFlags(from)
}

func (cpu *Cpu) Tax() {
	cpu.transfer(cpu.registers.A, &cpu.registers.X)
}

func (cpu *Cpu) Tay() {
	cpu.transfer(cpu.registers.A, &cpu.registers.Y)
}

func (cpu *Cpu) Txa() {
	cpu.transfer(cpu.registers.X, &cpu.registers.A)
}

func (cpu *Cpu) Tya() {
	cpu.transfer(cpu.registers.Y, &cpu.registers.A)
}

func (cpu *Cpu) Tsx() {
	cpu.transfer(cpu.registers.SP, &cpu.registers.X)
}

func (cpu *Cpu) Txs() {
	cpu.transfer(cpu.registers.X, &cpu.registers.SP)
}

func (cpu *Cpu) push(value uint8) {
	cpu.memory.store(0x0100|uint16(cpu.registers.SP), value)
	cpu.registers.SP--
}

func (cpu *Cpu) pull() (value uint8) {
	cpu.registers.SP++
	value = cpu.memory.fetch(0x0100 | uint16(cpu.registers.SP))
	return
}

func (cpu *Cpu) Pha() {
	cpu.push(cpu.registers.A)
}

func (cpu *Cpu) Php() {
	cpu.push(uint8(cpu.registers.P))
}

func (cpu *Cpu) Pla() {
	cpu.registers.A = cpu.setZNFlags(cpu.pull())
}

func (cpu *Cpu) Plp() {
	cpu.registers.P = Status(cpu.pull())
}

func (cpu *Cpu) And(address uint16) {
	cpu.registers.A = cpu.setZNFlags(cpu.registers.A & cpu.memory.fetch(address))
}

func (cpu *Cpu) Eor(address uint16) {
	cpu.registers.A = cpu.setZNFlags(cpu.registers.A ^ cpu.memory.fetch(address))
}

func (cpu *Cpu) Ora(address uint16) {
	cpu.registers.A = cpu.setZNFlags(cpu.registers.A | cpu.memory.fetch(address))
}
