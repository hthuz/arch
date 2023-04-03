
import threading


def square(num):
   
    for i in range(num):
        print(threading.get_ident())
    return

def cube(num):
    for i in range(num):
        print(threading.get_ident())
    return

if __name__ == "__main__":

    t1 = threading.Thread(target = square, args = (10,))
    t2 = threading.Thread(target = cube, args = (10,))

    t1.start()
    t2.start()

    for i in range(40):
        print(threading.get_native_id())
