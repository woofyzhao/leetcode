package main

type Level struct {
	prev int
	sum  int
}

// brute force 按层枚举
func trap1(height []int) (res int) {
	maxHeight := 0
	for _, h := range height {
		if h > maxHeight {
			maxHeight = h
		}
	}
	tower := make([]Level, maxHeight+1)
	for i, h := range height {
		for k := 1; k <= h; k++ {
			if tower[k].prev > 0 {
				tower[k].sum += i - tower[k].prev
			}
			tower[k].prev = i + 1
		}
	}
	for i := range tower {
		res += tower[i].sum
	}
	return
}

// 双指针
func trap(height []int) int {
	res, left, right, maxLeft, maxRight := 0, 0, len(height)-1, 0, 0
	for left <= right {
		if height[left] <= height[right] {
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				res += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				res += maxRight - height[right]
			}
			right--
		}
	}
	return res
}

func main() {

}
