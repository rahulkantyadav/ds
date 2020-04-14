package main

import (
	"fmt"
	"strings"
)

func main() {

	//var arr = [][]int{{0, 1}, {1, 2}}
	var arr = [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}}
	//var arr = [][]int{{1, 4}, {0, 5}, {0, 4}, {1, 1}, {1, 5}}

	fmt.Println(stringShift("abcdefg", arr))
	//fmt.Println(stringShift("mecsk", arr))
}

func stringShift(s string, shift [][]int) string {
	count := 0
	for i := range shift {
		if shift[i][0] == 1 {
			count -= shift[i][1]
		} else {
			count += shift[i][1]
		}
	}
	count %= len(s)
	if count < 0 {
		count = len(s) + count
	}

	index := 0

	var str strings.Builder

	for index < len(s) {
		if count == len(s) {
			count = 0
		}
		str.WriteString(string(s[count]))
		count++
		index++
	}
	return str.String()
}
