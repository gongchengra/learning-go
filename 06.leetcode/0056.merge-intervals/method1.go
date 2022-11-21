package main

func merge(intervals [][]int) [][]int {
	s := make([]bool, 10000)
	for _, v := range intervals {
		for i := v[0]; i < v[1]+1; i++ {
			s[i] = true
		}
	}
	start, end := 0, 0
	res := [][]int{}
	for i := 1; i < 10000; i++ {
		if s[i-1] == false && s[i] == true {
			start = i
		}
		if s[i-1] == true && s[i] == false {
			end = i - 1
			res = append(res, []int{start, end})
		}
	}
	return res
}
