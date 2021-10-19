# Complexity O(n), Space O(1)
def validate_subseq(arr, sub):
    j = 0
    i = 0
    while i < len(arr) and j < len(sub):
        if arr[i] == sub[j]:
            j += 1
        i += 1
    return j == len(sub)


if __name__ == '__main__':
    print(validate_subseq([5, 1, 22, 25, 6, -1, 8, 10], [1, 6, -1, 10]))
