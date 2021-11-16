# Complexity O(n^2) n - length of the input array
# Space O(n)
def threeNumberSum(array, targetSum):
    array.sort()
    res = []
    n = len(array)
    for i in range(n - 2):
        left = i + 1
        right = n - 1
        while left < right:
            currSum = array[i] + array[left] + array[right]
            if currSum == targetSum:
                res.append([array[i], array[left], array[right]])
                left += 1
                right -= 1
            elif currSum > targetSum:
                right -= 1
            else:
                left += 1
    return res

if __name__ == "__main__":
    print(threeNumberSum([12, 3, 1, 2, -6, 5, -8, 6], 0), "vs", [[-8, 2, 6], [-8, 3, 5], [-6, 1, 5]])
