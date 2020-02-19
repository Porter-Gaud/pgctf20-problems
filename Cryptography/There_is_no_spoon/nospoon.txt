[< Go to Cryptography](/Cryptography) OR [<< Go to Home](/)
# There is no Spoon - ?? Points
## Problem
These weird numbers were emailed to me...

`[19  278   2  209   194   -53] 
 |94  172   -69  69  158   -17|
 [-19 255   36  263  186   -79]` 

But then I found this scribbled on a piece of paper crumpled in the trash next to the librarians computer...

`[ 4  6 -1  2   5  1] 
 |-3  3 10  19 -2  0|
 | 5  3  2  5   3  2|
 | 2  5 -5 -2   3 -1|
 |-2  5 -10 0   9 -5|
 [ 1  -2  2  0  -1 1]` 

Can you figure out what the email said? (Hint: you need to insert the brackets into the flag when you submit your answer, just replace the appropriate spaces)

## Solution
Find the inverse of the encoding matrix and multilply it by the message matrix.


## Flag
`pgctf{agentsmith}`
