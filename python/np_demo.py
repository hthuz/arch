
import numpy as np


A = np.array([[1,2],[3,4]])
def reset():
    global A
    A = np.arange(1,21).reshape(4,5)
    return
    
def show(msg):
    print(f"==={msg}===")


show("original matrix")
reset()
print(A)


show("transponse a matrix")
reset()
print(A.transpose())
print(np.transpose(A))
print(A.T)

show("flatten a matrix")
reset()
A = A.flatten()
print(A)


show("sum a matrix")
reset()
print(np.sum(A))
print(np.sum(A,axis=0))
print(np.sum(A,axis=1))

show("indexing a matrix")
reset()
print(A[0])
print(A[0:1])
print(A[0,1])
print(A[0][1])
print(A[:,1])
print(A[:,0:2])


show("multiplying two matrices")
reset()
B = A.transpose()
print(A @ B)
print(np.dot(A,B))
print(A.dot(B))
print(A * B.transpose())


show("list to matrix")
reset()
a = [1,2,3,4]
A = np.array(a)
print(A)

a = [[1,2,3],[4,5,6]]
A = np.array(a)
print(A)

show("stacking two matrics")
reset()
B = A[0,:]
print(np.vstack((A,B)))
B = A[:,0:1]
print(np.hstack((A,B)))









