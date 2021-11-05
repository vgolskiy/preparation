# Complexity O(n), Space O(n) n - number of nodes in tree
class BinaryTree:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None

    def insert(self, values, i=0):
        if i >= len(values):
            return
        queue = [self]
        while len(queue) > 0:
            current = queue.pop(0)
            if current.left is None:
                current.left = BinaryTree(values[i])
                break
            queue.append(current.left)
            if current.right is None:
                current.right = BinaryTree(values[i])
                break
            queue.append(current.right)
        self.insert(values, i + 1)
        return self

def branchSumsHelper(node, sums, runSum):
    runSum += node.value
    if not node.left and not node.right:
        return sums.append(runSum)
    if node.left:
        branchSumsHelper(node.left, sums, runSum)
    if node.right:
        branchSumsHelper(node.right, sums, runSum)


def branchSums(root):
    sums = []
    branchSumsHelper(root, sums, 0)
    return sums


if __name__ == '__main__':
    tree = BinaryTree(1).insert([2, 3, 4, 5, 6, 7, 8, 9, 10])
    print("{} vs {}".format(branchSums(tree), [15, 16, 18, 10, 11]))
