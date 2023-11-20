package trees

func PostOrderTraversal(root BinaryNode) []int {
	return walkPostOrder(&root, []int{})
}

func walkPostOrder(node *BinaryNode, path []int) []int {
	if node == nil {
		return path
	}

	path = walkPostOrder(node.Left, path)
	path = walkPostOrder(node.Right, path)

	path = append(path, node.Value)
	return path
}
