package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	if len(s) < 4 || len(s) > 12 {
		return nil
	}

	isIPByte := func(s string) bool {
		n, err := strconv.Atoi(s)
		if err != nil {
			return false
		}
		if s[0] == '0' && len(s) > 1 {
			return false
		}
		return n <= 255
	}

	var restore func(string, int) []string
	restore = func(remain string, parts int) (res []string) {
		if parts == 1 {
			if isIPByte(remain) {
				return []string{remain}
			} else {
				return nil
			}
		}

		n := min(len(remain), 3)
		for k := 1; k <= n; k++ {
			if isIPByte(remain[0:k]) {
				left := restore(remain[k:], parts-1)
				if left == nil {
					continue
				}
				for _, s := range left {
					res = append(res, remain[0:k]+"."+s)
				}
			}
		}
		return res
	}

	return restore(s, 4)
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}
