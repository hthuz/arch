
# Find longest subsequence of a string

import numpy as np

def LISdp(A):
    n = len(A)
    A.append(np.inf)
    subseq = []

    # Initialize base case
    LIS = np.zeros((n + 1,n + 1))
    LIS[0][:] = 0

    # Note the matrix's index are not the same
    # This really makes the matrix really confusing
    # i = 1 to n
    for i in range(1, n + 1):
        # j = i to n
        for j in range(i - 1, n + 1):
            if A[i - 1] >= A[j]:
                LIS[i][j] = LIS[i - 1][j]
            else:
                LIS[i][j] = max(LIS[i - 1][j], LIS[i - 1,i - 1] + 1)

    # Get exact subsequence. Get a little crazy about the index
    j = n
    for i in range(n, 0, -1):
        if LIS[i,j] == LIS[i - 1, i - 1] + 1:
            subseq.append(A[i - 1])
            j = i - 1

    subseq.reverse()

    return subseq, LIS[n, n]

if __name__ == "__main__":
    subseq, val = LISdp([6,3,5,2,7,8,1])
    print(subseq, val)

