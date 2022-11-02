package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}

/*
	func twoSum(nums []int, target int) []int {
		ret := []int{0, 0}
		for k, v := range nums {
			for i := k + 1; i < len(nums); i++ {
				if v+nums[i] == target {
					ret[0], ret[1] = k, i
				}
			}
		}
		return ret
	}
*/
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if j, ok := m[target-v]; ok {
			return []int{j, i}
		}
		m[v] = i
	}
	return nil
}
