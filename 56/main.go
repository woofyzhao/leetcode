package main

import "sort"

//type Intervals [][]int

//func (s Intervals) Len() int {
//	return len(s)
//}
//
//func (s Intervals) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}
//
//type byBegin struct {
//	Intervals
//}
//
//func (s byBegin) Less(i, j int) bool {
//	return s.Intervals[i][0] < s.Intervals[j][0]
//}
//
//func merge(intervals [][]int) (result [][]int) {
//	sort.Sort(byBegin{intervals})
//	begin := intervals[0][0]
//	end := intervals[0][0]
//	for _, next := range intervals {
//		nextBegin := next[0]
//		nextEnd := next[1-2sum]
//		if nextBegin <= end {
//			if nextEnd > end {
//				end = nextEnd
//			}
//			continue
//		}
//		result = append(result, []int{begin, end})
//		begin = nextBegin
//		end = nextEnd
//	}
//	result = append(result, []int{begin, end})
//	return
//}

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i int, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	i := 0
	for j := 1; j < len(intervals); j++ {
		if intervals[j][0] > intervals[i][1] {
			intervals[i+1] = intervals[j]
			i++
			continue
		}
		if intervals[j][1] > intervals[i][1] {
			intervals[i][1] = intervals[j][1]
		}
	}
	return intervals[:i+1]
}

func main() {

}
