package main

import "fmt"

func main() {
	fmt.Println(insert([][]int{[]int{1, 3}, []int{6, 9}}, []int{2, 5}))
	fmt.Println(insert([][]int{[]int{1, 2}, []int{3, 5}, []int{6, 7}, []int{8, 10}, []int{12, 16}}, []int{4, 8}))
}

type Interval struct {
	Start int
	End   int
}

func insert(intervals [][]int, newInterval []int) [][]int {
	ilist := []Interval{}
	for _, v := range intervals {
		tmp := Interval{v[0], v[1]}
		ilist = append(ilist, tmp)
	}
	newItv := Interval{newInterval[0], newInterval[1]}
	resList := insertIntervals(ilist, newItv)
	res := [][]int{}
	for _, v := range resList {
		tmp := []int{v.Start, v.End}
		res = append(res, tmp)
	}
	return res
}
