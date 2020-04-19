package main

func rotatematrix(matrix [][]int) {
	maxRow := len(matrix)
	maxCol := len(matrix[0])
	iteration := len(matrix) / 2

	for i := 0; i < iteration; i++ {
		startRow := i
		endRow := maxRow - i
		startCol := i
		endCol := maxCol - i

		for j := startCol; j < endCol; j++ {
			matrix[i][j] = doRotate(matrix, i, j, startRow, endRow, startCol, endCol, 0)
		}
	}

}

func doRotate(matrix [][]int, i int, j int, startRow int, endRow int, startCol int, endCol int, rotationCount int) int {
	var nextRow int
	var nextCol int
	var temp int

	temp = matrix[i][j]

	if rotationCount == 0 {
		nextRow = endRow
		nextCol = startCol + (i - startRow)
	} else if rotationCount == 1 {
		nextRow = endRow - (j - startCol)
		nextCol = endCol
	} else if rotationCount == 2 {
		nextRow = startRow
		nextCol = startCol + (i - startRow)
	} else if rotationCount == 3 {
		nextRow = startCol + (endCol - j)
		nextCol = startCol
		matrix[i][j] = matrix[nextRow][nextCol]
		return temp
	}
	matrix[i][j] = matrix[nextRow][nextCol]
	return doRotate(matrix, nextRow, nextCol, startRow, endRow, startCol, endCol, rotationCount+1)
}

func main() {

}
