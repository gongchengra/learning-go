package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var recur func(l, r *TreeNode) bool
	recur = func(l, r *TreeNode) bool {
		if l == nil && r == nil {
			return true
		}
		if l == nil || r == nil {
			return false
		}
		return l.Val == r.Val && recur(l.Left, r.Right) && recur(l.Right, r.Left)
	}
	return recur(root.Left, root.Right)
}
