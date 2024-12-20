package main

type MinStack struct {
	elements []int
	minimals []int
}

func Constructor() MinStack {
	return MinStack{
		elements: []int{},
		minimals: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.elements = append(this.elements, val)
	if len(this.minimals) == 0 || val <= this.minimals[len(this.minimals)-1] {
		this.minimals = append(this.minimals, val)
	}
}

func (this *MinStack) Pop() {
	last := this.elements[len(this.elements)-1]
	this.elements = this.elements[0 : len(this.elements)-1]
	if last == this.minimals[len(this.minimals)-1] {
		this.minimals = this.minimals[0 : len(this.minimals)-1]
	}
}

func (this *MinStack) Top() int {
	return this.elements[len(this.elements)-1]
}

func (this *MinStack) GetMin() int {
	return this.minimals[len(this.minimals)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
