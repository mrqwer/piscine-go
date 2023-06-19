package piscine

func Compact(ptr *[]string) int {
	cnt := 0
	for i := range *ptr {
		if len((*ptr)[i]) != 0 {
			cnt++
		}
	}
	temp := make([]string, cnt)

	j := 0
	for _, v := range *ptr {
		if len(v) != 0 {
			temp[j] = v
			j++
		}
	}
	(*ptr) = temp
	// for i := range temp {
	// 	(*ptr)[i] = temp[i]
	// }

	// for i := j; i < len((*ptr)); i++ {
	// 	(*ptr)[i] = ""
	// }
	return len((*ptr))
}
