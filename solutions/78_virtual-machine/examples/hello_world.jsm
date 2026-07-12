# Jacek Assembly Language (JSM)
SET 72 # H
STORE 0
SET 101 # e
STORE 1
SET 108 # l
STORE 2
SET 108 # l
STORE 3
SET 111 # o
STORE 4
SET 32 # space
STORE 5
SET 87 # W
STORE 6
SET 111 # o
STORE 7
SET 114 # r
STORE 8
SET 108 # l
STORE 9
SET 100 # d
STORE 10
SET 33 # !
STORE 11

# 12 -> 0, end of string
# 13 -> additional address register

# Load our character into memory
LOADI 13
# 0 means end of string, so we halt if we reach it
JZ 42
# Print the character
WRITC
# Increment the address register
LOAD 13
ADDI 1
STORE 13
# Jump back to the beginning of the loop
JMP 30

HALT
