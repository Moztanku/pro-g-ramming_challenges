package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Println("Usage: vm <program_file>")
		os.Exit(1)
	}

	filename := args[0]

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var instructions []Instruction
	for scanner.Scan() {
		line := scanner.Text()

		if idx := len(line); idx > 0 {
			findCommentIndex := func(s string) int {
				for i := 0; i < len(s); i++ {
					if s[i] == '#' {
						return i
					}
				}
				return -1
			}

			if commentIdx := findCommentIndex(line); commentIdx != -1 {
				line = line[:commentIdx]
			}
		}

		if len(line) == 0 {
			line = "NOP"
		}

		words := strings.Fields(line)

		instruction := Instruction{}

		opcode, ok := opcodeMap[words[0]]
		if !ok {
			panic(fmt.Sprintf("Unknown instruction: %s", words[0]))
		}
		instruction.Opcode = opcode

		if len(words) > 1 {
			var operand int16
			_, err := fmt.Sscanf(words[1], "%d", &operand)
			if err != nil {
				panic(fmt.Sprintf("Invalid operand: %s", words[1]))
			}
			instruction.Operand = operand
		}

		instructions = append(instructions, instruction)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	const memorySize = 256
	vm := NewVirtualMachine(memorySize)

	vm.Run(instructions)
}
