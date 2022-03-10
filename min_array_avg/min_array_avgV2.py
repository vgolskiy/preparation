# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def sliceAvg2(a, i):
    return (a[i] + a[i + 1])/2

def sliceAvg3(a, i):
    return (a[i] + a[i + 1] + a[i + 2])/3

def solution(a):
    n = len(a)
    res = sliceAvg2(a, 0)
    pos = 0
    for i in range(n - 2):
        tmp2 = sliceAvg2(a, i)
        if tmp2 < res:
            res = tmp2
            pos = i
        tmp3 = sliceAvg3(a, i)
        if tmp3 < res:
            res = tmp3
            pos = i
    tmp2 = sliceAvg2(a, n - 2)
    if tmp2 < res:
        pos = n - 2
    return pos

print(solution([2, 2, 5]))
