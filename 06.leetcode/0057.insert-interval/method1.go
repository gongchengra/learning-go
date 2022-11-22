package main

func insertIntervals(a []Interval, newInterval Interval) []Interval {
	res := []Interval{}
	if len(a) == 0 {
		res = append(res, newInterval)
		return res
	}
	cur := 0
	for cur < len(a) && a[cur].End < newInterval.Start {
		res = append(res, a[cur])
		cur++
	}
	for cur < len(a) && a[cur].Start <= newInterval.End {
		newInterval = Interval{Start: min(newInterval.Start, a[cur].Start), End: max(newInterval.End, a[cur].End)}
		cur++
	}
	res = append(res, newInterval)
	for cur < len(a) {
		res = append(res, a[cur])
		cur++
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
