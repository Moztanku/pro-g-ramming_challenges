READ      # Input n
STORE 0    # memory[0] = n
SET 0
STORE 3    # memory[3] = 0 (constant zero)
SET 2
STORE 4    # memory[4] = 2 (constant two)
LOAD 0
CMP 4
JN 35      # n < 2 -> not prime
JZ 32      # n == 2 -> prime
SET 2
STORE 1    # memory[1] = divisor d
LOAD 1
CMP 0
JZ 32      # d == n -> prime
JP 32      # d > n -> prime
LOAD 0
STORE 2    # memory[2] = remainder candidate rem
LOAD 2
CMP 1
JN 25      # if rem < d, stop subtracting
LOAD 2
SUB 1
STORE 2
JMP 18     # keep subtracting while rem >= d
LOAD 2
CMP 3
JZ 35      # rem == 0 -> divisible -> not prime
LOAD 1
ADDI 1
STORE 1
JMP 12     # try next divisor
SET 1
WRITE      # prime => 1
HALT
SET 0
WRITE      # not prime => 0
HALT
