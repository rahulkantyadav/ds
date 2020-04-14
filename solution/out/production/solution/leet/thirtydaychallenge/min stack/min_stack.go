package main

import "fmt"

type Item struct {
	value int
	min   int
	prev  *Item
}

type MinStack struct {
	top  *Item
	size int
}

func Constructor() MinStack {
	return MinStack{top: nil, size: 0}
}

func (this *MinStack) Push(x int) {
	minVal := x
	if this.top != nil && this.top.min >= x {
		minVal = x
	} else if this.top != nil {
		minVal = this.top.min
	}

	prev := this.top
	newItem := Item{x, minVal, prev}
	this.top = &newItem
	this.size++
}

func (this *MinStack) Pop() {
	if this.size == 0 {
		return
	}
	prev := this.top.prev
	this.top = prev
	this.size--
}

func (this *MinStack) Top() int {
	if this.size == 0 {
		return 0
	}
	return this.top.value
}

func (this *MinStack) GetMin() int {
	if this.size == 0 {
		return 0
	}
	return this.top.min
}

func main() {
	obj := Constructor()
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-1)
	obj.Pop()

	param_4 := obj.GetMin()

	fmt.Println(param_4)
}
