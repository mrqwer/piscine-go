package piscine

import "piscine/quest11/structs"

func IsPositiveNode(node *structs.NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *structs.NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return false
	default:
		return true
	}
}

func ListForEachIf(l *structs.List, f func(*structs.NodeL), cond func(*structs.NodeL) bool) {
	tmp := l.Head

	for tmp != nil {
		if cond(tmp) {
			f(tmp)
		}
		tmp = tmp.Next
	}
}
