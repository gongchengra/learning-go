package main

func maxProfit(prices []int) int {
	max, tmp := 0, 0
	for i := 1; i < len(prices); i++ {
		tmp += prices[i] - prices[i-1]
		if tmp < 0 {
			tmp = 0
		}
		if max < tmp {
			max = tmp
		}
	}
	return max
}
