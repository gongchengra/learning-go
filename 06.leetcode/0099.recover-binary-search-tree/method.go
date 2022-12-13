package main

func recoverTree(root *TreeNode) {
	// 用于记录树中的两个错误节点
	var x, y, pred *TreeNode
	// 中序遍历二叉搜索树，找到两个错误节点
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		if pred != nil && root.Val < pred.Val {
			y = root
			if x == nil {
				x = pred
			} else {
				return
			}
		}
		pred = root
		dfs(root.Right)
	}
	dfs(root)
	// 交换两个错误节点的值
	x.Val, y.Val = y.Val, x.Val
}
