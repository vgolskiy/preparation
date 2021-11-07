# Complexity O(nlog(n)) nlog(n) - complexity of the best sorting algorithm (Heap/Merge Sort),
# Space O(1)
def maximal_speed(reds, blues):
    n = len(reds)
    tandem = 0
    for i in range(n):
        tandem += max(reds[n - 1 - i], blues[i])
    return tandem


def minimal_speed(reds, blues):
    n = len(reds)
    tandem = 0
    for i in range(n):
        tandem += max(reds[i], blues[i])
    return tandem


def tandemBicycle(redShirtSpeeds, blueShirtSpeeds, fastest):
    redShirtSpeeds.sort()
    blueShirtSpeeds.sort()
    if fastest:
        return maximal_speed(redShirtSpeeds, blueShirtSpeeds)
    return minimal_speed(redShirtSpeeds, blueShirtSpeeds)


if __name__ == "__main__":
    redShirtSpeeds = [5, 5, 3, 9, 2]
    blueShirtSpeeds = [3, 6, 7, 2, 1]
    fastest = False
    expected = 32 if fastest else 25
    actual = tandemBicycle(redShirtSpeeds, blueShirtSpeeds, fastest)
    print(actual, "vs", expected)
