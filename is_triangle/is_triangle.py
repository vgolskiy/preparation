#Complexity O(N^3)
def verify_triangle(a, b, c):
    if a + b > c and a + c > b and c + b > a:
        return True
    return False

def solution(a):
    n = len(a)
    for i in range(n - 2):
        for j in range(i + 1, n - 1):
            for k in range(j + 1, n):
                if verify_triangle(a[i], a[j], a[k]):
                    return 1
    return 0
