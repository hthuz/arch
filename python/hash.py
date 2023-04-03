
from os import urandom
import hashlib


def myprint():
    print("hello world")

if __name__ == "__main__":
    cnt = 0
    while cnt < 10:

        print(hashlib.sha256(urandom(20)).hexdigest())
        cnt += 1

        
    
