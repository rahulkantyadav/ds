package main

import "fmt"

//Incomplete solution:
//https://www.geeksforgeeks.org/largest-subarray-with-equal-number-of-0s-and-1s/

func main() {
	//arr := []int{0, 1, 0}
	arr := []int{0, 1, 0, 1}
	//arr := []int{0, 1, 1, 0, 0, 1, 1, 1, 0, 0}
	fmt.Println(findMaxLength(arr))

}

func findMaxLength(nums []int) int {
	for index, item := range nums {
		if item == 0 {
			item = -1
		}
		nums[index] = item
		if index != 0 {
			nums[index] = nums[index-1] + item
		}
	}

	ans := 0
	m := make(map[int]int)

	for index, item := range nums {
		if _, ok := m[item]; ok {
			crSize := (index - m[item])
			if crSize > ans {
				ans = crSize
			}
		} else {
			m[item] = index
		}
	}

	if ans == 0 {
		for index, item := range nums {
			if item == 0 {
				return index + 1
			}
		}
	}

	return ans
}
