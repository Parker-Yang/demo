package quick_sort

func QuickSort(arr []int, left, right int) []int {
	if left < right {
		partitionIndex := partition(arr, left, right)
		QuickSort(arr, left, partitionIndex-1)
		QuickSort(arr, partitionIndex+1, right)
	}
	return arr
}

/*func partition(arr []int, left, right int) int {
	pivot := left
	index := left + 1
	for i := index; i <= right; i++ {
		if arr[i] < arr[pivot] {
			swap(arr, index, i)
			index++
		}
	}
	swap(arr, pivot, index-1)
	return index - 1
}*/

func partition(arr []int, left, right int) int {
	pivot := left
	i := left + 1
	j := right
	for ;i<=j; {
		if arr[j]<arr[pivot] {
			if arr[i]>arr[pivot] {
				swap(arr,i,j)
			}else {
				i++
			}
		}else {
			j--
		}
	}
	swap(arr,pivot,i-1)
	return i-1
}

func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
