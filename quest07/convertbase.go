package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	n := atoi(nbr, baseFrom)
	res := printNbrBase(n, baseTo)
	return res
}

func printNbrBase(nbr int, base string) string {
	negative := false
	if nbr < 0 {
		negative = true
		nbr = -nbr
	}
	result := ""
	b := len(base)
	for nbr > 0 {
		remainder := nbr % b
		result = string(base[remainder]) + result
		nbr = nbr / b
	}
	if negative {
		result = "-" + result
	}
	return result
}

func atoi(s string, base string) int {
	result := 0
	for i := 0; i < len(s); i++ {
		idx := -1
		for j := 0; j < len(base); j++ {
			if s[i] == base[j] {
				idx = j
				break
			}
		}
		if idx == -1 {
			return 0
		}
		result = result*len(base) + idx
	}
	return result
}
