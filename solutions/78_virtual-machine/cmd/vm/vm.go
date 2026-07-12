package main

import "fmt"

type VirtualMachine struct {
	ACC    int16  // Accumulator
	PC     uint16 // Program Counter
	Memory []int16
}

func NewVirtualMachine(memorySize int) *VirtualMachine {
	return &VirtualMachine{
		ACC:    0,
		PC:     0,
		Memory: make([]int16, memorySize),
	}
}

func (vm *VirtualMachine) Run(instructions []Instruction) {
	var running bool = true

	for running {
		var instruction Instruction = instructions[vm.PC]

		switch instruction.Opcode {
		case OP_SET:
			vm.ACC = instruction.Operand
		case OP_LOAD:
			vm.ACC = vm.Memory[instruction.Operand]
		case OP_STORE:
			vm.Memory[instruction.Operand] = vm.ACC
		case OP_LOADI:
			vm.ACC = vm.Memory[vm.Memory[instruction.Operand]]
		case OP_STOREI:
			vm.Memory[vm.Memory[instruction.Operand]] = vm.ACC
		case OP_ADD:
			vm.ACC += vm.Memory[instruction.Operand]
		case OP_SUB:
			vm.ACC -= vm.Memory[instruction.Operand]
		case OP_ADDI:
			vm.ACC += instruction.Operand
		case OP_SUBI:
			vm.ACC -= instruction.Operand
		case OP_CMP:
			if vm.ACC > vm.Memory[instruction.Operand] {
				vm.ACC = 1
			} else if vm.ACC < vm.Memory[instruction.Operand] {
				vm.ACC = -1
			} else {
				vm.ACC = 0
			}
		case OP_JMP:
			vm.PC = uint16(instruction.Operand)
			continue
		case OP_JZ:
			if vm.ACC == 0 {
				vm.PC = uint16(instruction.Operand)
				continue
			}
		case OP_JNZ:
			if vm.ACC != 0 {
				vm.PC = uint16(instruction.Operand)
				continue
			}
		case OP_JP:
			if vm.ACC > 0 {
				vm.PC = uint16(instruction.Operand)
				continue
			}
		case OP_JN:
			if vm.ACC < 0 {
				vm.PC = uint16(instruction.Operand)
				continue
			}
		case OP_READ:
			var input int16
			fmt.Scanf("%d", &input)
			vm.ACC = input
		case OP_WRITE:
			fmt.Println(vm.ACC)
		case OP_WRITC:
			fmt.Printf("%c", vm.ACC)
		case OP_HALT:
			running = false
		case OP_NOP:
			// Do nothing
		default:
			panic(fmt.Sprintf("Unknown opcode: %d", instruction.Opcode))
		}

		vm.PC++
	}
}
