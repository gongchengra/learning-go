package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	if target < matrix[0][0] || target > matrix[len(matrix)-1][len(matrix[0])-1] {
		return false
	}
	r := 0
	for r < len(matrix) && matrix[r][0] <= target {
		r++
	}
	r--
	i, j := 0, len(matrix[0])-1
	for i <= j {
		mid := (i + j) / 2
		switch {
		case matrix[r][mid] < target:
			i = mid + 1
		case matrix[r][mid] > target:
			j = mid - 1
		default:
			return true
		}
	}
	return matrix[r][j] == target
}
