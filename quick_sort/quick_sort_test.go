package quicksort_test

import (
	quicksort "algos/quick_sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSort(t *testing.T) {

	arr := []int{9, 3, 7, 4, 69, 420, 42}

	quicksort.QuickSort(arr)

	assert.EqualValues(t,
		[]int{3, 4, 7, 9, 42, 69, 420},
		arr,
	)
}
