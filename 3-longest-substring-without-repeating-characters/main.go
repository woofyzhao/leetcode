package main

// 003
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}

	left, right, maxLen := 0, 0, 0
	pos := make(map[byte]int)
	for right < len(s) {
		p, ok := pos[s[right]]
		if ok && p >= left {
			left = p + 1
		}
		pos[s[right]] = right
		maxLen = max(maxLen, right-left+1)
		right += 1
	}
	return maxLen
}

//

func main() {
	println(lengthOfLongestSubstring(""))
	println(lengthOfLongestSubstring("a"))
	println(lengthOfLongestSubstring("aaa"))
	println(lengthOfLongestSubstring("abcabcbb"))
	println(lengthOfLongestSubstring("fddf"))
	println(lengthOfLongestSubstring("dvdf"))
}
