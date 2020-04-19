package main

var min int

func minPathSum(matrix [][]int) int {
	min = 1000000000000000000
	minSum(matrix, 0, 0, len(matrix), len(matrix[0]), matrix[0][0])
	return min
}

func minSum(matrix [][]int, i int, j int, row int, col int, sum int) {
	if i == row-1 && j == col-1 {
		if sum < min {
			min = sum
		}
		return
	}
	if isValid(matrix, i, j+1, row, col, sum) {
		minSum(matrix, i, j+1, row, col, sum+matrix[i][j+1])
	}

	if isValid(matrix, i+1, j, row, col, sum) {
		minSum(matrix, i+1, j, row, col, sum+matrix[i+1][j])
	}
}

func isValid(matrix [][]int, i int, j int, row int, col int, sum int) bool {
	if i >= 0 && i < row && j >= 0 && j < col && sum < min {
		return true
	}
	return false
}

func main() {

}
