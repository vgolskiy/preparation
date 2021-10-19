# Complexity O(n), Space O(n)
def sorted_squared_array(arr):
    pos = len(arr) - 1
    res = [0] * len(arr)
    begin = 0
    end = len(arr) - 1
    while begin <= end:
        if abs(arr[begin]) < abs(arr[end]):
            res[pos] = pow(arr[end], 2)
            end -= 1
        else:
            res[pos] = pow(arr[begin], 2)
            begin += 1
        pos -= 1
    return res


if __name__ == '__main__':
    print(sorted_squared_array([-7, -5, -4, 3, 6, 8, 9]))
