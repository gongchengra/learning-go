package main

func generate(numRows int) [][]int {
	res := [][]int{}
	for i := 0; i < numRows; i++ {
		row := []int{}
		for j := 0; j < i+1; j++ {
			if j == 0 || j == i {
				row = append(row, 1)
			} else if i > 1 {
				row = append(row, res[i-1][j-1]+res[i-1][j])
			}
		}
		res = append(res, row)
	}
	return res
}
