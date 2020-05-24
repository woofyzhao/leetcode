package main

// no division, O(n), no extra space
func productExceptSelf(nums []int) (result []int) {
	result = make([]int, len(nums))

	//prefix
	p := 1
	for i, v := range nums {
		result[i] = p * v
		p = result[i]
	}

	//suffix
	p = 1
	for i := len(nums) - 1; i >= 0; i-- {
		if i == 0 {
			result[i] = p
		} else {
			result[i] = result[i-1] * p
		}
		p *= nums[i]
	}
	return
}

func main() {

}

// #数组 #数学运算
// #array #math #arithmetics
