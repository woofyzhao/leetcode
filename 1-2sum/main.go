package main

func twoSum0(nums []int, target int) []int {
	lookup := make(map[int]map[int]bool)
	for i, n := range nums {
		set, ok := lookup[n]
		if !ok {
			set = make(map[int]bool)
			lookup[n] = set
		}
		set[i] = true
	}
	for i, n := range nums {
		set := lookup[target-n]
		for k := range set {
			if k != i {
				return []int{i, k}
			}
		}
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		another := target - nums[i]
		if p, ok := m[another]; ok {
			return []int{p, i}
		}
		m[nums[i]] = i
	}
	return nil
}
