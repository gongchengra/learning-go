package main

func pathSum(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}
	path := []int{}
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, level int, sum int) {
		if node == nil {
			return
		}
		if level >= len(path) {
			path = append(path, node.Val)
		} else {
			path[level] = node.Val
		}
		sum -= node.Val
		if node.Left == nil && node.Right == nil && sum == 0 {
			tmp := make([]int, level+1)
			copy(tmp, path)
			res = append(res, tmp)
		}
		dfs(node.Left, level+1, sum)
		dfs(node.Right, level+1, sum)
	}
	dfs(root, 0, targetSum)
	return res
}
