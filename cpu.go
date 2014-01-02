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

func (cpu *Cpu) setZFlag(value uint8) uint8 {
	if value == 0 {
		cpu.registers.P |= Z
	} else {
		cpu.registers.P &= ^Z
	}

	return value
}

func (cpu *Cpu) setNFlag(value uint8) uint8 {
	if value&(uint8(1)<<7) != 0 {
		cpu.registers.P |= N
	} else {
		cpu.registers.P &= ^N
	}

	return value
}

func (cpu *Cpu) setZNFlags(value uint8) uint8 {
	cpu.setZFlag(value)
	cpu.setNFlag(value)
	return value
}

func (cpu *Cpu) setCFlagAddition(value int16) uint16 {
	if value > 0xff {
		cpu.registers.P |= C
	} else {
		cpu.registers.P &= ^C
	}

	return uint16(value)
}

func (cpu *Cpu) setCFlagSubtraction(value int16) uint16 {
	if value >= 0 {
		cpu.registers.P |= C
	} else {
		cpu.registers.P &= ^C
	}

	return uint16(value)
}

func (cpu *Cpu) setVFlagAddition(term1 uint16, term2 uint16, result uint16) uint16 {
	if ((term1^term2)&uint16(N) == 0) && ((term1^result)&uint16(N) == uint16(N)) {
		cpu.registers.P |= V
	} else {
		cpu.registers.P &= ^V
	}

	return result
}

func (cpu *Cpu) setVFlagSubtraction(term1 uint16, term2 uint16, result uint16) uint16 {
	if ((term1^result)&uint16(N)) != 0 && ((term1^term2)&uint16(N)) != 0 {
		cpu.registers.P |= V
	} else {
		cpu.registers.P &= ^V
	}

	return result
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

func (cpu *Cpu) relativeAddress() (result uint16) {
	value := uint16(cpu.memory.fetch(cpu.registers.PC))
	cpu.registers.PC++

	if value > 0x7f {
		result = cpu.registers.PC - (0x0100 - value)
	} else {
		result = cpu.registers.PC + value
	}

	return
}

func (cpu *Cpu) absoluteAddress() (result uint16) {
	low := cpu.memory.fetch(cpu.registers.PC)
	high := cpu.memory.fetch(cpu.registers.PC + 1)
	cpu.registers.PC += 2

	result = (uint16(high) << 8) | uint16(low)
	return
}

func (cpu *Cpu) indirectAddress() (result uint16) {
	low := cpu.memory.fetch(cpu.registers.PC)
	high := cpu.memory.fetch(cpu.registers.PC + 1)
	cpu.registers.PC += 2

	// XXX: The 6502 had a bug in which it incremented only the
	// high byte instead of the whole 16-bit address when
	// computing the address.
	//
	// See http://www.obelisk.demon.co.uk/6502/reference.html#JMP
	// and http://www.6502.org/tutorials/6502opcodes.html#JMP for
	// details
	a_high := (uint16(high) << 8) | uint16(low+1)
	a_low := (uint16(high) << 8) | uint16(low)

	low = cpu.memory.fetch(a_low)
	high = cpu.memory.fetch(a_high)

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

func (cpu *Cpu) Bit(address uint16) {
	value := cpu.memory.fetch(address)
	cpu.setZFlag(value & cpu.registers.A)
	cpu.registers.P = Status(uint8(cpu.registers.P) | (value & 0xc0))
}

func (cpu *Cpu) Adc(address uint16) {
	orig := uint16(cpu.registers.A)
	value := uint16(cpu.memory.fetch(address))

	result := cpu.setCFlagAddition(int16(orig) + int16(value) + int16(cpu.registers.P&C))
	cpu.registers.A = cpu.setZNFlags(uint8(cpu.setVFlagAddition(orig, value, result)))
}

func (cpu *Cpu) Sbc(address uint16) {
	orig := uint16(cpu.registers.A)
	value := uint16(cpu.memory.fetch(address))

	result := cpu.setCFlagSubtraction(int16(orig) - int16(value) - int16(cpu.registers.P&C))
	cpu.registers.A = cpu.setZNFlags(uint8(cpu.setVFlagSubtraction(orig, value, result)))
}

func (cpu *Cpu) compare(address uint16, register uint8) {
	value := cpu.memory.fetch(address)
	cpu.setZNFlags(uint8(cpu.setCFlagSubtraction(int16(register) - int16(value))))
}

func (cpu *Cpu) Cmp(address uint16) {
	cpu.compare(address, cpu.registers.A)
}

func (cpu *Cpu) Cpx(address uint16) {
	cpu.compare(address, cpu.registers.X)
}

func (cpu *Cpu) Cpy(address uint16) {
	cpu.compare(address, cpu.registers.Y)
}

func (cpu *Cpu) Inc(address uint16) {
	cpu.memory.store(address, cpu.setZNFlags(cpu.memory.fetch(address)+1))
}

func (cpu *Cpu) increment(register *uint8) {
	*register = cpu.setZNFlags(*register + 1)
}

func (cpu *Cpu) Inx() {
	cpu.increment(&cpu.registers.X)
}

func (cpu *Cpu) Iny() {
	cpu.increment(&cpu.registers.Y)
}

func (cpu *Cpu) Dec(address uint16) {
	cpu.memory.store(address, cpu.setZNFlags(cpu.memory.fetch(address)-1))
}

func (cpu *Cpu) decrement(register *uint8) {
	*register = cpu.setZNFlags(*register - 1)
}

func (cpu *Cpu) Dex() {
	cpu.decrement(&cpu.registers.X)
}

func (cpu *Cpu) Dey() {
	cpu.decrement(&cpu.registers.Y)
}

type Direction int

const (
	left Direction = iota
	right
)

func (cpu *Cpu) shift(direction Direction, value uint8, store func(uint8)) {
	bit := uint8(0)

	switch direction {
	case left:
		bit = value & uint8(N)
		value <<= 1
	case right:
		bit = value & uint8(C)
		value >>= 1
	}

	if bit == 0 {
		cpu.registers.P &= ^C
	} else {
		cpu.registers.P |= C
	}

	store(cpu.setZNFlags(value))
}

func (cpu *Cpu) AslA() {
	cpu.shift(left, cpu.registers.A, func(value uint8) { cpu.registers.A = value })
}

func (cpu *Cpu) Asl(address uint16) {
	cpu.shift(left, cpu.memory.fetch(address), func(value uint8) { cpu.memory.store(address, value) })
}

func (cpu *Cpu) LsrA() {
	cpu.shift(right, cpu.registers.A, func(value uint8) { cpu.registers.A = value })
}

func (cpu *Cpu) Lsr(address uint16) {
	cpu.shift(right, cpu.memory.fetch(address), func(value uint8) { cpu.memory.store(address, value) })
}

func (cpu *Cpu) rotate(direction Direction, value uint8, store func(uint8)) {
	bit := uint8(0)

	switch direction {
	case left:
		bit = value & uint8(N)
		value = ((value << 1) & uint8(^C)) | uint8(cpu.registers.P&C)
	case right:
		bit = value & uint8(C)
		value = ((value >> 1) & uint8(^N)) | uint8((cpu.registers.P&C)<<7)
	}

	if bit == 0 {
		cpu.registers.P &= ^C
	} else {
		cpu.registers.P |= C
	}

	store(cpu.setZNFlags(value))
}

func (cpu *Cpu) RolA() {
	cpu.rotate(left, cpu.registers.A, func(value uint8) { cpu.registers.A = value })
}

func (cpu *Cpu) Rol(address uint16) {
	cpu.rotate(left, cpu.memory.fetch(address), func(value uint8) { cpu.memory.store(address, value) })
}

func (cpu *Cpu) RorA() {
	cpu.rotate(right, cpu.registers.A, func(value uint8) { cpu.registers.A = value })
}

func (cpu *Cpu) Ror(address uint16) {
	cpu.rotate(right, cpu.memory.fetch(address), func(value uint8) { cpu.memory.store(address, value) })
}

func (cpu *Cpu) Jmp(address uint16) {
	cpu.registers.PC = address
}

func (cpu *Cpu) Jsr(address uint16) {
	value := cpu.registers.PC - 1

	cpu.push(uint8(value >> 8))
	cpu.push(uint8(value))

	cpu.registers.PC = address
}

func (cpu *Cpu) Rts() {
	low := cpu.pull()
	high := cpu.pull()

	cpu.registers.PC = (uint16(high) << 8) | uint16(low)
}

func (cpu *Cpu) branch(address uint16, condition func() bool, cycles *uint16) {
	if condition() {
		*cycles++

		if !SamePage(cpu.registers.PC, address) {
			*cycles++
		}

		cpu.registers.PC = address
	}
}

func (cpu *Cpu) Bcc(address uint16, cycles *uint16) {
	cpu.branch(address, func() bool { return cpu.registers.P&C == 0 }, cycles)
}

func (cpu *Cpu) Bcs(address uint16, cycles *uint16) {
	cpu.branch(address, func() bool { return cpu.registers.P&C != 0 }, cycles)
}

func (cpu *Cpu) Beq(address uint16, cycles *uint16) {
	cpu.branch(address, func() bool { return cpu.registers.P&Z != 0 }, cycles)
}

func (cpu *Cpu) Bmi(address uint16, cycles *uint16) {
	cpu.branch(address, func() bool { return cpu.registers.P&N != 0 }, cycles)
}

func (cpu *Cpu) Bne(address uint16, cycles *uint16) {
	cpu.branch(address, func() bool { return cpu.registers.P&Z == 0 }, cycles)
}

func (cpu *Cpu) Bpl(address uint16, cycles *uint16) {
	cpu.branch(address, func() bool { return cpu.registers.P&N == 0 }, cycles)
}

func (cpu *Cpu) Bvc(address uint16, cycles *uint16) {
	cpu.branch(address, func() bool { return cpu.registers.P&V == 0 }, cycles)
}

func (cpu *Cpu) Bvs(address uint16, cycles *uint16) {
	cpu.branch(address, func() bool { return cpu.registers.P&V != 0 }, cycles)
}
