package piscine

import "piscine/quest11/structs"

func ListSort(l *structs.NodeI) *structs.NodeI {
	curr := l
	for curr != nil {
		currNext := curr.Next
		for currNext != nil {
			if curr.Data > currNext.Data {
				curr.Data, currNext.Data = currNext.Data, curr.Data
			}
			currNext = currNext.Next
		}
		curr = curr.Next
	}
	return l
}
