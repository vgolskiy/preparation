def getPermutations(array):
    perms = []
    permutationHelper(array, [], perms)
    return perms


def permutationHelper(array, currPerm, perms):
    if not len(array) and len(currPerm):
        perms.append(currPerm)
    else:
        for i in range(len(array)):
            newArr = array[:i] + array[i+1:]
            newPerm = currPerm + [array[i]]
            permutationHelper(newArr, newPerm, perms)

if __name__ == "__main__":
    print(getPermutations([1, 2, 3]))
