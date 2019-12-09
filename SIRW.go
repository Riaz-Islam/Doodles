package main

import "fmt"

type dsu struct {
	par  []int
	size []int
}

func dsuInit(n int) *dsu {
	d := &dsu{
		par:  make([]int, n+1),
		size: make([]int, n+1),
	}
	for i := 0; i <= n; i++ {
		d.par[i] = i
		d.size[i] = 1
	}
	return d
}

func (d *dsu) Find(u int) int {
	if d.par[u] == u {
		return u
	}
	d.par[u] = d.Find(d.par[u])
	return d.par[u]
}

func (d *dsu) Join(u, v int) bool {
	u = d.Find(u)
	v = d.Find(v)
	if u == v {
		return false
	}
	d.par[v] = u
	d.size[u] += d.size[v]
	return true
}

func (d *dsu) InSameSet(u, v int) bool {
	u = d.Find(u)
	v = d.Find(v)
	return u == v
}

func sirw(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	r := len(grid)
	c := len(grid[0])
	n := r * c

	d := dsuInit(n)
	pos := make([]int, n+1)
	for i := range grid {
		for j := range grid[i] {
			pos[grid[i][j]] = i*c + j
		}
	}
	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}

	for i := 0; i < n; i++ {
		x, y := pos[i]/c, pos[i]%c
		for k := 0; k < 4; k++ {
			p, q := x+dx[k], y+dy[k]
			if p >= 0 && p < r && q >= 0 && q < c && grid[p][q] < grid[x][y] {
				d.Join(pos[i], p*c+q)
			}
		}
		if d.InSameSet(0, (r-1)*c+c-1) {
			return i
		}
	}
	return n
}

func main() {
	grid := [][]int{
		{0, 2, 3, 4, 5},
		{18, 19, 20, 24, 1},
		{12, 13, 14, 15, 16},
		{11, 17, 21, 22, 23},
		{6, 7, 8, 9, 10},
	}
	fmt.Println(sirw(grid))
}
