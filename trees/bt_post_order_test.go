package trees_test

import (
	"algos/trees"
	tree_mocks "algos/trees/__mocks__"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBTPostOrder(t *testing.T) {

	assert.EqualValues(
		t,
		trees.PostOrderTraversal(tree_mocks.Tree),
		[]int{
			7,
			5,
			15,
			10,
			29,
			45,
			30,
			100,
			50,
			20,
		},
	)

}
