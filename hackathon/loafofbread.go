package piscine

func trim(str string) string {
	i := 0
	for ; i < len(str) && str[i] == ' '; i++ {
	}
	j := len(str) - 1
	for ; j >= 0 && str[j] == ' '; j-- {
	}
	if i >= j {
		return ""
	}
	return str[i : j+1]
}

func LoafOfBread(str string) string {
	str = trim(str)
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	ans := ""
	count := 0
	for i, ch := range str {
		if count%6 == 5 {
			count++
			if i != len(str)-1 {
				ans += " "
			}
		} else if ch == ' ' {
			continue
		} else {
			count++
			ans += string(ch)
		}
	}
	ans += "\n"
	return ans
}
