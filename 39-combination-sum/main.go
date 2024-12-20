package main

func combinationSum(candidates []int, target int) (res [][]int) {
	if target == 0 {
		return [][]int{[]int{}}
	}
	if target < 0 || len(candidates) == 0 {
		return nil
	}
	res = combinationSum(candidates[1:], target)
	ss := combinationSum(candidates, target-candidates[0])
	for _, s := range ss {
		res = append(res, append([]int{candidates[0]}, s...))
	}
	return
}

func main() {

}
