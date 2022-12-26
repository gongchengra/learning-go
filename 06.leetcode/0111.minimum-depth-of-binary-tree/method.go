package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil {
		return 1 + minDepth(root.Right)
	}
	if root.Right == nil {
		return 1 + minDepth(root.Left)
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	return 1 + min(minDepth(root.Left), minDepth(root.Right))
}
