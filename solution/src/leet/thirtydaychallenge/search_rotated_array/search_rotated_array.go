package main

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 1 && nums[0] != target {
		return -1
	}
	return bSearch(nums, target, 0, len(nums)-1)
}

func bSearch(nums []int, target int, start int, end int) int {
	if start > end {
		return -1
	}

	mid := (start + end) / 2

	if nums[mid] == target {
		return mid
	}
	// if 1st half is sorted
	if nums[start] <= nums[mid] {
		if target >= nums[start] && target <= nums[mid] {
			return bSearch(nums, target, start, mid-1)
		}
		return bSearch(nums, target, mid+1, end)

	}

	// 2nd hafl is sorted
	if target >= nums[mid] && target <= nums[end] {
		return bSearch(nums, target, mid+1, end)
	}
	return bSearch(nums, target, start, mid-1)

}

func main() {
	//nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{3, 1}
	fmt.Println(search(nums, 1))
}
