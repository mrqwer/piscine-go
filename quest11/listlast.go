package piscine

import "piscine/quest11/structs"

func ListLast(l *structs.List) interface{} {
	if l.Head == nil {
		return nil
	}

	return l.Tail.Data
}
