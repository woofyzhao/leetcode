package main

import (
	"fmt"
	"sort"
)

// two pointers, O(n^2)
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		// c 对应的指针初始指向数组的最右端
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second >= third {
				break
			}
			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return ans
}

//func twoSum(nums []int, target int) (res [][]int) {
//	n := len(nums)
//	if n < 2 {
//		return nil
//	}
//
//	if nums[0]+nums[1] > target {
//		return nil
//	}
//
//	if nums[n-1]+nums[n-2] < target {
//		return nil
//	}
//
//	left := 0
//	right := n - 1
//	for left < right {
//		sum := nums[left] + nums[right]
//		if sum < target {
//			left += 1
//		} else if sum > target {
//			right -= 1
//		} else {
//			res = append(res, []int{nums[left], nums[right]})
//			left += 1
//			right -= 1
//		}
//	}
//	return
//}
//
//type Thriple struct {
//	a, b, c int
//}
//
//func threeSum(nums []int) (res [][]int) {
//	sort.Ints(nums)
//	seen := make(map[Thriple]bool)
//	for i := range nums {
//		target := -nums[i]
//		pairs := twoSum(nums[i+1:], target)
//		for _, pair := range pairs {
//			entry := []int{nums[i], pair[0], pair[1]}
//			tp := Thriple{a: nums[i], b: pair[0], c: pair[1]}
//			if !seen[tp] {
//				res = append(res, entry)
//				seen[tp] = true
//			}
//		}
//	}
//	return
//}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
	//a := []int{1-2sum, 2, 3-longest-substring-without-repeating-characters}
	//b := []int{1-2sum, 2, 3-longest-substring-without-repeating-characters}
}
