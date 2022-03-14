import bisect

def solution(a):
    target = sum(a)/2
    a.sort()
    delta = 0
    qty = 0
    n = len(a)
    while delta < target:
        maximal = a.pop(n -1)
        delta += maximal/2
        bisect.insort(a, maximal/2)
        qty += 1
    return qty

print(solution([0]))
