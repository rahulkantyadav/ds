package main

import "fmt"

func main() {
	//arr := [][]byte{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}
	arr := [][]byte{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}}
	fmt.Println(numIslands(arr))
}

var dir = [4][2]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func numIslands(grid [][]byte) int {
	count := 0
	row := len(grid)
	col := len(grid[0])

	rowQueue := make([]int, 0, row*col)
	colQueue := make([]int, 0, row*col)

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == 1 {
				count++
				rowQueue = append(rowQueue, i)
				colQueue = append(colQueue, j)
				updateConnectedLands(grid, rowQueue, colQueue)
			}
		}
	}
	return count
}

func updateConnectedLands(grid [][]byte, rowQueue []int, colQueue []int) {
	if len(rowQueue) <= 0 {
		return
	}

	row := rowQueue[0]
	col := colQueue[0]

	rowQueue = rowQueue[1:]
	colQueue = colQueue[1:]

	grid[row][col] = 0

	for i := 0; i < len(dir); i++ {
		newCol := col + dir[i][0]
		newRow := row + dir[i][1]
		if valid(newRow, newCol, grid) {
			rowQueue = append(rowQueue, newRow)
			colQueue = append(colQueue, newCol)
			updateConnectedLands(grid, rowQueue, colQueue)
		}

	}

}

func valid(newRow int, newCol int, grid [][]byte) bool {
	row := len(grid)
	col := len(grid[0])
	if newRow >= 0 && newRow < row && newCol >= 0 && newCol < col && grid[newRow][newCol] == 1 {
		return true
	}
	return false
}

func enqueue(queue []int, item int) {
	queue = append(queue, item)
}

func dequeue(queue []int) int {
	val := queue[0]
	queue = queue[1:]
	return val
}
