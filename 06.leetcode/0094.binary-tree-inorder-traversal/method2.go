package main

func inorderTraversal(root *TreeNode) (res []int) {
	var inorder func(root *TreeNode)
	inorder = func(root *TreeNode) {
		if root != nil {
			inorder(root.Left)
			res = append(res, root.Val)
			inorder(root.Right)
		}
	}
	inorder(root)
	return res
}
