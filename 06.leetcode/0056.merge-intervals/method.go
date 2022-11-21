package main

import "sort"

type Interval struct {
	Start int
	End   int
}

type Ilist []Interval

func (il Ilist) Len() int { return len(il) }

func (il Ilist) Less(i int, j int) bool {
	if il[i].Start == il[j].Start {
		return il[i].End < il[j].End
	}
	return il[i].Start < il[j].Start
}

func (il Ilist) Swap(i int, j int) { il[i], il[j] = il[j], il[i] }

func merge(intervals [][]int) [][]int {
	ilist := []Interval{}
	for _, v := range intervals {
		tmp := Interval{v[0], v[1]}
		ilist = append(ilist, tmp)
	}
	resList := mergeIntervals(ilist)
	res := [][]int{}
	for _, v := range resList {
		tmp := []int{v.Start, v.End}
		res = append(res, tmp)
	}
	return res
}

func mergeIntervals(intervals []Interval) []Interval {
	if len(intervals) < 2 {
		return intervals
	}
	var il Ilist
	ret := []Interval{}
	il = intervals
	sort.Sort(il)
	tmp := il[0]
	for i := 1; i < len(il); i++ {
		if il[i].End < tmp.Start || il[i].Start > tmp.End {
			ret = append(ret, tmp)
			tmp = il[i]
		} else {
			if tmp.Start >= il[i].Start {
				tmp.Start = il[i].Start
			}
			if tmp.End <= il[i].End {
				tmp.End = il[i].End
			}
		}
	}
	ret = append(ret, tmp)
	return ret
}
