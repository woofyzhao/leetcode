package main

func searchRange(nums []int, target int) (res []int) {
	res = []int{-1, -1}
	if len(nums) == 0 {
		return
	}

	a, b := 0, len(nums)-1
	for a <= b {
		m := (a + b) / 2
		if nums[m] >= target {
			b = m - 1
			if nums[m] == target {
				res[0] = m
			}
		} else {
			a = m + 1
		}
	}

	a, b = 0, len(nums)-1
	for a <= b {
		m := (a + b) / 2
		if nums[m] <= target {
			a = m + 1
			if nums[m] == target {
				res[1] = m
			}
		} else {
			b = m - 1
		}
	}
	return
}

func main() {

}
