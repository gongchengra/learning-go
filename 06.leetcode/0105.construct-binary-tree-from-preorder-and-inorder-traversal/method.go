package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	res := &TreeNode{Val: preorder[0]}
	if len(preorder) == 1 {
		return res
	}
	idx := 0
	for preorder[0] != inorder[idx] {
		idx++
	}
	res.Left = buildTree(preorder[1:idx+1], inorder[:idx])
	res.Right = buildTree(preorder[idx+1:], inorder[idx+1:])
	return res
}
