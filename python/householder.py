
import numpy as np

N = 3
v = np.array([
    [2],
    [3 + np.sqrt(29)],
    [4]
])

a = np.array([
    [2],
    [3],
    [4]
])

def householder(v):
    return np.eye(N) - 2 * v @ np.transpose(v) / (np.transpose(v) @ v)

print(v)
print(householder(v))
print(householder(v) @ a)

