package main

func insertIntervals(a []Interval, newInterval Interval) []Interval {
	la := len(a)
	if la == 0 {
		return []Interval{newInterval}
	}
	if newInterval.End < a[0].Start {
		return append([]Interval{newInterval}, a...)
	}
	if a[la-1].End < newInterval.Start {
		return append(a, newInterval)
	}
	res := make([]Interval, 0, la)
	for i := range a {
		if isOverlap(a[i], newInterval) {
			newInterval = merge(a[i], newInterval)
			if i == la-1 {
				res = append(res, newInterval)
			}
			continue
		}
		if a[i].End < newInterval.Start {
			res = append(res, a[i])
			continue
		}
		if newInterval.End <= a[i].Start {
			res = append(res, newInterval)
			res = append(res, a[i:]...)
			break
		}
	}
	return res
}

func isOverlap(a, b Interval) bool {
	return (a.Start <= b.Start && b.Start <= a.End) ||
		(a.Start <= b.End && b.End <= a.End) ||
		(b.Start <= a.Start && a.End <= b.End)
}

func merge(a, b Interval) Interval {
	start, end := a.Start, a.End
	if b.Start < start {
		start = b.Start
	}
	if b.End > end {
		end = b.End
	}
	return Interval{Start: start, End: end}
}
