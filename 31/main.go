package main

import "slices"

func nextPermutation(nums []int) {
	if len(nums) < 2 {
		return
	}

	i := len(nums) - 1
	for i > 0 && nums[i] <= nums[i-1] {
		i--
	}
	if i == 0 {
		slices.Reverse(nums)
		return
	}

	j := len(nums) - 1
	for nums[j] <= nums[i-1] {
		j--
	}
	nums[i-1], nums[j] = nums[j], nums[i-1]
	slices.Reverse(nums[i:])
	return
}

func main() {

}
