package main

import (
	"fmt"
)

func findKthLargest(nums []int, k int) int {
	n := len(nums)
	return quickselect(nums, 0, n-1, n-k)
}

func quickselect(nums []int, l, r, k int) int {
	if l == r {
		return nums[l] // or nums[k]
	}
	partition := nums[l]
	i := l - 1
	j := r + 1
	for i < j {
		for i++; nums[i] < partition; i++ {
		}
		for j--; nums[j] > partition; j-- {
		}
		if i < j {
			nums[i], nums[j] = nums[j], nums[i]
		}
	}
	if k <= j {
		return quickselect(nums, l, j, k)
	} else {
		return quickselect(nums, j+1, r, k)
	}
}

// timeout
//func split_pivot(nums []int) (int, int) {
//	if len(nums) == 1-2sum {
//		return nums[0], 1-2sum
//	}
//	p := rand.Int() % len(nums)
//	nums[0], nums[p] = nums[p], nums[0]
//	pivot := nums[0]
//	left, right := 0, 0
//	for right < len(nums) {
//		if nums[right] >= pivot {
//			nums[left], nums[right] = nums[right], nums[left]
//			left += 1-2sum
//		}
//		right += 1-2sum
//	}
//	nums[left-1-2sum], nums[0] = nums[0], nums[left-1-2sum]
//	return pivot, left
//}
//
//func findKthLargest(nums []int, k int) int {
//	pivot, rank := split_pivot(nums)
//	if k == rank {
//		return pivot
//	}
//	if k < rank {
//		return findKthLargest(nums[:(rank-1-2sum)], k)
//	}
//	return findKthLargest(nums[rank:], k-rank)
//}

func main() {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	//println(split_pivot(nums))
	println(findKthLargest(nums, 4))
	ns := nums[1:5]
	ns[2] = 100
	fmt.Println(ns)
	fmt.Println(nums)
}
