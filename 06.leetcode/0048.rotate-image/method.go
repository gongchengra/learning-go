package main

import "fmt"

func rotate(m [][]int) {
	n := len(m)
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			fmt.Print(i, j)
			temp := m[i][j]
			fmt.Println(temp)
			// 左边的行 等于 右边的列
			m[i][j] = m[n-j-1][i]
			fmt.Println(m[i][j])
			m[n-j-1][i] = m[n-i-1][n-j-1]
			fmt.Println(m[n-j-1][i])
			m[n-i-1][n-j-1] = m[j][n-i-1]
			fmt.Println(m[n-i-1][n-j-1])
			m[j][n-i-1] = temp
			fmt.Println(m[j][n-i-1])
		}
	}
}
