package main

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	if n == 0 {
		return 0
	}
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for j := 0; j < n; j++ {
		dp[0][j] = int(matrix[0][j] - '0')
		for i := 1; i < m; i++ {
			if matrix[i][j] == '1' {
				dp[i][j] = dp[i-1][j] + 1
			}
		}
	}
	max := 0
	for i := 0; i < m; i++ {
		tmp := largestRectangleArea(dp[i])
		if max < tmp {
			max = tmp
		}
	}
	return max
}

func largestRectangleArea(heights []int) int {
	h := append(heights, -1)
	n := len(h)
	var max, height, left, right, area int
	var stack []int
	for right < n {
		if len(stack) == 0 || h[stack[len(stack)-1]] <= h[right] {
			stack = append(stack, right)
			right++
			continue
		}
		height = h[stack[len(stack)-1]]
		stack = stack[:len(stack)-1]
		if len(stack) == 0 {
			left = -1
		} else {
			left = stack[len(stack)-1]
		}
		area = (right - left - 1) * height
		if max < area {
			max = area
		}
	}
	return max
}
