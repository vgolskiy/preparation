# Complexity O(n*log(n) + m*log(m)) n, m - quantities of elements in arrays,
# nlog(n) - complexity of the best sorting algorithm (Heap/Merge Sort),
# Space O(1)
def smallestDifference(arrayOne, arrayTwo):
    arrayOne.sort()
    arrayTwo.sort()
    i = 0
    j = 0
    delta = float('inf')
    while i < len(arrayOne) and j < len(arrayTwo):
        if arrayOne[i] == arrayTwo[j]:
            return [arrayOne[i], arrayTwo[j]]
        elif arrayOne[i] < arrayTwo[j]:
            if delta > arrayTwo[j] - arrayOne[i]:
                delta = arrayTwo[j] - arrayOne[i]
                res = [arrayOne[i], arrayTwo[j]]
            i += 1
        else:
            if delta > arrayOne[i] - arrayTwo[j]:
                delta = arrayOne[i] - arrayTwo[j]
                res = [arrayOne[i], arrayTwo[j]]
            j += 1
    return res


if __name__ == "__main__":
    print(smallestDifference([-1, 5, 10, 20, 28, 3], [26, 134, 135, 15, 17]), "vs", [28, 26])
