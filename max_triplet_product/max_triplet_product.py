# Complexity O(N^3)

import sys

def solution(a):
    res = -(sys.maxsize - 1)
    n = len(a)
    for i in range(n - 2):
        for j in range(i + 1, n - 1):
            for k in range(j + 1, n):
                res = max(res, a[i] * a[j] * a[k])
    return res
