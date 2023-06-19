package piscine

// func Unmatch(a []int) int {
// 	m := make(map[int]int)
// 	for i := 0; i < len(a); i++ {
// 		m[a[i]]++
// 	}
// 	for i := 0; i < len(a); i++ {
// 		if v := m[a[i]]; v%2 == 1 {
// 			return a[i]
// 		}
// 	}

// 	return -1
// }

func Unmatch(a []int) int {
	m := make(map[int]int)
	for _, num := range a {
		m[num]++
	}
	for _, num := range a {
		if m[num]%2 == 1 {
			return num
		}
	}

	return -1
}
