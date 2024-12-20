package main

func maxProfit(prices []int) (res int) {
	if len(prices) < 2 {
		return
	}
	dp := make([]int, len(prices))
	for i := 1; i < len(prices); i++ {
		delta := prices[i] - prices[i-1]
		if dp[i-1]+delta > 0 {
			dp[i] = dp[i-1] + delta
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return
}

func maxProfit1(prices []int) int {
	if len(prices) < 1 {
		return 0
	}
	min, maxProfit := prices[0], 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
		if prices[i] < min {
			min = prices[i]
		}
	}
	return maxProfit
}

func main() {

}
