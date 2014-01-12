package m65go2

// Represents opcodes for the 6502 CPU
type OpCode uint8

// Represents an instruction for the 6502 CPU.  The Exec field
// implements the instruction and returns the total clock cycles to be
// consumed by the instruction.
type Instruction struct {
	OpCode OpCode
	Exec   func(CPUer) (cycles uint16)
}

// Stores instructions understood by the 6502 CPU, indexed by opcode.
type InstructionTable map[OpCode]Instruction

// Returns a new, empty InstructionTable
func NewInstructionTable() InstructionTable {
	instructions := make(map[OpCode]Instruction)
	return instructions
}

// Adds an instruction to the InstructionTable
func (instructions InstructionTable) AddInstruction(inst Instruction) {
	instructions[inst.OpCode] = inst
}

// Removes any instruction with the given opcode
func (instructions InstructionTable) RemoveInstruction(opcode OpCode) {
	delete(instructions, opcode)
}

// Adds the 6502 CPU's instruction set to the InstructionTable.
func (instructions InstructionTable) InitInstructions() {
	// LDA

	//     Immediate
	instructions.AddInstruction(Instruction{
		OpCode: 0xa9,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Lda(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0xa5,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Lda(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		OpCode: 0xb5,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Lda(cpu.zeroPageIndexedAddress(X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0xad,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Lda(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		OpCode: 0xbd,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Lda(cpu.absoluteIndexedAddress(X, &cycles))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		OpCode: 0xb9,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Lda(cpu.absoluteIndexedAddress(Y, &cycles))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		OpCode: 0xa1,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Lda(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		OpCode: 0xb1,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 5
			cpu.Lda(cpu.indirectIndexedAddress(&cycles))
			return
		}})

	// LDX

	//     Immediate
	instructions.AddInstruction(Instruction{
		OpCode: 0xa2,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Ldx(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0xa6,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Ldx(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,Y
	instructions.AddInstruction(Instruction{
		OpCode: 0xb6,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Ldx(cpu.zeroPageIndexedAddress(Y))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0xae,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Ldx(cpu.absoluteAddress())
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		OpCode: 0xbe,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Ldx(cpu.absoluteIndexedAddress(Y, &cycles))
			return
		}})

	// LDY

	//     Immediate
	instructions.AddInstruction(Instruction{
		OpCode: 0xa0,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Ldy(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0xa4,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Ldy(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		OpCode: 0xb4,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Ldy(cpu.zeroPageIndexedAddress(X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0xac,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Ldy(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		OpCode: 0xbc,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Ldy(cpu.absoluteIndexedAddress(X, &cycles))
			return
		}})

	// STA

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0x85,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Sta(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x95,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Sta(cpu.zeroPageIndexedAddress(X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0x8d,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Sta(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x9d,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 5
			cpu.Sta(cpu.absoluteIndexedAddress(X, nil))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		OpCode: 0x99,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 5
			cpu.Sta(cpu.absoluteIndexedAddress(Y, nil))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		OpCode: 0x81,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Sta(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		OpCode: 0x91,
		Exec: func(cpu CPUer) (cycles uint16) {
			cpu.Sta(cpu.indirectIndexedAddress(&cycles))
			cycles = 6
			return
		}})

	// STX

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0x86,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Stx(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,Y
	instructions.AddInstruction(Instruction{
		OpCode: 0x96,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Stx(cpu.zeroPageIndexedAddress(Y))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0x8e,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Stx(cpu.absoluteAddress())
			return
		}})

	// STY

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0x84,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Sty(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x94,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Sty(cpu.zeroPageIndexedAddress(X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0x8c,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Sty(cpu.absoluteAddress())
			return
		}})

	// TAX

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0xaa,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Tax()
			return
		}})

	// TAY

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0xa8,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Tay()
			return
		}})

	// TXA

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0x8a,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Txa()
			return
		}})

	// TYA

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0x98,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Tya()
			return
		}})

	// TSX

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0xba,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Tsx()
			return
		}})

	// TXS

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0x9a,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Txs()
			return
		}})

	// PHA

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0x48,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Pha()
			return
		}})

	// PHP

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0x08,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Php()
			return
		}})

	// PLA

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0x68,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Pla()
			return
		}})

	// PLP

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0x28,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Plp()
			return
		}})

	// AND

	//     Immediate
	instructions.AddInstruction(Instruction{
		OpCode: 0x29,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.And(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0x25,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.And(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x35,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.And(cpu.zeroPageIndexedAddress(X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0x2d,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.And(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x3d,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.And(cpu.absoluteIndexedAddress(X, &cycles))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		OpCode: 0x39,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.And(cpu.absoluteIndexedAddress(Y, &cycles))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		OpCode: 0x21,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.And(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		OpCode: 0x31,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 5
			cpu.And(cpu.indirectIndexedAddress(&cycles))
			return
		}})

	// EOR

	//     Immediate
	instructions.AddInstruction(Instruction{
		OpCode: 0x49,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Eor(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0x45,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Eor(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x55,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Eor(cpu.zeroPageIndexedAddress(X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0x4d,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Eor(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x5d,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Eor(cpu.absoluteIndexedAddress(X, &cycles))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		OpCode: 0x59,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Eor(cpu.absoluteIndexedAddress(Y, &cycles))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		OpCode: 0x41,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Eor(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		OpCode: 0x51,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 5
			cpu.Eor(cpu.indirectIndexedAddress(&cycles))
			return
		}})

	// ORA

	//     Immediate
	instructions.AddInstruction(Instruction{
		OpCode: 0x09,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Ora(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0x05,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Ora(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x15,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Ora(cpu.zeroPageIndexedAddress(X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0x0d,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Ora(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x1d,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Ora(cpu.absoluteIndexedAddress(X, &cycles))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		OpCode: 0x19,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Ora(cpu.absoluteIndexedAddress(Y, &cycles))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		OpCode: 0x01,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Ora(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		OpCode: 0x11,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 5
			cpu.Ora(cpu.indirectIndexedAddress(&cycles))
			return
		}})

	// BIT

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0x24,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Bit(cpu.zeroPageAddress())
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0x2c,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Bit(cpu.absoluteAddress())
			return
		}})

	// ADC

	//     Immediate
	instructions.AddInstruction(Instruction{
		OpCode: 0x69,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Adc(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0x65,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Adc(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x75,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Adc(cpu.zeroPageIndexedAddress(X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0x6d,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Adc(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x7d,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Adc(cpu.absoluteIndexedAddress(X, &cycles))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		OpCode: 0x79,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Adc(cpu.absoluteIndexedAddress(Y, &cycles))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		OpCode: 0x61,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Adc(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		OpCode: 0x71,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 5
			cpu.Adc(cpu.indirectIndexedAddress(&cycles))
			return
		}})

	// SBC

	//     Immediate
	instructions.AddInstruction(Instruction{
		OpCode: 0xe9,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Sbc(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0xe5,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Sbc(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		OpCode: 0xf5,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Sbc(cpu.zeroPageIndexedAddress(X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0xed,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Sbc(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		OpCode: 0xfd,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Sbc(cpu.absoluteIndexedAddress(X, &cycles))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		OpCode: 0xf9,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Sbc(cpu.absoluteIndexedAddress(Y, &cycles))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		OpCode: 0xe1,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Sbc(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		OpCode: 0xf1,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 5
			cpu.Sbc(cpu.indirectIndexedAddress(&cycles))
			return
		}})

	// CMP

	//     Immediate
	instructions.AddInstruction(Instruction{
		OpCode: 0xc9,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Cmp(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0xc5,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Cmp(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		OpCode: 0xd5,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Cmp(cpu.zeroPageIndexedAddress(X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0xcd,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Cmp(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		OpCode: 0xdd,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Cmp(cpu.absoluteIndexedAddress(X, &cycles))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		OpCode: 0xd9,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Cmp(cpu.absoluteIndexedAddress(Y, &cycles))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		OpCode: 0xc1,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Cmp(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		OpCode: 0xd1,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 5
			cpu.Cmp(cpu.indirectIndexedAddress(&cycles))
			return
		}})

	// CPX

	//     Immediate
	instructions.AddInstruction(Instruction{
		OpCode: 0xe0,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Cpx(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0xe4,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Cpx(cpu.zeroPageAddress())
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0xec,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Cpx(cpu.absoluteAddress())
			return
		}})

	// CPY

	//     Immediate
	instructions.AddInstruction(Instruction{
		OpCode: 0xc0,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Cpy(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0xc4,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Cpy(cpu.zeroPageAddress())
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0xcc,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 4
			cpu.Cpy(cpu.absoluteAddress())
			return
		}})

	// INC

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0xe6,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 5
			cpu.Inc(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		OpCode: 0xf6,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Inc(cpu.zeroPageIndexedAddress(X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0xee,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Inc(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		OpCode: 0xfe,
		Exec: func(cpu CPUer) (cycles uint16) {
			cpu.Inc(cpu.absoluteIndexedAddress(X, &cycles))
			cycles = 7
			return
		}})

	// INX

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0xe8,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Inx()
			return
		}})

	// INY

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0xc8,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Iny()
			return
		}})

	// DEC

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0xc6,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 5
			cpu.Dec(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		OpCode: 0xd6,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Dec(cpu.zeroPageIndexedAddress(X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0xce,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Dec(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		OpCode: 0xde,
		Exec: func(cpu CPUer) (cycles uint16) {
			cpu.Dec(cpu.absoluteIndexedAddress(X, &cycles))
			cycles = 7
			return
		}})

	// DEX

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0xca,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Dex()
			return
		}})

	// DEY

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0x88,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Dey()
			return
		}})

	// ASL

	//     Accumulator
	instructions.AddInstruction(Instruction{
		OpCode: 0x0a,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.AslA()
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0x06,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 5
			cpu.Asl(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x16,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Asl(cpu.zeroPageIndexedAddress(X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0x0e,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Asl(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x1e,
		Exec: func(cpu CPUer) (cycles uint16) {
			cpu.Asl(cpu.absoluteIndexedAddress(X, &cycles))
			cycles = 7
			return
		}})

	// LSR

	//     Accumulator
	instructions.AddInstruction(Instruction{
		OpCode: 0x4a,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.LsrA()
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0x46,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 5
			cpu.Lsr(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x56,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Lsr(cpu.zeroPageIndexedAddress(X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0x4e,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Lsr(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x5e,
		Exec: func(cpu CPUer) (cycles uint16) {
			cpu.Lsr(cpu.absoluteIndexedAddress(X, &cycles))
			cycles = 7
			return
		}})

	// ROL

	//     Accumulator
	instructions.AddInstruction(Instruction{
		OpCode: 0x2a,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.RolA()
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0x26,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 5
			cpu.Rol(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x36,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Rol(cpu.zeroPageIndexedAddress(X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0x2e,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Rol(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x3e,
		Exec: func(cpu CPUer) (cycles uint16) {
			cpu.Rol(cpu.absoluteIndexedAddress(X, &cycles))
			cycles = 7
			return
		}})

	// ROR

	//     Accumulator
	instructions.AddInstruction(Instruction{
		OpCode: 0x6a,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.RorA()
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		OpCode: 0x66,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 5
			cpu.Ror(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x76,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Ror(cpu.zeroPageIndexedAddress(X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0x6e,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Ror(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		OpCode: 0x7e,
		Exec: func(cpu CPUer) (cycles uint16) {
			cpu.Ror(cpu.absoluteIndexedAddress(X, &cycles))
			cycles = 7
			return
		}})

	// JMP

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0x4c,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 3
			cpu.Jmp(cpu.absoluteAddress())
			return
		}})

	//     Indirect
	instructions.AddInstruction(Instruction{
		OpCode: 0x6c,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 5
			cpu.Jmp(cpu.indirectAddress())
			return
		}})

	// JSR

	//     Absolute
	instructions.AddInstruction(Instruction{
		OpCode: 0x20,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Jsr(cpu.absoluteAddress())
			return
		}})

	// RTS

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0x60,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Rts()
			return
		}})

	// BCC

	//     Relative
	instructions.AddInstruction(Instruction{
		OpCode: 0x90,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Bcc(cpu.relativeAddress(), &cycles)
			return
		}})

	// BCS

	//     Relative
	instructions.AddInstruction(Instruction{
		OpCode: 0xb0,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Bcs(cpu.relativeAddress(), &cycles)
			return
		}})

	// BEQ

	//     Relative
	instructions.AddInstruction(Instruction{
		OpCode: 0xf0,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Beq(cpu.relativeAddress(), &cycles)
			return
		}})

	// BMI

	//     Relative
	instructions.AddInstruction(Instruction{
		OpCode: 0x30,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Bmi(cpu.relativeAddress(), &cycles)
			return
		}})

	// BNE

	//     Relative
	instructions.AddInstruction(Instruction{
		OpCode: 0xd0,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Bne(cpu.relativeAddress(), &cycles)
			return
		}})

	// BPL

	//     Relative
	instructions.AddInstruction(Instruction{
		OpCode: 0x10,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Bpl(cpu.relativeAddress(), &cycles)
			return
		}})

	// BVC

	//     Relative
	instructions.AddInstruction(Instruction{
		OpCode: 0x50,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Bvc(cpu.relativeAddress(), &cycles)
			return
		}})

	// BVS

	//     Relative
	instructions.AddInstruction(Instruction{
		OpCode: 0x70,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Bvs(cpu.relativeAddress(), &cycles)
			return
		}})

	// CLC

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0x18,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Clc()
			return
		}})

	// CLD

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0xd8,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Cld()
			return
		}})

	// CLI

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0x58,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Cli()
			return
		}})

	// CLV

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0xb8,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Clv()
			return
		}})

	// SEC

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0x38,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Sec()
			return
		}})

	// SED

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0xf8,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Sed()
			return
		}})

	// SEI

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0x78,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			cpu.Sei()
			return
		}})

	// BRK

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0x00,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 7
			cpu.Brk()
			return
		}})

	// NOP

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0xea,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 2
			return
		}})

	// RTI

	//     Implied
	instructions.AddInstruction(Instruction{
		OpCode: 0x40,
		Exec: func(cpu CPUer) (cycles uint16) {
			cycles = 6
			cpu.Rti()
			return
		}})

}
