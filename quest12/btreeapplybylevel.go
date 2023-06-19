package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	var stackNode []*TreeNode

	if root == nil {
		return
	}

	stackNode = append(stackNode, root)

	for len(stackNode) > 0 {
		n := len(stackNode)

		for i := 0; i < n; i++ {
			curr := stackNode[0]
			stackNode = stackNode[1:]
			f(curr.Data)
			if curr.Left != nil {
				stackNode = append(stackNode, curr.Left)
			}

			if curr.Right != nil {
				stackNode = append(stackNode, curr.Right)
			}
		}
	}
}
