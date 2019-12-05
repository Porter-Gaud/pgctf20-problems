#!/usr/bin/env python3
import socket
import sys
from threading import Thread

HOST = '127.0.0.1'
PORT = 44444

flag = open("flag.txt", "r").read().rstrip('\n')  # Strip the end if there's a newline
validnum = open("ans.txt", "r").read().rstrip('\n')

def greeting():
    ret =  "  **************************************\n"
    ret += "  *    Welcome to VerySecure atm.      *\n"
    ret += "  **************************************\n"
    ret += "\n\n"
    ret += "Please enter a valid card number: "
    return ret

def is_valid_card(cardnumber):
    if cardnumber == validnum:
        return True
    else:
        return False

def client_handle(conn, ip, port, MAX_BUFFER = 4096):

    conn.send(greeting().encode("utf8"))
    client_input = conn.recv(MAX_BUFFER).decode("utf8").rstrip()

    size = len(client_input)

    if size >= MAX_BUFFER:
        print("Length of input from {} is too long.".format(ip))
    
    result = is_valid_card(str(client_input))

    if result:
        conn.sendall(flag.encode("utf8"))
    else:
        conn.sendall("Sorry card number is invalid :(".encode("utf8"))
    
    conn.close()
    print("Connection from {} closed".format(ip))

def server():
    s = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    s.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
    print("Socket created")

    try:
        s.bind((HOST, PORT))
        print("Socket binding successful")
    except socket.error as msg:
        print("Bind Error")
        sys.exit()

    s.listen(10)
    print("Socket Listening...")

    while True:
        conn, addr = s.accept()
        ip, port = str(addr[0]), str(addr[1])
        print("Connection from {}".format(ip))
        try:
            Thread(target=client_handle, args=(conn, ip, port)).start()
        except:
            print("Threading failed")

    s.close()
 
if __name__ == '__main__':
    server()
