package piscine

func BTreeIsBinary(root *TreeNode) bool {
	return isValid(root, "-2147483648", "2147483647")
}

func isValid(root *TreeNode, min, max string) bool {
	if root == nil {
		return true
	}

	if max <= root.Data {
		return false
	}
	if min >= root.Data {
		return false
	}
	return isValid(root.Left, min, root.Data) && isValid(root.Right, root.Data, max)
}
