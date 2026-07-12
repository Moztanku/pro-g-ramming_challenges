package main

type Opcode byte

const (
	// Storage
	OP_SET    Opcode = 0x01
	OP_LOAD   Opcode = 0x02
	OP_STORE  Opcode = 0x03
	OP_LOADI  Opcode = 0x04
	OP_STOREI Opcode = 0x05

	// Arithmetic
	OP_ADD  Opcode = 0x10
	OP_SUB  Opcode = 0x11
	OP_ADDI Opcode = 0x12
	OP_SUBI Opcode = 0x13

	// Control Flow
	OP_CMP Opcode = 0x20
	OP_JMP Opcode = 0x21
	OP_JZ  Opcode = 0x22
	OP_JNZ Opcode = 0x23
	OP_JP  Opcode = 0x24
	OP_JN  Opcode = 0x25

	// Input/Output
	OP_READ  Opcode = 0x30
	OP_WRITE Opcode = 0x31
	OP_WRITC Opcode = 0x32
	OP_HALT  Opcode = 0x33
	OP_NOP   Opcode = 0x00
)

var opcodeMap = map[string]Opcode{
	"SET":    OP_SET,
	"LOAD":   OP_LOAD,
	"STORE":  OP_STORE,
	"LOADI":  OP_LOADI,
	"STOREI": OP_STOREI,

	"ADD":  OP_ADD,
	"SUB":  OP_SUB,
	"ADDI": OP_ADDI,
	"SUBI": OP_SUBI,

	"CMP": OP_CMP,
	"JMP": OP_JMP,
	"JZ":  OP_JZ,
	"JNZ": OP_JNZ,
	"JP":  OP_JP,
	"JN":  OP_JN,

	"READ":  OP_READ,
	"WRITE": OP_WRITE,
	"WRITC": OP_WRITC,
	"HALT":  OP_HALT,
	"NOP":   OP_NOP,
}
