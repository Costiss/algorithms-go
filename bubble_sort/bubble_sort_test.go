package bubblesort_test

import (
	bubblesort "algos/bubble_sort"
	"testing"

	"gotest.tools/v3/assert"
)

func TestBubbleSort(t *testing.T) {

	unsorted := []int{9, 5, 0, 1, 2, 8, 12}

	sorted := bubblesort.BubbleSort(unsorted)

	assert.DeepEqual(
		t,
		[]int{0, 1, 2, 5, 8, 9, 12},
		sorted,
	)
}
