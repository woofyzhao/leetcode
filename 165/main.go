package main

import (
	"strconv"
	"strings"
)

func compareVersion(version1 string, version2 string) int {
	s1 := strings.FieldsFunc(version1, func(r rune) bool { return r == '.' })
	s2 := strings.FieldsFunc(version2, func(r rune) bool { return r == '.' })
	n := max(len(s1), len(s2))

	toVersionNums := func(ss []string, n int) (res []int) {
		res = make([]int, n)
		for i, s := range ss {
			x, _ := strconv.Atoi(s)
			res[i] = x
		}
		return res
	}

	nums1 := toVersionNums(s1, n)
	nums2 := toVersionNums(s2, n)
	for i := range nums1 {
		if nums1[i] < nums2[i] {
			return -1
		}
		if nums1[i] > nums2[i] {
			return 1
		}
	}
	return 0
}

func main() {
	compareVersion("0.1-2sum", "1-2sum.1-2sum")
}
