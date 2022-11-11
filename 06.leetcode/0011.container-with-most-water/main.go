package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

/* method 1

func maxArea(height []int) int {
	max := min(height[0], height[1])
	for i := 0; i < len(height); i++ {
		for j := i + 1; j < len(height); j++ {
			tmp := min(height[i], height[j]) * (j - i)
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}
*/

func maxArea(height []int) (max int) {
	for i, j := 0, len(height)-1; i < j; {
		a, b := height[i], height[j]
		h := min(a, b)
		area := h * (j - i)
		if max < area {
			max = area
		}
		if a < b {
			i++
		} else {
			j--
		}
	}
	return max
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
