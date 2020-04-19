package main

func main() {

}

func squareSortedArray(arr []int) {
	startIndex := -1
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 0 {
			startIndex = i
			break
		}
	}

	if startIndex == -1 {
		for i := 0; i < len(arr); i++ {
			arr[i] *= arr[i]
		}
	} else {
		j := startIndex
		for i := 0; i < startIndex; i++ {
			temp := arr[i]
			temp *= temp
			arr[i] = arr[j] * arr[j]
			arr[j] = arr[temp]
			j++
		}
	}

}
