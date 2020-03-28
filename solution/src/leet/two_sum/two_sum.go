package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{3, 3}, 6))
}

func twoSum(nums []int, target int) []int {
	var startIndex, endIndex int
	elementMap := make(map[int]int)

	for index, element := range nums {
		if val, exist := elementMap[target-element]; exist {
			startIndex = val
			endIndex = index
			break
		}
		elementMap[element] = index
	}
	ret := []int{startIndex, endIndex}
	return ret
}
