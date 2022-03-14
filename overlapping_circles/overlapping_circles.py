# you can write to stdout for debugging purposes, e.g.
# print("this is a debug message")

# def solution(a):
#     n = len(a)
#     pic = []
#     for i in range(n):
#         pic.append((i - a[i], 0))
#         pic.append((i + a[i], 1))
#     pic.sort()
#     print(pic)
#     qty = 0
#     res = 0
#     for el in pic:
#         if not el[1]:
#             qty += 1
#         if el[1]:
#             qty -= 1
#     return res

from collections import defaultdict

a = [1, 10, 100, 1]

def solution(a):
    start = defaultdict(int)
    stop = defaultdict(int)
    for i in range(len(a)):
        start[i - a[i]] += 1
        stop[i + a[i]] += 1
    active = 0
    intersections = 0
    for i in range(-len(a), len(a)):
        intersections += active * start[i] + (start[i] * (start[i] - 1)) / 2
        active += start[i]
        active -= stop[i]
        if intersections > 10000000:
            return -1
    return int(intersections)

print(solution(a))

# print(solution([1, 0]))
