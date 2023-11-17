package trees_test

import (
	"algos/trees"
	tree_mocks "algos/trees/__mocks__"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBTPreOrder(t *testing.T) {

	assert.EqualValues(
		t,
		trees.PreOrderTraversal(tree_mocks.Tree),
		[]int{
			20,
			10,
			5,
			7,
			15,
			50,
			30,
			29,
			45,
			100,
		},
	)

}
