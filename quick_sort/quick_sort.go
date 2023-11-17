package quicksort

func QuickSort(array []int) {
	qs(array, 0, len(array)-1)
}

func qs(arr []int, low int, high int) {
	if low >= high {
		return
	}

	pivotIdx := partition(arr, low, high)

	qs(arr, low, pivotIdx-1)
	qs(arr, pivotIdx+1, high)
}

func partition(arr []int, low int, high int) int {

	pivot := arr[high]

	idx := low - 1

	for i := low; i < high; i++ {
		if arr[i] <= pivot {
			idx += 1
			tmp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = tmp
		}
	}

	idx += 1
	arr[high] = arr[idx]
	arr[idx] = pivot

	return idx
}
