import socket
import pickle

class Message:
    def __init__(self, size):
        self.size = size
    def increase(self):
        self.size += 1


if __name__ == "__main__":
    server = socket.socket(socket.AF_INET,socket.SOCK_STREAM)

    server.bind(('127.0.0.1',1101))
    server.listen()

    while True:
        c,addr = server.accept()

        print("Got connection from ",addr)
        data = c.recv(1024)
        print(data)
        print(len(data))
        print(pickle.loads(data).size)

        # c.send("Message from arch".encode())
