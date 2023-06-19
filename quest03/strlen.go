package piscine

func StrLen(s string) int {
	cnt := 0
	for _, r := range s {
		_ = r
		cnt++
	}
	return cnt
}
