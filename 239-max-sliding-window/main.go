package main

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return nil
	}
	deq := make([]int, 0, k)
	res := make([]int, 0, len(nums)-k+1)

	for i, v := range nums {
		if i >= k && deq[0] <= i-k {
			deq = deq[1:]
		}
		for len(deq) > 0 && v >= nums[deq[len(deq)-1]] {
			deq = deq[:len(deq)-1]
		}
		deq = append(deq, i)
		if i >= k-1 {
			res = append(res, nums[deq[0]])
		}
	}

	return res
}

func main() {

}
