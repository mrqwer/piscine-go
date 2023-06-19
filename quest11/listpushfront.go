package piscine

import "piscine/quest11/structs"

func ListPushFront(l *structs.List, data interface{}) {
	newNode := &structs.NodeL{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
}
