package main

const (
	MinValue = -100001
	MaxValue = +100001
)

//first attempt, sliding window, AC
func smallestRange_1(nums [][]int) []int {
	k := len(nums)
	minValue := MaxValue
	maxValue := MinValue

	//value to list of array ids
	group := make(map[int][]int)
	for id, array := range nums {
		for i, value := range array {
			if i > 0 && value == array[i-1] {
				continue
			}
			if value < minValue {
				minValue = value
			}
			if value > maxValue {
				maxValue = value
			}
			group[value] = append(group[value], id)
		}
	}

	//array id to value count
	count := make(map[int]int)

	minStart := minValue
	minEnd := maxValue + 1
	minLen := minEnd - minStart

	//two pointers
	start := minValue
	end := minValue

	for end <= maxValue+1 {
		//march end pointer
		if len(count) < k {
			ids := group[end]
			for _, id := range ids {
				count[id] += 1
			}
			end++
			continue
		}

		//already covered all, record the best result so far
		if end-start < minLen {
			minLen = end - start
			minStart = start
			minEnd = end
		}

		//then try squeezing from front
		ids := group[start]
		for _, id := range ids {
			count[id] -= 1
			if count[id] == 0 {
				delete(count, id)
			}
		}
		start++
	}
	return []int{minStart, minEnd - 1}
}

func main() {

}

// #堆 #滑动窗口
// #heap #sliding_windows #map
