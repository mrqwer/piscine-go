package piscine

func ConcatParams(args []string) string {
	res := ""

	for i := range args {
		res += args[i]
		if i != len(args)-1 {
			res += "\n"
		}
	}

	return res
}
