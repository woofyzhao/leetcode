package main

import "fmt"

//second attempt, dp
func wordBreak(s string, wordDict []string) bool {
	n := len(s)
	splitable := make([]bool, n)

	for i := range s {
		for _, w := range wordDict {
			k := i - len(w)
			if k < -1 {
				continue
			}
			if s[k+1:i+1] == w {
				splitable[i] = (k == -1 || splitable[k])
				if splitable[i] {
					break
				}
			}
		}
	}
	return splitable[n-1]
}

func main() {
	s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	dict := []string{"a", "aa", "aaa", "aaaa", "aaaaa", "aaaaaa", "aaaaaaa", "aaaaaaaa", "aaaaaaaaa", "aaaaaaaaaa"}
	fmt.Println(wordBreak(s, dict))

	fmt.Println(wordBreak("dogs", []string{"dog", "s", "gs"}))
}
