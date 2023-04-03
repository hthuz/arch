import socket



if __name__ == "__main__":
    s = socket.socket()

    port = 1200

    
    s.connect(('127.0.0.1',port))

    print (s.recv(1024).decode())

    s.close()

