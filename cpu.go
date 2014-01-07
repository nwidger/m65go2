// Package go6502 simulates the 6502 CPU
package go6502

import (
	"fmt"
	"os"
)

// Flags used by P (Status) register
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

// The 6502's registers, all registers are 8-bit values except for PC
// which is 16-bits.
type Registers struct {
	A  uint8  // accumulator
	X  uint8  // index register X
	Y  uint8  // index register Y
	P  Status // processor status
	SP uint8  // stack pointer
	PC uint16 // program counter
}

// Creates a new set of Registers.  All registers are initialized to
// 0.
func NewRegisters() Registers {
	return Registers{}
}

// Resets all registers.  Register P is initialized with only the I
// bit set, SP is initialized to 0xfd, PC is initialized to 0xfffc
// (the RESET vector) and all other registers are initialized to 0.
func (reg *Registers) Reset() {
	reg.A = 0
	reg.X = 0
	reg.Y = 0
	reg.P = I
	reg.SP = 0xfd
	reg.PC = 0xfffc
}

// Prints the values of each register to os.Stderr.
func (reg *Registers) String() {
	fmt.Fprintf(os.Stderr, "A:  %#02x (%03dd) (%08bb)\n", reg.A, reg.A, reg.A)
	fmt.Fprintf(os.Stderr, "X:  %#02x (%03dd) (%08bb)\n", reg.X, reg.X, reg.X)
	fmt.Fprintf(os.Stderr, "Y:  %#02x (%03dd) (%08bb)\n", reg.Y, reg.Y, reg.Y)
	fmt.Fprintf(os.Stderr, "SP: %#02x (%03dd) (%08bb)\n", reg.SP, reg.SP, reg.SP)

	f := ""

	getFlag := func(flag Status, set string) string {
		if reg.P&flag != 0 {
			return set
		}

		return "-"
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

// Represents the 6502 CPU.
type CPU struct {
	decode       bool
	divisor      uint16
	clock        *Clock
	Registers    Registers
	memory       Memory
	instructions InstructionTable
}

// Returns a pointer to a new CPU with the given Memory, clock divisor
// and clock.
func NewCPU(mem Memory, divisor uint16, clock *Clock) *CPU {
	return &CPU{decode: false, divisor: divisor, clock: clock, Registers: NewRegisters(), memory: mem, instructions: NewInstructionTable()}
}

// Resets the CPU by resetting both the registers and memory.
func (cpu *CPU) Reset() {
	cpu.Registers.Reset()
	cpu.memory.Reset()
}

// Error type used to indicate that the CPU attempted to execute an
// invalid opcode
type BadOpCodeError OpCode

func (b BadOpCodeError) Error() string {
	return fmt.Sprintf("No such opcode %#02x", b)
}

// Executes the instruction pointed to by the PC register in the
// number of clock cycles as returned by the instruction's Exec
// function.  Returns the number of cycles executed and any error
// (such as BadOpCodeError).
func (cpu *CPU) Execute() (cycles uint16, error error) {
	ticks := cpu.clock.ticks

	// fetch
	opcode := OpCode(cpu.memory.Fetch(cpu.Registers.PC))
	inst, ok := cpu.instructions[opcode]

	if !ok {
		return 0, BadOpCodeError(opcode)
	}

	// execute
	cpu.Registers.PC++
	cycles = inst.Exec(cpu)

	// count cycles
	cpu.clock.Await(ticks + uint64(cycles*cpu.divisor))

	return cycles, nil
}

// Executes instruction until Execute() returns an error.
func (cpu *CPU) Run() (error error) {
	for {
		if _, error := cpu.Execute(); error != nil {
			return error
		}
	}

	return nil
}

func (cpu *CPU) setZFlag(value uint8) uint8 {
	if value == 0 {
		cpu.Registers.P |= Z
	} else {
		cpu.Registers.P &= ^Z
	}

	return value
}

func (cpu *CPU) setNFlag(value uint8) uint8 {
	cpu.Registers.P &= ^N
	cpu.Registers.P |= Status(value & (uint8(1) << 7))
	return value
}

func (cpu *CPU) setZNFlags(value uint8) uint8 {
	cpu.setZFlag(value)
	cpu.setNFlag(value)
	return value
}

func (cpu *CPU) setCFlagAddition(value uint16) uint16 {
	cpu.Registers.P &= ^C
	cpu.Registers.P |= Status(value >> 8 & 0x1)
	return value
}

func (cpu *CPU) setVFlagAddition(term1 uint16, term2 uint16, result uint16) uint16 {
	cpu.Registers.P &= ^V
	cpu.Registers.P |= Status((^(term1 ^ term2) & (term1 ^ result) & 0x80) >> 1)
	return result
}

func (cpu *CPU) load(address uint16, register *uint8) {
	*register = cpu.setZNFlags(cpu.memory.Fetch(address))
}

func (cpu *CPU) immediateAddress() (result uint16) {
	result = cpu.Registers.PC
	cpu.Registers.PC++
	return
}

func (cpu *CPU) zeroPageAddress() (result uint16) {
	result = uint16(cpu.memory.Fetch(cpu.Registers.PC))
	cpu.Registers.PC++
	return
}

func (cpu *CPU) zeroPageIndexedAddress(index uint8) (result uint16) {
	result = uint16(cpu.memory.Fetch(cpu.Registers.PC) + index)
	cpu.Registers.PC++
	return
}

func (cpu *CPU) relativeAddress() (result uint16) {
	value := uint16(cpu.memory.Fetch(cpu.Registers.PC))
	cpu.Registers.PC++

	if value > 0x7f {
		result = cpu.Registers.PC - (0x0100 - value)
	} else {
		result = cpu.Registers.PC + value
	}

	return
}

func (cpu *CPU) absoluteAddress() (result uint16) {
	low := cpu.memory.Fetch(cpu.Registers.PC)
	high := cpu.memory.Fetch(cpu.Registers.PC + 1)
	cpu.Registers.PC += 2

	result = (uint16(high) << 8) | uint16(low)
	return
}

func (cpu *CPU) indirectAddress() (result uint16) {
	low := cpu.memory.Fetch(cpu.Registers.PC)
	high := cpu.memory.Fetch(cpu.Registers.PC + 1)
	cpu.Registers.PC += 2

	// XXX: The 6502 had a bug in which it incremented only the
	// high byte instead of the whole 16-bit address when
	// computing the address.
	//
	// See http://www.obelisk.demon.co.uk/6502/reference.html#JMP
	// and http://www.6502.org/tutorials/6502opcodes.html#JMP for
	// details
	aHigh := (uint16(high) << 8) | uint16(low+1)
	aLow := (uint16(high) << 8) | uint16(low)

	low = cpu.memory.Fetch(aLow)
	high = cpu.memory.Fetch(aHigh)

	result = (uint16(high) << 8) | uint16(low)
	return
}

func (cpu *CPU) absoluteIndexedAddress(index uint8, cycles *uint16) (result uint16) {
	low := cpu.memory.Fetch(cpu.Registers.PC)
	high := cpu.memory.Fetch(cpu.Registers.PC + 1)
	cpu.Registers.PC += 2

	address := (uint16(high) << 8) | uint16(low)
	result = address + uint16(index)

	if cycles != nil && !SamePage(address, result) {
		*cycles++
	}

	return
}

func (cpu *CPU) indexedIndirectAddress() (result uint16) {
	address := uint16(cpu.memory.Fetch(cpu.Registers.PC) + cpu.Registers.X)
	cpu.Registers.PC++

	low := cpu.memory.Fetch(address)
	high := cpu.memory.Fetch(address + 1)

	result = (uint16(high) << 8) | uint16(low)
	return
}

func (cpu *CPU) indirectIndexedAddress(cycles *uint16) (result uint16) {
	address := uint16(cpu.memory.Fetch(cpu.Registers.PC))
	cpu.Registers.PC++

	low := cpu.memory.Fetch(address)
	high := cpu.memory.Fetch(address + 1)

	address = (uint16(high) << 8) | uint16(low)

	result = address + uint16(cpu.Registers.Y)

	if cycles != nil && !SamePage(address, result) {
		*cycles++
	}

	return
}

// Loads a byte of memory into the accumulator setting the zero and
// negative flags as appropriate.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if A = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of A is set
func (cpu *CPU) Lda(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: LDA $%04x\n", cpu.Registers.PC, address)
	}

	cpu.load(address, &cpu.Registers.A)
}

// Loads a byte of memory into the X register setting the zero and
// negative flags as appropriate.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if X = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of X is set
func (cpu *CPU) Ldx(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: LDX $%04x\n", cpu.Registers.PC, address)
	}

	cpu.load(address, &cpu.Registers.X)
}

// Loads a byte of memory into the Y register setting the zero and
// negative flags as appropriate.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if Y = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of Y is set
func (cpu *CPU) Ldy(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: LDY $%04x\n", cpu.Registers.PC, address)
	}

	cpu.load(address, &cpu.Registers.Y)
}

func (cpu *CPU) store(address uint16, value uint8) {
	cpu.memory.Store(address, value)
}

// Stores the contents of the accumulator into memory.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Sta(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: STA $%04x\n", cpu.Registers.PC, address)
	}

	cpu.store(address, cpu.Registers.A)
}

// Stores the contents of the X register into memory.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Stx(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: STX $%04x\n", cpu.Registers.PC, address)
	}

	cpu.store(address, cpu.Registers.X)
}

// Stores the contents of the Y register into memory.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Sty(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: STY $%04x\n", cpu.Registers.PC, address)
	}

	cpu.store(address, cpu.Registers.Y)
}

func (cpu *CPU) transfer(from uint8, to *uint8) {
	*to = cpu.setZNFlags(from)
}

// Copies the current contents of the accumulator into the X register
// and sets the zero and negative flags as appropriate.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if X = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of X is set
func (cpu *CPU) Tax() {
	if cpu.decode {
		fmt.Printf("  %04x: TAX\n", cpu.Registers.PC)
	}

	cpu.transfer(cpu.Registers.A, &cpu.Registers.X)
}

// Copies the current contents of the accumulator into the Y register
// and sets the zero and negative flags as appropriate.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if Y = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of Y is set
func (cpu *CPU) Tay() {
	if cpu.decode {
		fmt.Printf("  %04x: TAY\n", cpu.Registers.PC)
	}

	cpu.transfer(cpu.Registers.A, &cpu.Registers.Y)
}

// Copies the current contents of the X register into the accumulator
// and sets the zero and negative flags as appropriate.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if A = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of A is set
func (cpu *CPU) Txa() {
	if cpu.decode {
		fmt.Printf("  %04x: TXA\n", cpu.Registers.PC)
	}

	cpu.transfer(cpu.Registers.X, &cpu.Registers.A)
}

// Copies the current contents of the Y register into the accumulator
// and sets the zero and negative flags as appropriate.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if A = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of A is set
func (cpu *CPU) Tya() {
	if cpu.decode {
		fmt.Printf("  %04x: TYA\n", cpu.Registers.PC)
	}

	cpu.transfer(cpu.Registers.Y, &cpu.Registers.A)
}

// Copies the current contents of the stack register into the X
// register and sets the zero and negative flags as appropriate.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if X = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of X is set
func (cpu *CPU) Tsx() {
	if cpu.decode {
		fmt.Printf("  %04x: TSX\n", cpu.Registers.PC)
	}

	cpu.transfer(cpu.Registers.SP, &cpu.Registers.X)
}

// Copies the current contents of the X register into the stack
// register.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Txs() {
	if cpu.decode {
		fmt.Printf("  %04x: TXS\n", cpu.Registers.PC)
	}

	cpu.transfer(cpu.Registers.X, &cpu.Registers.SP)
}

func (cpu *CPU) push(value uint8) {
	cpu.memory.Store(0x0100|uint16(cpu.Registers.SP), value)
	cpu.Registers.SP--
}

func (cpu *CPU) push16(value uint16) {
	cpu.push(uint8(value >> 8))
	cpu.push(uint8(value))
}

func (cpu *CPU) pull() (value uint8) {
	cpu.Registers.SP++
	value = cpu.memory.Fetch(0x0100 | uint16(cpu.Registers.SP))
	return
}

func (cpu *CPU) pull16() (value uint16) {
	low := cpu.pull()
	high := cpu.pull()

	value = (uint16(high) << 8) | uint16(low)
	return
}

// Pushes a copy of the accumulator on to the stack.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Pha() {
	if cpu.decode {
		fmt.Printf("  %04x: PHA\n", cpu.Registers.PC)
	}

	cpu.push(cpu.Registers.A)
}

// Pushes a copy of the status flags on to the stack.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Php() {
	if cpu.decode {
		fmt.Printf("  %04x: PHP\n", cpu.Registers.PC)
	}

	cpu.push(uint8(cpu.Registers.P | B))
}

// Pulls an 8 bit value from the stack and into the accumulator. The
// zero and negative flags are set as appropriate.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if A = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of A is set
func (cpu *CPU) Pla() {
	if cpu.decode {
		fmt.Printf("  %04x: PLA\n", cpu.Registers.PC)
	}

	cpu.Registers.A = cpu.setZNFlags(cpu.pull())
}

// Pulls an 8 bit value from the stack and into the processor
// flags. The flags will take on new states as determined by the value
// pulled.
//
// C 	Carry Flag 	  Set from stack
// Z 	Zero Flag 	  Set from stack
// I 	Interrupt Disable Set from stack
// D 	Decimal Mode Flag Set from stack
// B 	Break Command 	  Set from stack
// V 	Overflow Flag 	  Set from stack
// N 	Negative Flag 	  Set from stack
func (cpu *CPU) Plp() {
	if cpu.decode {
		fmt.Printf("  %04x: PLP\n", cpu.Registers.PC)
	}

	cpu.Registers.P = Status(cpu.pull())
}

// A logical AND is performed, bit by bit, on the accumulator contents
// using the contents of a byte of memory.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if A = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 set
func (cpu *CPU) And(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: AND $%04x\n", cpu.Registers.PC, address)
	}

	cpu.Registers.A = cpu.setZNFlags(cpu.Registers.A & cpu.memory.Fetch(address))
}

// An exclusive OR is performed, bit by bit, on the accumulator
// contents using the contents of a byte of memory.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if A = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 set
func (cpu *CPU) Eor(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: EOR $%04x\n", cpu.Registers.PC, address)
	}

	cpu.Registers.A = cpu.setZNFlags(cpu.Registers.A ^ cpu.memory.Fetch(address))
}

// An inclusive OR is performed, bit by bit, on the accumulator
// contents using the contents of a byte of memory.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if A = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 set
func (cpu *CPU) Ora(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: ORA $%04x\n", cpu.Registers.PC, address)
	}

	cpu.Registers.A = cpu.setZNFlags(cpu.Registers.A | cpu.memory.Fetch(address))
}

// This instructions is used to test if one or more bits are set in a
// target memory location. The mask pattern in A is ANDed with the
// value in memory to set or clear the zero flag, but the result is
// not kept. Bits 7 and 6 of the value from memory are copied into the
// N and V flags.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if the result if the AND is zero
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Set to bit 6 of the memory value
// N 	Negative Flag 	  Set to bit 7 of the memory value
func (cpu *CPU) Bit(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BIT $%04x\n", cpu.Registers.PC, address)
	}

	value := cpu.memory.Fetch(address)
	cpu.setZFlag(value & cpu.Registers.A)
	cpu.Registers.P = Status(uint8(cpu.Registers.P) | (value & 0xc0))
}

func (cpu *CPU) addition(value uint16) {
	orig := uint16(cpu.Registers.A)
	result := cpu.setCFlagAddition(orig + value + uint16(cpu.Registers.P&C))
	cpu.Registers.A = cpu.setZNFlags(uint8(cpu.setVFlagAddition(orig, value, result)))
}

// This instruction adds the contents of a memory location to the
// accumulator together with the carry bit. If overflow occurs the
// carry bit is set, this enables multiple byte addition to be
// performed.
//
// C 	Carry Flag 	  Set if overflow in bit 7
// Z 	Zero Flag 	  Set if A = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Set if sign bit is incorrect
// N 	Negative Flag 	  Set if bit 7 set
func (cpu *CPU) Adc(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: ADC $%04x\n", cpu.Registers.PC, address)
	}

	value := uint16(cpu.memory.Fetch(address))
	cpu.addition(value)
}

// This instruction subtracts the contents of a memory location to the
// accumulator together with the not of the carry bit. If overflow
// occurs the carry bit is clear, this enables multiple byte
// subtraction to be performed.
//
// C 	Carry Flag 	  Clear if overflow in bit 7
// Z 	Zero Flag 	  Set if A = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Set if sign bit is incorrect
// N 	Negative Flag 	  Set if bit 7 set
func (cpu *CPU) Sbc(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: SBC $%04x\n", cpu.Registers.PC, address)
	}

	value := uint16(cpu.memory.Fetch(address)) ^ 0xff
	cpu.addition(value)
}

func (cpu *CPU) compare(address uint16, register uint8) {
	value := uint16(cpu.memory.Fetch(address)) ^ 0xff + 1
	cpu.setZNFlags(uint8(cpu.setCFlagAddition(uint16(register) + value)))
}

// This instruction compares the contents of the accumulator with
// another memory held value and sets the zero and carry flags as
// appropriate.
//
// C 	Carry Flag 	  Set if A >= M
// Z 	Zero Flag 	  Set if A = M
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of the result is set
func (cpu *CPU) Cmp(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: CMP $%04x\n", cpu.Registers.PC, address)
	}

	cpu.compare(address, cpu.Registers.A)
}

// This instruction compares the contents of the X register with
// another memory held value and sets the zero and carry flags as
// appropriate.
//
// C 	Carry Flag 	  Set if X >= M
// Z 	Zero Flag 	  Set if X = M
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of the result is set
func (cpu *CPU) Cpx(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: CPX $%04x\n", cpu.Registers.PC, address)
	}

	cpu.compare(address, cpu.Registers.X)
}

// This instruction compares the contents of the Y register with
// another memory held value and sets the zero and carry flags as
// appropriate.
//
// C 	Carry Flag 	  Set if Y >= M
// Z 	Zero Flag 	  Set if Y = M
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of the result is set
func (cpu *CPU) Cpy(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: CPY $%04x\n", cpu.Registers.PC, address)
	}

	cpu.compare(address, cpu.Registers.Y)
}

// Adds one to the value held at a specified memory location setting
// the zero and negative flags as appropriate.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if result is zero
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of the result is set
func (cpu *CPU) Inc(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: INC $%04x\n", cpu.Registers.PC, address)
	}

	cpu.memory.Store(address, cpu.setZNFlags(cpu.memory.Fetch(address)+1))
}

func (cpu *CPU) increment(register *uint8) {
	*register = cpu.setZNFlags(*register + 1)
}

// Adds one to the X register setting the zero and negative flags as
// appropriate.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if X is zero
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of X is set
func (cpu *CPU) Inx() {
	if cpu.decode {
		fmt.Printf("  %04x: INX\n", cpu.Registers.PC)
	}

	cpu.increment(&cpu.Registers.X)
}

// Adds one to the Y register setting the zero and negative flags as
// appropriate.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if Y is zero
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of Y is set
func (cpu *CPU) Iny() {
	if cpu.decode {
		fmt.Printf("  %04x: INY\n", cpu.Registers.PC)
	}

	cpu.increment(&cpu.Registers.Y)
}

// Subtracts one from the value held at a specified memory location
// setting the zero and negative flags as appropriate.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if result is zero
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of the result is set
func (cpu *CPU) Dec(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: DEC $%04x\n", cpu.Registers.PC, address)
	}

	cpu.memory.Store(address, cpu.setZNFlags(cpu.memory.Fetch(address)-1))
}

func (cpu *CPU) decrement(register *uint8) {
	*register = cpu.setZNFlags(*register - 1)
}

// Subtracts one from the X register setting the zero and negative
// flags as appropriate.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if X is zero
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of X is set
func (cpu *CPU) Dex() {
	if cpu.decode {
		fmt.Printf("  %04x: DEX\n", cpu.Registers.PC)
	}

	cpu.decrement(&cpu.Registers.X)
}

// Subtracts one from the Y register setting the zero and negative
// flags as appropriate.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Set if Y is zero
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of Y is set
func (cpu *CPU) Dey() {
	if cpu.decode {
		fmt.Printf("  %04x: DEY\n", cpu.Registers.PC)
	}

	cpu.decrement(&cpu.Registers.Y)
}

type direction int

const (
	left direction = iota
	right
)

func (cpu *CPU) shift(direction direction, value uint8, store func(uint8)) {
	c := Status(0)

	switch direction {
	case left:
		c = Status((value & uint8(N)) >> 7)
		value <<= 1
	case right:
		c = Status(value & uint8(C))
		value >>= 1
	}

	cpu.Registers.P &= ^C
	cpu.Registers.P |= c

	store(cpu.setZNFlags(value))
}

// This operation shifts all the bits of the accumulator one bit
// left. Bit 0 is set to 0 and bit 7 is placed in the carry flag. The
// effect of this operation is to multiply the memory contents by 2
// (ignoring 2's complement considerations), setting the carry if the
// result will not fit in 8 bits.
//
// C 	Carry Flag 	  Set to contents of old bit 7
// Z 	Zero Flag 	  Set if A = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of the result is set
func (cpu *CPU) AslA() {
	if cpu.decode {
		fmt.Printf("  %04x: ASL A\n", cpu.Registers.PC)
	}

	cpu.shift(left, cpu.Registers.A, func(value uint8) { cpu.Registers.A = value })
}

// This operation shifts all the bits of the memory contents one bit
// left. Bit 0 is set to 0 and bit 7 is placed in the carry flag. The
// effect of this operation is to multiply the memory contents by 2
// (ignoring 2's complement considerations), setting the carry if the
// result will not fit in 8 bits.
//
// C 	Carry Flag 	  Set to contents of old bit 7
// Z 	Zero Flag 	  Set if A = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of the result is set
func (cpu *CPU) Asl(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: ASL $%04x\n", cpu.Registers.PC, address)
	}

	cpu.shift(left, cpu.memory.Fetch(address), func(value uint8) { cpu.memory.Store(address, value) })
}

// Each of the bits in A is shift one place to the right. The bit that
// was in bit 0 is shifted into the carry flag. Bit 7 is set to zero.
//
// C 	Carry Flag 	  Set to contents of old bit 0
// Z 	Zero Flag 	  Set if result = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of the result is set
func (cpu *CPU) LsrA() {
	if cpu.decode {
		fmt.Printf("  %04x: LSR A\n", cpu.Registers.PC)
	}

	cpu.shift(right, cpu.Registers.A, func(value uint8) { cpu.Registers.A = value })
}

// Each of the bits in M is shift one place to the right. The bit that
// was in bit 0 is shifted into the carry flag. Bit 7 is set to zero.
//
// C 	Carry Flag 	  Set to contents of old bit 0
// Z 	Zero Flag 	  Set if result = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of the result is set
func (cpu *CPU) Lsr(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: LSR $%04x\n", cpu.Registers.PC, address)
	}

	cpu.shift(right, cpu.memory.Fetch(address), func(value uint8) { cpu.memory.Store(address, value) })
}

func (cpu *CPU) rotate(direction direction, value uint8, store func(uint8)) {
	c := Status(0)

	switch direction {
	case left:
		c = Status(value & uint8(N) >> 7)
		value = ((value << 1) & uint8(^C)) | uint8(cpu.Registers.P&C)
	case right:
		c = Status(value & uint8(C))
		value = ((value >> 1) & uint8(^N)) | uint8((cpu.Registers.P&C)<<7)
	}

	cpu.Registers.P &= ^C
	cpu.Registers.P |= c

	store(cpu.setZNFlags(value))
}

// Move each of the bits in A one place to the left. Bit 0 is filled
// with the current value of the carry flag whilst the old bit 7
// becomes the new carry flag value.
//
// C 	Carry Flag 	  Set to contents of old bit 7
// Z 	Zero Flag 	  Set if A = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of the result is set
func (cpu *CPU) RolA() {
	if cpu.decode {
		fmt.Printf("  %04x: ROL A\n", cpu.Registers.PC)
	}

	cpu.rotate(left, cpu.Registers.A, func(value uint8) { cpu.Registers.A = value })
}

// Move each of the bits in A one place to the left. Bit 0 is filled
// with the current value of the carry flag whilst the old bit 7
// becomes the new carry flag value.
//
// C 	Carry Flag 	  Set to contents of old bit 7
// Z 	Zero Flag 	  Set if A = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of the result is set
func (cpu *CPU) Rol(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: ROL $%04x\n", cpu.Registers.PC, address)
	}

	cpu.rotate(left, cpu.memory.Fetch(address), func(value uint8) { cpu.memory.Store(address, value) })
}

// Move each of the bits in A one place to the right. Bit 7 is filled
// with the current value of the carry flag whilst the old bit 0
// becomes the new carry flag value.
//
// C 	Carry Flag 	  Set to contents of old bit 0
// Z 	Zero Flag 	  Set if A = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of the result is set
func (cpu *CPU) RorA() {
	if cpu.decode {
		fmt.Printf("  %04x: ROR A\n", cpu.Registers.PC)
	}

	cpu.rotate(right, cpu.Registers.A, func(value uint8) { cpu.Registers.A = value })
}

// Move each of the bits in M one place to the right. Bit 7 is filled
// with the current value of the carry flag whilst the old bit 0
// becomes the new carry flag value.
//
// C 	Carry Flag 	  Set to contents of old bit 0
// Z 	Zero Flag 	  Set if A = 0
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Set if bit 7 of the result is set
func (cpu *CPU) Ror(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: ROR $%04x\n", cpu.Registers.PC, address)
	}

	cpu.rotate(right, cpu.memory.Fetch(address), func(value uint8) { cpu.memory.Store(address, value) })
}

// Sets the program counter to the address specified by the operand.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Jmp(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: JMP $%04x\n", cpu.Registers.PC, address)
	}

	cpu.Registers.PC = address
}

// The JSR instruction pushes the address (minus one) of the return
// point on to the stack and then sets the program counter to the
// target memory address.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Jsr(address uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: JSR $%04x\n", cpu.Registers.PC, address)
	}

	value := cpu.Registers.PC - 1

	cpu.push16(value)

	cpu.Registers.PC = address
}

// The RTS instruction is used at the end of a subroutine to return to
// the calling routine. It pulls the program counter (minus one) from
// the stack.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Rts() {
	if cpu.decode {
		fmt.Printf("  %04x: RTS\n", cpu.Registers.PC)
	}

	cpu.Registers.PC = cpu.pull16() + 1
}

func (cpu *CPU) branch(address uint16, condition func() bool, cycles *uint16) {
	if condition() {
		*cycles++

		if !SamePage(cpu.Registers.PC, address) {
			*cycles++
		}

		cpu.Registers.PC = address
	}
}

// If the carry flag is clear then add the relative displacement to
// the program counter to cause a branch to a new location.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Bcc(address uint16, cycles *uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BCC $%04x\n", cpu.Registers.PC, address)
	}

	cpu.branch(address, func() bool { return cpu.Registers.P&C == 0 }, cycles)
}

// If the carry flag is set then add the relative displacement to the
// program counter to cause a branch to a new location.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Bcs(address uint16, cycles *uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BCS $%04x\n", cpu.Registers.PC, address)
	}

	cpu.branch(address, func() bool { return cpu.Registers.P&C != 0 }, cycles)
}

// If the zero flag is set then add the relative displacement to the
// program counter to cause a branch to a new location.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Beq(address uint16, cycles *uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BEQ $%04x\n", cpu.Registers.PC, address)
	}

	cpu.branch(address, func() bool { return cpu.Registers.P&Z != 0 }, cycles)
}

// If the negative flag is set then add the relative displacement to
// the program counter to cause a branch to a new location.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Bmi(address uint16, cycles *uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BMI $%04x\n", cpu.Registers.PC, address)
	}

	cpu.branch(address, func() bool { return cpu.Registers.P&N != 0 }, cycles)
}

// If the zero flag is clear then add the relative displacement to the
// program counter to cause a branch to a new location.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Bne(address uint16, cycles *uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BNE $%04x\n", cpu.Registers.PC, address)
	}

	cpu.branch(address, func() bool { return cpu.Registers.P&Z == 0 }, cycles)
}

// If the negative flag is clear then add the relative displacement to
// the program counter to cause a branch to a new location.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Bpl(address uint16, cycles *uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BPL $%04x\n", cpu.Registers.PC, address)
	}

	cpu.branch(address, func() bool { return cpu.Registers.P&N == 0 }, cycles)
}

// If the overflow flag is clear then add the relative displacement to
// the program counter to cause a branch to a new location.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Bvc(address uint16, cycles *uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BVC $%04x\n", cpu.Registers.PC, address)
	}

	cpu.branch(address, func() bool { return cpu.Registers.P&V == 0 }, cycles)
}

// If the overflow flag is set then add the relative displacement to
// the program counter to cause a branch to a new location.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Bvs(address uint16, cycles *uint16) {
	if cpu.decode {
		fmt.Printf("  %04x: BVS $%04x\n", cpu.Registers.PC, address)
	}

	cpu.branch(address, func() bool { return cpu.Registers.P&V != 0 }, cycles)
}

// Set the carry flag to zero.
//
// C 	Carry Flag 	  Set to 0
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Clc() {
	if cpu.decode {
		fmt.Printf("  %04x: CLC\n", cpu.Registers.PC)
	}

	cpu.Registers.P &^= C
}

// Set the decimal mode flag to zero.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Set to 0
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Cld() {
	if cpu.decode {
		fmt.Printf("  %04x: CLD\n", cpu.Registers.PC)
	}

	cpu.Registers.P &^= D
}

// Clears the interrupt disable flag allowing normal interrupt
// requests to be serviced.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Set to 0
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Cli() {
	if cpu.decode {
		fmt.Printf("  %04x: CLI\n", cpu.Registers.PC)
	}

	cpu.Registers.P &^= I
}

// Clears the interrupt disable flag allowing normal interrupt
// requests to be serviced.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Set to 0
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Clv() {
	if cpu.decode {
		fmt.Printf("  %04x: CLV\n", cpu.Registers.PC)
	}

	cpu.Registers.P &^= V
}

// Set the carry flag to one.
//
// C 	Carry Flag 	  Set to 1
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Sec() {
	if cpu.decode {
		fmt.Printf("  %04x: SEC\n", cpu.Registers.PC)
	}

	cpu.Registers.P |= C
}

// Set the decimal mode flag to one.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Set to 1
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Sed() {
	if cpu.decode {
		fmt.Printf("  %04x: SED\n", cpu.Registers.PC)
	}

	cpu.Registers.P |= D
}

// Set the interrupt disable flag to one.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Set to 1
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Not affected
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Sei() {
	if cpu.decode {
		fmt.Printf("  %04x: SEI\n", cpu.Registers.PC)
	}

	cpu.Registers.P |= I
}

// The BRK instruction forces the generation of an interrupt
// request. The program counter and processor status are pushed on the
// stack then the IRQ interrupt vector at $FFFE/F is loaded into the
// PC and the break flag in the status set to one.
//
// C 	Carry Flag 	  Not affected
// Z 	Zero Flag 	  Not affected
// I 	Interrupt Disable Not affected
// D 	Decimal Mode Flag Not affected
// B 	Break Command 	  Set to 1
// V 	Overflow Flag 	  Not affected
// N 	Negative Flag 	  Not affected
func (cpu *CPU) Brk() {
	if cpu.decode {
		fmt.Printf("  %04x: BRK\n", cpu.Registers.PC)
	}

	cpu.Registers.PC++

	cpu.push16(cpu.Registers.PC)
	cpu.push(uint8(cpu.Registers.P | B))

	cpu.Registers.P |= I

	low := cpu.memory.Fetch(0xfffe)
	high := cpu.memory.Fetch(0xffff)

	cpu.Registers.PC = (uint16(high) << 8) | uint16(low)
}

// The RTI instruction is used at the end of an interrupt processing
// routine. It pulls the processor flags from the stack followed by
// the program counter.
//
// C 	Carry Flag 	  Set from stack
// Z 	Zero Flag 	  Set from stack
// I 	Interrupt Disable Set from stack
// D 	Decimal Mode Flag Set from stack
// B 	Break Command 	  Set from stack
// V 	Overflow Flag 	  Set from stack
// N 	Negative Flag 	  Set from stack
func (cpu *CPU) Rti() {
	if cpu.decode {
		fmt.Printf("  %04x: RTI\n", cpu.Registers.PC)
	}

	cpu.Registers.P = Status(cpu.pull())
	cpu.Registers.PC = cpu.pull16()
}
