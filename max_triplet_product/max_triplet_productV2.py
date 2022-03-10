# Complexity O(N*log(N))

import sys


def placeBig(n, big):
    if big[2] < n:
        if big[1] < n:
            if big[0] < n:
                return [n, big[0], big[1]]
            return [big[0], n, big[1]]
        return [big[0], big[1], n]
    return big


def placeLow(n, low):
    if low[1] > n:
        if low[0] > n:
            return [n, low[0]]
        return [low[0], n]
    return low


def solution(arr):
    big = [-(sys.maxsize - 1)] * 3
    low = [sys.maxsize] * 2
    n = len(arr)
    for i in range(n):
        big = placeBig(arr[i], big)
        low = placeLow(arr[i], low)
    return max(big[0] * big[1] * big[2], big[0] * low[0] * low[1])

print(solution([-5, 5, -5, 4]))
