package piscine

func sort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
}

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	sort(arr)
	return arr[2]
}
