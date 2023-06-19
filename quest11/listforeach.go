package piscine

import "piscine/quest11/structs"

func ListForEach(l *structs.List, f func(*structs.NodeL)) {
	tmp := l.Head

	for tmp != nil {
		f(tmp)
		tmp = tmp.Next
	}
}

func Add2_node(node *structs.NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) + 2
	case string:
		node.Data = node.Data.(string) + "2"
	}
}

func Subtract3_node(node *structs.NodeL) {
	switch node.Data.(type) {
	case int:
		node.Data = node.Data.(int) - 3
	case string:
		node.Data = node.Data.(string) + "-3"
	}
}
