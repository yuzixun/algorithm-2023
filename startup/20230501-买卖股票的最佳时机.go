package main

func maxProfit(prices []int) int {
	minPrice := 1 << 31
	maxProfitValue := 0

	for i := 0; i < len(prices); i++ {
		minPrice = min(minPrice, prices[i])
		maxProfitValue = max(maxProfitValue, prices[i]-minPrice)
	}

	return maxProfitValue
}
