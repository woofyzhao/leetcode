package main

import (
	"container/heap"
)

type Item struct {
	value      int
	arrayIndex int
	elemIndex  int
}
type Heap []*Item

func (h Heap) Len() int           { return len(h) }
func (h Heap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h Heap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *Heap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Item))
}

func (h *Heap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func smallestRange(nums [][]int) []int {
	k := len(nums)
	hp := &Heap{}
	currMax := -100001
	minLen := 200001
	minStart := 0
	minEnd := 0
	for i := 0; i < k; i++ {
		item := &Item{
			value:      nums[i][0],
			arrayIndex: i,
			elemIndex:  0,
		}
		heap.Push(hp, item)
		if item.value > currMax {
			currMax = item.value
		}
	}
	for {
		top := heap.Pop(hp).(*Item)
		r := currMax - top.value + 1
		if r < minLen {
			minLen = r
			minStart = top.value
			minEnd = currMax
		}
		if top.elemIndex == len(nums[top.arrayIndex])-1 {
			break
		}
		next := &Item{
			value:      nums[top.arrayIndex][top.elemIndex+1],
			arrayIndex: top.arrayIndex,
			elemIndex:  top.elemIndex + 1,
		}
		heap.Push(hp, next)
		if next.value > currMax {
			currMax = next.value
		}
	}
	return []int{minStart, minEnd}
}
