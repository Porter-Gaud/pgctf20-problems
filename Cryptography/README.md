[< Go to Cryptography](/Cryptography) OR [<< Go to Home](/)
# There is no Spoon - ?? Points
## Problem
These weird numbers were emailed to me...

`[103 89   41  157   86   167] 
 |3  228   68  208  188   327|
 |107 172  36  168  132   175|
 |52  -65  12  33   -29   16 |
 |-27 -280 -17 -118 -181 -151|
 [ 8   62  -3  12    35    8 ]` 

But then I found this scribbled on a piece of paper crumpled in the trash next to the librarians computer...

`[ 4  6 -1 ] 
 | -3 3 10 |
 | 5  3  2 |
 | 2  5 -5 |
 |-2  5 -10|
 [1  -2  2 ]` 

Can you figure out what the email said? (Hint: you need to insert the brackets into the flag when you submit your answer, just replace the appropriate spaces)

## Solution
Find the inverse of the encoding matrix and multilply it by the message matrix.


## Flag
`pgctf{agentsmith}`
