package piscine

import "piscine/quest11/structs"

func ListAt(l *structs.NodeL, pos int) *structs.NodeL {
	if l == nil || pos < 0 {
		return nil
	}

	cnt := 0

	tmp := l
	for tmp != nil {
		if cnt == pos {
			return tmp
		}
		cnt++
		tmp = tmp.Next
	}
	return nil
}
