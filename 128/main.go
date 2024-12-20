package main

func longestConsecutive(nums []int) (res int) {
	cnt := make(map[int]int) // 0 means not present
	for _, n := range nums {
		if cnt[n] > 0 {
			continue
		}
		cnt[n] = 1
		left := n - cnt[n-1]
		right := n + cnt[n+1]
		sum := 1 + cnt[n-1] + cnt[n+1]
		cnt[left], cnt[right] = sum, sum
		res = max(res, sum)
		//fmt.Println("n = {}, left = {}, right = {}, sum = {}, max = {}", n, left, right, sum, res)
	}
	return
}

func main() {
	longestConsecutive([]int{4, 0, -4, -2, 2, 5, 2, 0, -8, -8, -8, -8, -1, 7, 4, 5, 5, -4, 6, 6, -3})
}
