package _65go2

type OpCode uint8

type Instruction struct {
	opcode OpCode
	exec   func(*Cpu) (cycles uint16)
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
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 2
			cpu.Lda(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0xa5,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 3
			cpu.Lda(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0xb5,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 4
			cpu.Lda(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0xad,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 4
			cpu.Lda(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0xbd,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 4
			cpu.Lda(cpu.absoluteIndexedAddress(cpu.registers.X, &cycles))
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		opcode: 0xb9,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 4
			cpu.Lda(cpu.absoluteIndexedAddress(cpu.registers.Y, &cycles))
			return
		}})

	//     (Indirect,X)
	instructions.AddInstruction(Instruction{
		opcode: 0xa1,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 6
			cpu.Lda(cpu.indexedIndirectAddress())
			return
		}})

	//     (Indirect),Y
	instructions.AddInstruction(Instruction{
		opcode: 0xb1,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 5
			cpu.Lda(cpu.indirectIndexedAddress(&cycles))
			return
		}})

	// LDX

	//     Immediate
	instructions.AddInstruction(Instruction{
		opcode: 0xa2,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 2
			cpu.Ldx(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0xa6,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 3
			cpu.Ldx(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,Y
	instructions.AddInstruction(Instruction{
		opcode: 0xb6,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 4
			cpu.Ldx(cpu.zeroPageIndexedAddress(cpu.registers.Y))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0xae,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 4
			cpu.Ldx(cpu.absoluteAddress())
			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		opcode: 0xbe,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 4
			cpu.Ldx(cpu.absoluteIndexedAddress(cpu.registers.Y, &cycles))
			return
		}})

	// LDY

	//     Immediate
	instructions.AddInstruction(Instruction{
		opcode: 0xa0,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 2
			cpu.Ldy(cpu.immediateAddress())
			return
		}})

	//     Zero Page
	instructions.AddInstruction(Instruction{
		opcode: 0xa4,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 3
			cpu.Ldy(cpu.zeroPageAddress())
			return
		}})

	//     Zero Page,X
	instructions.AddInstruction(Instruction{
		opcode: 0xb4,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 4
			cpu.Ldy(cpu.zeroPageIndexedAddress(cpu.registers.X))
			return
		}})

	//     Absolute
	instructions.AddInstruction(Instruction{
		opcode: 0xac,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 4
			cpu.Ldy(cpu.absoluteAddress())
			return
		}})

	//     Absolute,X
	instructions.AddInstruction(Instruction{
		opcode: 0xbc,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 4
			cpu.Ldy(cpu.absoluteIndexedAddress(cpu.registers.X, &cycles))
			return
		}})
}
