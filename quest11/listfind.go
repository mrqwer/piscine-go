package piscine

import "piscine/quest11/structs"

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *structs.List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	tmp := l.Head

	for tmp != nil {
		if comp(tmp.Data, ref) {
			return &tmp.Data
		}
		tmp = tmp.Next
	}
	return nil
}
