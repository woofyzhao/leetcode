package main

// first attempt, not good
func findPeakElement_1(nums []int) int {
	n := len(nums)
	if n < 2 {
		return 0
	}
	if nums[0] > nums[1] {
		return 0
	}
	if nums[n-1] > nums[n-2] {
		return n - 1
	}
	low := 0
	high := n - 1
	for low <= high {
		mid := (low + high) / 2
		if mid > 0 && nums[mid-1] > nums[mid] {
			high = mid - 1
			continue
		}
		if mid < n-1 && nums[mid] < nums[mid+1] {
			low = mid + 1
			continue
		}
		return mid
	}
	return -1
}

func findPeakElement(nums []int) int {
	n := len(nums)
	low := 0
	high := n - 1
	for low < high {
		mid := (low + high) / 2
		if nums[mid+1] > nums[mid] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return low
}

func main() {

}

// #二分
// #binary_search
