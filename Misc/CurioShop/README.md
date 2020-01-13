[< Go to Misc](/Misc) OR [<< Go to Home](/)

# The Curio Shop - 50 Points

## Problem

Welcome to the Curio Shop. Hope you find what you're looking for.

## Solution

The money input function is vulnerable to integer overflow.  
Computers store integers using binary. This raises a problem: how do we represent negative numbers? **Two's Complement** allows computers to represent negative numbers by using the first bit. `0001`<sub>2</sub> = `1`<sub>10</sub>, and `1001`<sub>2</sub> = `-7`<sub>10</sub>. This means that using addition and multiplication, there can be a case where a positive number + a positive number = a negative number–– we call this **overflow**. `0111 + 0011 = 1010`. Because the first bit is reserved for the +/- sign, we've just written `8 + 2 = -6`.

In the Curio Shop, users are allowed to enter how many of an item they want. We store prices as a 32-bit integer, so the largest positive value is `2^31 - 1 = 2147483647`. Since flags cost $1000000, we can overflow this 32-bit integer by saying we want at least `2147483648 / 1000000 = 2,148` flags. This will cause the price to be negative, and buying a flag will **give** the user money since we say `money -= price` after the transaction.

## Flag

`pgctf{2147483648_is_a_mag1c_numb3r}`
