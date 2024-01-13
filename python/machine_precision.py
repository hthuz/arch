
import numpy as np

print(np.finfo(np.float16))

# half precision: 16 bits float
# single precison: 32 bits float
# double precison: 64 bits float
# quad precision: 128 bits float
# no 256 bits float now

# 1 sign, 5 exponent, 10 mantissa
# max 0 11110 1111111111 => 2^(30 - 15) * (1 + 1023 / 1024) = 65504
# epsilon fl(1 + eps) > 1. i.e. (1 + 1 / 1024) - 1 = 1 / 1024 = 9.76e-4
# underflow level(with 0 as leading bit): 2^(-14)*(1/1024) = 2^(-24) = 5.96e-8 
# underflow level(): 2^(-14)*(1 + 0) = 6.10e-5

eps = np.finfo(np.float16).eps
var1 = np.float16(1)
print("eps:", eps)
print("var1", var1)
print("type eps", type(eps))
print("type var1", type(var1))
var2 = np.add(var1, eps)
# below doesn't work, as python will force it to be float64
# var3 = np.add(var1, eps / 2) 
var3 = np.float16(np.add(var1, eps / 2))

print("type var2", type(var2))
print("type var3", type(var3))
print(var2)
print(var3)
