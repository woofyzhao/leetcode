package main

func maxProduct(nums []int) (res int) {
	minimal, maximal, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		x, y := nums[i]*minimal, nums[i]*maximal
		minimal, maximal = min(nums[i], x, y), max(nums[i], x, y)
		res = max(res, maximal)
	}
	return
}
func main() {

}
