# Complexity O(nlog(n)) nlog(n) - complexity of the best sorting algorithm (Heap/Merge Sort),
# Space O(1)
def minimumWaitingTime(queries):
    queries.sort()
    length = 0
    n = len(queries)
    for i, q in enumerate(queries):
        length += q * (n - (i + 1))
    return length


if __name__ == '__main__':
    queries = [3, 2, 1, 2, 6]
    expected = 17
    actual = minimumWaitingTime(queries)
    print(actual, "vs", expected)
