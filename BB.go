package main

import "fmt"

func intMax4(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func call2(values []int, l, r int, mem [][]int) int {
	if r-l <= 1 {
		return 0
	}
	if mem[l][r] != -1 {
		return mem[l][r]
	}

	mem[l][r] = 0
	for i := l + 1; i < r; i++ {
		tmp := values[l]*values[i]*values[r] + call2(values, l, i, mem) + call2(values, i, r, mem)
		mem[l][r] = intMax4(mem[l][r], tmp)
	}

	return mem[l][r]
}

func bb(values []int) int {
	newval := []int{}
	newval = append(newval, 1)
	for _, v := range values {
		newval = append(newval, v)
	}
	newval = append(newval, 1)
	mem := make([][]int, len(newval))
	for i := range mem {
		mem[i] = make([]int, len(newval))
		for j := range mem[i] {
			mem[i][j] = -1
		}
	}
	return call2(newval, 0, len(newval)-1, mem)
}

func main() {
	values := []int{3, 1, 5, 8}
	fmt.Println(bb(values))
}
