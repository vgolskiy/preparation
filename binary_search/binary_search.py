# Complexity O(log(n)) n - array elements quantity,
# Space O(log(n))
def BS_proc(array, target, left, right):
    if left > right:
        return -1
    pos = (right + left) // 2
    if target == array[pos]:
        return pos
    elif target < array[pos]:
        return BS_proc(array, target, left, pos - 1)
    return BS_proc(array, target, pos + 1, right)


def binarySearch(array, target):
    left = 0
    right = len(array) - 1
    return BS_proc(array, target, left, right)


if __name__ == "__main__":
    test = [0, 1, 21, 33, 45, 45, 61, 71, 72, 73]
    print(binarySearch(test, 33), "vs", 3)
