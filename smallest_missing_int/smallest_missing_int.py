# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

def solution(a):
    a = list(set(a))
    a.sort()
    res = 0
    prev = 0
    start = True
    for n in a:
        if n > 0:
            if prev:
                if n - prev > 1:
                    return prev + 1
                prev = n
            if start:
                if n != 1:
                    return 1
                start = False
                prev = n
    if start:
        return 1
    return n + 1
