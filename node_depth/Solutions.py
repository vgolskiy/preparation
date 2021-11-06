#Complexity: O(n^2) n - number of nodes
#Space: O(n)
def nodeDepthsHelper(node, lvls, lvl):
    if not node:
        return
    if lvl:
        lvls.append(lvl)
    nodeDepthsHelper(node.left, lvls, lvl + 1)
    nodeDepthsHelper(node.right, lvls, lvl + 1)


def nodeDepths(root):
    lvls = []
    nodeDepthsHelper(root, lvls, lvl=0)
    return sum(lvls)


class BinaryTree:
    def __init__(self, value):
        self.value = value
        self.left = None
        self.right = None

    def insert(self, value):
        if value < self.value:
            if not self.left:
                self.left = BinaryTree(value)
            else:
                self.left.insert(value)
        else:
            if not self.right:
                self.right = BinaryTree(value)
            else:
                self.right.insert(value)

    def __str__(self, level=0):
        res = "\t"*level+str(self.value)+"\n"
        if self.left:
            res += self.left.__str__(level+1)
        if self.right:
            res += self.right.__str__(level+1)
        return res


if __name__ == '__main__':
    tree = BinaryTree(5)
    tree.insert(4)
    tree.insert(3)
    tree.insert(2)
    tree.insert(1)
    tree.insert(6)
    tree.insert(7)
    tree.insert(8)
    tree.insert(9)
    tree.insert(10)
    print(tree)
    print("25 vs", nodeDepths(tree))
