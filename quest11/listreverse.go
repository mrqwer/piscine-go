package piscine

import "piscine/quest11/structs"

func ListReverse(l *structs.List) {
	curr := l.Head
	var prev *structs.NodeL

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head, l.Tail = l.Tail, l.Head
}

// head            tail
// [1]->[2]->[3]->[4]->[5]
// [1]->[2]->[3]->[4]->[5]
// slow
// fast
