#Complexity: O(n + m) n - length of document, m - number of characters
#Space: O(c) c - number of unique characters
def generateDocument(characters, document):
    if len(document):
        buf = dict()
        for c in characters:
            buf[c] = buf.get(c, 0) + 1
        for c in document:
            if c not in buf or buf[c] - 1 < 0:
                return False
            buf[c] -= 1
    return True


if __name__ == "__main__":
    characters = "Bste!hetsi ogEAxpelrt x "
    document = "AlgoExpert is the Best!"
    print(True, "vs", generateDocument(characters, document))
