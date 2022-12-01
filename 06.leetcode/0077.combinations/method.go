package main

func combine(n int, k int) [][]int {
	combination := make([]int, k)
	res := [][]int{}
	var dfs func(int, int)
	dfs = func(idx, begin int) {
		if idx == k {
			tmp := make([]int, k)
			copy(tmp, combination)
			res = append(res, tmp)
			return
		}
		for i := begin; i <= n+1-k+idx; i++ {
			combination[idx] = i
			dfs(idx+1, i+1)
		}
	}
	dfs(0, 1)
	return res
}
