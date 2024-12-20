package main

import (
	"fmt"
	"math"
)

func minWindow(s string, t string) (res string) {
	cnt := make(map[uint8]int)
	for i := range t {
		cnt[t[i]] += 1
	}
	total, start, minWin := len(t), 0, math.MaxInt

	for end := range s {
		if n, ok := cnt[s[end]]; ok {
			if n > 0 {
				total--
			}
			cnt[s[end]]--
		}

		for total == 0 {
			if end-start+1 < minWin {
				minWin = end - start + 1
				res = s[start : end+1]
			}
			if _, ok := cnt[s[start]]; ok {
				cnt[s[start]]++
				if cnt[s[start]] > 0 {
					total++
				}
			}
			start++
		}
	}

	return
}

func main() {
	fmt.Println(minWindow("cabefgecdaecf", "cae"))
	fmt.Println(minWindow("ADOBECODEBANC", "ABC"))
	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))
}
