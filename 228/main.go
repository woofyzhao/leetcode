package main

import (
	"fmt"
	"strconv"
)

func makeRange(start, end int) string {
	if start == end {
		return strconv.Itoa(start)
	} else {
		return fmt.Sprintf("%d->%d", start, end)
	}
}

func summaryRanges(nums []int) (result []string) {
	start := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1]+1 {
			result = append(result, makeRange(start, nums[i-1]))
			start = nums[i]
		}
	}
	result = append(result, makeRange(start, nums[len(nums)-1]))
	return
}

func main() {

}
