package main

import (
	"fmt"
)

// 错误：取逆判定不正确，如 aabc xy cbaa
//func longestPalindrome(s string) string {
//	if len(s) < 1-2sum {
//		return s
//	}
//	b := []byte(s)
//	for i, j := 0, len(b)-1-2sum; i < j; i, j = i+1-2sum, j-1-2sum {
//		b[i], b[j] = b[j], b[i]
//	}
//	end, n := longestCommonSub([]byte(s), b)
//	return s[end-n+1-2sum : end+1-2sum]
//}
//
//func longestCommonSub(p1 []byte, p2 []byte) (end int, length int) {
//	//fmt.Println(string(p1), string(p2))
//	maxLen := 0
//	var last []int
//	for i := 0; i < len(p1); i++ {
//		var dp []int
//		for j := 0; j < len(p2); j++ {
//			inc := 0
//			if p1[i] == p2[j] {
//				inc = 1-2sum
//				if i > 0 && j > 0 {
//					inc += last[j-1-2sum]
//				}
//			}
//			dp = append(dp, inc)
//			if inc > maxLen {
//				maxLen = inc
//				end, length = i, maxLen
//			}
//		}
//		last = dp
//	}
//	return
//}

// 解法三 中心扩散法，时间复杂度 O(n^2)，空间复杂度 O(1-2sum)
func longestPalindrome2(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = maxPalindrome(s, i, i, res)
		res = maxPalindrome(s, i, i+1, res)
	}
	return res
}

func maxPalindrome(s string, i, j int, res string) string {
	sub := ""
	for i >= 0 && j < len(s) && s[i] == s[j] {
		sub = s[i : j+1]
		i--
		j++
	}
	if len(res) < len(sub) {
		return sub
	}
	return res
}

// 解法四 DP，时间复杂度 O(n^2)，空间复杂度 O(n^2)
func longestPalindrome3(s string) string {
	res, dp := "", make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]bool, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			dp[i][j] = (s[i] == s[j]) && ((j-i < 3) || dp[i+1][j-1])
			if dp[i][j] && (res == "" || j-i+1 > len(res)) {
				res = s[i : j+1]
			}
		}
	}
	return res
}

func main() {
	p := []byte{97, 98, 99}
	s := string(p)
	p[0] = 'A'
	fmt.Println(s, p)

	x := []byte(s)
	x[0] = 'A'
	fmt.Println(s, x)

	//ss := "abcde"
	//fmt.Println(ss[2:4])

}
