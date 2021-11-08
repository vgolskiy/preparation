# Complexity O(n) n - input number,
# Space O(1)
def getNthFib(n):
    prev = 0
    curr = 1
    if n > 1:
        while (n > 2):
            tmp = curr
            curr = curr + prev
            prev = tmp
            n -= 1
        return curr
    return prev


if __name__ == "__main__":
    print(getNthFib(3))
