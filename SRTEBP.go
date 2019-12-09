package main

import "fmt"

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

func intMax3(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func intMin3(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(grid [][]int, x, y int) (int, int, int, int) {
	xmin, xmax, ymin, ymax := x, x, y, y
	grid[x][y] = 0
	for i := 0; i < 4; i++ {
		p, q := x+dx[i], y+dy[i]
		if p >= 0 && p < len(grid) && q >= 0 && q < len(grid[0]) && grid[p][q] == 1 {
			xn, xm, yn, ym := dfs(grid, p, q)
			xmin = intMin3(xmin, xn)
			xmax = intMax3(xmax, xm)
			ymin = intMin3(ymin, yn)
			ymax = intMax3(ymax, ym)
		}
	}
	return xmin, xmax, ymin, ymax
}

func srtemp(grid [][]int, x, y int) int {
	xmin, xmax, ymin, ymax := dfs(grid, x, y)
	return (xmax - xmin + 1) * (ymax - ymin + 1)
}

func main() {
	grid := [][]int{
		{0, 0, 1, 1, 0},
		{1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0},
	}
	fmt.Println(srtemp(grid, 1, 2))
}
