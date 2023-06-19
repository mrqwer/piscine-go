package piscine

func minValueNode(node *TreeNode) *TreeNode {
	curr := node
	for curr.Left != nil {
		curr = curr.Left
	}
	return curr
}

func BTreeDeleteNode(root, node *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	if root.Data > node.Data {
		root.Left = BTreeDeleteNode(root.Left, node)
	} else if root.Data < node.Data {
		root.Right = BTreeDeleteNode(root.Right, node)
	} else {
		var temp *TreeNode
		if root.Left == nil {
			temp = root.Right
			root = nil
			return temp
		} else {
			temp = root.Left
			root = nil
			return temp
		}
		temp = minValueNode(root.Right)

		root.Data = temp.Data

		root.Right = BTreeDeleteNode(root.Right, temp)
	}
	return root
}
