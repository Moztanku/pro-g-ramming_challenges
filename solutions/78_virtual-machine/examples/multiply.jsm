READ      # First number (multiplicand)
STORE 0
READ      # Second number (multiplier, expected >= 0)
STORE 1
SET 0     # result = 0
STORE 2
SET 0     # zero constant for CMP
STORE 3
LOAD 1
CMP 3
JZ 21     # If multiplier == 0, skip loop
JN 21     # If multiplier < 0, skip loop (not handled)
LOAD 2
ADD 0
STORE 2
LOAD 1
SUBI 1
STORE 1
LOAD 1
CMP 3
JP 12     # Continue while multiplier > 0
LOAD 2
WRITE
HALT
