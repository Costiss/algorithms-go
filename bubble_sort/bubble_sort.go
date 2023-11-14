package bubblesort

func BubbleSort(array []int) []int {

	size := len(array)
	new_array := make([]int, size)
	copy(new_array, array)

	for i := 0; i < size; i++ {
		for j := 0; j < size-i-1; j++ {
			current := j
			next := j + 1

			if new_array[current] > new_array[next] {
				buff := new_array[next]
				new_array[next] = new_array[current]
				new_array[current] = buff
			}
		}

	}

	return new_array
}
