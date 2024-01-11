
import numpy as np

A = np.array(([7,2],
               [-4,1]))
B = np.transpose(A) @ A
print(A)
print(B)
print(np.linalg.eig(A))
print(np.linalg.eig(B))
