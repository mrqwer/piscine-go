package piscine

import "piscine/quest11/structs"

func SortedListMerge(n1 *structs.NodeI, n2 *structs.NodeI) *structs.NodeI {
	dummy := &structs.NodeI{Data: 0, Next: nil}
	curr := dummy
	for n1 != nil && n2 != nil {
		if n1.Data > n2.Data {
			curr.Next = &structs.NodeI{Data: n2.Data, Next: nil}
			n2 = n2.Next
		} else {
			curr.Next = &structs.NodeI{Data: n1.Data, Next: nil}
			n1 = n1.Next
		}
		curr = curr.Next
	}

	for n1 != nil {
		curr.Next = &structs.NodeI{Data: n1.Data, Next: nil}
		n1 = n1.Next
		curr = curr.Next
	}

	for n2 != nil {
		curr.Next = &structs.NodeI{Data: n2.Data, Next: nil}
		n2 = n2.Next
		curr = curr.Next
	}
	return dummy.Next
}
