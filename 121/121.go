package main

func maxProfit(prices []int) int {
	if len(prices) < 2  {
		return 0
	}
	low := prices[0]
	profit := 0
	for i := 1; i < len(prices); i++ {
		tmp := prices[i] - low
		if tmp > profit {
			profit = tmp
		} else if tmp < 0 {
			low = prices[i]
		}
	}
	return profit
}
