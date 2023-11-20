package trees

func PreOrderTraversal(root BinaryNode) []int {
	return walkPreOrder(&root, []int{})
}

func walkPreOrder(node *BinaryNode, path []int) []int {
	if node == nil {
		return path
	}

	path = append(path, node.Value)

	path = walkPreOrder(node.Left, path)
	path = walkPreOrder(node.Right, path)

	return path
}
