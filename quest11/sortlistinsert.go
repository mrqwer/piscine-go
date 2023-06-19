package piscine

import "piscine/quest11/structs"

func SortListInsert(l *structs.NodeI, data_ref int) *structs.NodeI {
	curr := l
	for curr.Next != nil {
		curr = curr.Next
	}

	curr.Next = &structs.NodeI{Data: data_ref}
	return ListSort(l)
}
