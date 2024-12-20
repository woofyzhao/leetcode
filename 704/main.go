package main

import "fmt"

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1
	for low <= high {
		mid := (low + high) / 2
		if target < nums[mid] {
			high = mid - 1
		} else if target > nums[mid] {
			low = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func search2(nums []int, target int) int {
	a, b := 0, len(nums)-1
	for a <= b {
		m := (a + b) >> 1
		if nums[m] == target {
			return m
		}
		if nums[m] > target {
			b = m - 1
		} else {
			a = m + 1
		}
	}
	return -1
}

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(nums, 9))
	fmt.Println(search(nums, -1))
	fmt.Println(search(nums, 12))
	fmt.Println(search(nums, -12))
	fmt.Println(search(nums, 99))
	nums = []int{-1}
	fmt.Println(search(nums, -1))
	fmt.Println(search(nums, -10))
	fmt.Println(search(nums, 0))
	nums = []int{5, 6}
	fmt.Println(search(nums, 5))
	fmt.Println(search(nums, 6))
	fmt.Println(search(nums, 4))
	fmt.Println(search(nums, 7))
}

// #二分查找
// #binary_search
