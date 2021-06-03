package main

import (
	"math"
)

// 最小栈
type MinStack struct {
	list []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		list: []int{math.MaxInt64},
	}
}

func (this *MinStack) Push(val int) {
	min := this.list[len(this.list)-1]
	if val < min {
		min = val
	}
	this.list = append(this.list, val, min)
}

func (this *MinStack) Pop() {
	this.list = append(this.list[:0], this.list[:len(this.list)-2]...)
}

func (this *MinStack) Top() int {
	return this.list[len(this.list)-2]
}

func (this *MinStack) GetMin() int {
	return this.list[len(this.list)-1]
}

func main() {
	data := Constructor()
	data.Push(6)
	data.Push(-2)
	data.Push(4)
	data.Push(-3)
	data.Top()
	data.Pop()
	data.Top()
	data.GetMin()
}
