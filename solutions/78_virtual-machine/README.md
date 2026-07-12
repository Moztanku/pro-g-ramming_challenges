# Virtual Machine in Go

#### This is a simple virtual machine implemented in Go. It supports basic arithmetic operations, memory management, control flow, and IO operations.

### Description:
- **Accumulator (ACC)**: Used for arithmetic operations and storing intermediate results, 16-bit signed integer.
- **Program Counter (PC)**: Points to the current instruction in memory, 16-bit unsigned inteI've checked it for you, it works, now, I challenge you to change the output to say for example `97ger.
- **Memory**: A fixed-size array of 16-bit signed integers (256 words by default), which can be used to store data and instructions.

### Instruction set:
| Instruction | Effect |
| --: | --- |
| | **Storage** |
| `SET <value>` | `ACC = value` |
| `LOAD <address>` | `ACC = memory[address]` |
| `STORE <address>` | `memory[address] = ACC` |
| `LOADI <address>` | `ACC = memory[memory[address]]` |
| `STOREI <address>` | `memory[memory[address]] = ACC` |
|  | **Arithmetic** |
| `ADD <address>` | `ACC += memory[address]` |
| `SUB <address>` | `ACC -= memory[address]` |
| `ADDI <value>` | `ACC += value` |
| `SUBI <value>` | `ACC -= value` |
| | **Control Flow** |
| `CMP <address>` | `ACC = (ACC > memory[address]) ? 1 : (ACC < memory[address] ? -1 : 0)` |
| `JMP <address>` | `PC = address` |
| `JZ <address>` | `PC = (ACC == 0) ? address : PC + 1` |
| `JNZ <address>` | `PC = (ACC != 0) ? address : PC + 1` |
| `JP <address>` | `PC = (ACC > 0) ? address : PC + 1` |
| `JN <address>` | `PC = (ACC < 0) ? address : PC + 1` |
| | **Input/Output** |
| `READ` | Read a value from the input and store it in the accumulator. |
| `WRITE` | Write the value from the accumulator to the output. |
| `WRITC` | Write ASCII character from the accumulator to the output. |
| | **Other** |
| `NOP` | No operation, does nothing. |
| `HALT` | Stop execution of the program. |

### Example program:
```
# Comments are ignored by the virtual machine.
# Only one instruction per line is allowed.
# Empty lines, or lines with only comments are treated as NOPs.
# This program reads two numbers, adds them, and writes the result.
READ        # Read first number into ACC
STORE 0     # Store it in memory[0]
READ        # Read second number into ACC
STORE 1     # Store it in memory[1]
LOAD 0      # Load first number into ACC
ADD 1       # Add second number to ACC
WRITE       # Write the result to output
HALT        # Stop execution
```
