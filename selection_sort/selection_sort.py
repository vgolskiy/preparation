def find_smaller(array):
    minimal = 0
    for i in range(1, len(array)):
        if array[minimal] > array[i]:
            minimal = i
    return minimal

def swap(i, j, array):
    array[i], array[j] = array[j], array[i]

def selectionSort(array):
    for i in range(len(array)):
        j = find_smaller(array[i:]) + i
        swap(i, j, array)
    return array

if __name__ == "__main__":
    test = [8, 5, 2, 9, 5, 6, 3]
    example = [2, 3, 5, 5, 6, 8, 9]
    print(selectionSort(test), "vs", example)
