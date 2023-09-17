

import time

def fib(n):
    if(n == 0 or n == 1):
        return 1
    return fib(n - 1) + fib(n - 2)

def fib_mem(n,seq):
    if n == 0 or n == 1:
        seq[n] = 1;
        return 1;
    if seq[n] == 0:
        seq[n] = fib_mem(n - 1, seq) + fib_mem(n - 2, seq);
        return seq[n];
    return seq[n];

def compute_time(num):
    prev = time.time()
    print(fib(num))
    naive_time = time.time() - prev;
    print("Naive time consumed: ", naive_time)

    seq = []
    for i in range(100):
        seq.append(0);
    prev = time.time()
    print(fib_mem(num,seq))
    mem_time = time.time() - prev;
    print("Memorization time consumed: ", mem_time)
    print("Improvement Ratio is ", naive_time / mem_time)



# The larger the input num is, the higher the ratio
compute_time(30)


