package main

import (
	"fmt"
	"sort"
)

//binary search with more conise sorting api
func findRightInterval(intervals [][]int) (result []int) {
	n := len(intervals)
	sorted := make([]int, n)
	for i := range intervals {
		sorted[i] = i
	}
	sort.Slice(sorted, func(i, j int) bool {
		return intervals[sorted[i]][0] < intervals[sorted[j]][0]
	})
	result = make([]int, n)
	for i, interval := range intervals {
		right := sort.Search(n, func(i int) bool {
			return intervals[sorted[i]][0] >= interval[1]
		})
		if right == n {
			result[i] = -1
		} else {
			result[i] = sorted[right]
		}
	}
	return
}

func main() {
	fmt.Println(findRightInterval([][]int{{1, 4}, {2, 3}, {3, 4}}))
}

// #区间 #排序 #二分
// #interval #sort #binary_search
