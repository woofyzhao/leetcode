package main

import "math"

// https://www.youtube.com/watch?v=LPFhl65R7ww
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1)
	}

	x, y := len(nums1), len(nums2)
	low, high := 0, x
	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (x+y+1)/2 - partitionX

		maxLeftX := math.MinInt
		if partitionX > 0 {
			maxLeftX = nums1[partitionX-1]
		}
		minRightX := math.MaxInt
		if partitionX < x {
			minRightX = nums1[partitionX]
		}

		maxLeftY := math.MinInt
		if partitionY > 0 {
			maxLeftY = nums2[partitionY-1]
		}
		minRightY := math.MaxInt
		if partitionY < y {
			minRightY = nums2[partitionY]
		}

		if maxLeftX <= minRightY && maxLeftY <= minRightX {
			if (x+y)%2 == 0 {
				return float64(max(maxLeftX, maxLeftY)+min(minRightX, minRightY)) / 2
			} else {
				return float64(max(maxLeftX, maxLeftY))
			}
		}
		if maxLeftX > minRightY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}

	return -1.0
}
