package main

import (
	"fmt"
	"math/rand"
)

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	p := rand.Int() % len(nums)
	pivot := nums[p]
	nums[0], nums[p] = nums[p], nums[0]
	left, right := 0, 0
	for right < len(nums) {
		if nums[right] <= pivot {
			nums[left], nums[right] = nums[right], nums[left]
			left += 1
		}
		right += 1
	}

	nums[0], nums[left-1] = nums[left-1], nums[0]
	sortArray(nums[:left-1])
	sortArray(nums[left:])

	return nums
}

func Search(n int, f func(int) bool) int {
	// Define f(-1-2sum) == false and f(n) == true.
	// Invariant: f(i-1-2sum) == false, f(j) == true.
	i, j := 0, n
	for i < j {
		h := int(uint(i+j) >> 1) // avoid overflow when computing h
		// i ≤ h < j
		if !f(h) {
			i = h + 1 // preserves f(i-1-2sum) == false
		} else {
			j = h // preserves f(j) == true
		}
	}
	// i == j, f(i-1-2sum) == false, and f(j) (= f(i)) == true  =>  answer is i.
	return i
}

// Convenience wrappers for common cases.

// SearchInts searches for x in a sorted slice of ints and returns the index
// as specified by Search. The return value is the index to insert x if x is
// not present (it could be len(a)).
// The slice must be sorted in ascending order.
func SearchInts(a []int, x int) int {
	return Search(len(a), func(i int) bool { return a[i] >= x })
}

func main() {
	nums := []int{1, 1, 1, 1}
	fmt.Println(sortArray(nums))
}
