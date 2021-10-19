# Complexity O(n*log(n)), Space O(1)
def sorted_squared_array(arr):
    arr = [n * n for n in arr]
    arr.sort()
    return arr


if __name__ == '__main__':
    print(sorted_squared_array([-7, -5, -4, 3, 6, 8, 9]))
