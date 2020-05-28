package main

import "fmt"

//first attempt, dfs, TLE
func search(s string, i int, dict map[string]bool) bool {
	if i == len(s) {
		return true
	}
	for j := len(s); j > i; j-- {
		if dict[s[i:j]] {
			r := search(s, j, dict)
			if r {
				return true
			}
		}
	}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	dict := make(map[string]bool)
	for _, w := range wordDict {
		dict[w] = true
	}
	return search(s, 0, dict)
}

func main() {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	dict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Println(wordBreak(s, dict))
}

// #dynamic_programmig
// #动态规划
