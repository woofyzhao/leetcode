package main

/*
func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	dp := make([][]int, len(text1))
	init := 0
	for i := 0; i < len(text1); i++ {
		dp[i] = make([]int, len(text2))
		if text1[i] == text2[0] {
			init = 1
		}
		dp[i][0] = init
	}

	init = 0
	for j := 0; j < len(text2); j++ {
		if text2[j] == text1[0] {
			init = 1
		}
		dp[0][j] = init
	}

	for i := 1; i < len(text1); i++ {
		for j := 1; j < len(text2); j++ {
			dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			if text1[i] == text2[j] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	return dp[len(text1)-1][len(text2)-1]
}
*/

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := 1; i < len(text1)+1; i++ {
		for j := 1; j < len(text2)+1; j++ {
			dp[i][j] = max(dp[i][j-1], dp[i-1][j])
			if text1[i-1] == text2[j-1] {
				dp[i][j] = max(dp[i][j], dp[i-1][j-1]+1)
			}
		}
	}
	return dp[len(text1)][len(text2)]
}

func main() {

}
