package main

import (
	"fmt"
	"math"
)

func maxSubArray(nums []int) int {
	res := math.MinInt
	sum := 0
	for _, n := range nums {
		sum += n
		if sum > res {
			res = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return res
}

func main() {
	fmt.Println(maxSubArray([]int{}))
}
