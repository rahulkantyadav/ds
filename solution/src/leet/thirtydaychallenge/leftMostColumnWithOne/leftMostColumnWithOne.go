package main

import "fmt"

type BinaryMatrix struct {
	matrix [][]int
}

func (binaryMatrix BinaryMatrix) Get(row int, col int) int {
	return binaryMatrix.matrix[row][col]
}

func (binaryMatrix BinaryMatrix) Dimensions() []int {
	ret := []int{len(binaryMatrix.matrix), len(binaryMatrix.matrix[0])}
	return ret
}

func main() {
	matrix := [][]int{{0, 0, 1, 1, 1, 1, 1, 1}, {0, 0, 0, 0, 0, 1, 1, 1}, {0, 0, 0, 1, 1, 1, 1, 1}, {0, 0, 0, 0, 0, 1, 1, 1}, {0, 0, 0, 1, 1, 1, 1, 1}, {0, 0, 1, 1, 1, 1, 1, 1}, {0, 0, 0, 0, 1, 1, 1, 1}, {0, 1, 1, 1, 1, 1, 1, 1}}

	obj := BinaryMatrix{matrix}
	fmt.Println(leftMostColumnWithOne(obj))
}

var index int

func leftMostColumnWithOne(binaryMatrix BinaryMatrix) int {
	dim := binaryMatrix.Dimensions()
	row := dim[0]
	col := dim[1]
	crRow := 0
	crCol := col - 1
	index = col
	traverse(binaryMatrix, crRow, crCol, row, col)
	if index == col {
		return -1
	}
	fmt.Println(count)
	return index
}

var count int

func traverse(binaryMatrix BinaryMatrix, crRow int, crCol int, row int, col int) {
	for crRow < row && crCol > -1 {
		val := binaryMatrix.Get(crRow, crCol)
		count++
		if val == 1 {
			if index > crCol {
				index = crCol
			}
			//can go left
			traverse(binaryMatrix, crRow, crCol-1, row, col)

		} else {
			// can go down only
			traverse(binaryMatrix, crRow+1, crCol, row, col)
		}
		break
	}
}
