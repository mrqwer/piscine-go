package piscine

func partition(arr []string, low, high int) ([]string, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []string, low, high int) []string {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}

func quickSortStart(arr []string) []string {
	return quickSort(arr, 0, len(arr)-1)
}

func SortWordArr(a []string) {
	quickSortStart(a)
}
