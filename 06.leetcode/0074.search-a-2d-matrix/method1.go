package main

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	n, l, r := len(matrix[0]), 0, len(matrix[0])*len(matrix)-1
	for l <= r {
		mid := l + (r-l)>>1
		if matrix[mid/n][mid%n] == target {
			return true
		} else if matrix[mid/n][mid%n] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return false
}
