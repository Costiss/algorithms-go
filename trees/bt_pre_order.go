package trees

type BinaryNode struct {
	Value int
	Left  *BinaryNode
	Right *BinaryNode
}

func PreOrderTraversal(root BinaryNode) []int {
	return walk(&root, []int{})
}

func walk(node *BinaryNode, path []int) []int {
	if node == nil {
		return path
	}

	path = append(path, node.Value)

	path = walk(node.Left, path)
	path = walk(node.Right, path)

	return path
}
