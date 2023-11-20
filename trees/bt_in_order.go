package trees

func InOrderTraversal(root BinaryNode) []int {
	return walkInOrder(&root, []int{})
}

func walkInOrder(node *BinaryNode, path []int) []int {
	if node == nil {
		return path
	}

	path = walkInOrder(node.Left, path)

	path = append(path, node.Value)

	path = walkInOrder(node.Right, path)

	return path
}
