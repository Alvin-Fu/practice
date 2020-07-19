package main

type Node struct {
	Val       int
	Neighbors []*Node
}

func main() {

}

func cloneGraph(node *Node) *Node {
	temp := make(map[int]*Node)
	return clone(node, temp)
}

func clone(node *Node, temp map[int]*Node) *Node {
	if node == nil {
		return nil
	}
	if v, ok := temp[node.Val]; ok {
		return v
	}
	newNode := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	temp[node.Val] = newNode
	for i := 0; i < len(node.Neighbors); i++ {
		newNode.Neighbors[i] = clone(node.Neighbors[i], temp)
	}
	return newNode
}
