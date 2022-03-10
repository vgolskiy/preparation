# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def prefixSum(a):
    n = len(a)
    p = [0]*(n + 1)
    for i in range(1, n + 1):
        p[i] = p[i - 1] + a[i - 1]
    return p

def sliceAvg(p, start, end):
    return (p[end] - p[start - 1])/(end - start + 1)

def solution(a):
    p = prefixSum(a)
    n = len(a)
    res = sliceAvg(p, 0, 1)
    pos = 0
    for i in range(n - 2):
        tmp2 = sliceAvg(p, i + 1, i + 2)
        if tmp2 < res:
            res = tmp2
            pos = i
        tmp3 = sliceAvg(p, i + 1, i + 3)
        if tmp3 < res:
            res = tmp3
            pos = i
    tmp2 = sliceAvg(p, n - 1, n)
    if tmp2 < res:
        pos = n - 2
    return pos

print(solution([2, 2, 5]))
