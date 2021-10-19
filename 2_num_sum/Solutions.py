# Complexity O(n^2), Space O(1)
def two_num_sum1(arr, target):
    for i in range(len(arr) - 1):
        for j in range(i +1, len(arr)):
            if arr[i] + arr[j] == target:
                return [arr[i], arr[j]]
    return []


# Complexity O(n*log(n)), Space O(1)
def two_num_sum2(arr, target):
    arr.sort()
    # starting positions of indexes
    begin = 0
    end = len(arr) - 1
    while begin < end:
        curr_sum = arr[begin] + arr[end]
        if curr_sum == target:
            return [arr[begin], arr[end]]
        elif curr_sum < target:
            begin += 1
        else:
            end -= 1
    return []
