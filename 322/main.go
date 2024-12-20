package main

func coinChange1(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = -1
	}
	for i := 0; i < amount; i++ {
		if dp[i] < 0 {
			continue
		}
		for _, c := range coins {
			if i+c <= amount && (dp[i+c] < 0 || dp[i]+1 < dp[i+c]) {
				dp[i+c] = dp[i] + 1
			}
		}
	}
	return dp[amount]
}

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 1; i < len(dp); i++ {
		dp[i] = amount + 1
	}
	for i := 1; i <= amount; i++ {
		for j := 0; j < len(coins); j++ {
			if coins[j] <= i {
				dp[i] = min(dp[i], dp[i-coins[j]]+1)
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func main() {

}
