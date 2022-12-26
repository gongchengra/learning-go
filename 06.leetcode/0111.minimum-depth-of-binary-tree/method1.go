package main

func minDepth(root *TreeNode) int {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	switch {
	case root == nil:
		return 0
	case root.Left == nil:
		return 1 + minDepth(root.Right)
	case root.Right == nil:
		return 1 + minDepth(root.Left)
	default:
		return 1 + min(minDepth(root.Left), minDepth(root.Right))
	}
}
