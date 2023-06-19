package piscine

import "piscine/quest11/structs"

func ListClear(l *structs.List) {
	if l.Head == nil {
		return
	}

	l.Head = nil
	l.Tail = nil
}
