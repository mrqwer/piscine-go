package piscine

// func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
// 	if root == nil || root.Data == elem {
// 		return root
// 	}

// 	if root.Data > elem {
// 		return BTreeSearchItem(root.Left, elem)
// 	} else {
// 		return BTreeSearchItem(root.Right, elem)
// 	}
// }

//                         [4]
//                   [2]          [7]
//                [0]   [3]   [5]      [12]
//   [7]

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	current := root
	parent := current
	for current != nil {
		if current.Data == elem {
			current.Parent = parent
			return current
		} else if current.Data > elem {
			parent = current
			current = current.Left
		} else {
			parent = current
			current = current.Right
		}
	}

	return nil
}
