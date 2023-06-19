package piscine

func check(c byte) bool {
	return !(c >= 48 && c <= 57) && !(c >= 65 && c <= 90) && !(c >= 97 && c <= 122)
}

func Capitalize(s string) string {
	res := ""
	capitalizeNext := true
	for i := 0; i < len(s); i++ {
		if capitalizeNext && i != 0 && !check(s[i-1]) {
			capitalizeNext = false
		}
		if s[i] >= 97 && s[i] <= 122 && capitalizeNext {
			res += string(s[i] - 32)
			capitalizeNext = false
		} else if s[i] >= 'A' && s[i] <= 'Z' && !capitalizeNext {
			res += string(s[i] + 32)
		} else {
			res += string(s[i])
		}

		if check(s[i]) {
			capitalizeNext = true
		}

	}
	return res
}
