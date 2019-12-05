#!/usr/bin/env python3
import random
from solve import luhn  # Used to remove multiple true numbers

# Generates the valid card number (ans.txt), and card dump (carddump.txt)
def generate(maxlines):
    hasvalid = False
    with open("carddump.txt", "w+") as f:
        for i in range(maxlines):
            number = random.randint(1000000000000000, 9999999999999999)
            if luhn(str(number)) and i > (maxlines / 2) and not hasvalid:
                f.write(str(number) + "\n")
                hasvalid = True
                validcard = str(number)
            elif luhn(str(number)):
                continue
            else:
                f.write(str(number) + "\n")
    f.close()
    return validcard

if __name__ == '__main__':
    validnumber = generate(100000)  # Best to use a relatively large number

    with open("ans.txt", "w+") as f:
        f.write(str(validnumber))
        print("Valid number is: {}".format(validnumber))
        print("Valid number is in ans.txt")
    f.close()
