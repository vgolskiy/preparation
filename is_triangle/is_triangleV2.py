#Complexity N*log(N)
def solution(a):
    n = len(a)
    if n >= 3:
        a.sort()
        for i in range(n - 2):
            if a[i] + a[i + 1] > a[i + 2]:
                return 1
    return 0
