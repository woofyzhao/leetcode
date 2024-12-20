package main

import "fmt"

func lengthOfLongestSubstring(s string) (res int) {
	if len(s) == 0 {
		return 0
	}
	m := make(map[byte]int)
	left, right := 0, 0
	for right < len(s) {
		if k, ok := m[s[right]]; ok && k >= left {
			left = k + 1
		}
		res = max(res, right-left+1)
		m[s[right]] = right
		right++
	}
	return
}

func main() {
	fmt.Println(lengthOfLongestSubstring("aaaabaccc"))
	fmt.Println(lengthOfLongestSubstring("a"))
	fmt.Println(lengthOfLongestSubstring("abcbacefgacceb"))
}
