# Complexity O(log(n)) log(n) - average complaxity of BST operations, O(n) is the worst (all leafs in one direction)
# Space O(D) D - tree depth
def findClosestValueInBst(tree, target):
    return findClosest(tree, target, tree.value)


def findClosest(tree, target, closest):
    if not tree:
        return closest
    if tree.value == target:
        return tree.value
    if abs(tree.value - target) < abs(closest - target):
        closest = tree.value
    if tree.value < target:
        return findClosest(tree.right, target, closest)
    else:
        return findClosest(tree.left, target, closest)


class BST:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None


if __name__ == '__main__':
    root = BST(10)
    root.left = BST(5)
    root.left.left = BST(2)
    root.left.left.left = BST(1)
    root.left.right = BST(5)
    root.right = BST(15)
    root.right.left = BST(13)
    root.right.left.right = BST(14)
    root.right.right = BST(22)
    expected = 22
    actual = findClosestValueInBst(root, 20)
    print("{} vs {}".format(expected, actual))
