# Complexity O(n) - n-number of nodes in list,
# Space O(1)
class LinkedList:
    def __init__(self, value):
        self.value = value
        self.next = None

    def addMany(self, values):
        current = self
        while current.next is not None:
            current = current.next
        for value in values:
            current.next = LinkedList(value)
            current = current.next
        return self

    def getNodesInArray(self):
        nodes = []
        current = self
        while current is not None:
            nodes.append(current.value)
            current = current.next
        return nodes


def removeDuplicatesFromLinkedList(linkedList):
    curr = linkedList
    while curr:
        nextDist = curr.next
        while nextDist and nextDist.value == curr.value:
            nextDist = nextDist.next
        curr.next = nextDist
        curr = nextDist
    return linkedList


if __name__ == "__main__":
    test = LinkedList(1).addMany([1, 3, 4, 4, 4, 5, 6, 6])
    expected = LinkedList(1).addMany([3, 4, 5, 6])
    actual = removeDuplicatesFromLinkedList(test)
    print(actual.getNodesInArray(), "vs", expected.getNodesInArray())
