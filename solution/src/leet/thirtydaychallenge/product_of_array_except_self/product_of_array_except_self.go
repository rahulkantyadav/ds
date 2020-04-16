package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(productExceptSelf(arr))
}

func productExceptSelf(nums []int) []int {
	getProduct(nums, 1, 0)
	return nums
}

func getProduct(nums []int, previousProduct int, index int) int {
	if index >= len(nums) {
		return 1
	}
	temp := nums[index]
	nextProduct := getProduct(nums, previousProduct*temp, index+1)
	nums[index] = nextProduct * previousProduct
	return nextProduct * temp
}
