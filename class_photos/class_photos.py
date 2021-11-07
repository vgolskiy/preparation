# Complexity O(nlog(n)) nlog(n) - complexity of the best sorting algorithm (Heap/Merge Sort),
# Space O(1)
def classPhotos(redShirtHeights, blueShirtHeights):
    redShirtHeights.sort()
    blueShirtHeights.sort()
    red = True if redShirtHeights[-1] < blueShirtHeights[-1] else False
    n = len(redShirtHeights)
    for i in range(n):
        if red:
            if redShirtHeights[n - (i + 1)] >= blueShirtHeights[n - (i + 1)]:
                return False
        else:
            if blueShirtHeights[n - (i + 1)] >= redShirtHeights[n - (i + 1)]:
                return False
    return True

if __name__ == "__main__":
    redShirtHeights = [5, 8, 1, 3, 4]
    blueShirtHeights = [6, 9, 2, 4, 5]
    expected = True
    actual = classPhotos(redShirtHeights, blueShirtHeights)
    print(actual, "vs", expected)
