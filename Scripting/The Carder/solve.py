#!/usr/bin/env python3
import sys

def luhn(number):
    summ = 0
    alt = 0
    i = len(number) - 1
    while i >= 0:
        num = int(number[i])
        if alt:
            num = num * 2
            if num > 9:
                num = (num % 10) + 1

        summ = summ + num
        alt = not alt
        i -= 1
    return summ % 10 == 0

if __name__ == "__main__":
    with open("carddump.txt", "r") as f:
        for line in f.read().split():
            n = str(line)
            if luhn(n):
                print(n)
            else:
                continue
    f.close()
