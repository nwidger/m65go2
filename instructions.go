package _65go2

type OpCode uint8

type Instruction struct {
	opcode OpCode
	exec   func(*Cpu) (cycles uint16)
}

type InstructionTable map[OpCode]Instruction

func NewInstructionTable() InstructionTable {
	instructions := make(map[OpCode]Instruction)
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
			address, pageCrossed := cpu.absoluteIndexedAddress(cpu.registers.Y)

			cpu.Lda(address)

			if pageCrossed {
				cycles++
			}

			return
		}})

	//     Absolute,Y
	instructions.AddInstruction(Instruction{
		opcode: 0xb9,
		exec: func(cpu *Cpu) (cycles uint16) {
			cycles = 4
			address, pageCrossed := cpu.absoluteIndexedAddress(cpu.registers.Y)

			cpu.Lda(address)

			if pageCrossed {
				cycles++
			}

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
			address, pageCrossed := cpu.indirectIndexedAddress()

			cpu.Lda(address)

			if pageCrossed {
				cycles++
			}

			return
		}})

}
