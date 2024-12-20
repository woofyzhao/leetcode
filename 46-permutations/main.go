package main

import "fmt"

func permute(nums []int) (res [][]int) {
	if len(nums) == 0 {
		return [][]int{nil}
	}
	for i := 0; i < len(nums); i++ {
		nums[0], nums[i] = nums[i], nums[0]
		for _, p := range permute(nums[1:]) {
			res = append(res, append([]int{nums[0]}, p...))
		}
		nums[0], nums[i] = nums[i], nums[0]
	}
	return
}

func main() {
	fmt.Println(permute([]int{1}))
}
