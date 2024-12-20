package main

import "fmt"

func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		}
		if target > nums[mid] {
			if nums[mid] < nums[right] && target > nums[right] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] > nums[right] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{1, 3}, 3))
}
