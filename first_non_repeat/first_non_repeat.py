# Complexity O(n) n - quantity of elements in string
# Space O(1) - as we are iterating over a lower case English alphabet 26 symbols
def firstNonRepeatingCharacter(string):
    alp = dict()
    for c in string:
        alp[c] = alp.get(c, 0) + 1
    for i in range(len(string)):
        if alp[string[i]] == 1:
            return i
    return -1


if __name__ == "__main__":
    print(1, "vs", firstNonRepeatingCharacter("abcdcaf"))
