package trees_test

import (
	"algos/trees"
	tree_mocks "algos/trees/__mocks__"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBTInOrder(t *testing.T) {

	assert.EqualValues(
		t,
		trees.InOrderTraversal(tree_mocks.Tree),
		[]int{
			5,
			7,
			10,
			15,
			20,
			29,
			30,
			45,
			50,
			100,
		},
	)

}
