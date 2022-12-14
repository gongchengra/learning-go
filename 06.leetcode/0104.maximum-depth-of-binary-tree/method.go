package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
