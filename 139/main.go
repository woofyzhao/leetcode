package main

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, w := range wordDict {
		dict[w] = true
	}
	breakable := make(map[string]int)
	return tryBreak(s, dict, breakable)
}

func tryBreak(s string, dict map[string]bool, breakable map[string]int) bool {
	if len(s) == 0 || breakable[s] == 1 {
		return true
	}
	if breakable[s] == -1 {
		return false
	}
	for i := 1; i <= len(s); i++ {
		if !dict[s[:i]] {
			continue
		}
		if tryBreak(s[i:], dict, breakable) {
			breakable[s] = 1
			return true
		}
	}
	breakable[s] = -1
	return false
}
