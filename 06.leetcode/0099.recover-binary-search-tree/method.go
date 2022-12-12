package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Ints2TreeNode(ints []int) *TreeNode {
	n := len(ints)
	if n == 0 {
		return nil
	}
	root := &TreeNode{
		Val: ints[0],
	}
	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root
	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]
		if i < n && ints[i] != 0 {
			node.Left = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Left)
		}
		i++
		if i < n && ints[i] != 0 {
			node.Right = &TreeNode{Val: ints[i]}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// Tree2Preorder 把 二叉树 转换成 preorder 的切片

func Tree2Preorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := []int{root.Val}
	res = append(res, Tree2Preorder(root.Left)...)
	res = append(res, Tree2Preorder(root.Right)...)
	return res
}

// Tree2Inorder 把 二叉树转换成 inorder 的切片

func Tree2Inorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := Tree2Inorder(root.Left)
	res = append(res, root.Val)
	res = append(res, Tree2Inorder(root.Right)...)
	return res
}

// Tree2Postorder 把 二叉树 转换成 postorder 的切片

func Tree2Postorder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}
	res := Tree2Postorder(root.Left)
	res = append(res, Tree2Postorder(root.Right)...)
	res = append(res, root.Val)
	return res
}

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
