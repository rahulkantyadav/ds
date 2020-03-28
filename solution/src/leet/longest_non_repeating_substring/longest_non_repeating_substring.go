package main

import "fmt"

func main() {
	//fmt.Println("", lengthOfLongestSubstring("pwwkew"))
	fmt.Println("", lengthOfLongestSubstring("au"))
}

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	crLen := 0
	var arr [255]int
	var starIndex = 0

	for pos, _ := range arr {
		arr[pos] = -2
	}

	for pos, c := range s {
		if arr[c] >= starIndex {
			starIndex = arr[c] + 1
		}
		crLen = pos - starIndex + 1
		if crLen > maxLen {
			maxLen = crLen
		}
		arr[c] = pos
	}
	return maxLen
}
