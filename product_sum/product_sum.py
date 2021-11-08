# Complexity O(n) n - number of all arrays elements,
# Space O(d) d - maximal depth of array
def productSum(array, d = 1):
    summa = 0
    for n in array:
        if type(n) is list:
            summa += productSum(n, d + 1)
        else:
            summa += n
    return summa * d


if __name__ == "__main__":
    test = [5, 2, [7, -1], 3, [6, [-13, 8], 4]]
    print(12, "vs", productSum(test))
