package main

import (
	"fmt"
	"strconv"
	"strings"
)

func splitIntoFibonacci(S string) []int {
	n := len(S)
	for i := 0; i <= n/2; i++ {
		if S[0] == '0' && i > 0 {
			return nil
		}
		attempt := make([]int, 0)
		a0, e := strconv.ParseInt(S[0:i+1], 10, 32)
		if e != nil {
			break
		}
		attempt = append(attempt, int(a0))
		for j := i + 1; j-i <= n/2; j++ {
			if S[i+1] == '0' && j > i+1 {
				break
			}
			b, e := strconv.ParseInt(S[i+1:j+1], 10, 32)
			if e != nil {
				break
			}
			attempt = append(attempt[0:1], int(b))
			k := j + 1
			a := a0
			for k < n {
				c := a + b
				if c > 0x7FFFFFFF {
					break
				}
				cs := strconv.FormatInt(c, 10)
				if strings.Index(S[k:], cs) != 0 {
					break
				}
				attempt = append(attempt, int(c))
				a = b
				b = c
				k += len(cs)
			}
			if k == n && len(attempt) > 2 {
				return attempt
			}
		}
	}
	return nil
}

func main() {
	fmt.Println(splitIntoFibonacci("1101111"))
}
