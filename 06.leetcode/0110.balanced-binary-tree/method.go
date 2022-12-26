package main

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	l, r := depth(root.Left), depth(root.Right)
	if l < r {
		l, r = r, l
	}
	abs := func(a int) int {
		if a > 0 {
			return a
		}
		return -a
	}
	return abs(l-r) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	return max(depth(root.Left), depth(root.Right)) + 1
}
