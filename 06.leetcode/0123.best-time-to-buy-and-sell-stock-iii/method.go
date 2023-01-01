package main

func maxProfit(prices []int) int {
	size := len(prices)
	if size <= 1 {
		return 0
	}
	profits := []int{}
	tmp := 0
	for i := 1; i < size; i++ {
		diff := prices[i] - prices[i-1]
		if tmp*diff >= 0 {
			tmp += diff
			continue
		}
		profits = append(profits, tmp)
		tmp = diff
	}
	profits = append(profits, tmp)
	res := 0
	for i := 0; i < len(profits); i++ {
		tmp = max(profits[:i]) + max(profits[i:])
		if res < tmp {
			res = tmp
		}
	}
	return res
}

func max(ps []int) int {
	max, tmp := 0, 0
	for _, p := range ps {
		if tmp < 0 {
			tmp = 0
		}
		tmp += p
		if max < tmp {
			max = tmp
		}
	}
	return max
}
