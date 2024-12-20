package main

func subsets(nums []int) (res [][]int) {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	res = subsets(nums[1:])
	n := len(res)
	for i := 0; i < n; i++ {
		res = append(res, append(res[i], nums[0]))
	}
	return
}

func main() {

}
