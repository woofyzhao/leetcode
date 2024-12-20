package main

func findPeakElement2(nums []int) int {
	a, b := 0, len(nums)-1
	for a <= b {
		m := (a + b) / 2
		if (m == 0 || nums[m] > nums[m-1]) && (m == len(nums)-1 || nums[m] < nums[m+1]) {
			return m
		}
		if nums[m] < nums[m+1] {
			a = m + 1
		} else {
			b = m - 1
		}
	}
	return -1
}
