package piscine

import "piscine/quest11/structs"

func ListRemoveIf(l *structs.List, data_ref interface{}) {
	var prev *structs.NodeL
	curr := l.Head

	for curr != nil {
		if curr.Data == data_ref {
			if prev != nil {
				prev.Next = curr.Next
			} else {
				l.Head = curr.Next
			}

			if curr == l.Tail {
				l.Tail = prev
			}
		} else {
			prev = curr
		}

		curr = curr.Next
	}
}
