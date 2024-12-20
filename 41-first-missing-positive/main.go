package main

func firstMissingPositive(nums []int) int {
	for i := range nums {
		k := nums[i]
		for k >= 1 && k <= len(nums) && k != nums[k-1] {
			k, nums[k-1] = nums[k-1], k
		}
	}

	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return len(nums) + 1
}

func main() {
	firstMissingPositive([]int{2, 1})
}
