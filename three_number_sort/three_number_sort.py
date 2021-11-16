# Complexity O(n) n - quantity of array elements,
# Space O(1)
def threeNumberSort(array, order):
    n = len(array)
    firstIdx = 0
    for i in range(n):
        if array[i] == order[0]:
            array[i], array[firstIdx] = array[firstIdx], array[i]
            firstIdx += 1
    lastIdx = n - 1
    for i in range(n - 1, firstIdx - 1, -1):
        if array[i] == order[2]:
            array[i], array[lastIdx] = array[lastIdx], array[i]
            lastIdx -= 1
    return array


if __name__ == "__main__":
    array = [1, 0, 0, -1, -1, 0, 1, 1]
    order = [0, 1, -1]
    expected = [0, 0, 0, 1, 1, 1, -1, -1]
    print(threeNumberSort(array, order), "vs", expected)
