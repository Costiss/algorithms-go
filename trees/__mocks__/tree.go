package tree_mocks

import "algos/trees"

var Tree = trees.BinaryNode{
	Value: 20,
	Left: &trees.BinaryNode{
		Value: 10,
		Left: &trees.BinaryNode{
			Value: 5,
			Left:  nil,
			Right: &trees.BinaryNode{
				Value: 7,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &trees.BinaryNode{
			Value: 15,
			Left:  nil,
			Right: nil,
		},
	},
	Right: &trees.BinaryNode{
		Value: 50,
		Left: &trees.BinaryNode{
			Value: 30,
			Left: &trees.BinaryNode{
				Value: 29,
				Left:  nil,
				Right: nil,
			},
			Right: &trees.BinaryNode{
				Value: 45,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &trees.BinaryNode{
			Value: 100,
			Left:  nil,
			Right: nil,
		},
	},
}
