package main

import (
	"fmt"
	"strconv"
)

type Queue struct {
	q    []interface{}
	head int
}

func NewQueue() *Queue {
	Q := &Queue{
		q:    []interface{}{},
		head: 0,
	}
	return Q
}

func (Q *Queue) Front() interface{} {
	return Q.q[Q.head]
}

func (Q *Queue) Empty() bool {
	return Q.head >= len(Q.q)
}

func (Q *Queue) Push(v interface{}) {
	Q.q = append(Q.q, v)
}

func (Q *Queue) Pop() interface{} {
	tmp := Q.q[Q.head]
	Q.head++
	return tmp
}

func decode(s string) (data [][]int) {
	if len(s) != 6 {
		return
	}
	data = make([][]int, 2)
	k := 0
	for i := 0; i < 2; i++ {
		data[i] = make([]int, 3)
		for j := 0; j < 3; j++ {
			data[i][j], _ = strconv.Atoi(string(s[k]))
			k++
		}
	}
	return
}

func encode(data [][]int) string {
	s := ""
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			s += strconv.Itoa(data[i][j])
		}
	}
	return s
}

func sp(data [][]int) int {

	targetG := [][]int{
		{1, 2, 3},
		{4, 5, 0},
	}
	target := encode(targetG)

	dx := []int{1, -1, 0, 0}
	dy := []int{0, 0, 1, -1}
	dist := map[string]int{}
	u := encode(data)
	dist[u] = 0
	que := NewQueue()
	que.Push(u)
	for !que.Empty() {
		u := que.Front().(string)
		que.Pop()

		if u == target {
			return dist[u]
		}

		g := decode(u)
		x := 0
		y := 0
		for i := 0; i < 2; i++ {
			for j := 0; j < 3; j++ {
				if g[i][j] == 0 {
					x, y = i, j
				}
			}
		}

		for k := 0; k < 4; k++ {
			p, q := x+dx[k], y+dy[k]
			if p >= 0 && p < 2 && q >= 0 && q < 3 {
				g[x][y], g[p][q] = g[p][q], g[x][y]
				v := encode(g)
				if _, ok := dist[v]; !ok {
					dist[v] = dist[u] + 1
					que.Push(v)
				}
				g[x][y], g[p][q] = g[p][q], g[x][y]
			}
		}
	}
	return -1
}

func main() {
	start := [][]int{
		{2, 3, 5},
		{1, 4, 0},
	}
	fmt.Println(sp(start))
}
