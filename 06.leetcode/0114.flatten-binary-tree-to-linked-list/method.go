package main

func flatten(root *TreeNode) {
	list := preorder(root)
	for i := 1; i < len(list); i++ {
		prev, cur := list[i-1], list[i]
		prev.Left, prev.Right = nil, cur
	}
	return
}

func preorder(root *TreeNode) (ans []*TreeNode) {
	if root != nil {
		ans = append(ans, root)
		ans = append(ans, preorder(root.Left)...)
		ans = append(ans, preorder(root.Right)...)
	}
	return
}
