package piscine

import (
	"piscine/quest11/structs"
)

func ListPushBack(l *structs.List, data interface{}) {
	newNode := &structs.NodeL{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		l.Tail.Next = newNode
		l.Tail = newNode
	}
}
