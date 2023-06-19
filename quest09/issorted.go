package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	ascendingCounter := 0
	descendingCounter := 0
	for i := 1; i < len(a); i++ {
		if f(a[i-1], a[i]) >= 0 {
			ascendingCounter++
		}
		if f(a[i-1], a[i]) <= 0 {
			descendingCounter++
		}
	}
	return ascendingCounter == len(a)-1 || descendingCounter == len(a)-1
}

// func f(a, b int) int {
// 	if a > b {
// 		return 1
// 	} else if a == b {
// 		return 0
// 	} else {
// 		return -1
// 	}
// }

// {0, 1, 2, 3, 4, 5}
// k   2k  3k  4k  5k  6k
//-k  -2k -3k -4k -5k -6k
// 0
// {0, 2, 1, 3}
//  k  2k k  2k
