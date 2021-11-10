# Complexity O(n^2) n - quantity of elements in array, O(n) is the best (all elements were sorted)
# Space O(1)
def insertion_sort(array):
    for i in range(1, len(array)):
        j = i
        while j > 0 and array[j] < array[j - 1]:
            array[j], array[j - 1] = array[j - 1], array[j]
            j -= 1
    return array


if __name__ == "__main__":
    print(insertion_sort([8, 5, 2, 9, 5, 6, 3]), "vs", [2, 3, 5, 5, 6, 8, 9])
