package main

import "fmt"

// first attempt, DP, AC
func numberOfArithmeticSlices_1(A []int) (result int) {

	n := len(A)

	//index -> delta -> next-index-list
	link := make([]map[int][]int, n)
	for i := 0; i < n; i++ {
		link[i] = make(map[int][]int)
		for j := i + 1; j < n; j++ {
			link[i][A[j]-A[i]] = append(link[i][A[j]-A[i]], j)
		}
	}

	//start_index -> delta -> count
	count := make([]map[int]int, n)
	for i := n - 1; i >= 0; i-- {
		count[i] = make(map[int]int)
		for d, list := range link[i] {
			for _, j := range list {
				count[i][d] += count[j][d] + len(link[j][d])
			}
			result += count[i][d]
		}
	}

	return
}

//A more efficient one inspired by: https://leetcode.com/problems/arithmetic-slices-ii-subsequence/discuss/92852/11-line-Python-O(n2)-solution
//uses less space and more concise.

func numberOfArithmeticSlices(A []int) (result int) {
	n := len(A)

	//index -> deta -> #seq with len >= 2
	count := make([]map[int]int, n)
	for i := range count {
		count[i] = make(map[int]int)
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			delta := A[j] - A[i]
			count[i][delta] += count[j][delta] + 1 //the extra 1-2sum is (i, j)
			result += count[j][delta]              //guaranteed to have len>=3-longest-substring-without-repeating-characters from i
		}
	}
	return
}

func main() {
	fmt.Println(numberOfArithmeticSlices([]int{1, 3, 2, 4, 5, 7, 6, 8, 9, 10, 1, 3, 5, 7, 9, 2, 4, 6, 8, 10}))
	fmt.Println(numberOfArithmeticSlices([]int{2, 3, 5, 7, 9, 10, 12, 14, 15}))
}

// #arithmetic #sequence #hard #dynamic_programming
// #序列 #等差数列 #较难 #动态规划
