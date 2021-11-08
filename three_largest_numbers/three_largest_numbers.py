# Complexity O(n) n - quantity of array elements,
# Space O(1)
def fillThree(num, three):
    three.append(num)
    n = len(three)
    if n == 1:
        return three
    if n == 3 and three[1] > three[2]:
        three[1], three[2] = three[2], three[1]
    if three[0] > three[1]:
        three[0], three[1] = three[1], three[0]
    return three


def setThreeNumbers(num, three):
    if three[2] < num:
        three[0], three[1], three[2] = three[1], three[2], num
    elif three[1] < num:
        three[0], three[1], three[2] = three[0], num, three[2]
    else:
        three[0], three[1], three[2] = num, three[1], three[2]


def findThreeLargestNumbers(array):
    three = []
    for i, num in enumerate(array):
        if i < 3:
            fillThree(num, three)
        if i >= 3 and three[0] < num:
            setThreeNumbers(num, three)
    return three

if __name__ == "__main__":
    array = [141, 1, 17, -7, -17, -27, 18, 541, 8, 7, 7]
    three = [18, 141, 541]
    print(findThreeLargestNumbers(array), "vs", three)
