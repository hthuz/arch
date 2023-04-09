import socket
import pickle


class Message:
    def __init__(self, size):
        self.size = size;
    def increase(self):
        self.size += 1


if __name__ == "__main__":
    s = socket.socket()
    port = 1101

    msg = Message(3)
    msg.increase()
    print(msg)
    
    s.connect(('127.0.0.1',port))

    s.send(pickle.dumps(msg))

    print (s.recv(1024).decode())

    s.close()

