package piscine

func IsPrintable(s string) bool {
	cnt := 0
	for _, r := range s {
		if r >= 32 && r <= 126 {
			cnt++
		}
	}
	return len(s) == cnt
}
