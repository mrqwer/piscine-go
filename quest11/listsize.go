package piscine

import "piscine/quest11/structs"

func ListSize(l *structs.List) int {
	cnt := 0

	tmp := l.Head

	for tmp != nil {
		cnt++
		tmp = tmp.Next
	}

	return cnt
}
