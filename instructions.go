package go6502

type OpCode uint8

type Instruction struct {
	opcode OpCode
	exec   func(*CPU) (cycles uint16)
}

type InstructionTable map[OpCode]Instruction

func NewInstructionTable() InstructionTable {
	instructions := make(map[OpCode]Instruction)
	InstructionTable(instructions).InitInstructions()
	return instructions
}

func (instructions InstructionTable) AddInstruction(inst Instruction) {
	instructions[inst.opcode] = inst
}

func (instructions InstructionTable) RemoveInstruction(opcode OpCode) {
	delete(instructions, opcode)
}

func (instructions InstructionTable) InitInstructions() {
	// LDA

	//     Immediate
	instructions.AddInstruction(Instruction{
		opcode: 0xa9,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Lda(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0xa5,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Lda(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0xb5,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Lda(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0xad,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Lda(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0xbd,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Lda(cpu.absoluteIndexedAddress(cpu.registers.X, &cycles))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		opcode: 0xb9,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Lda(cpu.absoluteIndexedAddress(cpu.registers.Y, &cycles))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		opcode: 0xa1,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Lda(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		opcode: 0xb1,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 5
			cpu.Lda(cpu.indirectIndexedAddress(&cycles))
			return
		}})

	// LDX

	//     Immediate
	instructions.AddInstruction(Instruction{
		opcode: 0xa2,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Ldx(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0xa6,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Ldx(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,Y
	instructions.AddInstruction(Instruction{
		opcode: 0xb6,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Ldx(cpu.zeroPageIndexedAddress(cpu.registers.Y))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0xae,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Ldx(cpu.absoluteAddress())
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		opcode: 0xbe,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Ldx(cpu.absoluteIndexedAddress(cpu.registers.Y, &cycles))
			return
		}})

	// LDY

	//     Immediate
	instructions.AddInstruction(Instruction{
		opcode: 0xa0,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Ldy(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0xa4,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Ldy(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0xb4,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Ldy(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0xac,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Ldy(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0xbc,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Ldy(cpu.absoluteIndexedAddress(cpu.registers.X, &cycles))
			return
		}})

	// STA

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0x85,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Sta(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0x95,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Sta(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0x8d,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Sta(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0x9d,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 5
			cpu.Sta(cpu.absoluteIndexedAddress(cpu.registers.X, nil))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		opcode: 0x99,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 5
			cpu.Sta(cpu.absoluteIndexedAddress(cpu.registers.Y, nil))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		opcode: 0x81,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Sta(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		opcode: 0x91,
		exec: func(cpu *CPU) (cycles uint16) {
			cpu.Sta(cpu.indirectIndexedAddress(&cycles))
			cycles = 6
			return
		}})

	// STX

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0x86,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Stx(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,Y
	instructions.AddInstruction(Instruction{
		opcode: 0x96,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Stx(cpu.zeroPageIndexedAddress(cpu.registers.Y))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0x8e,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Stx(cpu.absoluteAddress())
			return
		}})

	// STY

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0x84,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Sty(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0x94,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Sty(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0x8c,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Sty(cpu.absoluteAddress())
			return
		}})

	// TAX

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0xaa,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Tax()
			return
		}})

	// TAY

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0xa8,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Tay()
			return
		}})

	// TXA

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0x8a,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Txa()
			return
		}})

	// TYA

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0x98,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Tya()
			return
		}})

	// TSX

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0xba,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Tsx()
			return
		}})

	// TXS

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0x9a,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Txs()
			return
		}})

	// PHA

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0x48,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Pha()
			return
		}})

	// PHP

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0x08,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Php()
			return
		}})

	// PLA

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0x68,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Pla()
			return
		}})

	// PLP

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0x28,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Plp()
			return
		}})

	// AND

	//     Immediate
	instructions.AddInstruction(Instruction{
		opcode: 0x29,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.And(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0x25,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.And(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0x35,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.And(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0x2d,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.And(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0x3d,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.And(cpu.absoluteIndexedAddress(cpu.registers.X, &cycles))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		opcode: 0x39,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.And(cpu.absoluteIndexedAddress(cpu.registers.Y, &cycles))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		opcode: 0x21,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.And(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		opcode: 0x31,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 5
			cpu.And(cpu.indirectIndexedAddress(&cycles))
			return
		}})

	// EOR

	//     Immediate
	instructions.AddInstruction(Instruction{
		opcode: 0x49,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Eor(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0x45,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Eor(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0x55,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Eor(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0x4d,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Eor(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0x5d,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Eor(cpu.absoluteIndexedAddress(cpu.registers.X, &cycles))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		opcode: 0x59,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Eor(cpu.absoluteIndexedAddress(cpu.registers.Y, &cycles))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		opcode: 0x41,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Eor(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		opcode: 0x51,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 5
			cpu.Eor(cpu.indirectIndexedAddress(&cycles))
			return
		}})

	// ORA

	//     Immediate
	instructions.AddInstruction(Instruction{
		opcode: 0x09,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Ora(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0x05,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Ora(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0x15,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Ora(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0x0d,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Ora(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0x1d,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Ora(cpu.absoluteIndexedAddress(cpu.registers.X, &cycles))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		opcode: 0x19,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Ora(cpu.absoluteIndexedAddress(cpu.registers.Y, &cycles))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		opcode: 0x01,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Ora(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		opcode: 0x11,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 5
			cpu.Ora(cpu.indirectIndexedAddress(&cycles))
			return
		}})

	// BIT

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0x24,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Bit(cpu.zeroPageAddress())
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0x2c,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Bit(cpu.absoluteAddress())
			return
		}})

	// ADC

	//     Immediate
	instructions.AddInstruction(Instruction{
		opcode: 0x69,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Adc(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0x65,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Adc(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0x75,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Adc(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0x6d,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Adc(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0x7d,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Adc(cpu.absoluteIndexedAddress(cpu.registers.X, &cycles))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		opcode: 0x79,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Adc(cpu.absoluteIndexedAddress(cpu.registers.Y, &cycles))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		opcode: 0x61,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Adc(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		opcode: 0x71,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 5
			cpu.Adc(cpu.indirectIndexedAddress(&cycles))
			return
		}})

	// SBC

	//     Immediate
	instructions.AddInstruction(Instruction{
		opcode: 0xe9,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Sbc(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0xe5,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Sbc(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0xf5,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Sbc(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0xed,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Sbc(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0xfd,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Sbc(cpu.absoluteIndexedAddress(cpu.registers.X, &cycles))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		opcode: 0xf9,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Sbc(cpu.absoluteIndexedAddress(cpu.registers.Y, &cycles))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		opcode: 0xe1,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Sbc(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		opcode: 0xf1,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 5
			cpu.Sbc(cpu.indirectIndexedAddress(&cycles))
			return
		}})

	// CMP

	//     Immediate
	instructions.AddInstruction(Instruction{
		opcode: 0xc9,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Cmp(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0xc5,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Cmp(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0xd5,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Cmp(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0xcd,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Cmp(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0xdd,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Cmp(cpu.absoluteIndexedAddress(cpu.registers.X, &cycles))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		opcode: 0xd9,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Cmp(cpu.absoluteIndexedAddress(cpu.registers.Y, &cycles))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		opcode: 0xc1,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Cmp(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		opcode: 0xd1,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 5
			cpu.Cmp(cpu.indirectIndexedAddress(&cycles))
			return
		}})

	// CPX

	//     Immediate
	instructions.AddInstruction(Instruction{
		opcode: 0xe0,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Cpx(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0xe4,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Cpx(cpu.zeroPageAddress())
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0xec,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Cpx(cpu.absoluteAddress())
			return
		}})

	// CPY

	//     Immediate
	instructions.AddInstruction(Instruction{
		opcode: 0xc0,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Cpy(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0xc4,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Cpy(cpu.zeroPageAddress())
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0xcc,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 4
			cpu.Cpy(cpu.absoluteAddress())
			return
		}})

	// INC

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0xe6,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 5
			cpu.Inc(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0xf6,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Inc(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0xee,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Inc(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0xfe,
		exec: func(cpu *CPU) (cycles uint16) {
			cpu.Inc(cpu.absoluteIndexedAddress(cpu.registers.X, &cycles))
			cycles = 7
			return
		}})

	// INX

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0xe8,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Inx()
			return
		}})

	// INY

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0xc8,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Iny()
			return
		}})

	// DEC

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0xc6,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 5
			cpu.Dec(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0xd6,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Dec(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0xce,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Dec(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0xde,
		exec: func(cpu *CPU) (cycles uint16) {
			cpu.Dec(cpu.absoluteIndexedAddress(cpu.registers.X, &cycles))
			cycles = 7
			return
		}})

	// DEX

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0xca,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Dex()
			return
		}})

	// DEY

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0x88,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Dey()
			return
		}})

	// ASL

	//     Accumulator
	instructions.AddInstruction(Instruction{
		opcode: 0x0a,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.AslA()
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0x06,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 5
			cpu.Asl(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0x16,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Asl(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0x0e,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Asl(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0x1e,
		exec: func(cpu *CPU) (cycles uint16) {
			cpu.Asl(cpu.absoluteIndexedAddress(cpu.registers.X, &cycles))
			cycles = 7
			return
		}})

	// LSR

	//     Accumulator
	instructions.AddInstruction(Instruction{
		opcode: 0x4a,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.LsrA()
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0x46,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 5
			cpu.Lsr(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0x56,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Lsr(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0x4e,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Lsr(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0x5e,
		exec: func(cpu *CPU) (cycles uint16) {
			cpu.Lsr(cpu.absoluteIndexedAddress(cpu.registers.X, &cycles))
			cycles = 7
			return
		}})

	// ROL

	//     Accumulator
	instructions.AddInstruction(Instruction{
		opcode: 0x2a,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.RolA()
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0x26,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 5
			cpu.Rol(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0x36,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Rol(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0x2e,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Rol(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0x3e,
		exec: func(cpu *CPU) (cycles uint16) {
			cpu.Rol(cpu.absoluteIndexedAddress(cpu.registers.X, &cycles))
			cycles = 7
			return
		}})

	// ROR

	//     Accumulator
	instructions.AddInstruction(Instruction{
		opcode: 0x6a,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.RorA()
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0x66,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 5
			cpu.Ror(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0x76,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Ror(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0x6e,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Ror(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0x7e,
		exec: func(cpu *CPU) (cycles uint16) {
			cpu.Ror(cpu.absoluteIndexedAddress(cpu.registers.X, &cycles))
			cycles = 7
			return
		}})

	// JMP

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0x4c,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 3
			cpu.Jmp(cpu.absoluteAddress())
			return
		}})

	//     Indirect
	instructions.AddInstruction(Instruction{
		opcode: 0x6c,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 5
			cpu.Jmp(cpu.indirectAddress())
			return
		}})

	// JSR

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0x20,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Jsr(cpu.absoluteAddress())
			return
		}})

	// RTS

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0x60,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Rts()
			return
		}})

	// BCC

	//     Relative
	instructions.AddInstruction(Instruction{
		opcode: 0x90,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Bcc(cpu.relativeAddress(), &cycles)
			return
		}})

	// BCS

	//     Relative
	instructions.AddInstruction(Instruction{
		opcode: 0xb0,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Bcs(cpu.relativeAddress(), &cycles)
			return
		}})

	// BEQ

	//     Relative
	instructions.AddInstruction(Instruction{
		opcode: 0xf0,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Beq(cpu.relativeAddress(), &cycles)
			return
		}})

	// BMI

	//     Relative
	instructions.AddInstruction(Instruction{
		opcode: 0x30,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Bmi(cpu.relativeAddress(), &cycles)
			return
		}})

	// BNE

	//     Relative
	instructions.AddInstruction(Instruction{
		opcode: 0xd0,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Bne(cpu.relativeAddress(), &cycles)
			return
		}})

	// BPL

	//     Relative
	instructions.AddInstruction(Instruction{
		opcode: 0x10,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Bpl(cpu.relativeAddress(), &cycles)
			return
		}})

	// BVC

	//     Relative
	instructions.AddInstruction(Instruction{
		opcode: 0x50,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Bvc(cpu.relativeAddress(), &cycles)
			return
		}})

	// BVS

	//     Relative
	instructions.AddInstruction(Instruction{
		opcode: 0x70,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Bvs(cpu.relativeAddress(), &cycles)
			return
		}})

	// CLC

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0x18,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Clc()
			return
		}})

	// CLD

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0xd8,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Cld()
			return
		}})

	// CLI

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0x58,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Cli()
			return
		}})

	// CLV

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0xb8,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Clv()
			return
		}})

	// SEC

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0x38,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Sec()
			return
		}})

	// SED

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0xf8,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Sed()
			return
		}})

	// SEI

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0x78,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			cpu.Sei()
			return
		}})

	// BRK

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0x00,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 7
			cpu.Brk()
			return
		}})

	// NOP

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0xea,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 2
			return
		}})

	// RTI

	//     Implied
	instructions.AddInstruction(Instruction{
		opcode: 0x40,
		exec: func(cpu *CPU) (cycles uint16) {
			cycles = 6
			cpu.Rti()
			return
		}})

}
