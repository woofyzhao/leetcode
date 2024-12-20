package main

func rob(nums []int) (res int) {
	dp := make([]int, len(nums))
	res, dp[0] = nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = nums[i]
		var a, b int
		if i-2 >= 0 {
			a = dp[i-2]
		}
		if i-3 >= 0 {
			b = dp[i-3]
		}
		dp[i] = max(a, b) + dp[i]
		res = max(res, dp[i])
	}
	return
}

func main() {

}
