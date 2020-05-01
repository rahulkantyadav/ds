package main

import "fmt"

func main() {
	arr := []int{3, 2, 1, 0, 4}
	//arr := []int{2, 0, 0}
	fmt.Println(canJump(arr))
}

func canJump(nums []int) bool {
	crindex := len(nums) - 1
	if crindex == 0 {
		return true
	}

	zeroIndex := -1

	for crindex > -1 {
		if nums[crindex] == 0 {
			zeroIndex = crindex
			crindex--
			for i := crindex; i >= 0; i-- {
				if nums[i] != 0 {
					crindex = i
					break
				}
			}
			found := false
			for i := crindex; i >= 0; i-- {
				if nums[i] == 0 {
					crindex = i
					break
				} else {
						1 > 3 - 2
						1 > 1
					if nums[i] >= (zeroIndex - i) {
						if zeroIndex == len(nums)-1 {
							found = true
							break
						} else {
							if nums[i] > (zeroIndex - i) {
								found = true
								break
							}
						}

					}
				}
			}
			if found == false {
				return false
			}
		}
		crindex--
	}
	return true
}
