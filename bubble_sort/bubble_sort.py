# Complexity O(n^2) n - quantity of elements in array, O(n) is the best (all elements were sorted)
# Space O(1)
# qty is used for optimization (at the end of the array all elements are already sorted)
def bubbleSort(array):
    more = True
    qty = 0
    while more:
        more = False
        for i in range(len(array) - 1 - qty):
            if array[i] > array[i + 1]:
                array[i], array[i + 1] = array[i + 1], array[i]
                more = True
        qty += 1
    return array

if __name__ == "__main__":
    print(bubbleSort([8, 5, 2, 9, 5, 6, 3]), "vs", [2, 3, 5, 5, 6, 8, 9])
