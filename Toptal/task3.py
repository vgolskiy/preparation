def prefixSum(a):
    n = len(a)
    p = [0]*(n + 1)
    for i in range(1, n + 1):
        p[i] = p[i - 1] + a[i - 1]
    return p

def solution(p, s):
    total = sum(p)
    s.sort(reverse=True)
    capacity = prefixSum(s)
    for i in range(len(capacity)):
        if capacity[i] >= total:
            return i
    return 0

print(solution([1,2], [1,2]))
