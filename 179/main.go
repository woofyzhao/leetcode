package main

import (
	"slices"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}
	slices.SortFunc(strs, func(a, b string) int {
		if a+b > b+a {
			return -1
		}
		return 1
	})
	if strs[0] == "0" {
		return "0"
	}
	return strings.Join(strs, "")
}
