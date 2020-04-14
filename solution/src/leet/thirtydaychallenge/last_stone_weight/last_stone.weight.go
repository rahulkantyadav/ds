package main

import "fmt"

func LastStoneWeight() int {
	//arr := []int{2, 7, 4, 1, 8, 1}
	arr := []int{10, 10, 7, 2}
	ans := lastStoneWeight(arr)
	return ans
}

func lastStoneWeight(stones []int) int {

	h := new(heap)
	h.arr = make([]int, len(stones)+1)
	h.crIndex = 0
	h.capcacity = len(stones) + 1

	for _, stone := range stones {
		h.add(stone)
	}

	for h.crIndex >= 1 {
		num1 := h.getTop()
		num2 := h.getTop()
		if num2 == -1 {
			return num1
		}
		diff := num1 - num2
		if diff != 0 {
			h.add(diff)
		}
	}
	return 0
}

type heap struct {
	arr       []int
	capcacity int
	crIndex   int
}

func (h *heap) swim(index int) {
	arr := h.arr
	for index > 1 && arr[index] > arr[index/2] {
		temp := arr[index]
		arr[index] = arr[index/2]
		arr[index/2] = temp
		index = index / 2
	}
}

func (h *heap) sink(index int) {
	arr := h.arr
	for 2*index <= h.crIndex {
		swapIndex := 2 * index
		if swapIndex < h.crIndex && arr[swapIndex] < arr[swapIndex+1] {
			swapIndex++
		}
		if arr[swapIndex] > arr[index] {
			swapVal := arr[swapIndex]
			arr[swapIndex] = arr[index]
			arr[index] = swapVal
			index = swapIndex
		} else {
			break
		}

	}
}

func (h *heap) add(num int) {
	h.crIndex++
	h.arr[h.crIndex] = num
	if h.crIndex > 1 {
		h.swim(h.crIndex)
	}
}

func (h *heap) getTop() int {
	if h.crIndex == 0 {
		return -1
	}
	ans := h.arr[1]
	h.arr[1] = h.arr[h.crIndex]
	h.arr[h.crIndex] = -1
	h.crIndex--
	h.sink(1)
	return ans
}

func main() {
	fmt.Println(LastStoneWeight())
}
