import socket


if __name__ == "__main__":
    server = socket.socket(socket.AF_INET,socket.SOCK_STREAM)
    host = socket.gethostname()
    print(host)

    server.bind(('130.126.143.111',12000))

    server.listen()

    while True:
        c,addr = server.accept()

        print("GOt connection from ",addr)

        c.send("Message from arch".encode())
