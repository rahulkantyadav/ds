package main

import "fmt"

func main() {
	fmt.Println(checkValidString("(*))"))
}

func checkValidString(s string) bool {
	leftP := 0
	rightP := 0

	var str string

	for i := 0; i < len(s); i++ {
		str = string(s[i])
		if str == ")" {
			rightP++
		} else {
			leftP++
		}
		if rightP > leftP {
			return false
		}
	}

	leftP = 0
	rightP = 0
	for i := len(s) - 1; i > -1; i-- {
		str = string(s[i])
		if str == "(" {
			leftP++
		} else {
			rightP++
		}
		if leftP > rightP {
			return false
		}
	}

	return true
}
