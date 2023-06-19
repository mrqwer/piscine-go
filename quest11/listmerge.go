package piscine

import "piscine/quest11/structs"

func ListMerge(l1 *structs.List, l2 *structs.List) {
	if l1.Head == nil && l2.Head == nil {
		return
	} else if l1.Head == nil && l2.Head != nil {
		l1.Head = l2.Head
		return
	}
	curr := l1.Head

	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = l2.Head
}
