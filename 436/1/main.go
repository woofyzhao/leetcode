package main

import (
	"fmt"
	"sort"
)

type Item struct {
	interval []int
	index    int
}

type Intervals []*Item

func (s Intervals) Len() int {
	return len(s)
}

func (s Intervals) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

type byBegin struct {
	Intervals
}

func (s byBegin) Less(i, j int) bool {
	return s.Intervals[i].interval[0] < s.Intervals[j].interval[0]
}

type byEnd struct {
	Intervals
}

func (s byEnd) Less(i, j int) bool {
	return s.Intervals[i].interval[1] < s.Intervals[j].interval[1]
}

//attempt 1: two sorts + sliding window
func findRightInterval(intervals [][]int) (result []int) {
	result = make([]int, len(intervals))
	for i := range result {
		result[i] = -1
	}

	sortByBegin := make(Intervals, 0)
	sortByEnd := make(Intervals, 0)
	for i, interval := range intervals {
		item := &Item{
			interval: interval,
			index:    i,
		}
		sortByBegin = append(sortByBegin, item)
		sortByEnd = append(sortByEnd, item)
	}
	sort.Sort(byBegin{sortByBegin})
	sort.Sort(byEnd{sortByEnd})

	front := 0
	for _, item := range sortByBegin {
		for sortByEnd[front].interval[1] <= item.interval[0] {
			result[sortByEnd[front].index] = item.index
			front += 1
		}
	}
	return
}

func main() {
	fmt.Println(findRightInterval([][]int{{1, 4}, {2, 3}, {3, 4}}))
}

// #区间 #排序 #二分
// #interval #sort #binary_search
