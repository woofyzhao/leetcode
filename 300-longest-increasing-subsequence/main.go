package main

import "sort"

func lengthOfLIS1(nums []int) (res int) {
	if len(nums) < 1 {
		return
	}

	dp := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for k := 0; k < i; k++ {
			if nums[k] < nums[i] {
				if dp[k]+1 > dp[i] {
					dp[i] = dp[k] + 1
				}
			}
		}
		if dp[i] > res {
			res = dp[i]
		}
	}
	return
}

// 解法二 O(n log n) 贪心, DP[i] 表示长度i+1的上升子序列中结尾最小的数
func lengthOfLIS(nums []int) int {
	dp := []int{}
	for _, num := range nums {
		i := sort.SearchInts(dp, num)
		if i == len(dp) {
			dp = append(dp, num)
		} else {
			dp[i] = num
		}
	}
	return len(dp)
}

func main() {

}
