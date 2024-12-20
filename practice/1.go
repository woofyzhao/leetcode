package main

func twoSum(nums []int, target int) []int {
	if len(nums) <= 1 {
		return nil
	}

	m := make(map[int]int)
	for i, n := range nums {
		if k, ok := m[target-n]; ok {
			return []int{k, i}
		}
		m[n] = i
	}
	return nil
}

func main() {

}
