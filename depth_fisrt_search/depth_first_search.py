# Complexity O(v + e) - v - number of vertices, e - number of edges (need to write name to array + iterate thru the children)
# Space O(v) - length of returning array
class Node:
    def __init__(self, name):
        self.children = []
        self.name = name

    def addChild(self, name):
        self.children.append(Node(name))
        return self

    def depthFirstSearch(self, array):
        array.append(self.name)
        for child in self.children:
            child.depthFirstSearch(array)
        return array


if __name__ == '__main__':
    graph = Node("A")
    graph.addChild("B").addChild("C").addChild("D")
    graph.children[0].addChild("E").addChild("F")
    graph.children[2].addChild("G").addChild("H")
    graph.children[0].children[1].addChild("I").addChild("J")
    graph.children[2].children[0].addChild("K")
    print(graph.depthFirstSearch([]), "vs", ["A", "B", "E", "F", "I", "J", "C", "D", "G", "K", "H"])


