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
	reg.P = I
	reg.SP = 0xfd
	reg.PC = 0xfffc
}

func (reg *Registers) print() {
	fmt.Fprintf(os.Stderr, "A:  %#02x (%03dd) (%08bb)\n", reg.A, reg.A, reg.A)
	fmt.Fprintf(os.Stderr, "X:  %#02x (%03dd) (%08bb)\n", reg.X, reg.X, reg.X)
	fmt.Fprintf(os.Stderr, "Y:  %#02x (%03dd) (%08bb)\n", reg.Y, reg.Y, reg.Y)
	fmt.Fprintf(os.Stderr, "SP: %#02x (%03dd) (%08bb)\n", reg.SP, reg.SP, reg.SP)

	f := ""

	getFlag := func(flag Status, set string) string {
		if reg.P&flag != 0 {
			return set
		} else {
			return "-"
		}
	}

	f += getFlag(N, "N")
	f += getFlag(V, "V")
	f += "-" // -UNUSED-
	f += getFlag(B, "B")
	f += getFlag(D, "D")
	f += getFlag(I, "I")
	f += getFlag(Z, "Z")
	f += getFlag(C, "C")

	fmt.Fprintf(os.Stderr, "P:  %08bb (%s)\n", reg.P, f)
	fmt.Fprintf(os.Stderr, "PC: %#04x (%05dd) (%016bb)\n", reg.PC, reg.PC, reg.PC)
}

type Cpu struct {
	decode       bool
	clock        Clock
	registers    Registers
	memory       Memory
	instructions InstructionTable
}

func NewCpu(mem Memory, clock Clock) *Cpu {
	return &Cpu{decode: false, clock: clock, registers: NewRegisters(), memory: mem, instructions: NewInstructionTable()}
}

func (cpu *Cpu) Reset() {
	cpu.registers.reset()
	cpu.memory.reset()
}

type BadOpCodeError OpCode

func (b BadOpCodeError) Error() string {
	return fmt.Sprintf("No such opcode %#02x", b)
}

func (cpu *Cpu) Execute() (cycles uint16, error error) {
	ticks := cpu.clock.ticks

	// fetch
	opcode := OpCode(cpu.memory.fetch(cpu.registers.PC))
	inst, ok := cpu.instructions[opcode]

	if !ok {
		return 0, BadOpCodeError(opcode)
	}

	// execute
	cpu.registers.PC++
	cycles = inst.exec(cpu)

	// count cycles
	cpu.clock.await(ticks + uint64(cycles))

	return cycles, nil
}

func (cpu *Cpu) Run() (error error) {
	for {
		if _, error := cpu.Execute(); error != nil {
			fmt.Println(error)
			break
		}
	}

	return nil
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
	cpu.registers.P &= ^N
	cpu.registers.P |= Status(value & (uint8(1) << 7))
	return value
}

func (cpu *Cpu) setZNFlags(value uint8) uint8 {
	cpu.setZFlag(value)
	cpu.setNFlag(value)
	return value
}

func (cpu *Cpu) setCFlagAddition(value uint16) uint16 {
	cpu.registers.P &= ^C
	cpu.registers.P |= Status(value >> 8 & 0x1)
	return value
}

func (cpu *Cpu) setVFlagAddition(term1 uint16, term2 uint16, result uint16) uint16 {
	cpu.registers.P &= ^V
	cpu.registers.P |= Status((^(term1 ^ term2) & (term1 ^ result) & 0x80) >> 1)
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
	if cpu.decode {
		fmt.Printf("  %04x: LDA $%04x\n", cpu.registers.PC, address)
	}

	cpu.load(address, &cpu.registers.A)
}

func (cpu *Cpu) Ldx(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: LDX $%04x\n", cpu.registers.PC, address)
	}

	cpu.load(address, &cpu.registers.X)
}

func (cpu *Cpu) Ldy(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: LDY $%04x\n", cpu.registers.PC, address)
	}

	cpu.load(address, &cpu.registers.Y)
}

func (cpu *Cpu) store(address uint16, value uint8) {
	cpu.memory.store(address, value)
}

func (cpu *Cpu) Sta(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: STA $%04x\n", cpu.registers.PC, address)
	}

	cpu.store(address, cpu.registers.A)
}

func (cpu *Cpu) Stx(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: STX $%04x\n", cpu.registers.PC, address)
	}

	cpu.store(address, cpu.registers.X)
}

func (cpu *Cpu) Sty(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: STY $%04x\n", cpu.registers.PC, address)
	}

	cpu.store(address, cpu.registers.Y)
}

func (cpu *Cpu) transfer(from uint8, to *uint8) {
	*to = cpu.setZNFlags(from)
}

func (cpu *Cpu) Tax() {
	if cpu.decode {
		fmt.Printf("  %04x: TAX\n", cpu.registers.PC)
	}

	cpu.transfer(cpu.registers.A, &cpu.registers.X)
}

func (cpu *Cpu) Tay() {
	if cpu.decode {
		fmt.Printf("  %04x: TAY\n", cpu.registers.PC)
	}

	cpu.transfer(cpu.registers.A, &cpu.registers.Y)
}

func (cpu *Cpu) Txa() {
	if cpu.decode {
		fmt.Printf("  %04x: TXA\n", cpu.registers.PC)
	}

	cpu.transfer(cpu.registers.X, &cpu.registers.A)
}

func (cpu *Cpu) Tya() {
	if cpu.decode {
		fmt.Printf("  %04x: TYA\n", cpu.registers.PC)
	}

	cpu.transfer(cpu.registers.Y, &cpu.registers.A)
}

func (cpu *Cpu) Tsx() {
	if cpu.decode {
		fmt.Printf("  %04x: TSX\n", cpu.registers.PC)
	}

	cpu.transfer(cpu.registers.SP, &cpu.registers.X)
}

func (cpu *Cpu) Txs() {
	if cpu.decode {
		fmt.Printf("  %04x: TXS\n", cpu.registers.PC)
	}

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
	if cpu.decode {
		fmt.Printf("  %04x: PHA\n", cpu.registers.PC)
	}

	cpu.push(cpu.registers.A)
}

func (cpu *Cpu) Php() {
	if cpu.decode {
		fmt.Printf("  %04x: PHP\n", cpu.registers.PC)
	}

	cpu.push(uint8(cpu.registers.P | B))
}

func (cpu *Cpu) Pla() {
	if cpu.decode {
		fmt.Printf("  %04x: PLA\n", cpu.registers.PC)
	}

	cpu.registers.A = cpu.setZNFlags(cpu.pull())
}

func (cpu *Cpu) Plp() {
	if cpu.decode {
		fmt.Printf("  %04x: PLP\n", cpu.registers.PC)
	}

	cpu.registers.P = Status(cpu.pull())
}

func (cpu *Cpu) And(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: AND $%04x\n", cpu.registers.PC, address)
	}

	cpu.registers.A = cpu.setZNFlags(cpu.registers.A & cpu.memory.fetch(address))
}

func (cpu *Cpu) Eor(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: EOR $%04x\n", cpu.registers.PC, address)
	}

	cpu.registers.A = cpu.setZNFlags(cpu.registers.A ^ cpu.memory.fetch(address))
}

func (cpu *Cpu) Ora(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: ORA $%04x\n", cpu.registers.PC, address)
	}

	cpu.registers.A = cpu.setZNFlags(cpu.registers.A | cpu.memory.fetch(address))
}

func (cpu *Cpu) Bit(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BIT $%04x\n", cpu.registers.PC, address)
	}

	value := cpu.memory.fetch(address)
	cpu.setZFlag(value & cpu.registers.A)
	cpu.registers.P = Status(uint8(cpu.registers.P) | (value & 0xc0))
}

func (cpu *Cpu) addition(value uint16) {
	orig := uint16(cpu.registers.A)
	result := cpu.setCFlagAddition(orig + value + uint16(cpu.registers.P&C))
	cpu.registers.A = cpu.setZNFlags(uint8(cpu.setVFlagAddition(orig, value, result)))
}

func (cpu *Cpu) Adc(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: ADC $%04x\n", cpu.registers.PC, address)
	}

	value := uint16(cpu.memory.fetch(address))
	cpu.addition(value)
}

func (cpu *Cpu) Sbc(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: SBC $%04x\n", cpu.registers.PC, address)
	}

	value := uint16(cpu.memory.fetch(address)) ^ 0xff + 1
	cpu.addition(value)
}

func (cpu *Cpu) compare(address uint16, register uint8) {
	value := uint16(cpu.memory.fetch(address)) ^ 0xff + 1
	cpu.setZNFlags(uint8(cpu.setCFlagAddition(uint16(register) + value)))
}

func (cpu *Cpu) Cmp(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: CMP $%04x\n", cpu.registers.PC, address)
	}

	cpu.compare(address, cpu.registers.A)
}

func (cpu *Cpu) Cpx(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: CPX $%04x\n", cpu.registers.PC, address)
	}

	cpu.compare(address, cpu.registers.X)
}

func (cpu *Cpu) Cpy(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: CPY $%04x\n", cpu.registers.PC, address)
	}

	cpu.compare(address, cpu.registers.Y)
}

func (cpu *Cpu) Inc(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: INC $%04x\n", cpu.registers.PC, address)
	}

	cpu.memory.store(address, cpu.setZNFlags(cpu.memory.fetch(address)+1))
}

func (cpu *Cpu) increment(register *uint8) {
	*register = cpu.setZNFlags(*register + 1)
}

func (cpu *Cpu) Inx() {
	if cpu.decode {
		fmt.Printf("  %04x: INX\n", cpu.registers.PC)
	}

	cpu.increment(&cpu.registers.X)
}

func (cpu *Cpu) Iny() {
	if cpu.decode {
		fmt.Printf("  %04x: INY\n", cpu.registers.PC)
	}

	cpu.increment(&cpu.registers.Y)
}

func (cpu *Cpu) Dec(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: DEC $%04x\n", cpu.registers.PC, address)
	}

	cpu.memory.store(address, cpu.setZNFlags(cpu.memory.fetch(address)-1))
}

func (cpu *Cpu) decrement(register *uint8) {
	*register = cpu.setZNFlags(*register - 1)
}

func (cpu *Cpu) Dex() {
	if cpu.decode {
		fmt.Printf("  %04x: DEX\n", cpu.registers.PC)
	}

	cpu.decrement(&cpu.registers.X)
}

func (cpu *Cpu) Dey() {
	if cpu.decode {
		fmt.Printf("  %04x: DEY\n", cpu.registers.PC)
	}

	cpu.decrement(&cpu.registers.Y)
}

type Direction int

const (
	left Direction = iota
	right
)

func (cpu *Cpu) shift(direction Direction, value uint8, store func(uint8)) {
	c := Status(0)

	switch direction {
	case left:
		c = Status((value & uint8(N)) >> 7)
		value <<= 1
	case right:
		c = Status(value & uint8(C))
		value >>= 1
	}

	cpu.registers.P &= ^C
	cpu.registers.P |= c

	store(cpu.setZNFlags(value))
}

func (cpu *Cpu) AslA() {
	if cpu.decode {
		fmt.Printf("  %04x: ASL A\n", cpu.registers.PC)
	}

	cpu.shift(left, cpu.registers.A, func(value uint8) { cpu.registers.A = value })
}

func (cpu *Cpu) Asl(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: ASL $%04x\n", cpu.registers.PC, address)
	}

	cpu.shift(left, cpu.memory.fetch(address), func(value uint8) { cpu.memory.store(address, value) })
}

func (cpu *Cpu) LsrA() {
	if cpu.decode {
		fmt.Printf("  %04x: LSR A\n", cpu.registers.PC)
	}

	cpu.shift(right, cpu.registers.A, func(value uint8) { cpu.registers.A = value })
}

func (cpu *Cpu) Lsr(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: LSR $%04x\n", cpu.registers.PC, address)
	}

	cpu.shift(right, cpu.memory.fetch(address), func(value uint8) { cpu.memory.store(address, value) })
}

func (cpu *Cpu) rotate(direction Direction, value uint8, store func(uint8)) {
	c := Status(0)

	switch direction {
	case left:
		c = Status(value & uint8(N) >> 7)
		value = ((value << 1) & uint8(^C)) | uint8(cpu.registers.P&C)
	case right:
		c = Status(value & uint8(C))
		value = ((value >> 1) & uint8(^N)) | uint8((cpu.registers.P&C)<<7)
	}

	cpu.registers.P &= ^C
	cpu.registers.P |= c

	store(cpu.setZNFlags(value))
}

func (cpu *Cpu) RolA() {
	if cpu.decode {
		fmt.Printf("  %04x: ROL A\n", cpu.registers.PC)
	}

	cpu.rotate(left, cpu.registers.A, func(value uint8) { cpu.registers.A = value })
}

func (cpu *Cpu) Rol(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: ROL $%04x\n", cpu.registers.PC, address)
	}

	cpu.rotate(left, cpu.memory.fetch(address), func(value uint8) { cpu.memory.store(address, value) })
}

func (cpu *Cpu) RorA() {
	if cpu.decode {
		fmt.Printf("  %04x: ROR A\n", cpu.registers.PC)
	}

	cpu.rotate(right, cpu.registers.A, func(value uint8) { cpu.registers.A = value })
}

func (cpu *Cpu) Ror(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: ROR $%04x\n", cpu.registers.PC, address)
	}

	cpu.rotate(right, cpu.memory.fetch(address), func(value uint8) { cpu.memory.store(address, value) })
}

func (cpu *Cpu) Jmp(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: JMP $%04x\n", cpu.registers.PC, address)
	}

	cpu.registers.PC = address
}

func (cpu *Cpu) Jsr(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: JSR $%04x\n", cpu.registers.PC, address)
	}

	value := cpu.registers.PC - 1

	cpu.push(uint8(value >> 8))
	cpu.push(uint8(value))

	cpu.registers.PC = address
}

func (cpu *Cpu) Rts() {
	if cpu.decode {
		fmt.Printf("  %04x: RTS\n", cpu.registers.PC)
	}

	low := cpu.pull()
	high := cpu.pull()

	cpu.registers.PC = (uint16(high) << 8) | uint16(low) + 1
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
	if cpu.decode {
		fmt.Printf("  %04x: BCC $%04x\n", cpu.registers.PC, address)
	}

	cpu.branch(address, func() bool { return cpu.registers.P&C == 0 }, cycles)
}

func (cpu *Cpu) Bcs(address uint16, cycles *uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BCS $%04x\n", cpu.registers.PC, address)
	}

	cpu.branch(address, func() bool { return cpu.registers.P&C != 0 }, cycles)
}

func (cpu *Cpu) Beq(address uint16, cycles *uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BEQ $%04x\n", cpu.registers.PC, address)
	}

	cpu.branch(address, func() bool { return cpu.registers.P&Z != 0 }, cycles)
}

func (cpu *Cpu) Bmi(address uint16, cycles *uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BMI $%04x\n", cpu.registers.PC, address)
	}

	cpu.branch(address, func() bool { return cpu.registers.P&N != 0 }, cycles)
}

func (cpu *Cpu) Bne(address uint16, cycles *uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BNE $%04x\n", cpu.registers.PC, address)
	}

	cpu.branch(address, func() bool { return cpu.registers.P&Z == 0 }, cycles)
}

func (cpu *Cpu) Bpl(address uint16, cycles *uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BPL $%04x\n", cpu.registers.PC, address)
	}

	cpu.branch(address, func() bool { return cpu.registers.P&N == 0 }, cycles)
}

func (cpu *Cpu) Bvc(address uint16, cycles *uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BVC $%04x\n", cpu.registers.PC, address)
	}

	cpu.branch(address, func() bool { return cpu.registers.P&V == 0 }, cycles)
}

func (cpu *Cpu) Bvs(address uint16, cycles *uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BVS $%04x\n", cpu.registers.PC, address)
	}

	cpu.branch(address, func() bool { return cpu.registers.P&V != 0 }, cycles)
}

func (cpu *Cpu) Clc() {
	if cpu.decode {
		fmt.Printf("  %04x: CLC\n", cpu.registers.PC)
	}

	cpu.registers.P &^= C
}

func (cpu *Cpu) Cld() {
	if cpu.decode {
		fmt.Printf("  %04x: CLD\n", cpu.registers.PC)
	}

	cpu.registers.P &^= D
}

func (cpu *Cpu) Cli() {
	if cpu.decode {
		fmt.Printf("  %04x: CLI\n", cpu.registers.PC)
	}

	cpu.registers.P &^= I
}

func (cpu *Cpu) Clv() {
	if cpu.decode {
		fmt.Printf("  %04x: CLV\n", cpu.registers.PC)
	}

	cpu.registers.P &^= V
}

func (cpu *Cpu) Sec() {
	if cpu.decode {
		fmt.Printf("  %04x: SEC\n", cpu.registers.PC)
	}

	cpu.registers.P |= C
}

func (cpu *Cpu) Sed() {
	if cpu.decode {
		fmt.Printf("  %04x: SED\n", cpu.registers.PC)
	}

	cpu.registers.P |= D
}

func (cpu *Cpu) Sei() {
	if cpu.decode {
		fmt.Printf("  %04x: SEI\n", cpu.registers.PC)
	}

	cpu.registers.P |= I
}

func (cpu *Cpu) Brk() {
	if cpu.decode {
		fmt.Printf("  %04x: BRK\n", cpu.registers.PC)
	}

	cpu.registers.PC++

	cpu.push(uint8(cpu.registers.PC >> 8))
	cpu.push(uint8(cpu.registers.PC))
	cpu.push(uint8(cpu.registers.P | B))

	cpu.registers.P |= I

	low := cpu.memory.fetch(0xfffe)
	high := cpu.memory.fetch(0xffff)

	cpu.registers.PC = (uint16(high) << 8) | uint16(low)
}

func (cpu *Cpu) Rti() {
	if cpu.decode {
		fmt.Printf("  %04x: RTI\n", cpu.registers.PC)
	}

	cpu.registers.P = Status(cpu.pull())

	low := cpu.pull()
	high := cpu.pull()

	cpu.registers.PC = (uint16(high) << 8) | uint16(low)
}
