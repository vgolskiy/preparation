# Complexity O(n) n - quantity of elements in array,
# Space O(1)
def verify_peak(array, i):
    return array[i - 1] < array[i] and array[i] > array[i + 1]

def longestPeak(array):
    max_len = 0
    i = 1
    while i < len(array) - 1:
        if verify_peak(array, i):
            left = i - 2
            while left >= 0 and array[left] < array[left + 1]:
                left -= 1
            right = i + 2
            while right < len(array) and array[right] < array[right - 1]:
                right += 1
            max_len = max(max_len, right - left - 1)
            i = right
        else:
            i += 1
    return max_len


if __name__ == "__main__":
    array = [1, 2, 3, 3, 4, 0, 10, 6, 5, -1, -3, 2, 3]
    expected = 6
    print(longestPeak(array), "vs", expected)
